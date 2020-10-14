package dto

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ServiceError struct {
	Code    string
	Message string
}
