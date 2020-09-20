package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token shuould not be expired")
	assert.EqualValues(t, "", at.AccessToken, "New Access token should not have a defined Access token string")
	assert.True(t, at.UserID == 0, "New Access token should not have a defined USER ID")
}

func TestAccessTokenExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "Empty access token should be expired by default")
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "Any token should not expire before 24 hours")
}

func TestConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "wrong expiration time constant")
}
