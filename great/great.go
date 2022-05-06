package great

import (
	"log"
	"net/http"
)

// HandlerFunc 将请求和响应进行封装
type HandlerFunc func(*Context)

// 新增分组控制
type (
	RouterGroup struct {
		prefix      string        // 前缀
		middlewares []HandlerFunc // 支持中间件
		parent      *RouterGroup  // 父亲
		engine      *Engine
	}
	Engine struct {
		*RouterGroup // 为了让engine具有RouterGroup的功能
		router       *router
		groups       []*RouterGroup
	}
)

// Engine 里的router属性是一个路径对应一对请求响应
/*
type Engine struct {
	router *router
}
*/

// New 新建一个Engine
/*
func New() *Engine {
	return &Engine{router: newRouter()}
}
*/
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group 方法进行分组
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	return newGroup
}

// addRoute增加路径
/*
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}
*/
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s\n", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// GET POST
/*
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}
*/

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
