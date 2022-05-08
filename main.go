package main

import (
	"gproject/great"
	"net/http"
)

func main() {
	r := great.Default()
	r.GET("/", func(c *great.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *great.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}

/*
// HTML 功能测试
func main() {
	r := great.New()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	r.GET("/", func(c *great.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.Run(":9999")
}
*/

/*
// 模板测试
func main() {
	r := great.New()
	r.Static("/assets", "/usr/geektutu/blog/static")
	r.Run(":9999")
}
*/

/*
// 中间件测试
func onlyForV2() great.HandlerFunc {
	return func(c *great.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := great.New()
	r.Use(great.Logger()) // global midlleware
	r.GET("/", func(c *great.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *great.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
*/

/*
//分组测试
func main() {
	r := great.New()
	r.GET("/index", func(context *great.Context) {
		context.HTML(http.StatusOK, "<h1>Hello Great</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(context *great.Context) {
			context.HTML(http.StatusOK, "<h1>Hello Great</h1>")
		})
		v1.GET("/hello", func(c *great.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *great.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *great.Context) {
			c.JSON(http.StatusOK, great.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}
	r.Run(":9999")
}
*/

/*
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
*/

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
