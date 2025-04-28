package sale

import (
	"errors"
	"firestore-test/internal/core"
	"firestore-test/internal/core/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	controllerInstance *Controller
	once               sync.Once
)

type Controller struct {
	persistencePort core.SalePersistencePort
	useCase         core.SaleUseCaseHandler
}

func NewController(useCase core.SaleUseCaseHandler, persistencePort core.SalePersistencePort) *Controller {
	once.Do(func() {
		controllerInstance = &Controller{useCase: useCase, persistencePort: persistencePort}
	})
	return controllerInstance
}

func (c *Controller) RunController(r *gin.Engine) {
	r.GET("/sale/:orderNumber", func(gc *gin.Context) {
		c.getSaleByOrderNumber(gc)
	})

	r.POST("/sale", func(gc *gin.Context) {
		c.createSale(gc)
	})

	r.PUT("/sale", func(gc *gin.Context) {
		c.updateSale(gc)
	})

}

func (c *Controller) getSaleByOrderNumber(gc *gin.Context) {
	orderNumber := gc.Param("orderNumber")
	sale, err := c.persistencePort.FindByOrderNumber(orderNumber)
	if err != nil {
		gc.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if sale == nil {
		gc.JSON(404, gin.H{"error": "Sale not found"})
		return
	}

	gc.JSON(200, toResponseDTO(*sale))
}

func (c *Controller) createSale(gc *gin.Context) {
	var requestDTO RequestDTO
	err := gc.BindJSON(&requestDTO)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = c.useCase.Handle(toDomain(requestDTO))
	if err != nil {
		gc.JSON(500, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(201, gin.H{"message": fmt.Sprintf("Sale created with order number %s", requestDTO.OrderNumber)})
}

func (c *Controller) updateSale(gc *gin.Context) {
	var requestDTO UpdateRequestDTO
	err := gc.BindJSON(&requestDTO)
	if err != nil {
		gc.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = c.useCase.UpdateStatus(requestDTO.OrderNumber, requestDTO.Status)
	if err != nil {
		if errors.Is(err, domain.ErrResourceNotFound) {
			gc.JSON(404, gin.H{"error": "Sale not found"})
			return
		}

		gc.JSON(500, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(201, gin.H{"message": fmt.Sprintf("Sale updated for order number %s", requestDTO.OrderNumber)})
}
