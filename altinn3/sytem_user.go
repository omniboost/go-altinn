package altinn3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"
)

type SystemUserRequest struct {
	ExternalRef    string                           `json:"externalRef,omitempty"`
	SystemID       string                           `json:"systemId"`
	PartyOrgNo     string                           `json:"partyOrgNo"`
	AccessPackages []SystemUserRequestAccessPackage `json:"accessPackages"`
	RedirectURL    string                           `json:"redirectUrl"`
}

type SystemUserRequestAccessPackage struct {
	URN string `json:"urn"`
}

func (c *Client) CreateSystemUser() (string, error) {
	path := strings.TrimRight(GetAltinnBaseURL(c.signer.environment), "/") + "/authentication/api/v1/systemuser/request/vendor/agent"

	body := SystemUserRequest{
		ExternalRef: "omniboost_system_user_01",
		SystemID:    c.organizationID + "_omniboost",
		PartyOrgNo:  c.organizationID,
		AccessPackages: []SystemUserRequestAccessPackage{
			{
				URN: "urn:altinn:accesspackage:ansvarlig-revisor",
			},
		},
		//RedirectURL: "https://omniboost.io/",
	}

	jsonBody, _ := json.Marshal(body)

	request, err := http.NewRequest("POST", path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("could not create request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	accessToken, err := c.signer.GetAccessTokenForSystemUserRequest(nil)
	if err != nil {
		return "", fmt.Errorf("could not get access token: %w", err)
	}

	token, err := c.ExchangeToken(accessToken)
	if err != nil {
		return "", fmt.Errorf("could not exchange token: %w", err)
	}

	request.Header.Set("Authorization", "Bearer "+token)
	if c.signer.Debug {
		rr, _ := httputil.DumpRequest(request, true)
		fmt.Printf("%s\n", rr)
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

func (c *Client) ViewSystemUserRequest(requestID string) (string, error) {
	path := strings.TrimRight(GetAltinnBaseURL(c.signer.environment), "/") + "/authentication/api/v1/systemuser/request/vendor/agent/" + requestID
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return "", fmt.Errorf("could not create request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	accessToken, err := c.signer.GetAccessTokenForSystemUserRequest(nil)
	if err != nil {
		return "", fmt.Errorf("could not get access token: %w", err)
	}

	token, err := c.ExchangeToken(accessToken)
	if err != nil {
		return "", fmt.Errorf("could not exchange token: %w", err)
	}

	request.Header.Set("Authorization", "Bearer "+token)
	if c.signer.Debug {
		rr, _ := httputil.DumpRequest(request, true)
		fmt.Printf("%s\n", rr)
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
