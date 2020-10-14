package shortcode

import (
	dto "github.com/danushkaf/go-basic-samples/day02/modapp/dto"
	repository "github.com/danushkaf/go-basic-samples/day02/modapp/services/application/repository"
)

// Service provides shortcode related operations
type Service interface {
	AddApplication(dto.Application) dto.ServiceError
	Get(string, string) (dto.Application, dto.ServiceError)
	Delete(string, string) dto.ServiceError
}

type service struct {
	applicationRepository repository.Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(applicationRepository repository.Repository) Service {
	return &service{applicationRepository}
}

func (s *service) AddApplication(application dto.Application) dto.ServiceError {
	s.applicationRepository.Add(application)
	return dto.ServiceError{}
}

func (s *service) Get(partnerId string, id string) (dto.Application, dto.ServiceError) {
	_, application, _ := s.applicationRepository.Get(partnerId, id)
	return application, dto.ServiceError{}
}

func (s *service) Delete(partnerId string, id string) dto.ServiceError {
	_, application, _ := s.applicationRepository.Get(partnerId, id)
	s.applicationRepository.Delete(partnerId, id, application)
	return dto.ServiceError{}
}
