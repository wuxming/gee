package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMinGet(t *testing.T) {
	m := New()
	m.GET("/testGET", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		_, err := res.Write([]byte("GET 请求测试成功"))
		if err != nil {
			t.Error(err)
		}
	})
	//测试服务启动
	ts := httptest.NewServer(m)
	defer ts.Close()
	//发送 Get 请求
	res, err := http.Get(ts.URL + "/testGET")
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
	if string(body) != "GET 请求测试成功" {
		t.Errorf("body应该是`%s`，而不是`%s`", "GET 请求测试成功", string(body))
	}

}
