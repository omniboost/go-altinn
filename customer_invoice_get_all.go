package vismanet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-visma.net/utils"
)

func (c *Client) NewCustomerInvoiceGetAll() CustomerInvoiceGetAll {
	r := CustomerInvoiceGetAll{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerInvoiceGetAll struct {
	client      *Client
	queryParams *CustomerInvoiceGetAllQueryParams
	pathParams  *CustomerInvoiceGetAllPathParams
	method      string
	headers     http.Header
	requestBody CustomerInvoiceGetAllBody
}

func (r CustomerInvoiceGetAll) NewQueryParams() *CustomerInvoiceGetAllQueryParams {
	return &CustomerInvoiceGetAllQueryParams{}
}

type CustomerInvoiceGetAllQueryParams struct {
	// The field is deprecated for specific customer document endpoints. It will
	// only be usable from customer document endpoint.
	DocumentType string `schema:"documentType,omitempty"`

	// Parameter for showing if invoice has been released or not.
	Released int `schema:"released,omitempty"`

	// The dunning level of the document
	DunningLevel int `schema:"dunningLevel,omitempty"`

	// The date of the closing of the financial period.
	ClosedFinancialPeriod string `schema:"closedFinancialPeriod,omitempty"`

	// The date and time for when the document last released a dunning letter.
	DunningLetterDateTime string `schema:"dunningLetterDateTime,omitempty"`

	// Set time/date as before (<), after (>), before and including (=<) OR
	// after and including (=>) to filter on time frame.
	DunningLetterDateTimeCondition string `schema:"dunningLetterDateTimeCondition,omitempty"`

	// The project with which the document is associated.
	Project string `schema:"project,omitempty"`

	// True if you want to see all dunning information regarding this document.
	ExpandApplications bool `schema:"expandApplications,omitempty"`

	ExpandDunningInformation bool `schema:"expandDunningInformation,omitempty"`

	// True if you want to see all attachments regarding this document.
	ExpandAttachments bool `schema:"expandAttachments,omitempty"`

	// True if you want to see all VAT details regarding this document.
	ExpandTaxDetails bool `schema:"expandTaxDetails,omitempty"`

	// True if you want to see all information regarding the invoice address for
	// this document.
	ExpandInvoiceAddress bool `schema:"expandInvoiceAddress,omitempty"`

	// The financial period to which the transactions recorded in the document
	// is posted. Format YYYYMM.
	FinancialPeriod string `schema:"financialPeriod,omitempty"`

	// The date when payment for the document is due, in accordance with the
	// credit terms.
	DocumentDueDate string `schema:"documentDueDate,omitempty"`

	// The status of the document. Use the dropdown to select status.
	Status string `schema:"status,omitempty"`

	// This field has been deprecated and will be removed in future versions.
	// Use pagenumber and pagesize for pagination purposes. Pagenumber and
	// pagesize does not work with NumberToRead and SkipRecords.
	NumberToRead int `schema:"numberToRead,omitempty"`

	// This field has been deprecated and will be removed in future versions.
	// Use pagenumber and pagesize for pagination purposes. Pagenumber and
	// pagesize does not work with NumberToRead and SkipRecords.
	SkipRecords int `schema:"skipRecords,omitempty"`

	// The top part > External reference > The external reference used in
	// AutoInvoice.
	ExternalReference string `schema:"externalReference,omitempty"`

	// The top part > Payment ref. > The reference number of the document, as
	// automatically generated by the system in accordance with the number
	// series assigned to cash sales in the Customer ledger preferences window…
	PaymentReference string `schema:"paymentReference,omitempty"`

	// The top part > External reference > The external reference used in
	// AutoInvoice.
	CustomerRefNumber string `schema:"customerRefNumber,omitempty"`

	// Greater than value. The item which is the object for this, varies from
	// API to API.
	GreaterThanValue string `schema:"greaterThanValue,omitempty"`

	// System generated value for last modification of transaction/record. Use
	// format: YYYY-MM-DD HH:MM (date and time) to filter from date to present.
	LastModifiedDateTime string `schema:"lastModifiedDateTime,omitempty"`

	// System retrieved information for state/condition.
	LastModifiedDateTimeCondition string `schema:"lastModifiedDateTimeCondition,omitempty"`

	// Creation date and time.
	CreatedDateTime string `schema:"createdDateTime,omitempty"`

	// System-retrieved information for state/condition
	CreatedDateTimeCondition string `schema:"createdDateTimeCondition,omitempty"`

	// Pagination parameter. Page number.
	PageNumber int `schema:"pageNumber,omitempty"`

	// Pagination parameter. Number of items to be collected. Please use a page
	// size lower or equal to the allowed max page size which is returned as
	// part of the metadata information. If requested page size is greater than
	// allowed max page size, request will be limited to max page size.
	PageSize int `schema:"pageSize,omitempty"`
}

func (p CustomerInvoiceGetAllQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerInvoiceGetAll) QueryParams() *CustomerInvoiceGetAllQueryParams {
	return r.queryParams
}

func (r CustomerInvoiceGetAll) NewPathParams() *CustomerInvoiceGetAllPathParams {
	return &CustomerInvoiceGetAllPathParams{}
}

type CustomerInvoiceGetAllPathParams struct {
}

func (p *CustomerInvoiceGetAllPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomerInvoiceGetAll) PathParams() *CustomerInvoiceGetAllPathParams {
	return r.pathParams
}

func (r *CustomerInvoiceGetAll) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerInvoiceGetAll) SetMethod(method string) {
	r.method = method
}

func (r *CustomerInvoiceGetAll) Method() string {
	return r.method
}

func (r CustomerInvoiceGetAll) NewRequestBody() CustomerInvoiceGetAllBody {
	return CustomerInvoiceGetAllBody{}
}

type CustomerInvoiceGetAllBody struct {
}

func (r *CustomerInvoiceGetAll) RequestBody() *CustomerInvoiceGetAllBody {
	return nil
}

func (r *CustomerInvoiceGetAll) RequestBodyInterface() interface{} {
	return nil
}

func (r *CustomerInvoiceGetAll) SetRequestBody(body CustomerInvoiceGetAllBody) {
	r.requestBody = body
}

func (r *CustomerInvoiceGetAll) NewResponseBody() *CustomerInvoiceGetAllResponseBody {
	return &CustomerInvoiceGetAllResponseBody{}
}

type CustomerInvoiceGetAllResponseBody Invoices

func (r *CustomerInvoiceGetAll) URL() *url.URL {
	u := r.client.GetEndpointURL("/controller/api/v1/customerinvoice", r.PathParams())
	return &u
}

func (r *CustomerInvoiceGetAll) Do() (CustomerInvoiceGetAllResponseBody, error) {
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