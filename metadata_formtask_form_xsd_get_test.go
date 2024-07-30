package altinn_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMetadataFormtaskFormXSDGet(t *testing.T) {
	req := client.NewMetadataFormtaskFormXSDGet()
	req.PathParams().ServiceCode = 5492
	req.PathParams().ServiceEditionCode = 1
	req.PathParams().DataFormatID = 6400
	req.PathParams().DataFormatVersion = 45188
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))

}
