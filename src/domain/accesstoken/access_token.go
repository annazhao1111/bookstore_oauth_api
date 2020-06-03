package accesstoken

import (
	"fmt"
	"strings"
	"time"

	"github.com/annazhao/bookstore_oauth_api/src/utils/cryptos"
	"github.com/annazhao/bookstore_oauth_api/src/utils/errors"
)

const (
	expirationTime       = 24
	grantTypePassword    = "password"
	grantTypeCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// used for client_credentials grant type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Validate method is used to validate if the AccessTokenRequest is valid or not
func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeCredentials:
		break
	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}

	// if at.GrantType != grantTypePassword || at.GrantType != grantTypeCredentials {
	// 	return errors.NewBadRequestError("invalid grant_type parameter")
	// }
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// Validate method is used to validate if the AccessToken is valid or not
func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserID <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientID <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired method is used to check if the access token is expired or not
func (at AccessToken) IsExpired() bool {
	// now := time.Now().UTC()
	// expirationTime := time.Unix(at.Expires, 0)
	// return now.After(expirationTime)
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

// Generate method is used to generate a new access token
func (at *AccessToken) Generate() {
	at.AccessToken = cryptos.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
