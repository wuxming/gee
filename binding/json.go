package binding

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type jsonBinding struct {
}

func (jsonBinding) Name() string {
	return "json"
}
func (jsonBinding) Bind(req *http.Request, obj any) error {
	if req == nil || req.Body == nil {
		return errors.New("invalid request")
	}
	return decodeJSON(req.Body, obj)
}
func decodeJSON(r io.Reader, obj any) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(obj)
}
