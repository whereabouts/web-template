package hanlder

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

func Pre(engine interface{}, middlewares ...gin.HandlerFunc) {
	switch engine.(type) {
	case *gin.Engine:
		engine.(*gin.Engine).Use(middlewares...)
	case *gin.RouterGroup:
		engine.(*gin.RouterGroup).Use(middlewares...)
	}
}

func After(engine interface{}, middlewares ...gin.HandlerFunc) {
	finals := make([]gin.HandlerFunc, 0)
	for _, middleware := range middlewares {
		middleware := func(c *gin.Context) {
			c.Next()
			mwV := reflect.ValueOf(middleware)
			if mwV.IsValid() {
				mwV.Call([]reflect.Value{reflect.ValueOf(c)})
			}
		}
		finals = append(finals, middleware)
	}
	switch engine.(type) {
	case *gin.Engine:
		engine.(*gin.Engine).Use(finals...)
	case *gin.RouterGroup:
		engine.(*gin.RouterGroup).Use(finals...)
	}
}
