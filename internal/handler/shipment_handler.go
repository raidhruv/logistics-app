package handler

import (
	"github.com/gin-gonic/gin"
	"logistics-app/internal/service"
)

type ShipmentHandler struct {
	service *service.ShipmentService
}

func NewShipmentHandler(s *service.ShipmentService) *ShipmentHandler {
	return &ShipmentHandler{service: s}
}

func (h *ShipmentHandler) CreateShipment(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	shipment := h.service.CreateShipment(input.Name)

	c.JSON(200, shipment)
}

func (h *ShipmentHandler) GetShipment(c *gin.Context) {
	id := c.Param("id")

	shipment, ok := h.service.GetShipment(id)
	if !ok {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	c.JSON(200, shipment)
}

func (h *ShipmentHandler) UpdateShipment(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Status   string `json:"status"`
		Location string `json:"location"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	shipment, ok := h.service.UpdateShipment(id, input.Status, input.Location)
	if !ok {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	c.JSON(200, shipment)
}