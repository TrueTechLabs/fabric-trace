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
	}
	// 初始化数据库
	if err := pkg.MysqlInit(); err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
	}
	// 注册路由
	r := router.SetupRouter()

	// 启动服务
	r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))

}
