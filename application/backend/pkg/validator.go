package pkg

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// 验证错误类型
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// 验证规则常量
const (
	MinUsernameLength = 3
	MaxUsernameLength = 20
	MinPasswordLength = 6
	MaxPasswordLength = 50
	TraceCodeLength   = 18
	MaxFieldLength    = 100 // 通用字段最大长度
)

// ValidateUsername 验证用户名
func ValidateUsername(username string) *ValidationError {
	username = strings.TrimSpace(username)
	
	if len(username) < MinUsernameLength {
		return &ValidationError{
			Field:   "username",
			Message: fmt.Sprintf("用户名长度不能少于 %d 位", MinUsernameLength),
		}
	}
	
	if len(username) > MaxUsernameLength {
		return &ValidationError{
			Field:   "username",
			Message: fmt.Sprintf("用户名长度不能超过 %d 位", MaxUsernameLength),
		}
	}
	
	// 只允许字母、数字、下划线
	matched, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", username)
	if !matched {
		return &ValidationError{
			Field:   "username",
			Message: "用户名只能包含字母、数字和下划线",
		}
	}
	
	return nil
}

// ValidatePassword 验证密码强度
func ValidatePassword(password string) *ValidationError {
	if len(password) < MinPasswordLength {
		return &ValidationError{
			Field:   "password",
			Message: fmt.Sprintf("密码长度不能少于 %d 位", MinPasswordLength),
		}
	}
	
	if len(password) > MaxPasswordLength {
		return &ValidationError{
			Field:   "password",
			Message: fmt.Sprintf("密码长度不能超过 %d 位", MaxPasswordLength),
		}
	}
	
	// 检查密码复杂度（至少包含字母和数字）
	hasLetter := false
	hasDigit := false
	
	for _, char := range password {
		if unicode.IsLetter(char) {
			hasLetter = true
		}
		if unicode.IsDigit(char) {
			hasDigit = true
		}
	}
	
	if !hasLetter || !hasDigit {
		return &ValidationError{
			Field:   "password",
			Message: "密码必须同时包含字母和数字",
		}
	}
	
	return nil
}

// ValidateTraceCode 验证溯源码
func ValidateTraceCode(code string) *ValidationError {
	code = strings.TrimSpace(code)

	if len(code) != TraceCodeLength {
		return &ValidationError{
			Field:   "traceability_code",
			Message: fmt.Sprintf("溯源码长度必须为 %d 位", TraceCodeLength),
		}
	}

	// 检查是否为纯数字
	matched, _ := regexp.MatchString("^[0-9]+$", code)
	if !matched {
		return &ValidationError{
			Field:   "traceability_code",
			Message: "溯源码只能包含数字",
		}
	}

	return nil
}

// ValidateTraceabilityCode 验证溯源码（别名函数，兼容现有代码）
func ValidateTraceabilityCode(code string) *ValidationError {
	return ValidateTraceCode(code)
}

// ValidateUserType 验证用户类型
func ValidateUserType(userType string) *ValidationError {
	validTypes := map[string]bool{
		"种植户": true,
		"工厂":   true,
		"运输司机": true,
		"商店":   true,
		"消费者":  true,
	}
	
	if !validTypes[userType] {
		return &ValidationError{
			Field:   "userType",
			Message: "用户类型无效，必须为：种植户、工厂、运输司机、商店或消费者",
		}
	}
	
	return nil
}

// SanitizeInput 清理输入字符串，防止XSS攻击
func SanitizeInput(input string) string {
	// 移除危险的HTML标签
	input = strings.ReplaceAll(input, "<script>", "")
	input = strings.ReplaceAll(input, "</script>", "")
	input = strings.ReplaceAll(input, "<", "&lt;")
	input = strings.ReplaceAll(input, ">", "&gt;")
	
	// 去除首尾空格
	input = strings.TrimSpace(input)
	
	return input
}

// ValidateFormFile 验证上传的文件
func ValidateFormFile(filename string, maxSize int64) *ValidationError {
	// 检查文件扩展名
	ext := strings.ToLower(GetFileExt(filename))
	allowedExts := map[string]bool{
		"jpg":  true,
		"jpeg": true,
		"png":  true,
		"gif":  true,
		"webp": true,
	}
	
	if !allowedExts[ext] {
		return &ValidationError{
			Field:   "file",
			Message: "只支持上传图片格式：jpg、jpeg、png、gif、webp",
		}
	}
	
	return nil
}

// ValidateArg 验证通用参数
func ValidateArg(argName, value string, maxLength int) *ValidationError {
	value = strings.TrimSpace(value)

	if len(value) == 0 {
		return &ValidationError{
			Field:   argName,
			Message: fmt.Sprintf("%s 不能为空", argName),
		}
	}

	if len(value) > maxLength {
		return &ValidationError{
			Field:   argName,
			Message: fmt.Sprintf("%s 长度不能超过 %d 位", argName, maxLength),
		}
	}

	return nil
}

// TrimAndValidate 清理并验证字符串
func TrimAndValidate(value string, maxLength int) (string, error) {
	value = SanitizeInput(value)
	if err := ValidateArg("value", value, maxLength); err != nil {
		return "", err
	}
	return value, nil
}
