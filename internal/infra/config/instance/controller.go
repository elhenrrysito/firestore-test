package instance

import (
	"firestore-test/internal/infra/config/property"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.falabella.tech/rtl/logistics-corp/rso-portfolio/libraries/golang/lightms"
	"log"
)

type ControllerRunnable interface {
	RunController(r *gin.Engine)
}

type GinController struct {
	controllers []ControllerRunnable
}

func GetControllerInstance() lightms.PrimaryProcess {
	controllers := []ControllerRunnable{GetSaleControllerInstance()}
	return &GinController{controllers: controllers}
}

func (c *GinController) Start() {
	r := gin.Default()
	for _, o := range c.controllers {
		o.RunController(r)
	}

	port := fmt.Sprintf(":%s", property.GetServerProperty().Server.Port)
	err := r.Run(port)
	if err != nil {
		log.Fatalf("error trying to start server: %v", err)
		return
	}
}
