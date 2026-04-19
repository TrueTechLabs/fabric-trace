package router

// 路由文件
import (
	con "backend/controller"
	"backend/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	// 全局错误处理中间件
	r.Use(middleware.ErrorHandler())

	// 从配置文件读取 CORS 允许的来源
	allowedOrigins := viper.GetStringSlice("cors.allowed_origins")
	// 如果配置为空或只有 "*"，则在开发环境允许本地地址
	if len(allowedOrigins) == 0 || (len(allowedOrigins) == 1 && allowedOrigins[0] == "*") {
		if viper.GetString("app.mode") == "dev" {
			allowedOrigins = []string{
				"http://localhost:9528",
				"http://127.0.0.1:9528",
				"http://localhost:9090",
				"http://127.0.0.1:9090",
			}
		} else {
			// 生产环境不应使用 "*"，需要明确配置
			allowedOrigins = []string{}
		}
	}

	// 解决跨域问题
	r.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,                                    // 从配置读取允许的来源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},         // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                        // 暴露的响应头
		AllowCredentials: true,                                              // 允许传递凭据（例如 Cookie）
		MaxAge:           12 * time.Hour,                                    // 预检请求的有效期
	}))
	// 设置静态文件目录
	r.Static("/static", "./dist/static")
	r.LoadHTMLGlob("dist/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	// 测试GET请求
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 健康检查接口
	r.GET("/health", con.HealthCheck)
	r.GET("/ready", con.Readiness)
	r.GET("/live", con.Liveness)
	r.GET("/version", con.Version)

	//注册
	r.POST("/register", con.Register)
	//登录
	r.POST("/login", con.Login)
	//登出
	r.POST("/logout", con.Logout)
	//查询用户的类型
	r.POST("/getInfo", middleware.JWTAuthMiddleware(), con.GetInfo)
	//农产品上链
	r.POST("/uplink", middleware.JWTAuthMiddleware(), con.Uplink)
	// 获取农产品的上链信息
	r.POST("/getFruitInfo", con.GetFruitInfo)
	// 获取用户的农产品ID列表
	r.POST("/getFruitList", middleware.JWTAuthMiddleware(), con.GetFruitList)
	// 获取所有的农产品信息
	r.POST("/getAllFruitInfo", middleware.JWTAuthMiddleware(), con.GetAllFruitInfo)
	// 获取农产品上链历史(溯源)
	r.POST("/getFruitHistory", middleware.JWTAuthMiddleware(), con.GetFruitHistory)
	r.GET("/getImg/:filename", con.GetImg) // 获取图片
	return r
}
