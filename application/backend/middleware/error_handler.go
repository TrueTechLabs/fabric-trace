package middleware

import (
	"backend/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler 全局错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 检查是否有错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			
			// 尝试转换为 AppError
			if appErr, ok := err.Err.(*pkg.AppError); ok {
				c.JSON(appErr.HTTPStatus(), gin.H{
					"code":    appErr.Code,
					"message": appErr.Message,
				})
				return
			}

			// 默认错误响应
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    pkg.ErrCodeInternalError,
				"message": "服务器内部错误",
			})
		}
	}
}

// Response 统一成功响应
func Response(c *gin.Context, code int, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

// ErrorResponse 统一错误响应
func ErrorResponse(c *gin.Context, err error) {
	if appErr, ok := err.(*pkg.AppError); ok {
		c.JSON(appErr.HTTPStatus(), gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
		})
		return
	}

	// 默认错误响应
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    pkg.ErrCodeInternalError,
		"message": "服务器内部错误",
	})
}
