package api

import (
	"encoding/json"
	"forwardchaining/model"
	"net/http"
	"strconv"
)

func (api *API) FetchAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := api.questionsService.FetchAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(questions)
}

func (api *API) StoreQuestions(w http.ResponseWriter, r *http.Request) {
	var questions model.Questions
	json.NewDecoder(r.Body).Decode(&questions)
	err := json.NewDecoder(r.Body).Decode(&questions)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.questionsService.Store(&questions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(questions)
}

func (api *API) UpdateQuestions(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var questions model.Questions
	err = json.NewDecoder(r.Body).Decode(&questions)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.questionsService.Update(idInt, &questions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(questions)
}

func (api *API) DeleteQuestions(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = api.questionsService.Delete(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "Question deleted successfully"})
}
