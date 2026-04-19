package pkg

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// getBcryptCost 从配置获取 bcrypt 成本因子
func getBcryptCost() int {
	cost := viper.GetInt("security.bcrypt_cost")
	if cost < 4 {
		return 10 // 默认值，也是最小推荐值
	}
	if cost > 31 {
		return 31 // bcrypt 最大值
	}
	return cost
}

var (
	snowflakeNode *snowflake.Node
	snowflakeOnce sync.Once
)

// initSnowflakeNode 初始化雪花节点（单例）
func initSnowflakeNode() error {
	var err error
	snowflakeOnce.Do(func() {
		snowflakeNode, err = snowflake.NewNode(1)
	})
	return err
}

// GenerateID 生成雪花ID（使用单例节点，线程安全）
func GenerateID() (string, error) {
	if err := initSnowflakeNode(); err != nil {
		return "", fmt.Errorf("failed to init snowflake node: %w", err)
	}

	id := snowflakeNode.Generate()
	return id.String(), nil
}

// MustGenerateID 生成雪花ID，失败时 panic（仅用于初始化等关键路径）
func MustGenerateID() string {
	id, err := GenerateID()
	if err != nil {
		panic(err)
	}
	return id
}

// HashPassword 使用 bcrypt 对密码进行哈希
func HashPassword(password string) (string, error) {
	cost := getBcryptCost()
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// VerifyPassword 验证密码是否匹配
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// EncryptByMD5 已弃用：仅保留用于兼容性，新代码应使用 HashPassword
// Deprecated: 使用 HashPassword 替代
func EncryptByMD5(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	hashstr := hex.EncodeToString(hash[:])
	return hashstr
}

// 计算文件的 SHA-256 哈希值
func CalculateFileSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

// 获取文件后缀的辅助函数
func GetFileExt(filename string) string {
	// 查找最后一个.的位置
	lastDot := strings.LastIndex(filename, ".")
	if lastDot == -1 {
		return "" // 无扩展名
	}
	// 返回.之后的部分（转为小写，避免大小写问题）
	return strings.ToLower(filename[lastDot+1:])
}
