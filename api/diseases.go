package api

import (
	"encoding/json"
	"forwardchaining/model"
	"net/http"
	"strconv"
)

func (api *API) FetchAllDiseases(w http.ResponseWriter, r *http.Request) {
	diseases, err := api.diseasesService.FetchAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(diseases)
}

func (api *API) StoreDiseases(w http.ResponseWriter, r *http.Request) {
	var disease model.Diseases
	json.NewDecoder(r.Body).Decode(&disease)

	err := json.NewDecoder(r.Body).Decode(&disease)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.diseasesService.Store(&disease)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(disease)
}

func (api *API) UpdateDiseases(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var disease model.Diseases
	err = json.NewDecoder(r.Body).Decode(&disease)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.diseasesService.Update(idInt, &disease)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(disease)
}

func (api *API) DeleteDiseases(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = api.diseasesService.Delete(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "Disease deleted successfully"})

}
