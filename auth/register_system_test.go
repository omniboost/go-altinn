package auth_test

import "testing"

func TestRegisterSystem(t *testing.T) {
	systemID, err := client.RegisterSystem()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("SystemID: %s", systemID)
}
