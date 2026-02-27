package auth_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/omniboost/go-altinn/auth"
)

var client *auth.Client

func TestMain(m *testing.M) {
	privateKey := os.Getenv("PRIVATE_KEY")
	keyID := os.Getenv("KEY_ID")
	clientID := os.Getenv("CLIENT_ID")
	environment := os.Getenv("ENVIRONMENT")
	debug := os.Getenv("DEBUG")
	organizaionID := os.Getenv("ORGANIZATION_ID")

	var err error
	client, err = auth.NewClient(
		http.DefaultClient,
		[]byte(privateKey),
		keyID,
		environment,
		clientID,
	)
	client.SetOrganizationID(organizaionID)
	if err != nil {
		log.Fatal(err)
	}

	if debug != "" {
		client.SetDebug(true)
	}

	m.Run()
}
