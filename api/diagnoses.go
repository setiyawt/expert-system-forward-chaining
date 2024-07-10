package api

import (
	"encoding/json"
	"net/http"
)

func (api *API) FetchAllDiagnoses(w http.ResponseWriter, r *http.Request) {
	diagonoses, err := api.diagnosesService.FetchAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(diagonoses)
}
