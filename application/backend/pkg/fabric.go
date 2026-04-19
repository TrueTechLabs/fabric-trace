package pkg

// 本文件用于处理链码请求
import (
	"crypto/x509"
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"github.com/spf13/viper"
)

// FabricConfig Fabric 连接配置
type FabricConfig struct {
	MSPID        string
	CryptoPath   string
	CertPath     string
	KeyPath      string
	TLSCertPath  string
	PeerEndpoint string
	GatewayPeer  string
}

var (
	configOnce   sync.Once
	fabricConfig *FabricConfig
)

// GetFabricConfig 从环境变量或配置文件获取 Fabric 配置
func GetFabricConfig() *FabricConfig {
	cryptoPath := os.Getenv("FABRIC_CRYPTO_PATH")
	peerEndpoint := os.Getenv("FABRIC_PEER_ENDPOINT")

	if cryptoPath == "" {
		cryptoPath = "../../blockchain/network/organizations/peerOrganizations/org1.example.com"
	}
	if peerEndpoint == "" {
		peerEndpoint = viper.GetString("fabric.peer_endpoint")
		if peerEndpoint == "" {
			peerEndpoint = "127.0.0.1:7051"
		}
	}

	return &FabricConfig{
		MSPID:        "Org1MSP",
		CryptoPath:   cryptoPath,
		CertPath:     cryptoPath + "/users/User1@org1.example.com/msp/signcerts/User1@org1.example.com-cert.pem",
		KeyPath:      cryptoPath + "/users/User1@org1.example.com/msp/keystore/",
		TLSCertPath:  cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt",
		PeerEndpoint: peerEndpoint,
		GatewayPeer:  "peer0.org1.example.com",
	}
}

// getFabricConfig 获取 Fabric 配置（内部使用，单例）
func getFabricConfig() *FabricConfig {
	configOnce.Do(func() {
		fabricConfig = GetFabricConfig()
	})
	return fabricConfig
}

// 保留原有的常量以保持向后兼容
const (
	mspID        = "Org1MSP"
	cryptoPath   = "../../blockchain/network/organizations/peerOrganizations/org1.example.com"
	certPath     = cryptoPath + "/users/User1@org1.example.com/msp/signcerts/User1@org1.example.com-cert.pem"
	keyPath      = cryptoPath + "/users/User1@org1.example.com/msp/keystore/"
	tlsCertPath  = cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
	peerEndpoint = "127.0.0.1:7051"
	gatewayPeer  = "peer0.org1.example.com"
)

// 链码查询（使用连接池）
func ChaincodeQuery(fcn string, arg string) (string, error) {
	pool := GetGatewayPool()
	contract, err := pool.GetContract()
	if err != nil {
		return "", fmt.Errorf("failed to get contract: %w", err)
	}

	evaluateResult, err := contract.EvaluateTransaction(fcn, arg)
	if err != nil {
		Error("Chaincode query failed",
			String("function", fcn),
			String("arg", arg),
			Err(err))
		return "", fmt.Errorf("failed to evaluate transaction: %w", err)
	}

	Info("Chaincode query success",
		String("function", fcn),
		String("arg", arg),
		String("result", string(evaluateResult)))

	return string(evaluateResult), nil
}

// 链码调用，返回交易ID（使用连接池）
func ChaincodeInvoke(fcn string, args []string) (string, error) {
	pool := GetGatewayPool()
	contract, err := pool.GetContract()
	if err != nil {
		return "", fmt.Errorf("failed to get contract: %w", err)
	}

	Info("Submitting chaincode transaction",
		String("function", fcn),
		String("args", fmt.Sprintf("%v", args)))

	submitResult, commit, err := contract.SubmitAsync(fcn, client.WithArguments(args...))
	if err != nil {
		Error("Failed to submit transaction",
			String("function", fcn),
			Err(err))
		return "", fmt.Errorf("failed to submit transaction asynchronously: %w", err)
	}

	Info("Transaction submitted, waiting for commit",
		String("txid", string(submitResult)))

	if commitStatus, err := commit.Status(); err != nil {
		Error("Failed to get commit status",
			String("txid", commit.TransactionID()),
			Err(err))
		return "", fmt.Errorf("failed to get commit status: %w", err)
	} else if !commitStatus.Successful {
		Error("Transaction commit failed",
			String("txid", commitStatus.TransactionID),
			Int("status", int(commitStatus.Code)))
		return "", fmt.Errorf("transaction %s failed to commit with status: %d", commitStatus.TransactionID, int32(commitStatus.Code))
	}

	Info("Transaction committed successfully",
		String("txid", commit.TransactionID()))

	return commit.TransactionID(), nil
}

func GetContract() (*client.Contract, *grpc.ClientConn, *client.Gateway) {
	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := newGrpcConnection()
	// defer clientConnection.Close()

	id := newIdentity()
	sign := newSign()

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	// defer gw.Close()

	// Override default values for chaincode and channel name as they may differ in testing contexts.
	chaincodeName := "trace"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	channelName := "mychannel"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	return contract, clientConnection, gw
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func newIdentity() *identity.X509Identity {
	cfg := getFabricConfig()
	certificate, err := loadCertificate(cfg.CertPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(cfg.MSPID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign() identity.Sign {
	cfg := getFabricConfig()
	files, err := os.ReadDir(cfg.KeyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := os.ReadFile(path.Join(cfg.KeyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection() *grpc.ClientConn {
	cfg := getFabricConfig()
	certificate, err := loadCertificate(cfg.TLSCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, cfg.GatewayPeer)

	connection, err := grpc.Dial(cfg.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}
