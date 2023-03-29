package main

import (
	"net/http"
)

type Router struct {
	handlers map[string]HandlersChain
	//由 map 实现的路由器
}

func NewRouter() *Router {
	return &Router{handlers: make(map[string]HandlersChain)}
}

//路由器添加路由
func (r *Router) addroute(method, pattern string, handlers HandlersChain) {
	key := method + ":" + pattern
	r.handlers[key] = append(r.handlers[key], handlers...)
}

//路由处理
func (r *Router) handle(c *Context) {
	key := c.Method + ":" + c.Path
	//路由匹配成功，运行函数链
	if handlersChain, ok := r.handlers[key]; ok {
		for _, handlerFunc := range handlersChain {
			handlerFunc(c)
		}
	} else {
		c.Data(http.StatusNotFound, []byte("404 not found"))
	}
}
