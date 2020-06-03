package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	// if expirationTime != 24 {
	// 	t.Error("expiration time should be 24 hours")
	// }
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken(1)
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	// if at.IsExpired() {
	// 	// GetNewAccessToken() returns not a pointer, so we can't use at == nil
	// 	t.Error("brand new access token should not be expired")
	// }

	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")
	// if at.AccessToken != "" {
	// 	t.Error("new access token should not have defined access token id")
	// }

	assert.True(t, at.UserID == 0, "new access token should not have an associated user id")
	// if at.UserID != 0 {
	// 	t.Error("new access token should not have an associated user id")
	// }
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")
	// if !at.IsExpired() {
	// 	t.Error("empty access token should be expired by default")
	// }

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring 3 hours from now should not be expired")
	// if at.IsExpired() {
	// 	t.Error("access token expiring 3 hours from now should not be expired")
	// }
}
