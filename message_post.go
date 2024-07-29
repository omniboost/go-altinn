package altinn

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-altinn/utils"
)

func (c *Client) NewMessagePost() MessagePost {
	r := MessagePost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MessagePost struct {
	client      *Client
	queryParams *MessagePostQueryParams
	pathParams  *MessagePostPathParams
	method      string
	headers     http.Header
	requestBody MessagePostRequestBody
}

func (r MessagePost) NewQueryParams() *MessagePostQueryParams {
	return &MessagePostQueryParams{
		Complete: true,
		Sign:     true,
	}
}

type MessagePostQueryParams struct {
	Language string `schema:"language,omitempty"`
	Complete bool   `schema:"complete"`
	Sign     bool   `schema:"sign"`
}

func (p MessagePostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MessagePost) QueryParams() *MessagePostQueryParams {
	return r.queryParams
}

func (r MessagePost) NewPathParams() *MessagePostPathParams {
	return &MessagePostPathParams{
		OrganizationNumber: "my",
	}
}

type MessagePostPathParams struct {
	OrganizationNumber string `schema:"organization_number"`
}

func (p *MessagePostPathParams) Params() map[string]string {
	return map[string]string{
		"organization_number": p.OrganizationNumber,
	}
}

func (r *MessagePost) PathParams() *MessagePostPathParams {
	return r.pathParams
}

func (r *MessagePost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MessagePost) SetMethod(method string) {
	r.method = method
}

func (r *MessagePost) Method() string {
	return r.method
}

func (r MessagePost) NewRequestBody() MessagePostRequestBody {
	return MessagePostRequestBody{}
}

type MessagePostRequestBody struct {
	Type           string `json:"Type"`
	ServiceCode    string `json:"ServiceCode"`
	ServiceEdition string `json:"ServiceEdition"`
	Embedded       struct {
		Forms Forms `json:"forms"`
	} `json:"_embedded"`
}

func (b MessagePostRequestBody) MarshalJSON() ([]byte, error) {
	a := struct {
		Type           string `json:"Type"`
		ServiceCode    string `json:"ServiceCode"`
		ServiceEdition string `json:"ServiceEdition"`
		Embedded       struct {
			Forms []struct {
				Type              string `json:"Type"`
				DataFormatId      string `json:"DataFormatId"`
				DataFormatVersion string `json:"DataFormatVersion"`
				FormData          string `json:"FormData"`
			} `json:"forms"`
		} `json:"_embedded"`
	}{
		Type:           b.Type,
		ServiceCode:    b.ServiceCode,
		ServiceEdition: b.ServiceEdition,
	}
	for _, f := range b.Embedded.Forms {
		// convert form to xml
		rawXML, err := xml.Marshal(f.FormData)
		if err != nil {
			return nil, err
		}

		// buf := new(bytes.Buffer)
		// enc := json.NewEncoder(buf)
		// enc.SetEscapeHTML(false)
		// err = enc.Encode(string(rawXML))
		// if err != nil {
		// 	return nil, err
		// }

		a.Embedded.Forms = append(a.Embedded.Forms, struct {
			Type              string `json:"Type"`
			DataFormatId      string `json:"DataFormatId"`
			DataFormatVersion string `json:"DataFormatVersion"`
			FormData          string `json:"FormData"`
		}{
			Type:              f.Type,
			DataFormatId:      f.DataFormatId,
			DataFormatVersion: f.DataFormatVersion,
			FormData:          string(rawXML),
		})
	}
	return json.Marshal(a)
}

func (r *MessagePost) RequestBody() *MessagePostRequestBody {
	return &r.requestBody
}

func (r *MessagePost) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *MessagePost) SetRequestBody(body MessagePostRequestBody) {
	r.requestBody = body
}

func (r *MessagePost) NewResponseBody() *MessagePostResponseBody {
	return &MessagePostResponseBody{}
}

type MessagePostResponseBody []struct {
	Name               string `json:"Name"`
	Type               string `json:"Type"`
	OrganizationNumber string `json:"OrganizationNumber"`
	Status             string `json:"Status"`
}

func (r *MessagePost) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.organization_number}}/messages", r.PathParams())
	return &u
}

func (r *MessagePost) Do() (MessagePostResponseBody, error) {
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
