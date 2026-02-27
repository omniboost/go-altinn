package auth_test

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

func TestGetAccessTokenForSystemUserRequest(t *testing.T) {
	token, err := client.GetSigner().GetAccessTokenForSystemUserRequest(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("token: %s", token.AccessToken)
	t.Logf("token type: %s", token.TokenType)
	t.Logf("expires in: %d", token.ExpiresIn)
	t.Logf("token expires: %s", token.Expiry.Format(time.RFC3339))
	t.Logf("scope: %s", token.Scope)

	t2, err := client.ExchangeToken(token, "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("token: %s", t2)
}

func TestGetAccessTokenForUserInstance(t *testing.T) {
	token, err := client.GetSigner().GetAccessTokenForUserInstance(client.GetOrganizationID())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("token: %s", token.AccessToken)
	t.Logf("token type: %s", token.TokenType)
	t.Logf("expires in: %d", token.ExpiresIn)
	t.Logf("token expires: %s", token.Expiry.Format(time.RFC3339))
	t.Logf("scope: %s", token.Scope)

	t2, err := client.ExchangeToken(token, "HEIMRCollection", "SSBreporting123!")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("token: %s", t2)
}
