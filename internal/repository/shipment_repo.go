package repository

import ("logistics-app/internal/model"
"errors"
)
type ShipmentRepository struct {
	data map[string]model.Shipment
}

func NewShipmentRepository() *ShipmentRepository {
	return &ShipmentRepository{
		data: make(map[string]model.Shipment),
	}
}

func (r *ShipmentRepository) Save(s model.Shipment) {
	r.data[s.ID] = s
}

func (r *ShipmentRepository) GetByID(id string) (model.Shipment, bool) {
	s, ok := r.data[id]
	return s, ok
}

func (r *ShipmentRepository) Update(id string, s model.Shipment) {
	r.data[id] = s
}

func (r *ShipmentRepository) Delete(id string) error {
	if _, exists := r.data[id]; !exists {
		return errors.New("not found")
	}

	delete(r.data, id)
	return nil
}