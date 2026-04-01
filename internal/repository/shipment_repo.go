package repository

import "logistics-app/internal/model"

type ShipmentRepository struct {
	data map[string]model.Shipment
}

func NewShipmentRepository() *ShipmentRepository {
	return &ShipmentRepository{
		data: make(map[string]model.Shipment),
	}
}

func (r *ShipmentRepository) Create(s model.Shipment) {
	r.data[s.ID] = s
}

func (r *ShipmentRepository) Get(id string) (model.Shipment, bool) {
	s, ok := r.data[id]
	return s, ok
}

func (r *ShipmentRepository) Update(id string, s model.Shipment) {
	r.data[id] = s
}