package accesstoken

import (
	"strings"

	"github.com/pt-abhishek/oAuth-api/src/utils/errors"
)

//Service interface
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

//NewService Creates an instance of the service
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetByID(id string) (*AccessToken, *errors.RestErr) {
	if len(strings.TrimSpace(id)) == 0 {
		return nil, errors.NewBadRequestError("Invalid access token id")
	}
	token, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *service) Create(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
