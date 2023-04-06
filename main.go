package main

import (
	"fmt"
	"net/http"
)

func main() {
	m := Default()
	m.GET("/testGET", func(c *Context) {
		name := c.Query("name")
		c.JSON(http.StatusOK, H{
			"name": name[100],
			"msg":  "GET 请求测试成功",
		})
	})
	m.Static("/assets", "./a")
	err := m.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
