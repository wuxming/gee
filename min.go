package main

import (
	"net/http"
	"strings"
)

type HandlerFunc func(ctx *Context)
type HandlersChain []HandlerFunc
type Engine struct {
	*RouterGroup //继承后，可以调用RouterGroup的全部方法
	router       *Router
	groups       []*RouterGroup //存储所有的分组
}

func New() *Engine {
	engine := &Engine{router: NewRouter()}
	//作为根节点的 group,可以认为是最大的根分组
	rootGroup := &RouterGroup{engine: engine}
	engine.RouterGroup = rootGroup

	engine.groups = []*RouterGroup{rootGroup}

	return engine
}
func Default() *Engine {
	engine := New()
	engine.Use(Logger())
	return engine
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

//所有的 http 请求通过 ServeHTTP
func (e *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	c := NewContext(res, req)
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(c.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c.handlers = HandlersChain(middlewares)
	e.router.handle(c)
}
