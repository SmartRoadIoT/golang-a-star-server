package utils

import (
	"math"
	"smart-road/main/model/dto"
)

const EarthRadius = 6371

func GetDistance(startNode *dto.NodeDetailsDTO, endNode *dto.NodeDetailsDTO) float64 {
	latitude1, longitude1 := convertToRadians(startNode)
	latitude2, longitude2 := convertToRadians(endNode)

	latitudeDifference := latitude2 - latitude1
	longitudeDifference := longitude2 - longitude1

	return EarthRadius * 2 * math.Asin(math.Sqrt(	math.Pow(math.Sin(latitudeDifference / 2), 2) +
													math.Cos(latitude1) * math.Cos(latitude2) *
													math.Pow(math.Sin(longitudeDifference / 2), 2)))
}

func convertToRadians(node *dto.NodeDetailsDTO) (float64, float64) {
	return node.Latitude * math.Pi / 180, node.Longitude * math.Pi / 180
}