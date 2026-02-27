package altinn

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-altinn/utils"
)

func (c *Client) NewInstanceDataPost() InstanceDataPost {
	r := InstanceDataPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InstanceDataPost struct {
	client      *Client
	queryParams *InstanceDataPostQueryParams
	pathParams  *InstanceDataPostPathParams
	method      string
	headers     http.Header
	requestBody InstanceDataPostRequestBody
}

func (r InstanceDataPost) NewQueryParams() *InstanceDataPostQueryParams {
	return &InstanceDataPostQueryParams{}
}

type InstanceDataPostQueryParams struct {
	DataType string `schema:"dataType,omitempty"`
}

func (p InstanceDataPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *InstanceDataPost) QueryParams() *InstanceDataPostQueryParams {
	return r.queryParams
}

func (r InstanceDataPost) NewPathParams() *InstanceDataPostPathParams {
	return &InstanceDataPostPathParams{
		AppOrganization: "ssb",
		AppID:           "rs0297-01",
	}
}

type InstanceDataPostPathParams struct {
	AppOrganization string `schema:"organization"`
	AppID           string `schema:"appId"`
	ID              string `schema:"id"`
}

func (p *InstanceDataPostPathParams) Params() map[string]string {
	return map[string]string{
		"app_organization": p.AppOrganization,
		"app_id":           p.AppID,
		"id":               p.ID,
	}
}

func (r *InstanceDataPost) PathParams() *InstanceDataPostPathParams {
	return r.pathParams
}

func (r *InstanceDataPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InstanceDataPost) SetMethod(method string) {
	r.method = method
}

func (r *InstanceDataPost) Method() string {
	return r.method
}

func (r InstanceDataPost) NewRequestBody() InstanceDataPostRequestBody {
	return nil
}

type InstanceDataPostRequestBody any

func (r *InstanceDataPost) RequestBody() *InstanceDataPostRequestBody {
	return &r.requestBody
}

func (r *InstanceDataPost) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *InstanceDataPost) SetRequestBody(body InstanceDataPostRequestBody) {
	r.requestBody = body
}

func (r *InstanceDataPost) NewResponseBody() *InstanceDataPostResponseBody {
	return &InstanceDataPostResponseBody{}
}

type InstanceDataPostResponseBody struct {
	Id              string  `json:"id"`
	InstanceGuid    string  `json:"instanceGuid"`
	DataType        string  `json:"dataType"`
	Filename        *string `json:"filename"`
	ContentType     string  `json:"contentType"`
	BlobStoragePath string  `json:"blobStoragePath"`
	SelfLinks       struct {
		Apps     string `json:"apps"`
		Platform string `json:"platform"`
	} `json:"selfLinks"`
	Size                int           `json:"size"`
	ContentHash         *string       `json:"contentHash"`
	Locked              bool          `json:"locked"`
	Refs                interface{}   `json:"refs"`
	IsRead              bool          `json:"isRead"`
	Tags                []interface{} `json:"tags"`
	UserDefinedMetadata interface{}   `json:"userDefinedMetadata"`
	Metadata            interface{}   `json:"metadata"`
	DeleteStatus        interface{}   `json:"deleteStatus"`
	FileScanResult      string        `json:"fileScanResult"`
	References          interface{}   `json:"references"`
	Created             DateTime      `json:"created"`
	CreatedBy           string        `json:"createdBy"`
	LastChanged         DateTime      `json:"lastChanged"`
	LastChangedBy       string        `json:"lastChangedBy"`
}

func (r *InstanceDataPost) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.app_organization}}/{{.app_id}}/instances/{{.id}}/data", r.PathParams())
	return &u
}

func (r *InstanceDataPost) IsXML() bool {
	return true
}

func (r *InstanceDataPost) Do() (InstanceDataPostResponseBody, error) {
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
