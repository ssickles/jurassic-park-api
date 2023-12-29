package api

import (
	"github.com/gin-gonic/gin"
	"jurassic-park-api/data"
	"jurassic-park-api/models"
	"log"
	"net/http"
)

type CagesController struct {
	Store data.Store
}

func (cc CagesController) List(context *gin.Context) {
	cages, err := cc.Store.Cages.List()
	if err != nil {
		log.Printf("error listing cages: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"errors": []string{"An unexpected error occurred"}})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": cages})
}

func (cc CagesController) Create(context *gin.Context) {
	var cage models.Cage
	err := context.BindJSON(&cage)
	if err != nil {
		log.Printf("error binding cage json: %s\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Invalid json provided for a cage"}})
		return
	}

	createdCage, err := cc.Store.Cages.Create(cage)
	if err != nil {
		log.Printf("error creating cage: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"errors": []string{"An unexpected error occurred"}})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": createdCage})
}
