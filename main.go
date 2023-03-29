package main

import (
	"fmt"
	"net/http"
)

type Engine struct {
}

func (e *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	_, err := res.Write([]byte("请求的方法是：" + req.Method + "；请求的路径是：" + req.URL.Path))
	if err != nil {
		fmt.Println(err)
		return
	}
}
func main() {
	engine := &Engine{}
	err := http.ListenAndServe(":8080", engine)
	if err != nil {
		fmt.Println(err)
		return
	}
}
