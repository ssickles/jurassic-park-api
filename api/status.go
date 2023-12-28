package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getStatus(c *gin.Context) {
	c.String(http.StatusOK, "all systems operational")
}
