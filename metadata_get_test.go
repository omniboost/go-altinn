package altinn_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMetadataGet(t *testing.T) {
	req := client.NewMetadataGet()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))

}
