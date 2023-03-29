package main

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)
type HandlersChain []HandlerFunc
type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: NewRouter()}
}

func (e *Engine) GET(pattern string, handlers ...HandlerFunc) {
	e.router.addroute("GET", pattern, handlers)
}
func (e *Engine) POST(pattern string, handlers ...HandlerFunc) {
	e.router.addroute("POST", pattern, handlers)
}
func (e *Engine) PUT(pattern string, handlers ...HandlerFunc) {
	e.router.addroute("PUT", pattern, handlers)
}
func (e *Engine) DELETE(pattern string, handlers ...HandlerFunc) {
	e.router.addroute("DELETE", pattern, handlers)
}
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

//所有的 http 请求通过 ServeHTTP
func (e *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	c := NewContext(res, req)
	e.router.handle(c)
}
