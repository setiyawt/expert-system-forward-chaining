package repository

import (
	"database/sql"
	"forwardchaining/model"
)

type SymptomsRepository interface {
	FetchAll() ([]model.Symptoms, error)
	FetchByID(id int) (*model.Symptoms, error)
	Store(s *model.Symptoms) error
	Update(id int, s *model.Symptoms) error
	Delete(id int) error
}

type symptomsRepoImpl struct {
	db *sql.DB
}

func NewSymptomsRepo(db *sql.DB) *symptomsRepoImpl {
	return &symptomsRepoImpl{db}
}

func (s *symptomsRepoImpl) FetchAll() ([]model.Symptoms, error) {
	var symptomss []model.Symptoms
	query := "SELECT id, code, name FROM symptoms"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var symptoms model.Symptoms
		err := rows.Scan(&symptoms.ID, &symptoms.Code, &symptoms.Name)
		if err != nil {
			return nil, err
		}
		symptomss = append(symptomss, symptoms)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return symptomss, nil
}

func (s *symptomsRepoImpl) FetchByID(id int) (*model.Symptoms, error) {
	row := s.db.QueryRow("SELECT id, code, name FROM symptoms WHERE id = $1", id)

	var symptoms model.Symptoms
	err := row.Scan(&symptoms.ID, &symptoms.Code, &symptoms.Name)
	if err != nil {
		return nil, err
	}

	return &symptoms, nil
}

func (s *symptomsRepoImpl) Store(symptoms *model.Symptoms) error {

	_, err := s.db.Exec("INSERT INTO symptoms (code, name) VALUES ($1, $2)", symptoms.Code, symptoms.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *symptomsRepoImpl) Update(id int, symptoms *model.Symptoms) error {
	_, err := s.db.Exec("UPDATE symptoms SET code = $1, name= $2 WHERE id = $3", symptoms.Code, symptoms.Name, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *symptomsRepoImpl) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM symptoms WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
