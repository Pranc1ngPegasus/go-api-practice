package repository

import (
	"github.com/Pranc1ngPegasus/go-api-practice/internal/domain/model"
)

type (
	User interface {
		Get(id int64) (model.User, error)
	}
)
