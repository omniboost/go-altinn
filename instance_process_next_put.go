package altinn

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-altinn/utils"
)

func (c *Client) NewInstanceProcessNextPut() InstanceProcessNextPut {
	r := InstanceProcessNextPut{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InstanceProcessNextPut struct {
	client      *Client
	queryParams *InstanceProcessNextPutQueryParams
	pathParams  *InstanceProcessNextPutPathParams
	method      string
	headers     http.Header
	requestBody InstanceProcessNextPutRequestBody
}

func (r InstanceProcessNextPut) NewQueryParams() *InstanceProcessNextPutQueryParams {
	return &InstanceProcessNextPutQueryParams{}
}

type InstanceProcessNextPutQueryParams struct {
}

func (p InstanceProcessNextPutQueryParams) ToURLValues() (url.Values, error) {
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

func (r *InstanceProcessNextPut) QueryParams() *InstanceProcessNextPutQueryParams {
	return r.queryParams
}

func (r InstanceProcessNextPut) NewPathParams() *InstanceProcessNextPutPathParams {
	return &InstanceProcessNextPutPathParams{
		AppOrganization: "ssb",
		AppID:           "rs0297-01",
	}
}

type InstanceProcessNextPutPathParams struct {
	AppOrganization string `schema:"organization"`
	AppID           string `schema:"appId"`
	ID              string `schema:"id"`
}

func (p *InstanceProcessNextPutPathParams) Params() map[string]string {
	return map[string]string{
		"app_organization": p.AppOrganization,
		"app_id":           p.AppID,
		"id":               p.ID,
	}
}

func (r *InstanceProcessNextPut) PathParams() *InstanceProcessNextPutPathParams {
	return r.pathParams
}

func (r *InstanceProcessNextPut) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InstanceProcessNextPut) SetMethod(method string) {
	r.method = method
}

func (r *InstanceProcessNextPut) Method() string {
	return r.method
}

func (r InstanceProcessNextPut) NewRequestBody() InstanceProcessNextPutRequestBody {
	return InstanceProcessNextPutRequestBody{}
}

type InstanceProcessNextPutRequestBody struct {
}

func (r *InstanceProcessNextPut) RequestBody() *InstanceProcessNextPutRequestBody {
	return &r.requestBody
}

func (r *InstanceProcessNextPut) RequestBodyInterface() interface{} {
	return nil
}

func (r *InstanceProcessNextPut) SetRequestBody(body InstanceProcessNextPutRequestBody) {
	r.requestBody = body
}

func (r *InstanceProcessNextPut) NewResponseBody() *InstanceProcessNextPutResponseBody {
	return &InstanceProcessNextPutResponseBody{}
}

type InstanceProcessNextPutResponseBody struct {
	CurrentTask  interface{} `json:"currentTask"`
	ProcessTasks []struct {
		AltinnTaskType string `json:"altinnTaskType"`
		ElementId      string `json:"elementId"`
	} `json:"processTasks"`
	Started    DateTime `json:"started"`
	StartEvent string   `json:"startEvent"`
	Ended      DateTime `json:"ended"`
	EndEvent   string   `json:"endEvent"`
}

func (r *InstanceProcessNextPut) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.app_organization}}/{{.app_id}}/instances/{{.id}}/process/next", r.PathParams())
	return &u
}

func (r InstanceProcessNextPut) IsXML() bool {
	return false
}

func (r *InstanceProcessNextPut) Do() (InstanceProcessNextPutResponseBody, error) {
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
