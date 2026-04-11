package service

import (
	"logistics-app/internal/model"
	"logistics-app/internal/repository"

	"github.com/google/uuid"
)

type ShipmentService struct {
	repo *repository.ShipmentRepository
}

func NewShipmentService(r *repository.ShipmentRepository) *ShipmentService {
	return &ShipmentService{repo: r}
}

func (s *ShipmentService) CreateShipment(name string) model.Shipment {
	shipment := model.Shipment{
		ID:       uuid.New().String(),
		Status:   "created",
		Location: name,
	}

	s.repo.Save(shipment)

	return shipment
}

func (s *ShipmentService) GetShipment(id string) (model.Shipment, bool) {
	return s.repo.GetByID(id)
}

func (s *ShipmentService) UpdateShipment(id string, status string, location string) (model.Shipment, bool) {
	shipment, ok := s.repo.GetByID(id)
	if !ok {
		return model.Shipment{}, false
	}

	shipment.Status = status
	shipment.Location = location

	s.repo.Update(id, shipment)

	return shipment, true
}