package accesstoken

import "github.com/annazhao/bookstore_oauth_api/src/utils/errors"

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{}
}

func (s *service) GetByID(string) (*AccessToken, *errors.RestErr) {
	return nil, nil
}
