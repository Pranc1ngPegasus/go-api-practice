package repository

import (
	modelImpl "github.com/Pranc1ngPegasus/go-api-practice/internal/adapter/domainimpl/model"
	"github.com/Pranc1ngPegasus/go-api-practice/internal/adapter/infrastructure"
	infraModel "github.com/Pranc1ngPegasus/go-api-practice/internal/adapter/infrastructure/model"
	"github.com/Pranc1ngPegasus/go-api-practice/internal/domain/model"
	"github.com/Pranc1ngPegasus/go-api-practice/internal/domain/repository"
)

var _ repository.User = (*userImpl)(nil)

type (
	userImpl struct {
		connector infrastructure.RDBConnector
	}
)

func NewUserImpl(
	connector infrastructure.RDBConnector,
) repository.User {
	return &userImpl{
		connector: connector,
	}
}

func (impl *userImpl) Get(id int64) (model.User, error) {
	record, err := infraModel.Users(
		infraModel.UserWhere.ID.EQ(id),
	).One(
		impl.connector.GetContext(),
		impl.connector.GetDB(),
	)
	if err != nil {
		return nil, err
	}

	user := modelImpl.UserFromRecord(record)

	return user, nil
}
