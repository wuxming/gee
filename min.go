package main

import (
	"net/http"
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

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

//所有的 http 请求通过 ServeHTTP
func (e *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	c := NewContext(res, req)
	e.router.handle(c)
}
