package dto

type NodeDTO struct {
	Id              uint64  	`json:"id"`
	Connected 		[]uint64 	`json:"connected"`
}