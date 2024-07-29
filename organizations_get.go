package altinn

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-altinn/utils"
)

func (c *Client) NewOrganizationsGet() OrganizationsGet {
	r := OrganizationsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type OrganizationsGet struct {
	client      *Client
	queryParams *OrganizationsGetQueryParams
	pathParams  *OrganizationsGetPathParams
	method      string
	headers     http.Header
	requestBody OrganizationsGetBody
}

func (r OrganizationsGet) NewQueryParams() *OrganizationsGetQueryParams {
	return &OrganizationsGetQueryParams{}
}

type OrganizationsGetQueryParams struct {
	QueryOptions             string `schema:"queryOptions"`
	IncludeInactiveReportees bool   `schema:"includeInactiveReportees"`
}

func (p OrganizationsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *OrganizationsGet) QueryParams() *OrganizationsGetQueryParams {
	return r.queryParams
}

func (r OrganizationsGet) NewPathParams() *OrganizationsGetPathParams {
	return &OrganizationsGetPathParams{}
}

type OrganizationsGetPathParams struct {
}

func (p *OrganizationsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *OrganizationsGet) PathParams() *OrganizationsGetPathParams {
	return r.pathParams
}

func (r *OrganizationsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *OrganizationsGet) SetMethod(method string) {
	r.method = method
}

func (r *OrganizationsGet) Method() string {
	return r.method
}

func (r OrganizationsGet) NewRequestBody() OrganizationsGetBody {
	return OrganizationsGetBody{}
}

type OrganizationsGetBody struct {
}

func (r *OrganizationsGet) RequestBody() *OrganizationsGetBody {
	return nil
}

func (r *OrganizationsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *OrganizationsGet) SetRequestBody(body OrganizationsGetBody) {
	r.requestBody = body
}

func (r *OrganizationsGet) NewResponseBody() *OrganizationsGetResponseBody {
	return &OrganizationsGetResponseBody{}
}

type OrganizationsGetResponseBody []struct {
	Name               string `json:"Name"`
	Type               string `json:"Type"`
	OrganizationNumber string `json:"OrganizationNumber"`
	Status             string `json:"Status"`
}

func (r *OrganizationsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("organizations", r.PathParams())
	return &u
}

func (r *OrganizationsGet) Do() (OrganizationsGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	err = r.client.Authenticate(req)
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
