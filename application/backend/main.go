package main

import (
	"backend/pkg"
	"backend/router"
	setting "backend/settings"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	// 加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Printf("init settings failed,err:%v\n", err)
		return
	}

	// 验证配置
	if err := pkg.ValidateConfig(); err != nil {
		fmt.Printf("config validation failed: %v\n", err)
		return
	}

	// 初始化日志系统
	if err := pkg.InitLoggerFromConfig(); err != nil {
		fmt.Printf("init logger failed: %v\n", err)
		return
	}
	defer pkg.CloseLogger()

	pkg.Info("Fabric-Trace server starting",
		pkg.String("mode", viper.GetString("app.mode")),
		pkg.String("port", viper.GetString("app.port")))

	// 初始化数据库
	if err := pkg.MysqlInit(); err != nil {
		pkg.Error("init mysql failed", pkg.Err(err))
		return
	}
	defer pkg.CloseDB()

	// 初始化 Fabric 连接池
	pkg.Info("Initializing Fabric gateway pool")
	if err := pkg.WarmupGatewayPool(); err != nil {
		pkg.Warn("Fabric gateway pool warmup failed, will retry on first request",
			pkg.Err(err))
	}
	defer pkg.CloseGatewayPool()

	// 注册路由
	r := router.SetupRouter()

	// 启动服务
	port := viper.GetInt("app.port")
	pkg.Info("Server started", pkg.String("address", fmt.Sprintf(":%d", port)))
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		pkg.Error("Server failed to start", pkg.Err(err))
	}

	// 清理资源
	pkg.Info("Server shutting down")
}
