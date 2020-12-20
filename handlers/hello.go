package handlers

import (
	"fmt"
	"github.com/whereabouts/web-template/engine/http_error"
	"github.com/whereabouts/web-template/proto"
	"net/http"
	//"mime/multipart"
)

// 普通请求：请求参数自动映射绑定到req；如果返回值为nil则将resp结构体以json格式映射到响应体，否则以json响应 *http_error.HttpError 内容
func SayHello(req *proto.SayHelloReq, resp *proto.SayHelloResp) *http_error.HttpError {
	fmt.Println("say hello")
	resp.Code = http.StatusOK
	resp.Message = fmt.Sprintf("hello, %s! your age is %d", req.Name, req.Age)
	return nil
}

// 上传单个文件：可直接封装到req结构体中实现解析，类型为 *multipart.FileHeader
func FileHello(req *proto.FileHelloReq, resp *proto.FileHelloResp) *http_error.HttpError {
	if file := req.File; file == nil {
		return http_error.Error(http.StatusBadRequest, "fail to find any file")
	}
	resp.Code = http.StatusOK
	resp.Message = fmt.Sprintf("success to upload file : %s", req.File.Filename)
	return nil
}

// 多文件上传：无法通过映射解析，需要让req继承 middleware.Context 结构体，使其具有获取 gin.context 的能力，通过gin原生操作获取多文件
func FilesHello(req *proto.FilesHelloReq, resp *proto.FilesHelloResp) *http_error.HttpError {
	form, err := req.GetContext().MultipartForm()
	if err != nil {
		return http_error.ErrToHttpError(err, http.StatusBadRequest)
	}
	files := form.File["files"]
	var message string
	for _, file := range files {
		message = fmt.Sprintf("%s[%s]", message, file.Filename)
	}
	resp.Code = http.StatusOK
	resp.Message = fmt.Sprintf("success to upload these files : %+v", message)
	return nil
}
