package dto

type NodeDetailsDTO struct {
	Id              uint64  `json:"id"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	HasCrossing     bool    `json:"hasCrossing"`
	HasTrafficLight bool    `json:"hasTrafficLight"`
}