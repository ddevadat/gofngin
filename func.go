package main

import (
	"func/pkg/sandbox"
	"net/http"
	"os"

	"github.com/fnproject/fdk-go"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func mountUserRoutes(router *gin.Engine) {
	api := router.Group("/sandbox/v1")
	{
		api.POST("/submit", sandbox.SubmitFile)
	}

	// Handle unmatched routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	})
}

func main() {
	r := InitializeRouter()
	mountUserRoutes(r)
	if os.Getenv("FN_FORMAT") == "" {
		// -- Local Run ---
		r.Run()
	} else {
		// --- OCI Function mode ---
		fdk.Handle(fdk.HTTPHandler(r))
	}
}
