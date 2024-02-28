package logutil

import (
	"fmt"
	"os"
)

func GetToken() string {
	token, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		fmt.Println("读取文件失败: /var/run/secrets/kubernetes.io/serviceaccount/token")
		panic(err)
	}
	return string(token)
}
