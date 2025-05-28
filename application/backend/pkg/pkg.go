package pkg

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"strings"

	"github.com/bwmarrin/snowflake"
)

// 本文件包含雪花ID、MD5函数
func GenerateID() string {
	node, _ := snowflake.NewNode(1)

	id := node.Generate()
	return id.String()
}

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
