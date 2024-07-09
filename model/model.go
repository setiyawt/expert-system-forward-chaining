package model

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
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
	CodeSymptoms string  `json:"code_symptoms"` // Kode GEjala
	Md           float32 `json:"md"`            //Measure of Doubt
	Mb           float32 `json:"mb"`            //Measure of Belief
}
