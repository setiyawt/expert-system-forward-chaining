package api

import (
	"fmt"
	"forwardchaining/service"
	"net/http"
)

type API struct {
	userService      service.UserService
	sessionService   service.SessionService
	diagnosesService service.DiagnosesService
	diseasesService  service.DiseasesService
	questionsService service.QuestionsService
	rulesService     service.RulesService
	symptomsService  service.SymptomsService
	mux              *http.ServeMux
}

func NewAPI(userService service.UserService, sessionService service.SessionService, diagnosesService service.DiagnosesService, diseasesService service.DiseasesService, questionsService service.QuestionsService, rulesService service.RulesService, symptomsService service.SymptomsService) API {
	mux := http.NewServeMux()
	api := API{
		userService,
		sessionService,
		diagnosesService,
		diseasesService,
		questionsService,
		rulesService,
		symptomsService,
		mux,
	}

	//USERS
	mux.Handle("/user/register", api.Post(http.HandlerFunc(api.Register)))
	mux.Handle("/user/login", api.Post(http.HandlerFunc(api.Login)))
	mux.Handle("/user/logout", api.Auth(http.HandlerFunc(api.Logout)))
	mux.Handle("/user/diagnoses", api.Get(api.Auth(http.HandlerFunc(api.FetchAllDiagnoses)))) //histori diagnose

	// ADMIN
	// Diseases
	mux.Handle("/diseases/get-all", api.Get(api.Auth(http.HandlerFunc(api.FetchAllDiseases))))
	mux.Handle("/diseases/add", api.Post(api.Auth(http.HandlerFunc(api.StoreDiseases))))
	mux.Handle("/diseases/update", api.Put(api.Auth(http.HandlerFunc(api.UpdateDiseases))))
	mux.Handle("/diseases/delete", api.Delete(api.Auth(http.HandlerFunc(api.DeleteDiseases))))

	// Question
	mux.Handle("/questions/get-all", api.Get(api.Auth(http.HandlerFunc(api.FetchAllQuestions))))
	mux.Handle("/questions/add", api.Post(api.Auth(http.HandlerFunc(api.StoreQuestions))))
	mux.Handle("/questions/update", api.Put(api.Auth(http.HandlerFunc(api.UpdateQuestions))))
	mux.Handle("/questions/delete", api.Delete(api.Auth(http.HandlerFunc(api.DeleteQuestions))))

	// Rules
	mux.Handle("/rules/get-all", api.Get(api.Auth(http.HandlerFunc(api.FetchAllRules))))
	mux.Handle("/rules/add", api.Post(api.Auth(http.HandlerFunc(api.StoreRules))))
	mux.Handle("/rules/update", api.Put(api.Auth(http.HandlerFunc(api.UpdateRules))))
	mux.Handle("/rules/delete", api.Delete(api.Auth(http.HandlerFunc(api.DeleteRules))))

	// Symptoms
	mux.Handle("/symptoms/get-all", api.Get(api.Auth(http.HandlerFunc(api.FetchAllSymptoms))))
	mux.Handle("/symptoms/add", api.Post(api.Auth(http.HandlerFunc(api.StoreSymptoms))))
	mux.Handle("/symptoms/update", api.Put(api.Auth(http.HandlerFunc(api.UpdateSymptoms))))
	mux.Handle("/symptoms/delete", api.Delete(api.Auth(http.HandlerFunc(api.DeleteSymptoms))))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}
