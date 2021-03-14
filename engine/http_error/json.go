package http_error

type DefaultBody struct {
	Code    interface{} `json:"code"`
	Message string      `json:"message"`
}

type JSON map[string]interface{}

func DefaultJSON(code interface{}, message string) DefaultBody {
	return DefaultBody{Code: code, Message: message}
}
