package test

import (
	"math"
	"smart-road/main/model/dto"
	"smart-road/main/utils"
	"testing"
)

const float64EqualityThreshold = 1e-9

func TestGetDistance(t *testing.T) {
	tables := []struct {
		start dto.NodeDetailsDTO
		end dto.NodeDetailsDTO
		expected float64
	} {
		{	dto.NodeDetailsDTO{Latitude: 53.32055555555556, Longitude: -1.729722222222221},
			dto.NodeDetailsDTO{Latitude: 53.31861111111111, Longitude: -1.6997222222222223},
			2.0043678382716137},
		{	dto.NodeDetailsDTO{Latitude: 53.32055555555556, Longitude: -1.729722222222221},
			dto.NodeDetailsDTO{Latitude: 53.32055555555556, Longitude: -1.729722222222221},
			0},
	}

	for _, table := range tables {
		distance := utils.GetDistance(table.start, table.end)
		if !almostEqual(distance, table.expected) {
			t.Errorf("Distance test failed, distance: %f, expected: %f.", distance, table.expected)
		}
	}
}

func almostEqual(a, b float64) bool {
	return math.Abs(a - b) <= float64EqualityThreshold
}