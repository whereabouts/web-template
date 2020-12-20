package server

import "github.com/whereabouts/web-template/engine/hanlder"

type Router func()

func SetRouter(s *Server, r Router) {
	r()
}

func Route(method string, path string, function interface{}) {
	gServer.GetEngine().Handle(method, path, hanlder.CreateHandlerFunc(function))
}
