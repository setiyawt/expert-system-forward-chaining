package main

import (
	"bufio"
	"fmt"
	"forwardchaining/model"
	repo "forwardchaining/repository"
	"os"
	"strings"
)

func RunCLI(userRepo *repo.UserRepository, sessionRepo *repo.SessionsRepository, questionsRepo *repo.QuestionsRepository, diagnosesRepo *repo.DiagnosesRepository, diseasesRepo *repo.DiseasesRepository, rulesRepo *repo.RulesRepository) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		fmt.Print("Pilih opsi: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			loginUser(reader, *userRepo, *&sessionRepo)
		case "2":
			registerUser(reader, userRepo)
		case "3":
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func loginUser(reader *bufio.Reader, userRepo repo.UserRepository, sessionRepo *repo.SessionsRepository) {
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	user := model.User{
		Username: username,
		Password: password,
	}

	// Check if the user exists and password is correct
	if err := userRepo.CheckPass(user, password); err != nil {
		fmt.Println("Username atau password salah.")
		return
	}

	session := model.Session{
		Username: username,
	}
	if err := sessionRepo.AddSession(session); err != nil {
		fmt.Println("Gagal memulai sesi.")
		return
	}

	fmt.Println("Login berhasil. Memulai proses diagnosis.")
	runDiagnosis(reader, diagnosesRepo, diseasesRepo, questionsRepo, rulesRepo)
}

func registerUser(reader *bufio.Reader, userRepo *repo.UserRepository) {
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	if err := userRepo.RegisterUser(username, password); err != nil {
		fmt.Println("Gagal mendaftar pengguna.")
		return
	}

	fmt.Println("Registrasi berhasil.")
}

func runDiagnosis(reader *bufio.Reader, diagnosesRepo *repo.DiagnosesRepository, diseasesRepo *repo.DiseasesRepository, questionsRepo *repo.QuestionsRepository, rulesRepo *repo.RulesRepository) {
	// Ambil pertanyaan
	questions, err := questionsRepo.GetAllQuestions()
	if err != nil {
		fmt.Println("Gagal mengambil pertanyaan.")
		return
	}

	var answers []model.Answer
	for _, question := range questions {
		fmt.Printf("%s: ", question.Question)
		answer := strings.TrimSpace(readInput(reader))
		answers = append(answers, model.Answer{
			QuestionID: question.ID,
			Response:   answer,
		})
	}

	// Terapkan forward chaining
	diagnosisResults, err := rulesRepo.ApplyForwardChaining(answers)
	if err != nil {
		fmt.Println("Gagal menerapkan forward chaining.")
		return
	}

	// Simpan hasil diagnosis
	for _, result := range diagnosisResults {
		if err := diagnosesRepo.SaveDiagnosis(result); err != nil {
			fmt.Println("Gagal menyimpan diagnosis.")
			return
		}
	}

	fmt.Println("Hasil Diagnosis:")
	for _, result := range diagnosisResults {
		fmt.Printf("Penyakit: %s\n", result.Name)
		fmt.Printf("Nilai: %.2f\n", result.Nilai)
		fmt.Printf("Deskripsi: %s\n", result.Description)
	}
}

func readInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
