package accesstoken

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pt-abhishek/oAuth-api/src/utils/errors"
)

const (
	expirationTime       = 24
	grantTypeClientCreds = "client_credentials"
)

//AccessToken is the actual access token
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ClientID    string `json:"client_id"`
	Expires     int64  `json:"expires"`
	JWT         string `json:"jwt"`
}

//TokenRequest is the modified request
type TokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	//For grant type client_credentials
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

//Repository is the DB
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(*TokenRequest) (*AccessToken, *errors.RestErr)
	// UpdateExpirationTime(AccessToken) *errors.RestErr
}

//GetNewAccessToken gets a new access token
func GetNewAccessToken(t *TokenRequest) *AccessToken {
	return &AccessToken{
		AccessToken: uuid.New().String(),
		Expires:     time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
		ClientID:    t.ClientID,
	}
}

//IsExpired checks if an access token is expired or not
func (t *AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(t.Expires, 0)
	return expirationTime.Before(now)
}

//Validate validates for empty access tokens
// func (t *AccessToken) Validate() *errors.RestErr {
// 	t.AccessToken = strings.TrimSpace(t.AccessToken)
// 	if len(t.AccessToken) == 0 {
// 		return errors.NewBadRequestError("Invalid access token in request")
// 	}
// 	if t.UserID <= 0 {
// 		return errors.NewBadRequestError("Invalid User ID")
// 	}
// 	if t.ClientID <= 0 {
// 		return errors.NewBadRequestError("Invalid Client ID")
// 	}
// 	if t.Expires <= 0 {
// 		return errors.NewBadRequestError("Invalid Expiration time")
// 	}
// 	return nil
// }

//Validate validates the tokenrequest
func (at *TokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypeClientCreds:
		break
	default:
		return errors.NewBadRequestError("Invalid grant type")
	}

	at.ClientSecret = strings.TrimSpace(at.ClientSecret)
	at.ClientID = strings.TrimSpace(at.ClientID)
	at.Scope = strings.TrimSpace(at.Scope)

	if at.ClientSecret == "" {
		return errors.NewBadRequestError("Client Secret mandatory")
	}
	if at.ClientID == "" {
		return errors.NewBadRequestError("Client ID mandatory")
	}
	if at.Scope == "" {
		return errors.NewBadRequestError("Scope mandatory")
	}
	return nil
}
