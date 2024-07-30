package altinn

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-altinn/utils"
)

func (c *Client) NewMetadataFormtaskFormXSDGet() MetadataFormtaskFormXSDGet {
	r := MetadataFormtaskFormXSDGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MetadataFormtaskFormXSDGet struct {
	client      *Client
	queryParams *MetadataFormtaskFormXSDGetQueryParams
	pathParams  *MetadataFormtaskFormXSDGetPathParams
	method      string
	headers     http.Header
	requestBody MetadataFormtaskFormXSDGetBody
}

func (r MetadataFormtaskFormXSDGet) NewQueryParams() *MetadataFormtaskFormXSDGetQueryParams {
	return &MetadataFormtaskFormXSDGetQueryParams{}
}

type MetadataFormtaskFormXSDGetQueryParams struct {
	Language int `schema:"language,omitempty"`
}

func (p MetadataFormtaskFormXSDGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MetadataFormtaskFormXSDGet) QueryParams() *MetadataFormtaskFormXSDGetQueryParams {
	return r.queryParams
}

func (r MetadataFormtaskFormXSDGet) NewPathParams() *MetadataFormtaskFormXSDGetPathParams {
	return &MetadataFormtaskFormXSDGetPathParams{}
}

type MetadataFormtaskFormXSDGetPathParams struct {
	ServiceCode        int `schema:"service_code"`
	ServiceEditionCode int `schema:"service_edition_code"`
	DataFormatID       int `schema:"data_format_id"`
	DataFormatVersion  int `schema:"data_format_version"`
}

func (p *MetadataFormtaskFormXSDGetPathParams) Params() map[string]string {
	return map[string]string{
		"service_code":         strconv.Itoa(p.ServiceCode),
		"service_edition_code": strconv.Itoa(p.ServiceEditionCode),
		"data_format_id":       strconv.Itoa(p.DataFormatID),
		"data_format_version":  strconv.Itoa(p.DataFormatVersion),
	}
}

func (r *MetadataFormtaskFormXSDGet) PathParams() *MetadataFormtaskFormXSDGetPathParams {
	return r.pathParams
}

func (r *MetadataFormtaskFormXSDGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MetadataFormtaskFormXSDGet) SetMethod(method string) {
	r.method = method
}

func (r *MetadataFormtaskFormXSDGet) Method() string {
	return r.method
}

func (r MetadataFormtaskFormXSDGet) NewRequestBody() MetadataFormtaskFormXSDGetBody {
	return MetadataFormtaskFormXSDGetBody{}
}

type MetadataFormtaskFormXSDGetBody struct {
}

func (r *MetadataFormtaskFormXSDGet) RequestBody() *MetadataFormtaskFormXSDGetBody {
	return nil
}

func (r *MetadataFormtaskFormXSDGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *MetadataFormtaskFormXSDGet) SetRequestBody(body MetadataFormtaskFormXSDGetBody) {
	r.requestBody = body
}

func (r *MetadataFormtaskFormXSDGet) NewResponseBody() *MetadataFormtaskFormXSDGetResponseBody {
	return &MetadataFormtaskFormXSDGetResponseBody{}
}

type MetadataFormtaskFormXSDGetResponseBody struct {
}

func (r *MetadataFormtaskFormXSDGet) URL() *url.URL {
	u := r.client.GetEndpointURL("metadata/formtask/{{.service_code}}/{{.service_edition_code}}/forms/{{.data_format_id}}/{{.data_format_version}}/xsd", r.PathParams())
	return &u
}

func (r *MetadataFormtaskFormXSDGet) Do() (MetadataFormtaskFormXSDGetResponseBody, error) {
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
