package repository

import (
	"database/sql"
	"forwardchaining/model"
)

type DiseasesRepository interface {
	FetchAll() ([]model.Diseases, error)
	FetchByID(id int) (*model.Diseases, error)
	Store(s *model.Diseases) error
	Update(id int, s *model.Diseases) error
	Delete(id int) error
}

type diseasesRepoImpl struct {
	db *sql.DB
}

func NewDiseasesRepo(db *sql.DB) *diseasesRepoImpl {
	return &diseasesRepoImpl{db}
}

func (s *diseasesRepoImpl) FetchAll() ([]model.Diseases, error) {
	var diseasess []model.Diseases
	query := "SELECT id, code, name FROM diseases"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var diseases model.Diseases
		err := rows.Scan(&diseases.ID, &diseases.Code, &diseases.Name)
		if err != nil {
			return nil, err
		}
		diseasess = append(diseasess, diseases)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return diseasess, nil
}

func (s *diseasesRepoImpl) FetchByID(id int) (*model.Diseases, error) {
	row := s.db.QueryRow("SELECT id, name, code FROM diseases WHERE id = $1", id)

	var diseases model.Diseases
	err := row.Scan(&diseases.ID, &diseases.Code, &diseases.Name)
	if err != nil {
		return nil, err
	}

	return &diseases, nil
}

func (s *diseasesRepoImpl) Store(diseases *model.Diseases) error {

	_, err := s.db.Exec("INSERT INTO diseases (code, name) VALUES ($1, $2)", diseases.Code, diseases.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *diseasesRepoImpl) Update(id int, diseases *model.Diseases) error {
	_, err := s.db.Exec("UPDATE diseases SET code = $1, name= $2 WHERE id = $3", diseases.Code, diseases.Name, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *diseasesRepoImpl) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM diseases WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
