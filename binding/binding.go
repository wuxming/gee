package binding

import "net/http"

const (
	MIMEJSON              = "application/json"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	// todo 其他类型
)

type Binding interface {
	Name() string
	Bind(*http.Request, any) error
}

var (
	JSON = jsonBinding{}
	// todo 其他类型
)

func Default(method, contentType string) Binding {
	switch contentType {
	case MIMEJSON:
		return JSON
	default:
		return JSON
	}
	// todo 其他类型
}
func validate(obj any) {

}
