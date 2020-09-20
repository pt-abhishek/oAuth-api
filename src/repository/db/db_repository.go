package db

import (
	"github.com/gocql/gocql"
	"github.com/pt-abhishek/oAuth-api/src/clients/cassandra"
	"github.com/pt-abhishek/oAuth-api/src/domain/accesstoken"
	"github.com/pt-abhishek/oAuth-api/src/utils/errors"
)

//DBRepository is an interface for db functions
type DBRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(accesstoken.AccessToken) *errors.RestErr
	UpdateExpirationTime(accesstoken.AccessToken) *errors.RestErr
}

type dbRepository struct{}

var (
	queryGetAccessToken       = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?"
	queryCreateAccessToken    = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?,?,?,?)"
	queryUpdateExpirationTime = "UPDATE access_tokens SET expires = ? WHERE access_token = ?"
)

//New returns a new repository
func New() DBRepository {
	return &dbRepository{}
}

func (db *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {

	var result accesstoken.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserID, &result.ClientID, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewResourceNotFoundError("No such access token found")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (db *dbRepository) Create(at accesstoken.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserID, at.ClientID, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (db *dbRepository) UpdateExpirationTime(at accesstoken.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryUpdateExpirationTime, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
