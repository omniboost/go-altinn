module github.com/omniboost/go-altinn

go 1.24.0

require (
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/google/uuid v1.6.0
	github.com/gorilla/schema v0.0.0-20171211162101-9fa3b6af65dc
	github.com/pkg/errors v0.9.1
	github.com/salrashid123/golang-jwt-signer v0.3.2
	golang.org/x/crypto v0.48.0
	golang.org/x/oauth2 v0.35.0
	gopkg.in/guregu/null.v3 v3.5.0
)

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20211111150515-2e872025e306
