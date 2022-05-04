package great

import (
	"net/http"
)

// HandlerFunc 将请求和响应进行封装
type HandlerFunc func(*Context)

// Engine 里的router属性是一个路径对应一对请求响应
type Engine struct {
	router *router
}

// New 新建一个Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute增加路径
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET POST
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// engine 需要实现ServeHTTP方法才能被当作实例作为ListenAndServe的第二个参数
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
	/*
		key := req.Method + "-" + req.URL.Path
			if handler, ok := engine.router[key]; ok {
				handler(w, req)
			} else {
				fmt.Fprintf(w, "404 NOT FOUND: %s", req.URL.Path)
			}
	*/
}
