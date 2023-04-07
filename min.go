package main

import (
	"html/template"
	"net/http"
	"strings"
)

type HandlerFunc func(ctx *Context)
type HandlersChain []HandlerFunc
type Engine struct {
	*RouterGroup  //继承后，可以调用RouterGroup的全部方法
	router        *Router
	groups        []*RouterGroup     //存储所有的分组
	htmlTemplates *template.Template //将所有模板加入内存
	funcMap       template.FuncMap   //自定义模板渲染函数
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
	engine:=New()
	engine.Use(Logger(), Recovery())
	return engine
}

// SetFuncMap 设置函数映射
func (e *Engine) SetFuncMap(funcMap template.FuncMap) {
	e.funcMap = funcMap
}

// LoadHTMLGlob 加载pattern 下的 .html
func (e *Engine) LoadHTMLGlob(pattern string) {
	//创建空白模板，添加函数映射，并加载文件
	e.htmlTemplates = template.Must(template.New("").Funcs(e.funcMap).ParseGlob(pattern))
}
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

// ServeHTTP 所有的 http 请求通过此函数
func (e *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	c := NewContext(res, req)
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(c.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c.engine = e
	c.handlers = HandlersChain(middlewares)
	e.router.handle(c)
}
