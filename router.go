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
func (r *Router) addRoute(method, pattern string, handlers HandlersChain) {
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
func (r *Router) getRoute(method, path string) (*node, map[string]string) {
	//找到该方法的前缀树根节点
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	searchParts := parsePatternAndPath(path)
	//查询 pattern
	n := root.search(searchParts, 0)
	params := make(map[string]string)
	if n != nil {
		//拆分 pattern
		parts := parsePatternAndPath(n.pattern)
		for i, part := range parts {
			if part[0] == ':' {
				// /test/:var1    /test/123
				params[part[1:]] = searchParts[i]
				// {"var" : "123"}
			}
			if part[0] == '*' && len(part) > 1 {
				// /test/*var1    /test/123/abc
				params[part[1:]] = strings.Join(searchParts[i:], "/")
				break
				// {"var" : "123/abc"}
			}
		}
		return n, params
	}
	return nil, nil
}

//路由处理
func (r *Router) handle(c *Context) {

	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.params = params
		key := c.Method + ":" + n.pattern
		//路由匹配成功，添加相应的函数链
		c.handlers = append(c.handlers, r.handlers[key]...)
	} else {
		//未匹配路由，则添加 404 handler
		c.handlers = append(c.handlers, func(ctx *Context) {
			ctx.Data(http.StatusNotFound, []byte("404 not found"))
		})
	}
	c.Next()
}

//pattern 和 path 以 / 分割成 parts 数组
func parsePatternAndPath(s string) []string {
	var ans []string
	parts := strings.Split(s, "/")
	for _, part := range parts {
		if part != "" {
			ans = append(ans, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return ans
}
