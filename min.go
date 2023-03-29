package main

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	//由 map 实现的路由器
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}
func (e *Engine) addroute(method, pattern string, handler HandlerFunc) {
	key := method + ":" + pattern
	e.router[key] = handler
}
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addroute("GET", pattern, handler)
}
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addroute("POST", pattern, handler)
}
func (e *Engine) PUT(pattern string, handler HandlerFunc) {
	e.addroute("PUT", pattern, handler)
}
func (e *Engine) DELETE(pattern string, handler HandlerFunc) {
	e.addroute("DELETE", pattern, handler)
}
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

//所有的 http 请求通过 ServeHTTP
func (e *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	key := req.Method + ":" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		c := NewContext(res, req)
		handler(c)
	} else {
		res.WriteHeader(http.StatusNotFound)
		_, err := res.Write([]byte("404 not found"))
		if err != nil {
			fmt.Println(err)
		}
	}
}
