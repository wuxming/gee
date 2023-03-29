package main

import (
	"fmt"
	"net/http"
)

func main() {
	m := New()
	m.GET("/testGET", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		_, err := res.Write([]byte("GET 请求测试成功"))
		if err != nil {
			fmt.Println(err)
		}
	})
	err := m.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
