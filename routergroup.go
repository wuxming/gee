package main

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
