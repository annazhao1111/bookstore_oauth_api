package rest

import (
	"time"

	"encoding/json"

	"github.com/annazhao/bookstore_oauth_api/src/domain/users"
	"github.com/annazhao/bookstore_oauth_api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var usersRestClient = rest.RequestBuilder{
	BaseURL: "http://localhost:8082",
	Timeout: 100 * time.Millisecond,
}

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {
}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		// if we have a time out
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}
	if response.StatusCode > 299 {
		// if we have an error situation
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			// if somebody change the RestError interface
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users login response")
	}
	return &user, nil
}
