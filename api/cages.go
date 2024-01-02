package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"jurassic-park-api/data"
	"jurassic-park-api/park"
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
	var payload park.CreateCagePayload
	err := context.BindJSON(&payload)
	if err != nil {
		log.Printf("error binding cage json: %s\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Invalid json provided for a cage"}})
		return
	}

	createdCage, err := park.CreateCage(cc.Store, payload)
	if err != nil {
		// We will return a different error message depending on the type of error
		switch {
		case errors.As(err, &park.CageNameAlreadyExistsError{}):
			context.JSON(http.StatusConflict, gin.H{"errors": []string{err.Error()}})
			return
		default:
			log.Printf("%s\n", err)
			context.JSON(http.StatusInternalServerError, gin.H{"errors": []string{"An unexpected error occurred"}})
			return
		}
	}

	context.JSON(http.StatusCreated, gin.H{"data": createdCage})
}
