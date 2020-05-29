package dto

import (
	"errors"
	"sync"
)

type NodeDTO struct {
	Id              uint64  	`json:"id"`
	Connected 		[]uint64 	`json:"connected"`
	maps 		map[uint64]int
	sync.RWMutex
}

func (n *NodeDTO) GetNeighbors() map[uint64]int {
	if n == nil {
		return nil
	}

	n.RLock()
	neighbors := n.maps
	n.RUnlock()

	return neighbors
}


// Returns the Nodees key.
func (n *NodeDTO) Key() uint64 {
	if n == nil {
		return -1
	}

	n.RLock()
	key := n.Id
	n.RUnlock()

	return key
}

func (n *NodeDTO) mapping() (maps  map[uint64]int) {




	for i := 0; i < len(n.Connected); i ++{
		n.RLock()
		maps[n.Connected[i]] = utils.GetDistance(startNode *dto.NodeDetailsDTO, endNode *dto.NodeDetailsDTO)
		n.RUnlock()
	}



	return
}