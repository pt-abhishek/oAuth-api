package db

import (
	"github.com/pt-abhishek/oAuth-api/src/clients/cassandra"
	"github.com/pt-abhishek/oAuth-api/src/domain/client"
	"github.com/pt-abhishek/oAuth-api/src/utils/errors"
)

const (
	queryCreateClient = "INSERT INTO clients(client_id, client_secret) VALUES(?,?)"
	queryFindByID     = "SELECT client_id, client_secret FROM clients WHERE client_id = ?"
)

//ClientRepository is the client DB interface
type ClientRepository interface {
	Create(*client.Client) *errors.RestErr
	Validate(*client.Client) *errors.RestErr
}

type clientRepo struct{}

//NewClient returns a new clientRepo instannce
func NewClient() ClientRepository {
	return &clientRepo{}
}

func (r *clientRepo) Create(c *client.Client) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateClient, c.ClientID, c.ClientSecret).Exec(); err != nil {
		return errors.NewInternalServerError("Unable To create client , DB error")
	}
	return nil
}

func (r *clientRepo) Validate(c *client.Client) *errors.RestErr {
	var existingClient client.Client
	if err := cassandra.GetSession().Query(queryFindByID, c.ClientID).Scan(&existingClient.ClientID, &existingClient.ClientSecret); err != nil {
		return errors.NewResourceNotFoundError("No such Client found")
	}
	if existingClient.ClientSecret != c.ClientSecret {
		return errors.NewBadRequestError("Invalid Client")
	}
	return nil
}
