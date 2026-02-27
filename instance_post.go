package altinn

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-altinn/utils"
)

func (c *Client) NewInstancePost() InstancePost {
	r := InstancePost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InstancePost struct {
	client      *Client
	queryParams *InstancePostQueryParams
	pathParams  *InstancePostPathParams
	method      string
	headers     http.Header
	requestBody InstancePostRequestBody
}

func (r InstancePost) NewQueryParams() *InstancePostQueryParams {
	return &InstancePostQueryParams{}
}

type InstancePostQueryParams struct {
}

func (p InstancePostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *InstancePost) QueryParams() *InstancePostQueryParams {
	return r.queryParams
}

func (r InstancePost) NewPathParams() *InstancePostPathParams {
	return &InstancePostPathParams{
		AppOrganization: "ssb",
		AppID:           "rs0297-01",
	}
}

type InstancePostPathParams struct {
	AppOrganization string `schema:"organization"`
	AppID           string `schema:"appId"`
}

func (p *InstancePostPathParams) Params() map[string]string {
	return map[string]string{
		"app_organization": p.AppOrganization,
		"app_id":           p.AppID,
	}
}

func (r *InstancePost) PathParams() *InstancePostPathParams {
	return r.pathParams
}

func (r *InstancePost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InstancePost) SetMethod(method string) {
	r.method = method
}

func (r *InstancePost) Method() string {
	return r.method
}

func (r InstancePost) NewRequestBody() InstancePostRequestBody {
	return InstancePostRequestBody{}
}

type InstancePostRequestBody struct {
	InstanceOwner struct {
		OrganisationNumber string `json:"organisationNumber,omitempty"`
		PartyID            string `json:"partyId,omitempty"`
		PersonNumber       string `json:"personNumber,omitempty"`
		Username           string `json:"username,omitempty"`
	} `json:"instanceOwner"`
	DueBefore    *DateTime `json:"dueBefore,omitempty"`
	VisibleAfter *DateTime `json:"visibleAfter,omitempty"`
}

func (r *InstancePost) RequestBody() *InstancePostRequestBody {
	return &r.requestBody
}

func (r *InstancePost) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *InstancePost) SetRequestBody(body InstancePostRequestBody) {
	r.requestBody = body
}

func (r *InstancePost) NewResponseBody() *InstancePostResponseBody {
	return &InstancePostResponseBody{}
}

type InstancePostResponseBody struct {
	ID            string `json:"id"`
	InstanceOwner struct {
		PartyID            string `json:"partyId"`
		PersonNumber       string `json:"personNumber"`
		OrganisationNumber string `json:"organisationNumber"`
		Username           string `json:"username"`
		Party              struct {
			PartyID       int    `json:"partyId"`
			PartyUuid     string `json:"partyUuid"`
			PartyTypeName int    `json:"partyTypeName"`
			Ssn           string `json:"ssn"`
			OrgNumber     string `json:"orgNumber"`
			UnitType      string `json:"unitType"`
			Name          string `json:"name"`
			IsDeleted     bool   `json:"isDeleted"`
		} `json:"party"`
	} `json:"instanceOwner"`
	AppID     string `json:"appId"`
	Org       string `json:"org"`
	SelfLinks struct {
		Apps     string `json:"apps"`
		Platform string `json:"platform"`
	} `json:"selfLinks"`
	DueBefore    *DateTime `json:"dueBefore"`
	VisibleAfter DateTime  `json:"visibleAfter"`
	Process      struct {
		Started     DateTime `json:"started"`
		StartEvent  string   `json:"startEvent"`
		CurrentTask struct {
			Flow           int       `json:"flow"`
			Started        DateTime  `json:"started"`
			ElementID      string    `json:"elementId"`
			Name           string    `json:"name"`
			AltinnTaskType string    `json:"altinnTaskType"`
			Ended          *DateTime `json:"ended"`
			Validated      *struct {
				Timestamp       DateTime `json:"timestamp"`
				CanCompleteTask bool     `json:"canCompleteTask"`
			} `json:"validated"`
			FlowType string `json:"flowType"`
		} `json:"currentTask"`
		Ended    interface{} `json:"ended"`
		EndEvent interface{} `json:"endEvent"`
	} `json:"process"`
	Status struct {
		IsArchived    bool      `json:"isArchived"`
		Archived      *DateTime `json:"archived"`
		IsSoftDeleted bool      `json:"isSoftDeleted"`
		SoftDeleted   *DateTime `json:"softDeleted"`
		IsHardDeleted bool      `json:"isHardDeleted"`
		HardDeleted   *DateTime `json:"hardDeleted"`
		ReadStatus    int       `json:"readStatus"`
		Substatus     *struct {
			Label       string `json:"label"`
			Description string `json:"description"`
		} `json:"substatus"`
	} `json:"status"`
	CompleteConfirmations []struct {
		StakeholderID string    `json:"stakeholderId"`
		ConfirmedOn   *DateTime `json:"confirmedOn"`
	} `json:"completeConfirmations"`
	Data              []interface{} `json:"data"`
	PresentationTexts interface{}   `json:"presentationTexts"`
	DataValues        interface{}   `json:"dataValues"`
	Created           DateTime      `json:"created"`
	CreatedBy         string        `json:"createdBy"`
	LastChanged       DateTime      `json:"lastChanged"`
	LastChangedBy     string        `json:"lastChangedBy"`
}

func (r *InstancePost) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.app_organization}}/{{.app_id}}/instances", r.PathParams())
	return &u
}

func (r InstancePost) IsXML() bool {
	return false
}

func (r *InstancePost) Do() (InstancePostResponseBody, error) {
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
