package pkg

import (
	"context"
	"crypto/x509"
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
)

// GatewayPool Fabric Gateway 连接池
type GatewayPool struct {
	mu             sync.RWMutex
	connection     *grpc.ClientConn
	gateway        *client.Gateway
	contract       *client.Contract
	chaincodeName  string
	channelName    string
	lastHealthCheck time.Time
	healthInterval  time.Duration
}

var (
	poolOnce sync.Once
	pool     *GatewayPool
)

// GetGatewayPool 获取 Gateway 连接池单例
func GetGatewayPool() *GatewayPool {
	poolOnce.Do(func() {
		pool = &GatewayPool{
			chaincodeName:  "trace",
			channelName:    "mychannel",
			healthInterval: 30 * time.Second,
		}
		pool.connect()
	})
	return pool
}

// connect 建立 Fabric 连接
func (p *GatewayPool) connect() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.connectLocked()
}

// connectLocked 在已持有锁的情况下建立连接
func (p *GatewayPool) connectLocked() error {
	cfg := getFabricConfig()

	// 创建 gRPC 连接
	clientConnection, err := p.newGrpcConnection()
	if err != nil {
		return fmt.Errorf("failed to create gRPC connection: %w", err)
	}

	// 创建身份和签名
	id, err := p.newIdentity(cfg.CertPath)
	if err != nil {
		clientConnection.Close()
		return fmt.Errorf("failed to create identity: %w", err)
	}

	sign, err := p.newSign(cfg.KeyPath)
	if err != nil {
		clientConnection.Close()
		return fmt.Errorf("failed to create sign: %w", err)
	}

	// 创建 Gateway 连接
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		clientConnection.Close()
		return fmt.Errorf("failed to connect to gateway: %w", err)
	}

	// 获取网络和合约
	network := gw.GetNetwork(p.channelName)
	contract := network.GetContract(p.chaincodeName)

	p.connection = clientConnection
	p.gateway = gw
	p.contract = contract
	p.lastHealthCheck = time.Now()

	Info("Gateway pool initialized",
		String("peer", cfg.PeerEndpoint),
		String("channel", p.channelName),
		String("chaincode", p.chaincodeName))

	return nil
}

// GetContract 获取链码合约（注意：返回的合约仅在当前调用期间有效）
func (p *GatewayPool) GetContract() (*client.Contract, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// 检查连接是否可用，如果为nil则尝试重连
	if p.gateway == nil {
		Warn("Gateway not initialized, attempting reconnect")
		if err := p.reconnectLocked(); err != nil {
			return nil, fmt.Errorf("gateway not initialized and reconnect failed: %w", err)
		}
		// 重连成功后继续获取contract
	}

	// 每次都从 gateway 获取新的 contract，避免竞态条件
	network := p.gateway.GetNetwork(p.channelName)
	contract := network.GetContract(p.chaincodeName)

	// 定期健康检查
	if time.Since(p.lastHealthCheck) > p.healthInterval {
		if err := p.healthCheckLocked(); err != nil {
			return nil, fmt.Errorf("health check failed: %w", err)
		}
		// 健康检查后重新获取 contract，因为 gateway 可能已重连
		if p.gateway == nil {
			return nil, fmt.Errorf("gateway not initialized after health check")
		}
		network = p.gateway.GetNetwork(p.channelName)
		contract = network.GetContract(p.chaincodeName)
	}

	return contract, nil
}

// healthCheck 健康检查
func (p *GatewayPool) healthCheck() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.healthCheckLocked()
}

// healthCheckLocked 在已持有锁的情况下进行健康检查
func (p *GatewayPool) healthCheckLocked() error {
	// 检查连接是否存在
	if p.contract == nil || p.gateway == nil || p.connection == nil {
		Warn("Gateway connection not initialized, attempting reconnect")
		return p.reconnectLocked()
	}

	// 检查 gRPC 连接状态
	state := p.connection.GetState()
	if state != connectivity.Ready && state != connectivity.Idle {
		Warn("Gateway connection not ready, attempting reconnect",
			String("state", state.String()))
		return p.reconnectLocked()
	}

	// 使用 WaitForStateChange 验证连接在短时间内保持稳定
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 等待连接状态变化（如果没有变化则超时，说明连接稳定）
	ch := make(chan struct{})
	go func() {
		p.connection.WaitForStateChange(ctx, state)
		close(ch)
	}()

	select {
	case <-ch:
		// 连接状态发生变化，可能不稳定
		Warn("Gateway connection state changed during health check",
			String("old_state", state.String()),
			String("new_state", p.connection.GetState().String()))
		return p.reconnectLocked()
	case <-ctx.Done():
		// 超时，说明连接状态稳定
	}

	p.lastHealthCheck = time.Now()

	Debug("Gateway health check passed",
		String("state", p.connection.GetState().String()))

	return nil
}

// reconnect 重新连接
func (p *GatewayPool) reconnect() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.reconnectLocked()
}

// reconnectLocked 在已持有锁的情况下重新连接
func (p *GatewayPool) reconnectLocked() error {
	Info("Attempting to reconnect to Fabric gateway")

	// 关闭旧连接
	if p.gateway != nil {
		p.gateway.Close()
		p.gateway = nil
	}
	if p.connection != nil {
		p.connection.Close()
		p.connection = nil
	}
	p.contract = nil

	return p.connectLocked()
}

// Close 关闭连接池
func (p *GatewayPool) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	var err error
	if p.gateway != nil {
		err = p.gateway.Close()
		p.gateway = nil
	}
	if p.connection != nil {
		closeErr := p.connection.Close()
		if err == nil {
			err = closeErr
		}
		p.connection = nil
	}
	p.contract = nil

	if err == nil {
		Info("Gateway pool closed")
	}

	return err
}

// WarmupGatewayPool 预热连接池（在应用启动时调用）
func WarmupGatewayPool() error {
	pool := GetGatewayPool()

	// 尝试获取合约来初始化连接
	contract, err := pool.GetContract()
	if err != nil {
		return fmt.Errorf("failed to warmup gateway pool: %w", err)
	}

	Info("Gateway pool warmed up successfully",
		String("channel", pool.channelName),
		String("chaincode", pool.chaincodeName))

	// contract 的引用只是为了触发初始化，不需要使用它
	_ = contract
	return nil
}

// CloseGatewayPool 关闭连接池（供 main 调用）
func CloseGatewayPool() error {
	pool := GetGatewayPool()
	return pool.Close()
}

// newGrpcConnection 创建 gRPC 连接
func (p *GatewayPool) newGrpcConnection() (*grpc.ClientConn, error) {
	cfg := getFabricConfig()
	certificate, err := p.loadCertificate(cfg.TLSCertPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, cfg.GatewayPeer)

	connection, err := grpc.Dial(cfg.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection: %w", err)
	}

	return connection, nil
}

// newIdentity 创建客户端身份
func (p *GatewayPool) newIdentity(certPath string) (*identity.X509Identity, error) {
	certificate, err := p.loadCertificate(certPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load certificate: %w", err)
	}

	cfg := getFabricConfig()
	id, err := identity.NewX509Identity(cfg.MSPID, certificate)
	if err != nil {
		return nil, fmt.Errorf("failed to create identity: %w", err)
	}

	return id, nil
}

// newSign 创建签名函数
func (p *GatewayPool) newSign(keyPath string) (identity.Sign, error) {
	cfg := getFabricConfig()
	files, err := os.ReadDir(cfg.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key directory: %w", err)
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("no private key files found in directory: %s", cfg.KeyPath)
	}

	privateKeyPEM, err := os.ReadFile(path.Join(cfg.KeyPath, files[0].Name()))
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create sign: %w", err)
	}

	return sign, nil
}

// loadCertificate 加载证书
func (p *GatewayPool) loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}
