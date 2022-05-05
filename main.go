package main

import (
	"gproject/great"
	"net/http"
)

// 测试动态路由
func main() {
	r := great.New()
	r.GET("/", func(c *great.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *great.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *great.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *great.Context) {
		c.JSON(http.StatusOK, great.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}

/*
测试JSON String HTML等方法
func main() {
	r := great.New()
	r.GET("/", func(context *great.Context) {
		context.HTML(http.StatusOK, "<h1>Hello Great</h1>")
	})
	r.GET("/hello", func(context *great.Context) {
		context.String(http.StatusOK, "hello %s, you're at %s\n", context.Query("name"), context.Path)
	})
	r.GET("/login", func(context *great.Context) {
		context.JSON(http.StatusOK, great.H{
			"username": context.PostForm("username"),
			"password": context.PostForm("password"),
		})
	})
	r.Run(":9999")
}
*/

/*
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
*/
