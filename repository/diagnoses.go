package repository

import (
	"database/sql"
	"forwardchaining/model"
)

type DiagnosesRepository interface {
	FetchAll() ([]model.Diagnoses, error)
	FetchByID(id int) (*model.Diagnoses, error)
	Store(s *model.Diagnoses) error
	Update(id int, s *model.Diagnoses) error
	Delete(id int) error
}

type diagnosesRepoImpl struct {
	db *sql.DB
}

func NewDiagnosesRepo(db *sql.DB) *diagnosesRepoImpl {
	return &diagnosesRepoImpl{db}
}

func (s *diagnosesRepoImpl) FetchAll() ([]model.Diagnoses, error) {
	var diagnosess []model.Diagnoses
	query := "SELECT id, name, nilai, description FROM diagnoses"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var diagnoses model.Diagnoses
		err := rows.Scan(&diagnoses.ID, &diagnoses.Name, &diagnoses.Nilai, &diagnoses.Description)
		if err != nil {
			return nil, err
		}
		diagnosess = append(diagnosess, diagnoses)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return diagnosess, nil
}

func (s *diagnosesRepoImpl) FetchByID(id int) (*model.Diagnoses, error) {
	row := s.db.QueryRow("SELECT id, name, nilai, description FROM diagnoses WHERE id = $1", id)

	var diagnoses model.Diagnoses
	err := row.Scan(&diagnoses.ID, &diagnoses.Name, &diagnoses.Nilai, &diagnoses.Description)
	if err != nil {
		return nil, err
	}

	return &diagnoses, nil
}

func (s *diagnosesRepoImpl) Store(diagnoses *model.Diagnoses) error {

	_, err := s.db.Exec("INSERT INTO diagnoses (name, nilai, description) VALUES ($1, $2, $3)", diagnoses.Name, diagnoses.Nilai, diagnoses.Description)
	if err != nil {
		return err
	}
	return nil
}

func (s *diagnosesRepoImpl) Update(id int, diagnoses *model.Diagnoses) error {
	_, err := s.db.Exec("UPDATE diagnoses SET name = $1, nilai = $2, description= $3 WHERE id = $4", diagnoses.Name, diagnoses.Nilai, diagnoses.Description, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *diagnosesRepoImpl) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM diagnoses WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
