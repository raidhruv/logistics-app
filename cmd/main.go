package main

import (
	"github.com/gin-gonic/gin"
	"logistics-app/internal/handler"
	"logistics-app/internal/repository"
	"logistics-app/internal/service"
)

func main() {
	r := gin.Default()

	repo := repository.NewShipmentRepository()
	service := service.NewShipmentService(repo)
	handler := handler.NewShipmentHandler(service)

	r.POST("/shipment", handler.CreateShipment)
	r.GET("/shipment/:id", handler.GetShipment)
	r.PUT("/shipment/:id", handler.UpdateShipment)

	r.Run(":8080")
}