package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getStatus(c *gin.Context) {
	// TODO: add checks to see if all our supporting services are up and running such as our Postgres connection
	c.JSON(http.StatusOK, gin.H{"data": "all systems operational"})
}
