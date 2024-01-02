package api

import (
	"github.com/gin-gonic/gin"
	"jurassic-park-api/data"
	"log"
	"net/http"
	"strconv"
)

type CageAssignmentsController struct {
	Store data.Store
}

func (cac CageAssignmentsController) List(context *gin.Context) {
	cageIdInput := context.Param("cage_id")
	cageId, err := strconv.ParseInt(cageIdInput, 10, 64)
	if err != nil {
		log.Printf("error parsing cage id %s: %s\n", cageIdInput, err)
		context.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Invalid cage id provided"}})
		return
	}

	dinosaurs, err := cac.Store.Dinosaurs.FindByCageId(cageId)
	if err != nil {
		log.Printf("error listing the dinosaurs assigned to cage %s: %s\n", cageId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"errors": []string{"An unexpected error occurred"}})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": dinosaurs})
}
