package dto

type Application struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PartnerID   string `json:"partner_id"`
}
