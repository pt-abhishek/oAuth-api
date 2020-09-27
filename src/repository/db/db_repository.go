package db

import (
	"github.com/gocql/gocql"
	"github.com/pt-abhishek/oAuth-api/src/clients/cassandra"
	"github.com/pt-abhishek/oAuth-api/src/domain/accesstoken"
	"github.com/pt-abhishek/oAuth-api/src/domain/client"
	"github.com/pt-abhishek/oAuth-api/src/utils/errors"
)

//DBRepository is an interface for db functions
type DBRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(*accesstoken.TokenRequest) (*accesstoken.AccessToken, *errors.RestErr)
	// UpdateExpirationTime(accesstoken.AccessToken) *errors.RestErr
}

type dbRepository struct {
	clientRepo ClientRepository
}

var (
	queryGetAccessToken    = "SELECT access_token, client_id, expires FROM access_tokens WHERE access_token=?"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?,?,?,?)"
)

//NewAT returns a new repository
func NewAT(repo ClientRepository) DBRepository {
	return &dbRepository{
		clientRepo: repo,
	}
}

func (db *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {

	var result accesstoken.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.ClientID, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewResourceNotFoundError("No such access token found")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (db *dbRepository) Create(t *accesstoken.TokenRequest) (*accesstoken.AccessToken, *errors.RestErr) {
	var c = &client.Client{
		ClientID:     t.ClientID,
		ClientSecret: t.ClientSecret,
	}
	if err := db.clientRepo.Validate(c); err != nil {
		return nil, err
	}
	token := accesstoken.GetNewAccessToken(t)
	//if token request contains the scope openID Connect then send JWT Too

	return token, nil
}
