package routes

import (
	"github.com/whereabouts/web-template/engine/server"
	"github.com/whereabouts/web-template/handlers"
	"net/http"
)

func Routes() {
	server.Route(http.MethodGet, "/sayHello", handlers.SayHello)
	// 子路由添加方式
	v1 := server.Group("v1")
	v1.Route(http.MethodGet, "/sayHello", handlers.SayHello)
}
