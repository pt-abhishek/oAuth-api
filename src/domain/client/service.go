package client

import "github.com/pt-abhishek/oAuth-api/src/utils/errors"

//Service is the interface for client
type Service interface {
	Create() (*Client, *errors.RestErr)
}

type clientService struct {
	repository Repository
}

//NewClientService ClientService
func NewClientService(repo Repository) Service {
	return &clientService{
		repository: repo,
	}
}

func (c *clientService) Create() (*Client, *errors.RestErr) {
	var client Client
	client.CreateNew()
	if err := c.repository.Create(&client); err != nil {
		return nil, err
	}
	return &client, nil
}
