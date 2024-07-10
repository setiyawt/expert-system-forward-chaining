package main

import (
	"bufio"
	"fmt"
	"forwardchaining/api"
	"forwardchaining/db"
	"forwardchaining/model"
	repo "forwardchaining/repository"
	"forwardchaining/service"
	"log"
	"os"
	"strings"
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

	// Start the web server in a separate goroutine
	go mainAPI.Start()

	// Run CLI
	RunCLI(userRepo, sessionRepo, diagnosesRepo, diseasesRepo, questionsRepo, symptomsRepo, rulesRepo)
}

func RunCLI(
	userRepo repo.UserRepository,
	sessionRepo repo.SessionsRepository,
	diagnosesRepo repo.DiagnosesRepository,
	diseasesRepo repo.DiseasesRepository,
	questionsRepo repo.QuestionsRepository,
	symptomsRepo repo.SymptomsRepository,
	rulesRepo repo.RulesRepository,
) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Question")
		// fmt.Println("4. Results")
		fmt.Println("4. Exit")
		fmt.Print("Select an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			login(userRepo, sessionRepo)
		case "2":
			register(userRepo)
		case "3":
			question(userRepo, sessionRepo, diseasesRepo, symptomsRepo, questionsRepo, rulesRepo, diagnosesRepo)
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func login(
	userRepo repo.UserRepository,
	sessionRepo repo.SessionsRepository,
) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	user, err := userRepo.FetchByID(1) // Gantilah metode FetchByID dengan metode yang tepat
	if err != nil || user.Password != password {
		fmt.Println("Invalid username or password.")
		return
	}

	fmt.Println("Login successful.")
	// Proceed to ask questions and analyze the answers
	// Add your logic here
}

func register(userRepo repo.UserRepository) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	user := model.User{
		Username: username,
		Password: password,
		// Fill other fields as necessary
	}

	err := userRepo.Add(user)
	if err != nil {
		fmt.Println("Failed to register user:", err)
		return
	}

	fmt.Println("User registered successfully.")
}

func question(
	userRepo repo.UserRepository,
	sessionRepo repo.SessionsRepository,
	diseasesRepo repo.DiseasesRepository,
	symptomsRepo repo.SymptomsRepository,
	questionsRepo repo.QuestionsRepository,
	rulesRepo repo.RulesRepository,
	diagnosesRepo repo.DiagnosesRepository,
) {
	reader := bufio.NewReader(os.Stdin)
	questionID := 1 // Mulai dari pertanyaan pertama

	for {
		question, err := questionsRepo.FetchByID(questionID)
		if err != nil {
			fmt.Println("No more questions.")
			break
		}

		fmt.Println(question.Question)
		fmt.Print("Answer: ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		// Lakukan sesuatu dengan jawaban di sini, misalnya menyimpan jawaban

		questionID++
	}

}
