package api

import (
	"encoding/json"
	"forwardchaining/model"
	"net/http"
	"strconv"
)

func (api *API) FetchAllRules(w http.ResponseWriter, r *http.Request) {
	rules, err := api.rulesService.FetchAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rules)
}

func (api *API) StoreRules(w http.ResponseWriter, r *http.Request) {
	var rules model.Rules
	json.NewDecoder(r.Body).Decode(&rules)
	err := json.NewDecoder(r.Body).Decode(&rules)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.rulesService.Store(&rules)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rules)
}

func (api *API) UpdateRules(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var rules model.Rules
	err = json.NewDecoder(r.Body).Decode(&rules)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.rulesService.Update(idInt, &rules)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rules)
}

func (api *API) DeleteRules(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = api.rulesService.Delete(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "Rule deleted successfully"})
}
