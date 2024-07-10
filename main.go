package main

import (
	"forwardchaining/api"
	"forwardchaining/db"
	"forwardchaining/model"
	repo "forwardchaining/repository"
	"forwardchaining/service"
	"log"
)

func main() {
	// Database credentials
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "forwardchaining",
		Port:         5432,
	}

	// Connect to the database
	dbConn, err := db.Connect(&dbCredential)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Execute SQL scripts
	err = db.SQLExecute(dbConn)
	if err != nil {
		log.Fatalf("Failed to execute SQL scripts: %v", err)
	}

	defer dbConn.Close()

	// Initialize repositories
	userRepo := repo.NewUserRepo(dbConn)
	sessionRepo := repo.NewSessionRepo(dbConn)
	diagnosesRepo := repo.NewDiagnosesRepo(dbConn)
	diseasesRepo := repo.NewDiseasesRepo(dbConn)
	questionsRepo := repo.NewQuestionsRepo(dbConn)
	rulesRepo := repo.NewRulesRepo(dbConn)
	symptomsRepo := repo.NewSymptomsRepo(dbConn)

	// Initialize services
	userService := service.NewUserService(userRepo)
	sessionService := service.NewSessionService(sessionRepo)
	diagnosesService := service.NewDiagnosesService(diagnosesRepo)
	diseasesService := service.NewDiseasesService(diseasesRepo)
	questionsService := service.NewQuestionsService(questionsRepo)
	rulesService := service.NewRulesService(rulesRepo)
	symptomsService := service.NewSymptomsService(symptomsRepo)

	// Create new API
	mainAPI := api.NewAPI(userService, sessionService, diagnosesService, diseasesService, questionsService, rulesService, symptomsService)
	mainAPI.Start()
	RunCLI(userRepo, sessionRepo, questionsRepo, diagnosesRepo)
}
