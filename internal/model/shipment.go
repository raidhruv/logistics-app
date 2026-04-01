package model

type Shipment struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	Location string `json:"location"`
}