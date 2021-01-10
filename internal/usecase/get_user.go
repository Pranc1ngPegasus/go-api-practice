package usecase

import (
	repoImpl "github.com/Pranc1ngPegasus/go-api-practice/internal/adapter/domainimpl/repository"
	"github.com/Pranc1ngPegasus/go-api-practice/internal/adapter/infrastructure"
	"github.com/Pranc1ngPegasus/go-api-practice/internal/domain/model"
	"github.com/Pranc1ngPegasus/go-api-practice/internal/domain/repository"
)

type (
	GetUser interface {
		Execute(input GetUserInput) (model.User, error)
	}

	getUser struct {
		repository repository.User
	}
)

func NewGetUser(
	connector infrastructure.RDBConnector,
) GetUser {
	return &getUser{
		repository: repoImpl.NewUserImpl(connector),
	}
}

type (
	GetUserInput struct {
		ID int64
	}
)

func (u *getUser) Execute(input GetUserInput) (model.User, error) {
	user, err := u.repository.Get(input.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
