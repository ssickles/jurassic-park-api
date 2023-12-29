package api

import (
	"github.com/gin-gonic/gin"
	"jurassic-park-api/data"
)

func SetupRouter(store data.Store) *gin.Engine {
	r := gin.Default()

	r.GET("/status", getStatus)

	dinosaurController := DinosaursController{Store: store}
	r.GET("/dinosaurs", dinosaurController.List)
	r.POST("/dinosaurs", dinosaurController.Create)

	return r
}
