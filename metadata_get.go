package altinn

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-altinn/utils"
)

func (c *Client) NewMetadataGet() MetadataGet {
	r := MetadataGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MetadataGet struct {
	client      *Client
	queryParams *MetadataGetQueryParams
	pathParams  *MetadataGetPathParams
	method      string
	headers     http.Header
	requestBody MetadataGetBody
}

func (r MetadataGet) NewQueryParams() *MetadataGetQueryParams {
	return &MetadataGetQueryParams{}
}

type MetadataGetQueryParams struct {
	QueryOptions             string `schema:"queryOptions"`
	IncludeInactiveReportees bool   `schema:"includeInactiveReportees"`
}

func (p MetadataGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MetadataGet) QueryParams() *MetadataGetQueryParams {
	return r.queryParams
}

func (r MetadataGet) NewPathParams() *MetadataGetPathParams {
	return &MetadataGetPathParams{}
}

type MetadataGetPathParams struct {
}

func (p *MetadataGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *MetadataGet) PathParams() *MetadataGetPathParams {
	return r.pathParams
}

func (r *MetadataGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MetadataGet) SetMethod(method string) {
	r.method = method
}

func (r *MetadataGet) Method() string {
	return r.method
}

func (r MetadataGet) NewRequestBody() MetadataGetBody {
	return MetadataGetBody{}
}

type MetadataGetBody struct {
}

func (r *MetadataGet) RequestBody() *MetadataGetBody {
	return nil
}

func (r *MetadataGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *MetadataGet) SetRequestBody(body MetadataGetBody) {
	r.requestBody = body
}

func (r *MetadataGet) NewResponseBody() *MetadataGetResponseBody {
	return &MetadataGetResponseBody{}
}

type MetadataGetResponseBody []struct {
	ServiceOwnerCode      string `json:"ServiceOwnerCode"`
	ServiceOwnerName      string `json:"ServiceOwnerName"`
	ServiceName           string `json:"ServiceName"`
	ServiceCode           string `json:"ServiceCode"`
	ServiceEditionName    string `json:"ServiceEditionName"`
	AltinnAppID           string `json:"AltinnAppId,omitempty"`
	ServiceEditionCode    int    `json:"ServiceEditionCode"`
	ValidFrom             string `json:"ValidFrom"`
	ValidTo               string `json:"ValidTo"`
	ServiceType           string `json:"ServiceType"`
	EnterpriseUserEnabled bool   `json:"EnterpriseUserEnabled,omitempty"`
	ConfidentialService   bool   `json:"ConfidentialService,omitempty"`
}

func (r *MetadataGet) URL() *url.URL {
	u := r.client.GetEndpointURL("metadata", r.PathParams())
	return &u
}

func (r *MetadataGet) Do() (MetadataGetResponseBody, error) {
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
