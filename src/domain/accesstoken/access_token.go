package accesstoken

import (
	"strings"
	"time"

	"github.com/pt-abhishek/oAuth-api/src/utils/errors"
)

const (
	expirationTime = 24
)

//AccessToken is the actual access token
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

//Repository is the DB
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

//GetNewAccessToken gets a new access token
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

//IsExpired checks if an access token is expired or not
func (t *AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(t.Expires, 0)
	return expirationTime.Before(now)
}

//Validate validates for empty access tokens
func (t *AccessToken) Validate() *errors.RestErr {
	t.AccessToken = strings.TrimSpace(t.AccessToken)
	if len(t.AccessToken) == 0 {
		return errors.NewBadRequestError("Invalid access token in request")
	}
	if t.UserID <= 0 {
		return errors.NewBadRequestError("Invalid User ID")
	}
	if t.ClientID <= 0 {
		return errors.NewBadRequestError("Invalid Client ID")
	}
	if t.Expires <= 0 {
		return errors.NewBadRequestError("Invalid Expiration time")
	}
	return nil
}
