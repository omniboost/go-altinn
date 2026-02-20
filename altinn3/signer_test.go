package altinn3

import (
	"os"
	"testing"
	"time"
)

func TestGetAcccessTokenForSystemRegister(t *testing.T) {
	privateKey := os.Getenv("PRIVATE_KEY")
	keyID := os.Getenv("KEY_ID")
	clientID := os.Getenv("CLIENT_ID")
	environment := os.Getenv("ENVIRONMENT")

	a, err := NewJWTSigner(
		[]byte(privateKey),
		keyID,
		environment,
		clientID,
	)
	if err != nil {
		t.Fatal(err)
	}

	a.Debug = false

	token, err := a.GetAcccessTokenForSystemRegister()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("token: %s", token.AccessToken)
	t.Logf("token type: %s", token.TokenType)
	t.Logf("expires in: %d", token.ExpiresIn)
	t.Logf("token expires: %s", token.Expiry.Format(time.RFC3339))
	t.Logf("scope: %s", token.Scope)
}
