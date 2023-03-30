package main

import (
	"net/http"
	"strings"
)

type Router struct {
	//map 的方式为每个方法都创建路由树，根节点存储 pattern
	roots map[string]*node
	//由 map 实现的路由器,前缀树找到 pattern，再匹配到函数链
	handlers map[string]HandlersChain
}

func NewRouter() *Router {
	return &Router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlersChain),
	}
}

//路由器添加路由
func (r *Router) addroute(method, pattern string, handlers HandlersChain) {
	key := method + ":" + pattern
	if _, ok := r.roots[method]; !ok {
		//为每个方法都创建一个前缀树。
		r.roots[method] = &node{}
	}
	parts := parsePatternAndPath(pattern)
	//将 pattern 插入到路由树
	r.roots[method].instert(pattern, parts, 0)
	r.handlers[key] = append(r.handlers[key], handlers...)
}

//路由处理
func (r *Router) handle(c *Context) {
	key := c.Method + ":"
	searchParts := parsePatternAndPath(c.Path)
	//找到该方法的前缀树
	n, ok := r.roots[c.Method]
	if !ok {
		c.Data(http.StatusNotFound, []byte("404 not found"))
		return
	}
	//查询 pattern
	resn := n.search(searchParts, 0)
	if resn != nil {
		key += resn.pattern
	}
	//路由匹配成功，运行函数链
	if handlersChain, ok := r.handlers[key]; ok {
		for _, handlerFunc := range handlersChain {
			handlerFunc(c)
		}
	} else {
		c.Data(http.StatusNotFound, []byte("404 not found"))
	}
}

//pattern 和 path 以 / 分割成 parts 数组
func parsePatternAndPath(s string) []string {
	ans := strings.Split(s, "/")
	return ans
}
