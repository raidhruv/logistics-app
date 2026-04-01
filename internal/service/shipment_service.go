package service

import (
	"logistics-app/internal/model"
	"logistics-app/internal/repository"
)

type ShipmentService struct {
	repo *repository.ShipmentRepository
}

func NewShipmentService(r *repository.ShipmentRepository) *ShipmentService {
	return &ShipmentService{repo: r}
}

func (s *ShipmentService) CreateShipment(shipment model.Shipment) {
	s.repo.Create(shipment)
}

func (s *ShipmentService) GetShipment(id string) (model.Shipment, bool) {
	return s.repo.Get(id)
}

func (s *ShipmentService) UpdateShipment(id string, shipment model.Shipment) {
	s.repo.Update(id, shipment)
}