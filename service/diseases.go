package service

import (
	"forwardchaining/model"
	"forwardchaining/repository"
)

type DiseasesService interface {
	FetchAll() ([]model.Diseases, error)
	FetchByID(id int) (*model.Diseases, error)
	Store(s *model.Diseases) error
	Update(id int, s *model.Diseases) error
	Delete(id int) error
}

type diseasesService struct {
	diseasesRepository repository.DiseasesRepository
}

func NewDiseasesService(diseasesRepository repository.DiseasesRepository) DiseasesService {
	return &diseasesService{diseasesRepository}
}

func (s *diseasesService) FetchAll() ([]model.Diseases, error) {
	diseasess, err := s.diseasesRepository.FetchAll()
	if err != nil {
		return nil, err
	}

	return diseasess, nil
}

func (s *diseasesService) FetchByID(id int) (*model.Diseases, error) {
	diseases, err := s.diseasesRepository.FetchByID(id)
	if err != nil {
		return nil, err
	}

	return diseases, nil
}

func (s *diseasesService) Store(diseases *model.Diseases) error {
	err := s.diseasesRepository.Store(diseases)
	if err != nil {
		return err
	}

	return nil
}

func (s *diseasesService) Update(id int, diseases *model.Diseases) error {
	err := s.diseasesRepository.Update(id, diseases)
	if err != nil {
		return err
	}

	return nil
}

func (s *diseasesService) Delete(id int) error {
	err := s.diseasesRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
