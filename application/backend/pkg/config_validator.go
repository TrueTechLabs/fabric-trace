package pkg

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// ValidateConfig 验证配置是否正确
func ValidateConfig() error {
	appMode := viper.GetString("app.mode")

	// 生产环境必须检查的配置
	if appMode == "prod" || appMode == "production" {
		// 检查 JWT 密钥
		jwtSecret := viper.GetString("jwt.secret")
		if jwtSecret == "" || jwtSecret == "testsecret" {
			return fmt.Errorf("production mode requires a strong jwt.secret (current: %q)", jwtSecret)
		}
		if len(jwtSecret) < 32 {
			return fmt.Errorf("production mode requires jwt.secret to be at least 32 characters (current: %d)", len(jwtSecret))
		}

		// 检查数据库密码
		mysqlPassword := viper.GetString("mysql.password")
		if mysqlPassword == "" || mysqlPassword == "fabrictrace" {
			return fmt.Errorf("production mode requires a strong mysql.password (current: %q)", mysqlPassword)
		}

		// 检查应用 salt
		appSalt := viper.GetString("app.salt")
		if appSalt == "" || appSalt == "testsalt" {
			return fmt.Errorf("production mode requires a strong app.salt (current: %q)", appSalt)
		}

		// 检查 CORS 配置
		corsOrigins := viper.GetStringSlice("cors.allowed_origins")
		if len(corsOrigins) == 0 {
			return fmt.Errorf("production mode requires cors.allowed_origins to be configured")
		}
		for _, origin := range corsOrigins {
			if origin == "*" || origin == "" {
				return fmt.Errorf("production mode should not use wildcard CORS origin (current: %q)", origin)
			}
		}
	}

	// 验证数据库名称（防止SQL注入）
	dbName := viper.GetString("mysql.db")
	if dbName == "" {
		return fmt.Errorf("mysql.db cannot be empty")
	}
	if !isValidIdentifier(dbName) {
		return fmt.Errorf("mysql.db contains invalid characters: %q", dbName)
	}

	// 验证 bcrypt 成本因子
	bcryptCost := viper.GetInt("security.bcrypt_cost")
	if bcryptCost < 4 {
		return fmt.Errorf("security.bcrypt_cost must be at least 4 (current: %d)", bcryptCost)
	}
	if bcryptCost > 31 {
		return fmt.Errorf("security.bcrypt_cost must be at most 31 (current: %d)", bcryptCost)
	}

	// 验证日志级别
	logLevel := viper.GetString("log.level")
	validLevels := map[string]bool{
		"debug": true,
		"info":  true,
		"warn":  true,
		"error": true,
	}
	if !validLevels[logLevel] {
		return fmt.Errorf("log.level must be one of: debug, info, warn, error (current: %q)", logLevel)
	}

	return nil
}

// isValidIdentifier 检查字符串是否是有效的 SQL 标识符
func isValidIdentifier(s string) bool {
	if s == "" {
		return false
	}
	// 只允许字母、数字和下划线，且不能以数字开头
	for i, r := range s {
		if i == 0 && (r >= '0' && r <= '9') {
			return false
		}
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_') {
			return false
		}
	}
	return true
}

// IsStrongPassword 检查密码强度
func IsStrongPassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}

	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, r := range password {
		switch {
		case r >= 'A' && r <= 'Z':
			hasUpper = true
		case r >= 'a' && r <= 'z':
			hasLower = true
		case r >= '0' && r <= '9':
			hasDigit = true
		case strings.ContainsAny("!@#$%^&*()_+-=[]{}|;:,.<>?`", string(r)):
			hasSpecial = true
		}
	}

	required := 0
	if hasUpper {
		required++
	}
	if hasLower {
		required++
	}
	if hasDigit {
		required++
	}
	if hasSpecial {
		required++
	}

	if required < 3 {
		return fmt.Errorf("password must contain at least 3 of: uppercase, lowercase, digits, special characters")
	}

	return nil
}
