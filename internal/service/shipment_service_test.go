package service

import (
	"testing"
	"logistics-app/internal/model"
	"logistics-app/internal/repository"
)

func TestCreateAndGetShipment(t *testing.T) {
	repo := repository.NewShipmentRepository()
	service := NewShipmentService(repo)

	shipment := model.Shipment{
		ID:       "1",
		Status:   "in_transit",
		Location: "Delhi",
	}

	service.CreateShipment(shipment)

	result, ok := service.GetShipment("1")

	if !ok {
		t.Errorf("Expected shipment to exist")
	}

	if result.ID != "1" {
		t.Errorf("Expected ID 1, got %s", result.ID)
	}
}