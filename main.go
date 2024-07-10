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
	diseaseCF := RunCLI(userRepo, sessionRepo, diagnosesRepo, diseasesRepo, questionsRepo, symptomsRepo, rulesRepo)
	for disease, cf := range diseaseCF {
		fmt.Printf("Disease: %s, CF: %.2f\n", disease, cf)
	}

}

func RunCLI(
	userRepo repo.UserRepository,
	sessionRepo repo.SessionsRepository,
	diagnosesRepo repo.DiagnosesRepository,
	diseasesRepo repo.DiseasesRepository,
	questionsRepo repo.QuestionsRepository,
	symptomsRepo repo.SymptomsRepository,
	rulesRepo repo.RulesRepository,
) map[string]float64 {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Question")
		// fmt.Println("4. Results")
		fmt.Println("4. Exit")
		fmt.Println("Select an option: ")

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
			os.Exit(0)
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

var confidenceLevels = map[string]float64{
	"1": 0.0, // Tidak Tahu
	"2": 0.2, // Tidak Yakin
	"3": 0.4, // Ragu-ragu
	"4": 0.6, // Cukup Yakin
	"5": 0.8, // Yakin
	"6": 1.0, // Sangat Yakin
}

func question(
	userRepo repo.UserRepository,
	sessionRepo repo.SessionsRepository,
	diseasesRepo repo.DiseasesRepository,
	symptomsRepo repo.SymptomsRepository,
	questionsRepo repo.QuestionsRepository,
	rulesRepo repo.RulesRepository,
	diagnosesRepo repo.DiagnosesRepository,
) map[string]float64 {
	reader := bufio.NewReader(os.Stdin)
	questionID := 1
	userAnswers := make(map[string]float64)

	fmt.Println("==================================")
	fmt.Println("Tidak Tahu: 1")
	fmt.Println("Tidak Yakin: 2")
	fmt.Println("Ragu-ragu: 3")
	fmt.Println("Cukup Yakin: 4")
	fmt.Println("Yakin: 5")
	fmt.Println("Sangat Yakin: 6")
	fmt.Println("==================================")

	// Fetch rules from repository
	rules, err := rulesRepo.FetchAll()
	if err != nil {
		fmt.Println("Error fetching rules:", err)
		return nil
	}

	// Debug: Print fetched rules
	fmt.Println("Fetched Rules:")
	for _, rule := range rules {
		fmt.Printf("Rule: %+v\n", rule)
	}

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

		if value, exists := confidenceLevels[answer]; exists {
			userAnswers[question.Code] = value
		} else {
			fmt.Println("Invalid answer, please try again.")
			continue
		}

		questionID++
	}

	// Debug: Print user answers
	fmt.Println("User Answers:")
	for code, cf := range userAnswers {
		fmt.Printf("Code: %s, CF: %.2f\n", code, cf)
	}

	// Calculate Certainty Factors
	diseaseCF := make(map[string]float64)

	for _, rule := range rules {
		if userCF, answered := userAnswers[rule.CodeSymptoms]; answered {
			// Forward chaining calculation
			cf := (float64(rule.Mb) - float64(rule.Md)) * userCF
			fmt.Printf("Calculating CF for disease %s: CF = (Mb - Md) * userCF = (%.2f - %.2f) * %.2f = %.2f\n",
				rule.CodeDeseases, float64(rule.Mb), float64(rule.Md), userCF, cf)

			if existingCF, exists := diseaseCF[rule.CodeDeseases]; exists {
				// Combine Certainty Factors using the provided formula
				diseaseCF[rule.CodeDeseases] = existingCF + cf*(1-existingCF)
			} else {
				diseaseCF[rule.CodeDeseases] = cf
			}
		}
	}

	// Debug: Print calculated Certainty Factors
	fmt.Println("Calculated Certainty Factors:")
	for disease, cf := range diseaseCF {
		fmt.Printf("Disease: %s, CF: %.2f\n", disease, cf)
	}

	return diseaseCF
}
