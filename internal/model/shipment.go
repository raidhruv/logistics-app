package model

type TransitPoint struct {
	Location string `json:"location"`
	Time     string `json:"time"`
}

type Shipment struct {
	ID            string         `json:"id"`
	Status        string         `json:"status"`
	Location      string         `json:"location"`
	TransitPoints []TransitPoint `json:"transit_points"`
}