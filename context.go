package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

const (
	jsonContentType = "application/json"
	plainContentTyp = "text/plain"
	htmlContentType = "text/html"
)
const abortIndex int = math.MaxInt8 >> 1

// H 使得数据构建更加简洁
type H map[string]interface{}

// Context 上下文，每个 http 请求都会生成一个 Context
type Context struct {
	//ServeHTTP 所需要的两个参数
	ResponseWriter http.ResponseWriter
	Request        *http.Request

	Method string
	Path   string

	StatusCode int

	params   map[string]string
	index    int           //函数链的下标指针
	handlers HandlersChain //请求访问的函数链

	engine *Engine //可以访问 html
}

func NewContext(ResponseWriter http.ResponseWriter, Request *http.Request) *Context {
	return &Context{
		Request:        Request,
		ResponseWriter: ResponseWriter,
		Method:         Request.Method,
		Path:           Request.URL.Path,
		index:          -1, //下标从-1 开始
		params:         make(map[string]string),
	}
}

//------------------------flow control-------------------------------------

// Next 开始运行函数链中当前下标后面的函数
func (c *Context) Next() {
	c.index++
	for ; c.index < len(c.handlers); c.index++ {
		c.handlers[c.index](c)
	}
}

// Abort 中止函数链的运行
func (c *Context) Abort() {
	c.index = abortIndex
}

//------------------------input-------------------------------------

func (c *Context) Postform(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Params(key string) string {
	return c.params[key]
}

//------------------------output-------------------------------------

// Status 写入请求头状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.ResponseWriter.WriteHeader(code)
}
func (c *Context) Fail(error string) {
	c.Abort()
	c.JSON(http.StatusInternalServerError, H{"message": error})
}

// SetHeader 重写请求头信息
func (c *Context) SetHeader(key, value string) {
	c.ResponseWriter.Header().Set(key, value)
}

// AddHeader 为该 Key 的请求头添加 value
func (c *Context) AddHeader(key, value string) {
	c.ResponseWriter.Header().Add(key, value)
}

// SetContentTpye 设置返回类型
func (c *Context) SetContentTpye(value string) {
	c.SetHeader("Content-Type", value)
}
func (c *Context) JSON(code int, obj any) {
	c.Status(code)
	c.SetContentTpye(jsonContentType)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		c.Fail(err.Error())
	}
	_, err = c.ResponseWriter.Write(jsonBytes)
	if err != nil {
		c.Fail(err.Error())
	}
}
func (c *Context) HTML(code int, name string, data interface{}) {
	c.Status(code)
	c.SetContentTpye(htmlContentType)
	//支持模板文件名选择模板渲染
	err := c.engine.htmlTemplates.ExecuteTemplate(c.ResponseWriter, name, data)
	if err != nil {
		c.Fail(err.Error())
	}
}
func (c *Context) String(code int, format string, values ...interface{}) {
	c.Status(code)
	c.SetContentTpye(plainContentTyp)
	_, err := c.ResponseWriter.Write([]byte(fmt.Sprintf(format, values...)))
	if err != nil {
		if err != nil {
			c.Fail(err.Error())
		}
	}
}
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	_, err := c.ResponseWriter.Write(data)
	if err != nil {
		c.Fail(err.Error())
	}
}
