package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

type SystemRegisterRequest struct {
	ID     string                      `json:"id"`
	Vendor SystemRegisterRequestVendor `json:"vendor"`

	Name        TranslatedString `json:"name"`
	Description TranslatedString `json:"description"`

	Rights              []SystemRegisterRight         `json:"rights,omitempty"`
	AccessPackages      []SystemRegisterAccessPackage `json:"accessPackages,omitempty"`
	ClientID            []string                      `json:"clientId"`
	AllowedRedirectURLs []string                      `json:"allowedredirecturls"`
	IsVisible           bool                          `json:"isVisible"`
}

type SystemRegisterRequestVendor struct {
	Authority string `json:"authority"`
	ID        string `json:"ID"`
}

type SystemRegisterAccessPackage struct {
	URN string `json:"urn"`
}
type SystemRegisterRight struct {
	Resource []SystemRegisterResource `json:"resource"`
}

type SystemRegisterResource struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type TranslatedString struct {
	Nb string `json:"nb,omitempty"`
	En string `json:"en,omitempty"`
	Nn string `json:"nn,omitempty"`
}

// RegisterSystem registers a system in Altinn. This is required to be able to use the system user functionality.
// The response is the system id, currently not used for anything, but it can be used to check if the system was registered successfully.
func (c *Client) RegisterSystem() (string, error) {
	path := strings.TrimRight(GetAltinnBaseURL(c.signer.environment), "/") + "/authentication/api/v1/systemregister/vendor/" + c.organizationID + "_omniboost"

	body := SystemRegisterRequest{
		ID: c.organizationID + "_omniboost",
		Vendor: SystemRegisterRequestVendor{
			Authority: "iso6523-actorid-upis",
			ID:        "0192:" + c.organizationID,
		},
		Name: TranslatedString{
			Nb: "Omniboost system",
			En: "Omniboost system",
			Nn: "Omniboost system",
		},
		Description: TranslatedString{
			Nb: "Omniboost system",
			En: "Omniboost system",
			Nn: "Omniboost system",
		},
		Rights: nil,
		AccessPackages: []SystemRegisterAccessPackage{
			{
				URN: "urn:altinn:accesspackage:overnatting",
			},
		},
		ClientID:            []string{c.signer.clientID},
		AllowedRedirectURLs: []string{"https://omniboost.io/"},
		IsVisible:           true,
	}

	jsonBody, _ := json.Marshal(body)

	request, err := http.NewRequest("PUT", path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("could not create request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	accessToken, err := c.signer.GetAccessTokenForSystemRegister()
	if err != nil {
		return "", fmt.Errorf("could not get access token: %w", err)
	}

	token, err := c.ExchangeToken(accessToken, "", "")
	if err != nil {
		return "", fmt.Errorf("could not exchange token: %w", err)
	}

	request.Header.Set("Authorization", "Bearer "+token)
	if c.signer.Debug {
		rr, _ := httputil.DumpRequest(request, true)
		log.Printf("%s\n", rr)
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("could not send request: %w", err)
	}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d, body: %s", response.StatusCode, string(bodyBytes))
	}

	return string(bodyBytes), nil
}
