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
	// 相应信息
	StatusCode int
	// 模糊匹配对应的值
	Params map[string]string
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

func (c *Context) HTML(code int, HTML string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	//c.StatusCode = code
	c.writer.Write([]byte(HTML))
}
