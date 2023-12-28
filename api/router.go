package api

import (
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/status", getStatus)

	return r
}
