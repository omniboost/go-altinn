package altinn3_test

import "testing"

func TestCreateSystemUser(t *testing.T) {
	token, err := client.CreateSystemUser()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("System user token: %s", token)
}

func TestViewSystemUserRequest(t *testing.T) {
	token, err := client.ViewSystemUserRequest("f10b4724-8b1c-4e3d-84e1-e3bd4861f3b0")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("System user token: %s", token)
}
