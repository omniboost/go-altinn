package altinn_test

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/pem"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	altinn "github.com/omniboost/go-altinn"
	"golang.org/x/crypto/pkcs12"
)

var (
	client *altinn.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	certPem := []byte(os.Getenv("CERT_PEM"))
	keyPem := []byte(os.Getenv("KEY_PEM"))
	p12 := os.Getenv("P12")
	debug := os.Getenv("DEBUG")

	client = altinn.NewClient(nil)
	if debug != "" {
		client.SetDebug(true)
	}

	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		log.Fatal(err)
	}

	data, err := base64.StdEncoding.DecodeString(p12)
	if err != nil {
		log.Fatal(err)
	}

	pems, err := pkcs12.ToPEM(data, "9Tu4Lp3Wac")
	if err != nil {
		log.Fatal(err)
	}

	var pemData []byte
	for _, b := range pems {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	cert, err = tls.X509KeyPair(pemData, pemData)
	if err != nil {
		panic(err)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:  []tls.Certificate{cert},
			Renegotiation: tls.RenegotiateOnceAsClient,
		},
	}
	httpClient := &http.Client{
		Transport: tr,
		Timeout:   30 * time.Second,
	}
	client.SetHTTPClient(httpClient)

	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}
	m.Run()
}
