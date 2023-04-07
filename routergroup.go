package min

import (
	"net/http"
	"path"
)

type RouterGroup struct {
	prefix      string       //该分组的前缀
	parent      *RouterGroup //父分组
	engine      *Engine      //可以通过 驱动器 engine 访问其他接口
	middlewares []HandlerFunc
}

func (g *RouterGroup) Use(middleware ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middleware...)
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine
	newGroup := &RouterGroup{
		prefix: g.prefix + "/" + prefix,
		parent: g,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}
func (g *RouterGroup) addRoute(method, pattern string, handlers HandlersChain) {
	pattern = g.prefix + pattern
	g.engine.router.addRoute(method, pattern, handlers)
}
func (g *RouterGroup) GET(pattern string, handlers ...HandlerFunc) {
	g.addRoute("GET", pattern, handlers)
}
func (g *RouterGroup) POST(pattern string, handlers ...HandlerFunc) {
	g.addRoute("POST", pattern, handlers)
}
func (g *RouterGroup) PUT(pattern string, handlers ...HandlerFunc) {
	g.addRoute("PUT", pattern, handlers)
}
func (g *RouterGroup) DELETE(pattern string, handlers ...HandlerFunc) {
	g.addRoute("DELETE", pattern, handlers)
}

// createStaticHandler 创建静态 handler
func (g *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	//拼成绝对路径
	absoluePath := path.Join(g.prefix, relativePath)
	//将 absoluePath 前缀去掉后，然后可以在 fs 目录下查找静态文件
	fileServer := http.StripPrefix(absoluePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Params("filepath")
		//查看文件是否存在
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		//静态文件服务器已经由原生 fileServer 实现
		fileServer.ServeHTTP(c.ResponseWriter, c.Request)
	}
}
func (g *RouterGroup) Static(relativePath, root string) {
	handler := g.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	g.GET(urlPattern, handler)
}
