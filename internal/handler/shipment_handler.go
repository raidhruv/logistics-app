package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"logistics-app/internal/model"
	"logistics-app/internal/service"
)

type ShipmentHandler struct {
	service *service.ShipmentService
}

func NewShipmentHandler(s *service.ShipmentService) *ShipmentHandler {
	return &ShipmentHandler{service: s}
}

func (h *ShipmentHandler) CreateShipment(c *gin.Context) {
	var shipment model.Shipment
	if err := c.ShouldBindJSON(&shipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.service.CreateShipment(shipment)
	c.JSON(http.StatusOK, shipment)
}

func (h *ShipmentHandler) GetShipment(c *gin.Context) {
	id := c.Param("id")
	shipment, ok := h.service.GetShipment(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, shipment)
}

func (h *ShipmentHandler) UpdateShipment(c *gin.Context) {
	id := c.Param("id")
	var shipment model.Shipment

	if err := c.ShouldBindJSON(&shipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.service.UpdateShipment(id, shipment)
	c.JSON(http.StatusOK, shipment)
}