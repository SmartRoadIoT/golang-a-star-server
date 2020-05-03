package dto

type DistanceDTO struct {
	Start *NodeDetailsDTO `json:"start"`
	End   *NodeDetailsDTO `json:"end"`
}
