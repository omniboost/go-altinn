package altinn_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestOrganizationsGet(t *testing.T) {
	req := client.NewOrganizationsGet()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
