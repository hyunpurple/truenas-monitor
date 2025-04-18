package utils

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// HTTPClient 封装的 HTTP 客户端
var HTTPClient = resty.New()

// Get 发起 GET 请求
func Get(url string) {
	resp, err := HTTPClient.R().Get(url)
	if err != nil {
		fmt.Printf("GET 请求失败: %v\n", err)
		return
	}
	fmt.Printf("GET 请求成功: %s\n", resp)
}
