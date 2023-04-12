package binding

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func requestWithBody(method, path, body string) *http.Request {
	req, _ := http.NewRequest(method, path, bytes.NewBufferString(body))
	return req
}

type BinStruct struct {
	Foo string `json:"foo"`
}

func TestBinding(t *testing.T) {
	bin := BinStruct{}
	b := JSON
	req := requestWithBody("POST", "/", `{"foo": "bar"}`)
	err := b.Bind(req, &bin)
	assert.NoError(t, err)
	assert.Equal(t, "bar", bin.Foo)
}
