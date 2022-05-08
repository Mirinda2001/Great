package great

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// 请求和响应
	Req    *http.Request
	writer http.ResponseWriter
	//请求方法和路径
	Method string
	Path   string
	// 响应信息
	StatusCode int
	// 模糊匹配对应的值
	Params map[string]string
	// 支持中间件
	handlers []HandlerFunc
	// 记录执行到第几个中间件
	index int
	// 新增engine   与HTML相关
	engine *Engine
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

// Param 获取模糊匹配对应的值
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Req:    req,
		writer: w,
		Method: req.Method,
		Path:   req.URL.Path,
		index:  -1,
	}
}

// Next 中间件相关的Next方法
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

// PostForm 使用FormValue方法
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 使用req.URL.Query().Get(key)
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 设置响应码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.writer.WriteHeader(code)
}

// SetHeader 往header里面写值
func (c *Context) SetHeader(key string, value string) {
	c.writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	//c.StatusCode = code
	c.Status(code)
	c.writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	//c.StatusCode = code
	encoder := json.NewEncoder(c.writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	//c.StatusCode = code
	c.writer.Write(data)
}

/*
func (c *Context) HTML(code int, HTML string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	//c.StatusCode = code
	c.writer.Write([]byte(HTML))
}

*/

// HTML 对原来的 (*Context).HTML()方法做了些小修改，使之支持根据模板文件名选择模板进行渲染。
func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}
