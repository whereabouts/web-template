package proto

import (
	"github.com/whereabouts/web-template/engine/hanlder"
	"mime/multipart"
)

type SayHelloReq struct {
	Name string `json:"name,default=tb" form:"name,default=tb"`
	Age  int    `json:"age" form:"age"`
}

type SayHelloResp struct {
	Code    int    `json:"code" form:"code"`
	Message string `json:"message" form:"message"`
}

type FileHelloReq struct {
	File *multipart.FileHeader `json:"file" form:"file"`
}

type FileHelloResp struct {
	Code    int    `json:"code" form:"code"`
	Message string `json:"message" form:"message"`
}

type FilesHelloReq struct {
	hanlder.Context
	Name string `json:"name" form:"name"`
}

type FilesHelloResp struct {
	Code    int    `json:"code" form:"code"`
	Message string `json:"message" form:"message"`
}
