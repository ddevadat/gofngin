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

func main() {
	r := gin.Default()

	// Gin route that uses same logic as myHandler
	r.POST("/ping", func(c *gin.Context) {
		p := &Person{Name: "World"}

		// Decode request body into Person (if provided)
		if err := c.ShouldBindJSON(p); err != nil {
			// fallback to default "World" if no name provided
			log.Printf("Error decoding body: %v", err)
		}

		// Create response
		msg := struct {
			Msg string `json:"message"`
		}{
			Msg: fmt.Sprintf("Hello %s", p.Name),
		}

		log.Print("Inside Go Hello World function (Gin style)")
		c.JSON(200, msg)
	})

	if os.Getenv("FN_FORMAT") == "" {
		// -- Local Run ---
		r.Run()
	} else {
		// --- OCI Function mode ---
		fdk.Handle(fdk.HTTPHandler(r))
	}
}
