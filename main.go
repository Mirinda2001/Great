package main

import (
	"fmt"
	"gproject/great"
	"net/http"
)

func main() {
	// 雏形代码测试
	r := great.New()
	r.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path : %s", request.URL.Path)
	})
	r.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		// 获取响应头的信息显示出来
		for k, v := range request.Header {
			//双引号围绕的字符串，由Go语法安全地转义
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
			//fmt.Fprintf(writer, "Header[%s] = %s", k, v)
		}
	})
	r.Run(":9999")
}
