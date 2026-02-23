package altinn3_test

import (
	"log"
	"os"
	"testing"

	"github.com/omniboost/go-altinn/altinn3"
)

var signer *altinn3.JWTSigner

func TestMain(m *testing.M) {
	privateKey := os.Getenv("PRIVATE_KEY")
	keyID := os.Getenv("KEY_ID")
	clientID := os.Getenv("CLIENT_ID")
	environment := os.Getenv("ENVIRONMENT")
	debug := os.Getenv("DEBUG")

	var err error
	signer, err = altinn3.NewJWTSigner(
		[]byte(privateKey),
		keyID,
		environment,
		clientID,
	)
	if err != nil {
		log.Fatal(err)
	}

	if debug != "" {
		signer.Debug = true
	}

	m.Run()
}
