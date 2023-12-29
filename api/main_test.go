package api

import (
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

// TestMain is a special function that runs before any tests are run in this package
// We are using it to ensure Gin is in test mode so we don't get all the debug output
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	code := m.Run()
	os.Exit(code)
}
