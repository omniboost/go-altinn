package altinn

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-altinn/utils"
)

func (c *Client) NewAuthenticateWithPassword() AuthenticateWithPassword {
	r := AuthenticateWithPassword{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AuthenticateWithPassword struct {
	client      *Client
	queryParams *AuthenticateWithPasswordQueryParams
	pathParams  *AuthenticateWithPasswordPathParams
	method      string
	headers     http.Header
	requestBody AuthenticateWithPasswordBody
}

func (r AuthenticateWithPassword) NewQueryParams() *AuthenticateWithPasswordQueryParams {
	return &AuthenticateWithPasswordQueryParams{}
}

type AuthenticateWithPasswordQueryParams struct {
	ForceEIAuthentication bool `schema:"ForceEIAuthentication"`
}

func (p AuthenticateWithPasswordQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AuthenticateWithPassword) QueryParams() *AuthenticateWithPasswordQueryParams {
	return r.queryParams
}

func (r AuthenticateWithPassword) NewPathParams() *AuthenticateWithPasswordPathParams {
	return &AuthenticateWithPasswordPathParams{}
}

type AuthenticateWithPasswordPathParams struct {
}

func (p *AuthenticateWithPasswordPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AuthenticateWithPassword) PathParams() *AuthenticateWithPasswordPathParams {
	return r.pathParams
}

func (r *AuthenticateWithPassword) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AuthenticateWithPassword) SetMethod(method string) {
	r.method = method
}

func (r *AuthenticateWithPassword) Method() string {
	return r.method
}

func (r AuthenticateWithPassword) NewRequestBody() AuthenticateWithPasswordBody {
	return AuthenticateWithPasswordBody{}
}

type AuthenticateWithPasswordBody struct {
	UserName     string `json:"UserName"`
	UserPassword string `json:"UserPassword"`
}

func (r *AuthenticateWithPassword) RequestBody() *AuthenticateWithPasswordBody {
	return &r.requestBody
}

func (r *AuthenticateWithPassword) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *AuthenticateWithPassword) SetRequestBody(body AuthenticateWithPasswordBody) {
	r.requestBody = body
}

func (r *AuthenticateWithPassword) NewResponseBody() *AuthenticateWithPasswordResponseBody {
	return &AuthenticateWithPasswordResponseBody{}
}

type AuthenticateWithPasswordResponseBody struct{}

func (r *AuthenticateWithPassword) URL() *url.URL {
	u := r.client.GetEndpointURL("authentication/authenticatewithpassword", r.PathParams())
	return &u
}

func (r *AuthenticateWithPassword) Do() (AuthenticateWithPasswordResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
