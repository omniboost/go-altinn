package altinn

import "net/url"

type Request interface {
	Method() string
	// QueryParams() QueryParams
	PathParamsInterface() PathParams
	RequestBodyInterface() interface{}
	URL() *url.URL
	IsXML() bool
}

type QueryParams interface {
	ToURLValues() (url.Values, error)
}

type PathParams interface {
	Params() map[string]string
}
