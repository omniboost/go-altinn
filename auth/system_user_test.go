package auth_test

import (
	"os"
	"testing"
)

func TestCreateSystemUser(t *testing.T) {
	token, err := client.CreateSystemUser()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("System user token: %s", token)
}

func TestViewSystemUserRequest(t *testing.T) {
	systemUserID := os.Getenv("SYSTEM_USER_ID")
	if systemUserID == "" {
		t.Fatal("SYSTEM_USER_ID environment variable is not set")
	}
	token, err := client.ViewSystemUserRequest(systemUserID)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("System user token: %s", token)
}
