package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestMinGet(t *testing.T) {
	m := New()
	g := m.Group("v1")
	g.Use(func(ctx *Context) {
		fmt.Println("123")
		ctx.Next()
		fmt.Println("321")
	})
	g.GET("/testGET/:var1", func(c *Context) {
		t.Log("/testGET/:var1")
		name := c.Query("name")
		var1 := c.Params("var1")
		c.JSON(http.StatusOK, H{
			"name": name,
			"var1": var1,
			"msg":  "GET 请求测试成功",
		})
	})
	//测试服务启动
	ts := httptest.NewServer(m)
	defer ts.Close()
	//发送 Get 请求
	res, err := http.Get(ts.URL + "/v1/testGET/123?name=张三")
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("状态码应该是%d，而不是%d", http.StatusOK, res.StatusCode)
	}
	resbody := make(map[string]interface{})
	_ = json.Unmarshal(body, &resbody)
	//比较
	if reflect.DeepEqual(resbody, map[string]interface{}{
		"name": "张三",
		"var1": "123",
		"msg":  "GET 请求测试成功",
	}) {
		t.Log(resbody)
	} else {
		t.Error("结果不符合")
	}
}
