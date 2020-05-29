package db

import (
	"github.com/annazhao/bookstore_oauth_api/src/domain/accesstoken"
	"github.com/annazhao/bookstore_oauth_api/src/utils/errors"
)

type dbRepository struct {
}

type DbRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
}

func New() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	return nil, nil
}
