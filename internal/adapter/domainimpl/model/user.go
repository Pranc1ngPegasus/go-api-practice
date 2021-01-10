package model

import (
	infraModel "github.com/Pranc1ngPegasus/go-api-practice/internal/adapter/infrastructure/model"
	"github.com/Pranc1ngPegasus/go-api-practice/internal/domain/model"
)

var _ model.User = (*userDTO)(nil)

type (
	userDTO struct {
		id        int64
		firstName string
		lastName  string
		email     string
	}
)

func NewUser() model.User {
	return newUserDTO()
}

func newUserDTO() *userDTO {
	return &userDTO{}
}

func UserFromRecord(
	record *infraModel.User,
) model.User {
	dto := newUserDTO()

	dto.id = record.ID
	dto.firstName = record.FirstName
	dto.lastName = record.LastName
	dto.email = record.Email

	return dto
}

func (dto *userDTO) ID() int64 {
	return dto.id
}

func (dto *userDTO) FirstName() string {
	return dto.firstName
}

func (dto *userDTO) LastName() string {
	return dto.lastName
}

func (dto *userDTO) Email() string {
	return dto.email
}
