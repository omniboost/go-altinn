package auth

import "net/http"

type Client struct {
	signer         *JWTSigner
	organizationID string
	httpClient     *http.Client
}

func NewClient(httpClient *http.Client, key []byte, keyID, environment, clientId string) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	signer, err := NewJWTSigner(key, keyID, environment, clientId)
	if err != nil {
		return nil, err
	}
	signer.SetHTTPClient(httpClient)
	return &Client{
		signer:     signer,
		httpClient: httpClient,
	}, nil
}

func (c *Client) GetOrganizationID() string {
	return c.organizationID
}

func (c *Client) SetOrganizationID(organizationID string) {
	c.organizationID = organizationID
}

func (c *Client) SetDebug(debug bool) {
	c.signer.SetDebug(debug)
}

func (c *Client) SetHTTPClient(httpClient *http.Client) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c.httpClient = httpClient
	c.signer.SetHTTPClient(httpClient)
}

func (c *Client) GetSigner() *JWTSigner {
	return c.signer
}
