package api

import (
	"github.com/gin-gonic/gin"
	"jurassic-park-api/data"
)

func SetupRouter(store data.Store) *gin.Engine {
	r := gin.Default()

	r.GET("/status", getStatus)

	cageController := CagesController{Store: store}
	r.GET("/cages", cageController.List)
	r.POST("/cages", cageController.Create)

	dinosaurController := DinosaursController{Store: store}
	r.GET("/dinosaurs", dinosaurController.List)
	r.POST("/dinosaurs", dinosaurController.Create)

	cageAssignmentController := CageAssignmentsController{Store: store}
	r.GET("/cages/:cage_id/assignments", cageAssignmentController.List)
	r.POST("/cages/:cage_id/assignments", cageAssignmentController.Create)

	return r
}
