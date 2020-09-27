package client

import (
	"github.com/google/uuid"
	"github.com/pt-abhishek/oAuth-api/src/utils/errors"
)

//Client is the Client subscribing to our auth service
type Client struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

//Repository is the interface for client Repository
type Repository interface {
	Create(*Client) *errors.RestErr
	Validate(*Client) *errors.RestErr
}

//CreateNew a new client model
func (c *Client) CreateNew() {
	clientID := uuid.New()
	clientSecret := uuid.New()
	c.ClientID = clientID.String()
	c.ClientSecret = clientSecret.String()
}
