package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"smart-road/main/model/dto"
)

func HandleFindRouteRequest(w http.ResponseWriter, r *http.Request) {
	var findRoute dto.FindRouteDTO

	err := json.NewDecoder(r.Body).Decode(&findRoute)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Distance: %+v", findRoute.Graph.Nodes)
}
