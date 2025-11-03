module github.com/omniboost/go-altinn

go 1.22

require (
	github.com/gorilla/schema v0.0.0-20171211162101-9fa3b6af65dc
	github.com/pkg/errors v0.9.1
	golang.org/x/crypto v0.19.0
	golang.org/x/oauth2 v0.17.0
	gopkg.in/guregu/null.v3 v3.5.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.21.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20211111150515-2e872025e306
