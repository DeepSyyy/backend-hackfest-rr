package request

type ClanCreateRequest struct {
	Name        string `json:"name" validate:"required,min=5,max=100"`
	Description string `json:"description" validate:"required,min=5,max=100"`
}
