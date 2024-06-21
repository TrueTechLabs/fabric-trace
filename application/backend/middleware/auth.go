package middleware

import (
	"backend/pkg"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer xxxxxxx.xxx.xxx  / X-TOKEN: xxx.xxx.xx
		// 这里的具体实现方式要依据你的实际业务情况决定
		// authHeader := c.Request.Header.Get("Authorization")
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "请求未携带token，无权限访问1",
			},
			)
			c.Abort()
			return
		}
		mc, err := pkg.ParseToken(authHeader)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "请求未携带token，无权限访问3",
				"data": err.Error(),
			},
			)
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("userID", mc.UserID)

		c.Next() // 后续的处理请求的函数中 可以用过c.Get(CtxUserIDKey) 来获取当前请求的用户信息
	}
}
