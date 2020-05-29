package dto

import (
	"errors"
	"sync"
)

type GraphDTO struct {
	mu    sync.RWMutex
	Nodes []*NodeDTO `json:"nodes"`
	nodes map[uint64]*NodeDTO

}

func (g *GraphDTO) Len() int {
	return len(g.Nodes)
}
func (g *GraphDTO) Get(key uint64) (n *NodeDTO, err error) {
	g.mu.RLock()
	n = g.get(key)
	g.mu.RUnlock()

	if n == nil {
		err = errors.New("graph: invalid key")
	}

	return
}

func (g *GraphDTO) get(key uint64) *NodeDTO {
	return g.nodes[key]
}
func (g *GraphDTO) mapping() ( nodes map[uint64]*NodeDTO , err error ) {




	for i := 0; i < len(g.Nodes); i ++{
		g.mu.RLock()
		nodes[g.Nodes[i].Id] = g.Nodes[i]
		g.mu.RUnlock()
	}

	if nodes == nil {
		err = errors.New("graph: invalid key")
	}

	return
}