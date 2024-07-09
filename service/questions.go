package service

import (
	"forwardchaining/model"
	"forwardchaining/repository"
)

type QuestionsService interface {
	FetchAll() ([]model.Questions, error)
	FetchByID(id int) (*model.Questions, error)
	Store(s *model.Questions) error
	Update(id int, s *model.Questions) error
	Delete(id int) error
}

type questionsService struct {
	questionsRepository repository.QuestionsRepository
}

func NewQuestionsService(questionsRepository repository.QuestionsRepository) QuestionsService {
	return &questionsService{questionsRepository}
}

func (s *questionsService) FetchAll() ([]model.Questions, error) {
	questionss, err := s.questionsRepository.FetchAll()
	if err != nil {
		return nil, err
	}

	return questionss, nil
}

func (s *questionsService) FetchByID(id int) (*model.Questions, error) {
	questions, err := s.questionsRepository.FetchByID(id)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (s *questionsService) Store(questions *model.Questions) error {
	err := s.questionsRepository.Store(questions)
	if err != nil {
		return err
	}

	return nil
}

func (s *questionsService) Update(id int, questions *model.Questions) error {
	err := s.questionsRepository.Update(id, questions)
	if err != nil {
		return err
	}

	return nil
}

func (s *questionsService) Delete(id int) error {
	err := s.questionsRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
