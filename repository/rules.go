package repository

import (
	"database/sql"
	"forwardchaining/model"
)

type RulesRepository interface {
	FetchAll() ([]model.Rules, error)
	FetchByID(id int) (*model.Rules, error)
	Store(s *model.Rules) error
	Update(id int, s *model.Rules) error
	Delete(id int) error
}

type rulesRepoImpl struct {
	db *sql.DB
}

func NewRulesRepo(db *sql.DB) *rulesRepoImpl {
	return &rulesRepoImpl{db}
}

func (s *rulesRepoImpl) FetchAll() ([]model.Rules, error) {
	var ruless []model.Rules
	query := "SELECT id, code_diseases, code_symptoms, md, mb FROM rules"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rules model.Rules
		err := rows.Scan(&rules.ID, &rules.CodeDeseases, &rules.CodeSymptoms, &rules.Md, &rules.Mb)
		if err != nil {
			return nil, err
		}
		ruless = append(ruless, rules)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ruless, nil
}

func (s *rulesRepoImpl) FetchByID(id int) (*model.Rules, error) {
	row := s.db.QueryRow("SELECT id, code_diseases, code_symptoms, md, mb FROM rules WHERE id = $1", id)

	var rules model.Rules
	err := row.Scan(&rules.ID, &rules.CodeDeseases, &rules.CodeSymptoms, &rules.Md, &rules.Mb)
	if err != nil {
		return nil, err
	}

	return &rules, nil
}

func (s *rulesRepoImpl) Store(rules *model.Rules) error {

	_, err := s.db.Exec("INSERT INTO rules (code_diseases, code_symptoms, md, mb) VALUES ($1, $2, $3, $4)", rules.CodeDeseases, rules.CodeSymptoms, rules.Md, rules.Mb)
	if err != nil {
		return err
	}
	return nil
}

func (s *rulesRepoImpl) Update(id int, rules *model.Rules) error {
	_, err := s.db.Exec("UPDATE rules SET code_diseases = $1, code_symptoms = $2, md= $3, mb= $4 WHERE id = $5", rules.CodeDeseases, rules.CodeSymptoms, rules.Md, rules.Mb, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *rulesRepoImpl) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM rules WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
