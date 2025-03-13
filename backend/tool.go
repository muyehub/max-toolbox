package backend

import (
	"encoding/base64"
	"fmt"
)

// ConvertBase64ToStr returns a base64 to str result
func (a *App) ConvertBase64ToStr(in string) string {
	// 解码 Base64
	decodedData, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		fmt.Println("Error decoding Base64:", err)
		return err.Error()
	}

	// 将解码后的字节切片转换为字符串
	return string(decodedData)
}

// ConvertStrToBase64 returns a str to base64 result
func (a *App) ConvertStrToBase64(in string) string {
	// 编码 Base64
	return base64.StdEncoding.EncodeToString([]byte(in))
}
