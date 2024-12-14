package pkg

import (
    "crypto/md5"
    "encoding/hex"

    "github.com/bwmarrin/snowflake"
)

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
