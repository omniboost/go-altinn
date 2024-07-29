package altinn_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMetadataFormtaskGet(t *testing.T) {
	req := client.NewMetadataFormtaskGet()
	req.PathParams().ServiceCode = 5492
	req.PathParams().ServiceEditionCode = 1
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))

}
