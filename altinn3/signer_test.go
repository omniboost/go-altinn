package altinn3_test

import (
	"testing"
	"time"
)

func TestGetAccessTokenForSystemRegister(t *testing.T) {
	token, err := client.GetSigner().GetAccessTokenForSystemRegister()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("token: %s", token.AccessToken)
	t.Logf("token type: %s", token.TokenType)
	t.Logf("expires in: %d", token.ExpiresIn)
	t.Logf("token expires: %s", token.Expiry.Format(time.RFC3339))
	t.Logf("scope: %s", token.Scope)
}
