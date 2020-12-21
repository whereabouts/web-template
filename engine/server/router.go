package server

import "github.com/whereabouts/web-template/engine/hanlder"

type Router func()

func (s *Server) Router(r Router) *Server {
	setRouter(s, r)
	return s
}

func Route(method string, path string, function interface{}) {
	gServer.GetEngine().Handle(method, path, hanlder.CreateHandlerFunc(function))
}

func setRouter(s *Server, r Router) {
	r()
}
