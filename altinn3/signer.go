package altinn3

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	jwtsigner "github.com/salrashid123/golang-jwt-signer"
)

type JWTSigner struct {
	keyContext  context.Context
	environment string
	clientID    string
	keyID       string
	Debug       bool
}

type AccessTokenResponse struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   int       `json:"expires_in"`
	Expiry      time.Time `json:"-"`
	Scope       string    `json:"scope"`
}

const SYSTEM_REGISTER_SCOPE = "altinn:authentication/systemregister.write"
const SYSTEM_REGISTER_USER_SCOPE = "altinn:authentication/systemuser.request.write altinn:authentication/systemuser.request.read"

// { "aud", GetAssertionAud(environment) },
// { "scope", scope },
// { "iss", clientId },
// { "exp", dateTimeOffset.ToUnixTimeSeconds() + 10 },
// { "iat", dateTimeOffset.ToUnixTimeSeconds() },
// { "jti", Guid.NewGuid().ToString() },
// { "authorization_details", new { type = "urn:altinn:systemuser", systemuser_org = new {
// {"authority","iso6523-actorid-upis"}
// {"ID", "0192:{systemUserOrgno}" }
// } } }

type TimeInt time.Time

func (t TimeInt) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", time.Time(t).Unix())), nil
}

type JWTPayload struct {
	Audience  string  `json:"aud"`
	Scope     string  `json:"scope"`
	Issuer    string  `json:"iss"`
	ExpiresAt TimeInt `json:"exp"`
	IssuedAt  TimeInt `json:"iat"`
	Jti       string  `json:"jti"`

	AuthorizationDetails *struct {
		Type          string `json:"type"`
		SystemUserOrg struct {
			Authority string `json:"authority"`
			ID        string `json:"ID"`
		} `json:"systemuser_org"`
	} `json:"authorization_details,omitempty"`
}

func (J *JWTPayload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Time(J.ExpiresAt)), nil
}

func (J *JWTPayload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Time(J.IssuedAt)), nil
}

func (J *JWTPayload) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

func (J *JWTPayload) GetIssuer() (string, error) {
	return J.Issuer, nil
}

func (J *JWTPayload) GetSubject() (string, error) {
	return "", nil
}

func (J *JWTPayload) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{J.Audience}, nil
}

var _ jwt.Claims = (*JWTPayload)(nil)

func NewJWTSigner(key []byte, keyID, environment, clientId string) (*JWTSigner, error) {
	if IsBase64(string(key)) {
		var err error
		key, err = base64.StdEncoding.DecodeString(string(key))
		if err != nil {
			return nil, err
		}
	}
	rblock, _ := pem.Decode(key)
	r, err := x509.ParsePKCS8PrivateKey(rblock.Bytes)
	if err != nil {
		return nil, err
	}
	var rsaPrivKey *rsa.PrivateKey
	var ok bool
	if rsaPrivKey, ok = r.(*rsa.PrivateKey); !ok {
		return nil, fmt.Errorf("RSA private key is not of type *rsa.PrivateKey")
	}

	jwtsigner.SigningMethodSignerRS256.Override()
	ctx := context.Background()
	keyctx, err := jwtsigner.NewSignerContext(ctx, &jwtsigner.SignerConfig{
		Signer: rsaPrivKey,
	})
	if err != nil {
		return nil, err
	}

	return &JWTSigner{
		keyContext:  keyctx,
		environment: environment,
		clientID:    clientId,
		keyID:       keyID,
	}, nil
}

func (j *JWTSigner) GetAccessTokenForSystemRegister() (*AccessTokenResponse, error) {
	token := jwt.NewWithClaims(jwtsigner.SigningMethodSignerRS256, &JWTPayload{
		Audience:             GetAssertionAud(j.environment),
		Scope:                SYSTEM_REGISTER_SCOPE,
		Issuer:               j.clientID,
		ExpiresAt:            TimeInt(time.Now().Add(10 * time.Second).UTC()),
		IssuedAt:             TimeInt(time.Now().UTC()),
		Jti:                  uuid.NewString(),
		AuthorizationDetails: nil,
	})
	token.Header["kid"] = j.keyID
	tokenString, err := token.SignedString(j.keyContext)
	if err != nil {
		return nil, err
	}

	if j.Debug {
		println("Generated JWT:")
		println(tokenString)
	}

	// Create form url encoded body
	values := url.Values{}
	values.Set("grant_type", "urn:ietf:params:oauth:grant-type:jwt-bearer")
	values.Set("assertion", tokenString)
	formBody := values.Encode()
	if j.Debug {
		println("Form URL Encoded Body:")
		println(formBody)
	}
	endpoint := GetTokenEndpoint(j.environment)

	// Create HTTP POST request
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(formBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if j.Debug {
		rr, _ := httputil.DumpRequest(req, true)
		fmt.Printf("%s\n", rr)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if j.Debug {
		rrr, _ := httputil.DumpResponse(resp, true)
		fmt.Printf("%s\n", rrr)
	}

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			log.Printf("Error closing response body: %v", cerr)
		}
	}()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Token endpoint returned non-200 status: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var tokenResponse AccessTokenResponse
	err = json.Unmarshal(bodyBytes, &tokenResponse)
	if err != nil {
		return nil, err
	}
	tokenResponse.Expiry = time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second).UTC()
	return &tokenResponse, nil
}

func IsBase64(input string) bool {
	r, _ := regexp.Compile("^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$")
	return r.MatchString(input)
}

func GetBaseURL(environment string) string {
	switch environment {
	case "prod":
		return "https://maskinporten.no/"
	case "test":
		return "https://test.maskinporten.no/"
	default:
		panic("Invalid environment setting. Valid values: prod, test")
	}
}

func GetAssertionAud(environment string) string {
	return strings.TrimRight(GetBaseURL(environment), "/") + "/token"
}

func GetTokenEndpoint(environment string) string {
	return strings.TrimRight(GetBaseURL(environment), "/") + "/token"
}
