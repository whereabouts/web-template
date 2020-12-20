package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HelloMiddlewarePre(context *gin.Context) {
	fmt.Println("hello pre")
}

func HelloMiddlewareAfter(context *gin.Context) {
	fmt.Println("hello after")
}
