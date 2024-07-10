package api

import (
	"encoding/json"
	"forwardchaining/model"
	"net/http"
	"strconv"
)

func (api *API) FetchAllSymptoms(w http.ResponseWriter, r *http.Request) {
	symptoms, err := api.symptomsService.FetchAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(symptoms)
}

func (api *API) StoreSymptoms(w http.ResponseWriter, r *http.Request) {
	var symptoms model.Symptoms
	json.NewDecoder(r.Body).Decode(&symptoms)
	err := json.NewDecoder(r.Body).Decode(&symptoms)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.symptomsService.Store(&symptoms)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(symptoms)
}

func (api *API) UpdateSymptoms(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var symptoms model.Symptoms
	err = json.NewDecoder(r.Body).Decode(&symptoms)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.symptomsService.Update(idInt, &symptoms)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(symptoms)
}

func (api *API) DeleteSymptoms(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = api.symptomsService.Delete(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "Symptoms deleted successfully"})
}
