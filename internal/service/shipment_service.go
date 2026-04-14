package service

import (
	"logistics-app/internal/model"
	"logistics-app/internal/repository"

	"github.com/google/uuid"
	"time"
)
const DefaultLocation = "Warehouse"
type ShipmentService struct {
	repo *repository.ShipmentRepository
}

func NewShipmentService(r *repository.ShipmentRepository) *ShipmentService {
	return &ShipmentService{repo: r}
}
// ShipmentService provides business logic for managing shipments. 
// It interacts with the ShipmentRepository to perform CRUD operations on shipments.
// its main responsibilities include creating new shipments, retrieving existing shipments by ID,
// and updating shipment details such as status and location. 
// The service ensures that the business rules are applied correctly when manipulating shipment data.
func (s *ShipmentService) CreateShipment(name string) model.Shipment {
	shipment := model.Shipment{
		ID:       uuid.New().String(),
		Status:   "created",
		Location: DefaultLocation,
		TransitPoints: []model.TransitPoint{
			{
				Location: DefaultLocation,
				Time:     time.Now().Format(time.RFC3339),
			},
		},
	}
// its also the place where i kept the default location for the shipment, which is "Warehouse".
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

	// update status safely
	if status != "" {
		shipment.Status = status
	}

	// update location + append transit history
	if location != "" && location != shipment.Location {
		shipment.Location = location

		shipment.TransitPoints = append(shipment.TransitPoints, model.TransitPoint{
			Location: location,
			Time:     time.Now().Format(time.RFC3339),
		})
	}

	s.repo.Update(id, shipment)

	return shipment, true
}
func (s *ShipmentService) DeleteShipment(id string) error {
    return s.repo.Delete(id)
}