package http_error

type JSON map[string]interface{}

func DefaultJSON(code interface{}, message string) interface{} {
	return struct {
		Code    interface{} `json:"err_code"`
		Message string      `json:"err_message"`
	}{code, message}
}
