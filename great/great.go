package great

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

// HandlerFunc 将请求和响应进行封装
type HandlerFunc func(*Context)

// 新增分组控制
type (
	RouterGroup struct {
		prefix      string        // 前缀
		middlewares []HandlerFunc // 用来存中间件
		parent      *RouterGroup  // 父亲
		engine      *Engine
	}
	Engine struct {
		*RouterGroup // 为了让engine具有RouterGroup的功能
		router       *router
		groups       []*RouterGroup
		//前者将所有的模板加载进内存，后者是所有的自定义模板渲染函数。
		htmlTemplates *template.Template
		funcMap       template.FuncMap
	}
)

// SetFuncMap 新增的和HTML有关的方法
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}
func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}

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

// Default 默认使用Logger和Recovery中间件
func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

/*
错误写法
func (engine *Engine) Default() {
	engine.Use(Logger(), Recovery())
	engine.Run(":9999")
}
*/

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	c.engine = engine
	engine.router.handle(c)
}

/*
// engine 需要实现ServeHTTP方法才能被当作实例作为ListenAndServe的第二个参数
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
*/

// Use 使用中间件
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

//  创建静态资源处理器
// 支持静态服务资源  FileSystem接口实现了对一系列命名文件的访问
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	// 先找到绝对路径
	absolutePath := group.prefix + relativePath
	// 根据路径和FileSystem生成文件服务器
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(context *Context) {
		//  获取文件路径
		file := context.Param("filepath")
		//  尝试打开文件  看文件是否存在
		if _, err := fs.Open(file); err != nil {
			context.Status(http.StatusNotFound)
			return
		}
		// 文件服务器接手处理请求
		fileServer.ServeHTTP(context.writer, context.Req)
	}
}

func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	group.GET(urlPattern, handler)
}
