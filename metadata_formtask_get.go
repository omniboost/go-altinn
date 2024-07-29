package altinn

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-altinn/utils"
)

func (c *Client) NewMetadataFormtaskGet() MetadataFormtaskGet {
	r := MetadataFormtaskGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MetadataFormtaskGet struct {
	client      *Client
	queryParams *MetadataFormtaskGetQueryParams
	pathParams  *MetadataFormtaskGetPathParams
	method      string
	headers     http.Header
	requestBody MetadataFormtaskGetBody
}

func (r MetadataFormtaskGet) NewQueryParams() *MetadataFormtaskGetQueryParams {
	return &MetadataFormtaskGetQueryParams{}
}

type MetadataFormtaskGetQueryParams struct {
	Language int `schema:"language,omitempty"`
}

func (p MetadataFormtaskGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MetadataFormtaskGet) QueryParams() *MetadataFormtaskGetQueryParams {
	return r.queryParams
}

func (r MetadataFormtaskGet) NewPathParams() *MetadataFormtaskGetPathParams {
	return &MetadataFormtaskGetPathParams{}
}

type MetadataFormtaskGetPathParams struct {
	ServiceCode        int `schema:"service_code"`
	ServiceEditionCode int `schema:"service_edition_code"`
}

func (p *MetadataFormtaskGetPathParams) Params() map[string]string {
	return map[string]string{
		"service_code":         strconv.Itoa(p.ServiceCode),
		"service_edition_code": strconv.Itoa(p.ServiceEditionCode),
	}
}

func (r *MetadataFormtaskGet) PathParams() *MetadataFormtaskGetPathParams {
	return r.pathParams
}

func (r *MetadataFormtaskGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MetadataFormtaskGet) SetMethod(method string) {
	r.method = method
}

func (r *MetadataFormtaskGet) Method() string {
	return r.method
}

func (r MetadataFormtaskGet) NewRequestBody() MetadataFormtaskGetBody {
	return MetadataFormtaskGetBody{}
}

type MetadataFormtaskGetBody struct {
}

func (r *MetadataFormtaskGet) RequestBody() *MetadataFormtaskGetBody {
	return nil
}

func (r *MetadataFormtaskGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *MetadataFormtaskGet) SetRequestBody(body MetadataFormtaskGetBody) {
	r.requestBody = body
}

func (r *MetadataFormtaskGet) NewResponseBody() *MetadataFormtaskGetResponseBody {
	return &MetadataFormtaskGetResponseBody{}
}

type MetadataFormtaskGetResponseBody struct {
	ServiceOwnerCode   string `json:"ServiceOwnerCode"`
	ServiceOwnerName   string `json:"ServiceOwnerName"`
	ServiceName        string `json:"ServiceName"`
	ServiceCode        string `json:"ServiceCode"`
	ServiceEditionName string `json:"ServiceEditionName"`
	ServiceEditionCode int    `json:"ServiceEditionCode"`
	ValidFrom          string `json:"ValidFrom"`
	ValidTo            string `json:"ValidTo"`
	ServiceType        string `json:"ServiceType"`
	RestEnabled        bool   `json:"RestEnabled"`
	FormsMetaData      []struct {
		FormID                 int    `json:"FormID"`
		FormName               string `json:"FormName"`
		DataFormatProviderType string `json:"DataFormatProviderType"`
		DataFormatID           string `json:"DataFormatID"`
		DataFormatVersion      int    `json:"DataFormatVersion"`
		IsOnlyXsdValidation    bool   `json:"IsOnlyXsdValidation"`
		FormType               string `json:"FormType"`
	} `json:"FormsMetaData"`
	EUSEnabled            bool `json:"EUSEnabled"`
	EnterpriseUserEnabled bool `json:"EnterpriseUserEnabled"`
	ProcessSteps          []struct {
		SequenceNumber int    `json:"SequenceNumber"`
		Name           string `json:"Name"`
		SecurityLevel  int    `json:"SecurityLevel"`
	} `json:"ProcessSteps"`
}

func (r *MetadataFormtaskGet) URL() *url.URL {
	u := r.client.GetEndpointURL("metadata/formtask/{{.service_code}}/{{.service_edition_code}}", r.PathParams())
	return &u
}

func (r *MetadataFormtaskGet) Do() (MetadataFormtaskGetResponseBody, error) {
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
