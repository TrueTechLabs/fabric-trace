package pkg

import (
	"fmt"
	"net/http"
)

// AppError 应用错误类型
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"error,omitempty"`
}

func (e AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// 错误码常量
const (
	ErrCodeSuccess           = 200
	ErrCodeBadRequest        = 400
	ErrCodeUnauthorized      = 401
	ErrCodeNotFound          = 404
	ErrCodeInternalError     = 500
	ErrCodeDatabaseError     = 5001
	ErrCodeBlockchainError   = 5002
	ErrCodeValidationError   = 4001
	ErrCodeUserExists        = 4002
	ErrCodeUserNotFound      = 4003
	ErrCodeInvalidPassword   = 4004
	ErrCodeInvalidTraceCode  = 4005
	ErrCodeTraceCodeNotFound = 4006
)

// NewError 创建新的应用错误
func NewError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// 预定义错误
var (
	ErrBadRequest        = NewError(ErrCodeBadRequest, "请求参数错误", nil)
	ErrUnauthorized      = NewError(ErrCodeUnauthorized, "未授权访问", nil)
	ErrNotFound          = NewError(ErrCodeNotFound, "资源不存在", nil)
	ErrInternalError     = NewError(ErrCodeInternalError, "服务器内部错误", nil)
	ErrDatabaseError     = NewError(ErrCodeDatabaseError, "数据库操作失败", nil)
	ErrBlockchainError   = NewError(ErrCodeBlockchainError, "区块链操作失败", nil)
	ErrUserExists        = NewError(ErrCodeUserExists, "用户已存在", nil)
	ErrUserNotFound      = NewError(ErrCodeUserNotFound, "用户不存在", nil)
	ErrInvalidPassword   = NewError(ErrCodeInvalidPassword, "密码错误", nil)
	ErrInvalidTraceCode  = NewError(ErrCodeInvalidTraceCode, "溯源码格式错误", nil)
	ErrTraceCodeNotFound = NewError(ErrCodeTraceCodeNotFound, "溯源码不存在", nil)
)

// WithError 为错误添加底层错误
func (e *AppError) WithError(err error) *AppError {
	return &AppError{
		Code:    e.Code,
		Message: e.Message,
		Err:     err,
	}
}

// WithMessage 为错误添加自定义消息
func (e *AppError) WithMessage(msg string) *AppError {
	return &AppError{
		Code:    e.Code,
		Message: msg,
		Err:     e.Err,
	}
}

// ValidationError 验证错误转应用错误
func ValidationErrorToAppError(err *ValidationError) *AppError {
	return &AppError{
		Code:    ErrCodeValidationError,
		Message: err.Error(),
	}
}

// HTTPStatus 获取 HTTP 状态码
func (e AppError) HTTPStatus() int {
	switch e.Code {
	case ErrCodeSuccess:
		return http.StatusOK
	case ErrCodeBadRequest, ErrCodeValidationError, ErrCodeUserExists,
	     ErrCodeUserNotFound, ErrCodeInvalidPassword, ErrCodeInvalidTraceCode, ErrCodeTraceCodeNotFound:
		return http.StatusBadRequest
	case ErrCodeUnauthorized:
		return http.StatusUnauthorized
	case ErrCodeNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
