package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fnproject/fdk-go"
	"github.com/gin-gonic/gin"
)

// Person struct
type Person struct {
	Name string `json:"name"`
}

func InitializeRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/ping", func(c *gin.Context) {
		p := &Person{Name: "World"}

		if err := c.ShouldBindJSON(p); err != nil {
			log.Printf("Error decoding body: %v", err)
		}

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello %s", p.Name),
		})
		log.Print("Inside Go Hello World function (Gin style)")
	})

	return r
}

func main() {
	r := InitializeRouter()

	if os.Getenv("FN_FORMAT") == "" {
		// -- Local Run ---
		r.Run()
	} else {
		// --- OCI Function mode ---
		fdk.Handle(fdk.HTTPHandler(r))
	}
}
