package repository

import (
	"database/sql"
	"forwardchaining/model"
)

type QuestionsRepository interface {
	FetchAll() ([]model.Questions, error)
	FetchByID(id int) (*model.Questions, error)
	Store(s *model.Questions) error
	Update(id int, s *model.Questions) error
	Delete(id int) error
}

type questionsRepoImpl struct {
	db *sql.DB
}

func NewQuestionsRepo(db *sql.DB) *questionsRepoImpl {
	return &questionsRepoImpl{db}
}

func (s *questionsRepoImpl) FetchAll() ([]model.Questions, error) {
	var questionss []model.Questions
	query := "SELECT id, code, question FROM questions"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var questions model.Questions
		err := rows.Scan(&questions.ID, &questions.Code, &questions.Question)
		if err != nil {
			return nil, err
		}
		questionss = append(questionss, questions)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return questionss, nil
}

func (s *questionsRepoImpl) FetchByID(id int) (*model.Questions, error) {
	row := s.db.QueryRow("SELECT id, code, question FROM questions WHERE id = $1", id)

	var questions model.Questions
	err := row.Scan(&questions.ID, &questions.Code, &questions.Question)
	if err != nil {
		return nil, err
	}

	return &questions, nil
}

func (s *questionsRepoImpl) Store(questions *model.Questions) error {

	_, err := s.db.Exec("INSERT INTO questions (code, question) VALUES ($1, $2, $3)", questions.Code, questions.Question)
	if err != nil {
		return err
	}
	return nil
}

func (s *questionsRepoImpl) Update(id int, questions *model.Questions) error {
	_, err := s.db.Exec("UPDATE questions SET code = $1, code= $2 WHERE id = $3", questions.Code, questions.Question, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *questionsRepoImpl) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM questions WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
