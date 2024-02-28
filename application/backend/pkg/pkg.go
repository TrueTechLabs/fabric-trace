package pkg

import (
	"crypto/md5"
	"encoding/hex"

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
