package accesstoken

import (
	"crypto/rsa"
	"io/ioutil"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/google/uuid"
	"github.com/pt-abhishek/oAuth-api/src/utils/errors"
)

const (
	expirationTime       = 24
	grantTypeClientCreds = "client_credentials"
)

var (
	privateKey *rsa.PrivateKey
)

func init() {
	pkBytes, err := ioutil.ReadFile("repository/keys/app.rsa")
	if err != nil {
		panic(err)
	}
	privateKey, _ = jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
}

//AccessToken is the actual access token
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ClientID    string `json:"client_id"`
	Expires     int64  `json:"expires"`
	JWT         string `json:"bearer_token"`
}

//TokenRequest is the modified request
type TokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	//For grant type client_credentials
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	//User Info for when scope contains openIDConnect
	UserID int64 `json:"user_id"`
}

//CustomClaims are the claims we send
type CustomClaims struct {
	Token        string `json:"access_token"`
	IsAuthorized bool   `json:"is_authorized"`
	UserID       int64  `json:"user_id"`
	jwt.StandardClaims
}

//Repository is the DB
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(*TokenRequest) (*AccessToken, *errors.RestErr)
}

//GetNewAccessToken gets a new access token
func GetNewAccessToken(t *TokenRequest) *AccessToken {
	return &AccessToken{
		AccessToken: uuid.New().String(),
		Expires:     time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
		ClientID:    t.ClientID,
	}
}

//AddJWTToken gets a new jwt token
func (at *AccessToken) AddJWTToken(t *TokenRequest) *errors.RestErr {
	claims := CustomClaims{
		at.AccessToken,
		true,
		t.UserID,
		jwt.StandardClaims{
			ExpiresAt: at.Expires,
			Issuer:    "Oauth Service",
		},
	}
	idToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	idTokenString, err := idToken.SignedString(privateKey)
	if err != nil {
		return errors.NewInternalServerError("Error signing the token")
	}
	at.JWT = idTokenString
	return nil
}

//IsExpired checks if an access token is expired or not
func (at *AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return expirationTime.Before(now)
}

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
	if at.UserID == 0 {
		return errors.NewBadRequestError("Send a Non Zero user id, and it is mandatory")
	}
	return nil
}
