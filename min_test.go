package min

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type header struct {
	Key   string
	Value string
}

func PreformRequset(r http.Handler, method, path string, body io.Reader, headers ...header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestMinGet(t *testing.T) {
	m := Default()
	var tmp string
	g := m.Group("v1")
	//中间件
	g.Use(func(ctx *Context) {
		tmp += "A"
		ctx.Next()
		tmp += "C"
	})
	//动态路由
	g.GET("/testGET/:var1", func(c *Context) {
		name := c.Query("name")
		var1 := c.Params("var1")
		tmp += var1
		c.JSON(http.StatusOK, H{
			"name": name,
			"tmp":  tmp,
			"msg":  "GET 请求测试成功",
		})
	})
	//测试服务启动 发送 Get 请求
	w := PreformRequset(m, "GET", "/v1/testGET/B?name=张三", nil)
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		return
	}
	if w.Code != http.StatusOK {
		t.Errorf("状态码应该是%d，而不是%d", http.StatusOK, w.Code)
	}
	resbody := make(map[string]interface{})
	_ = json.Unmarshal(body, &resbody)
	//比较
	if !reflect.DeepEqual(resbody, map[string]interface{}{
		"name": "张三",
		"tmp":  "AB",
		"msg":  "GET 请求测试成功",
	}) {
		t.Error("结果不符合")
	}
	t.Log(resbody)
	if tmp != "ABC" {
		t.Error("结果不符合")
	}
	t.Log(tmp)
}

type people struct {
	Name string `json:"name"`
	Nge  int    `json:"age"`
}

func TestBinding(t *testing.T) {
	m := Default()
	m.POST("/", func(c *Context) {
		p := people{}
		c.Bind(&p)
		assert.Equal(t, "sam", p.Name)
		assert.Equal(t, 20, p.Nge)
	})
	body := bytes.NewBufferString(`{"name":"sam","age":20}`)
	h := header{
		Key:   "Content-Type",
		Value: "application/json",
	}
	PreformRequset(m, "POST", "/", body, h)

}
