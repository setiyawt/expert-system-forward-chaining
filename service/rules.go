package service

import (
	"forwardchaining/model"
	"forwardchaining/repository"
)

type RulesService interface {
	FetchAll() ([]model.Rules, error)
	FetchByID(id int) (*model.Rules, error)
	Store(s *model.Rules) error
	Update(id int, s *model.Rules) error
	Delete(id int) error
}

type rulesService struct {
	rulesRepository repository.RulesRepository
}

func NewRulesService(rulesRepository repository.RulesRepository) RulesService {
	return &rulesService{rulesRepository}
}

func (s *rulesService) FetchAll() ([]model.Rules, error) {
	ruless, err := s.rulesRepository.FetchAll()
	if err != nil {
		return nil, err
	}

	return ruless, nil
}

func (s *rulesService) FetchByID(id int) (*model.Rules, error) {
	rules, err := s.rulesRepository.FetchByID(id)
	if err != nil {
		return nil, err
	}

	return rules, nil
}

func (s *rulesService) Store(rules *model.Rules) error {
	err := s.rulesRepository.Store(rules)
	if err != nil {
		return err
	}

	return nil
}

func (s *rulesService) Update(id int, rules *model.Rules) error {
	err := s.rulesRepository.Update(id, rules)
	if err != nil {
		return err
	}

	return nil
}

func (s *rulesService) Delete(id int) error {
	err := s.rulesRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
