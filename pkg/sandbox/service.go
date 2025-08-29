package sandbox

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Person struct (can also be in a shared package if used elsewhere)
type Person struct {
	Name string `json:"name"`
}

// SubmitFile handles POST /ping
func SubmitFile(c *gin.Context) {
	p := &Person{Name: "World"}

	// Bind JSON request body
	if err := c.ShouldBindJSON(p); err != nil {
		log.Printf("Error decoding body: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello %s", p.Name),
	})

	log.Print("Inside Go Hello World function (sandbox.SubmitFile)")
}
