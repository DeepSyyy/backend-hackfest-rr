package request

import "github.com/google/uuid"

type UpdateUsersRequest struct {
	Id       uuid.UUID `validate:"required"`
	Username string    `json:"username" validate:"required,min=5,max=100"`
	Password string    `json:"password" validate:"required,min=8,max=100"`
	Email    string    `json:"email" validate:"required,email"`
}
