package pkg

import (
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/spf13/viper"
)

// LogLevel 日志级别
type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

var (
	currentLogLevel int32 = int32(LevelInfo) // 使用 atomic 操作
	logMutex        sync.Mutex
	logFile         *os.File
)

// InitLogger 初始化日志系统
func InitLogger(level string, logFilePath string) error {
	logMutex.Lock()
	defer logMutex.Unlock()

	return initLoggerLocked(level, logFilePath)
}

// InitLoggerFromConfig 从配置文件初始化日志
func InitLoggerFromConfig() error {
	level := viper.GetString("log.level")
	logPath := viper.GetString("log.path")

	return InitLogger(level, logPath)
}

// initLoggerLocked 在已持有锁的情况下初始化日志
func initLoggerLocked(level string, logFilePath string) error {
	var newLevel LogLevel
	switch level {
	case "debug":
		newLevel = LevelDebug
	case "info":
		newLevel = LevelInfo
	case "warn":
		newLevel = LevelWarn
	case "error":
		newLevel = LevelError
	default:
		newLevel = LevelInfo
	}

	// 使用 atomic 操作更新日志级别
	atomic.StoreInt32(&currentLogLevel, int32(newLevel))

	if logFilePath != "" {
		// 关闭旧的日志文件（如果存在）
		if logFile != nil {
			if err := logFile.Close(); err != nil {
				// 记录到标准错误，因为日志文件可能已经不可用
				log.Printf("failed to close old log file: %v", err)
			}
			logFile = nil
		}

		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			// 重置输出到标准错误
			log.SetOutput(os.Stderr)
			return fmt.Errorf("failed to open log file: %w", err)
		}
		logFile = file
		log.SetOutput(file)
	} else {
		// 如果没有指定日志文件，确保输出到标准输出
		if logFile != nil {
			if err := logFile.Close(); err != nil {
				log.Printf("failed to close log file: %v", err)
			}
			logFile = nil
		}
		log.SetOutput(os.Stdout)
	}

	return nil
}

// CloseLogger 关闭日志文件
func CloseLogger() {
	logMutex.Lock()
	defer logMutex.Unlock()

	if logFile != nil {
		if err := logFile.Close(); err != nil {
			// 使用标准错误记录关闭失败
			log.Printf("failed to close log file: %v", err)
		}
		logFile = nil
		// 重置输出到标准输出
		log.SetOutput(os.Stdout)
	}
}

// Field 日志字段
type Field struct {
	Key   string
	Value interface{}
}

// String 创建字符串字段
func String(key, value string) Field {
	return Field{Key: key, Value: value}
}

// Int 创建整数字段
func Int(key string, value int) Field {
	return Field{Key: key, Value: value}
}

// Err 创建错误字段
func Err(err error) Field {
	return Field{Key: "error", Value: err}
}

// logMessage 内部日志函数
func logMessage(level LogLevel, msg string, fields ...Field) {
	// 使用 atomic 操作读取当前日志级别
	current := atomic.LoadInt32(&currentLogLevel)
	if int32(level) < current {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	levelStr := ""

	switch level {
	case LevelDebug:
		levelStr = "DEBUG"
	case LevelInfo:
		levelStr = "INFO"
	case LevelWarn:
		levelStr = "WARN"
	case LevelError:
		levelStr = "ERROR"
	}

	logMsg := fmt.Sprintf("[%s] %s: %s", timestamp, levelStr, msg)

	for _, field := range fields {
		logMsg += fmt.Sprintf(" %s=%v", field.Key, field.Value)
	}

	log.Println(logMsg)
}

// Debug 调试日志
func Debug(msg string, fields ...Field) {
	logMessage(LevelDebug, msg, fields...)
}

// Info 信息日志
func Info(msg string, fields ...Field) {
	logMessage(LevelInfo, msg, fields...)
}

// Warn 警告日志
func Warn(msg string, fields ...Field) {
	logMessage(LevelWarn, msg, fields...)
}

// Error 错误日志
func Error(msg string, fields ...Field) {
	logMessage(LevelError, msg, fields...)
}

// Int64 创建 int64 字段
func Int64(key string, value int64) Field {
	return Field{Key: key, Value: value}
}

// Bool 创建布尔字段
func Bool(key string, value bool) Field {
	return Field{Key: key, Value: value}
}

// Duration 创建时间段字段
func Duration(key string, value time.Duration) Field {
	return Field{Key: key, Value: value}
}
