package service

import (
	"forwardchaining/model"
	"forwardchaining/repository"
)

type DiagnosesService interface {
	FetchAll() ([]model.Diagnoses, error)
	FetchByID(id int) (*model.Diagnoses, error)
	Store(s *model.Diagnoses) error
	Update(id int, s *model.Diagnoses) error
	Delete(id int) error
}

type diagnosesService struct {
	diagnosesRepository repository.DiagnosesRepository
}

func NewDiagnosesService(diagnosesRepository repository.DiagnosesRepository) DiagnosesService {
	return &diagnosesService{diagnosesRepository}
}

func (s *diagnosesService) FetchAll() ([]model.Diagnoses, error) {
	diagnosess, err := s.diagnosesRepository.FetchAll()
	if err != nil {
		return nil, err
	}

	return diagnosess, nil
}

func (s *diagnosesService) FetchByID(id int) (*model.Diagnoses, error) {
	diagnoses, err := s.diagnosesRepository.FetchByID(id)
	if err != nil {
		return nil, err
	}

	return diagnoses, nil
}

func (s *diagnosesService) Store(diagnoses *model.Diagnoses) error {
	err := s.diagnosesRepository.Store(diagnoses)
	if err != nil {
		return err
	}

	return nil
}

func (s *diagnosesService) Update(id int, diagnoses *model.Diagnoses) error {
	err := s.diagnosesRepository.Update(id, diagnoses)
	if err != nil {
		return err
	}

	return nil
}

func (s *diagnosesService) Delete(id int) error {
	err := s.diagnosesRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
