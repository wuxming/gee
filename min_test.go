package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMinGet(t *testing.T) {
	m := New()
	m.GET("/testGET", func(c *Context) {
		name := c.Query("name")
		c.JSON(http.StatusOK, H{
			"name": name,
			"msg":  "GET 请求测试成功",
		})
	})
	//测试服务启动
	ts := httptest.NewServer(m)
	defer ts.Close()
	//发送 Get 请求
	res, err := http.Get(ts.URL + "/testGET?name=张三")
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

	if string(body) != `{"msg":"GET 请求测试成功","name":"张三"}` {
		t.Errorf("body应该是`%s`，而不是`%s`", `{"msg":"GET 请求测试成功","name":"张三"}`, string(body))
	} else {
		t.Log(string(body))
	}
}
