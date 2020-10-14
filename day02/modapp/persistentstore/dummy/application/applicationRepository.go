package application

import (
	dto "github.com/danushkaf/go-basic-samples/day02/modapp/dto"
	dummy "github.com/danushkaf/go-basic-samples/day02/modapp/persistentstore/dummy"
)

type DummyApplicationRepository struct {
	Session *dummy.Session
}

func (o *DummyApplicationRepository) Add(applicationDto dto.Application) error {
	return nil
}

func (o *DummyApplicationRepository) Get(partnerId string, id string) (bool, dto.Application, error) {
	return false, dto.Application{}, nil
}

func (o *DummyApplicationRepository) Delete(partnerId string, id string, applicationDto dto.Application) error {
	return nil
}
