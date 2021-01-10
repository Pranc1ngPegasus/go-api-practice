package model

type (
	User interface {
		ID() int64
		FirstName() string
		LastName() string
		Email() string
	}
)
