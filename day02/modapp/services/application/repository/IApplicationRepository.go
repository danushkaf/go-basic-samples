package repository

import dto "github.com/danushkaf/go-basic-samples/day02/modapp/dto"

type Repository interface {
	// Create a application.
	Add(dto.Application) error
	// Returns the application with given ID.
	Get(string, string) (bool, dto.Application, error)

	Delete(string, string, dto.Application) error
}
