package instance

import (
	// "firestore-test/internal/infra/config/property" // Remove this if not directly used
	// injectProps "firestore-test/internal/cmd/inject" // For ServerConfig type - REMOVED
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// ControllerRunnable defines the interface for controllers that can be run by GinController.
type ControllerRunnable interface {
	RunController(r *gin.Engine) // Simplified signature
}

// GinController manages the Gin engine and runs registered controllers.
type GinController struct {
	controllers []ControllerRunnable
	port        string // Store port directly
}

// NewGinController creates a new GinController.
func NewGinController(controllers []ControllerRunnable, port string) *GinController { // Accept port string
	return &GinController{
		controllers: controllers,
		port:        port,
	}
}

// Start initializes and starts the Gin server.
func (c *GinController) Start() {
	r := gin.Default()
	for _, o := range c.controllers {
		o.RunController(r) // Pass only the Gin engine
	}

	err := r.Run(c.port) // Use stored port
	if err != nil {
		log.Fatalf("error trying to start server: %v", err)
		// return // Not strictly necessary due to log.Fatalf
	}
}
