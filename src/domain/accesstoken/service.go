package accesstoken

import (
	"strings"

	"github.com/pt-abhishek/oAuth-api/src/utils/errors"
)

//Service interface
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(*TokenRequest) (*AccessToken, *errors.RestErr)
}

type tokenService struct {
	repository Repository
}

//NewTokenService Creates an instance of the service
func NewTokenService(repo Repository) Service {
	return &tokenService{
		repository: repo,
	}
}

func (s *tokenService) GetByID(id string) (*AccessToken, *errors.RestErr) {
	if len(strings.TrimSpace(id)) == 0 {
		return nil, errors.NewBadRequestError("Invalid access token id")
	}
	token, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *tokenService) Create(req *TokenRequest) (*AccessToken, *errors.RestErr) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	at, err := s.repository.Create(req)
	if err != nil {
		return nil, err
	}
	return at, nil
}
