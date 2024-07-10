package model

import "time"

type User struct {
	ID       int    `json:"id"`
	Username string `gorm:"type:varchar(100);unique"`
	Password string `json:"password"`
}

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

type Session struct {
	ID       int       `json:"id"`
	Token    string    `json:"token"`
	Username string    `json:"username"`
	Expiry   time.Time `json:"expiry"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type Diagnoses struct { //sebagai jawaban hasil tes
	ID          int     `json:"id"`
	Name        string  `json:"name"`        // Nama penyakitnya
	Nilai       float32 `json:"nilai"`       // Nilai dari Certainty Factor
	Description string  `json:"description"` //Deskripsi dari penyakit tersebut
}

type Symptoms struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"` // Nama gejala
}

type Diseases struct {
	ID   int    `json:"id"`
	Code string `json:"code"` // Kode penyakit
	Name string `json:"name"` // Nama penyakit
}

type Questions struct {
	ID       int    `json:"id"`
	Code     string `json:"code"`     // Kode Question
	Question string `json:"question"` // Pertanyaan yang dijawab oleh user
}

type Rules struct {
	ID           int     `json:"id"`
	CodeDeseases string  `json:"code_deseases"` // Kode Penyakit
	CodeSymptoms string  `json:"code_symptoms"` // Kode Gejala
	Md           float32 `json:"md"`            //Measure of Doubt
	Mb           float32 `json:"mb"`            //Measure of Belief
}
