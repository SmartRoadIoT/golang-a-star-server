package dto

type FindRouteDTO struct {
	Graph           *GraphDTO  `json:"graph"`
	Nodes        	*[]NodeDetailsDTO `json:"nodes"`
	Start       	uint64 `json:"start"`
	End       		uint64 `json:"end"`
}