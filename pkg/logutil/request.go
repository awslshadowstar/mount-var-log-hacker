package logutil

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

var LOGS_URL string
var HOST_FS_BASE string

func init() {
	LOGS_URL = fmt.Sprintf("https://%s:10250/logs", DefaultGateway())
	HOST_FS_BASE = LOGS_URL + "/root_link"
}

func DoRequest(url string) (io.ReadCloser, error) {
	// 创建一个 HTTP 客户端
	client := &http.Client{}

	// 创建一个 GET 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 设置请求头部，包括 Authorization
	req.Header.Set("Authorization", "Bearer "+GetToken())

	// 忽略证书验证
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
