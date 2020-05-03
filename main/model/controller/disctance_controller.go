package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"smart-road/main/model/dto"
	"smart-road/main/utils"
)

func HandleDistanceRequest(w http.ResponseWriter, r *http.Request) {
	var distance dto.DistanceDTO

	err := json.NewDecoder(r.Body).Decode(&distance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Distance: %+v", utils.GetDistance(distance.Start, distance.End))
}
