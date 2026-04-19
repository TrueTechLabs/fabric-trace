package controller

import (
	"backend/pkg"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// HealthResponse 健康检查响应
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp string    `json:"timestamp"`
	Uptime    string    `json:"uptime"`
	Services  ServiceStatus `json:"services"`
}

// ServiceStatus 服务状态
type ServiceStatus struct {
	Database   string `json:"database"`
	Blockchain string `json:"blockchain"`
}

var startTime = time.Now()

// Health 健康检查接口
func Health(c *gin.Context) {
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
		Uptime:    time.Since(startTime).String(),
		Services: ServiceStatus{
			Database:   checkDatabase(),
			Blockchain: checkBlockchain(),
		},
	}

	// 如果有服务不健康，返回 503
	if response.Services.Database == "down" || response.Services.Blockchain == "down" {
		response.Status = "unhealthy"
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

// Readiness 就绪检查接口
func Readiness(c *gin.Context) {
	dbStatus := checkDatabase()
	blockchainStatus := checkBlockchain()

	if dbStatus == "up" && blockchainStatus == "up" {
		c.JSON(http.StatusOK, gin.H{
			"status": "ready",
		})
		return
	}

	c.JSON(http.StatusServiceUnavailable, gin.H{
		"status": "not_ready",
		"database":   dbStatus,
		"blockchain": blockchainStatus,
	})
}

// Liveness 存活检查接口
func Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "alive",
	})
}

// checkDatabase 检查数据库连接
func checkDatabase() string {
	db := pkg.DB()
	if db == nil {
		return "down"
	}

	// 尝试 ping 数据库
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := db.PingContext(ctx)
	if err != nil {
		pkg.Error("Database health check failed", pkg.Err(err))
		return "down"
	}

	return "up"
}

// checkBlockchain 检查区块链连接
func checkBlockchain() string {
	// 尝试进行一个简单的链码查询，带5秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	type result struct {
		data string
		err  error
	}

	resultChan := make(chan result, 1)

	go func() {
		data, err := pkg.ChaincodeQuery("GetAllFruitInfo", "")
		resultChan <- result{data, err}
	}()

	select {
	case <-ctx.Done():
		pkg.Error("Blockchain health check timed out")
		return "down"
	case res := <-resultChan:
		if res.err != nil {
			pkg.Error("Blockchain health check failed", pkg.Err(res.err))
			return "down"
		}
		return "up"
	}
}

// GetVersion 获取系统版本信息
func GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":       "fabric-trace",
		"version":    "1.0.0",
		"environment": viper.GetString("app.mode"),
		"go_version": "1.19+",
	})
}

// 别名函数，用于路由兼容
func HealthCheck(c *gin.Context) {
	Health(c)
}

func Version(c *gin.Context) {
	GetVersion(c)
}
