package altinn_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestAuthenticateWithPassword(t *testing.T) {
	req := client.NewAuthenticateWithPassword()
	req.RequestBody().UserName = os.Getenv("USER_NAME")
	req.RequestBody().UserPassword = os.Getenv("USER_PASSWORD")
	req.QueryParams().ForceEIAuthentication = false
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
