package service

import (
	"forwardchaining/model"
	"forwardchaining/repository"
)

type SymptomsService interface {
	FetchAll() ([]model.Symptoms, error)
	FetchByID(id int) (*model.Symptoms, error)
	Store(s *model.Symptoms) error
	Update(id int, s *model.Symptoms) error
	Delete(id int) error
}

type symptomsService struct {
	symptomsRepository repository.SymptomsRepository
}

func NewSymptomsService(symptomsRepository repository.SymptomsRepository) SymptomsService {
	return &symptomsService{symptomsRepository}
}

func (s *symptomsService) FetchAll() ([]model.Symptoms, error) {
	symptomss, err := s.symptomsRepository.FetchAll()
	if err != nil {
		return nil, err
	}

	return symptomss, nil
}

func (s *symptomsService) FetchByID(id int) (*model.Symptoms, error) {
	symptoms, err := s.symptomsRepository.FetchByID(id)
	if err != nil {
		return nil, err
	}

	return symptoms, nil
}

func (s *symptomsService) Store(symptoms *model.Symptoms) error {
	err := s.symptomsRepository.Store(symptoms)
	if err != nil {
		return err
	}

	return nil
}

func (s *symptomsService) Update(id int, symptoms *model.Symptoms) error {
	err := s.symptomsRepository.Update(id, symptoms)
	if err != nil {
		return err
	}

	return nil
}

func (s *symptomsService) Delete(id int) error {
	err := s.symptomsRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
