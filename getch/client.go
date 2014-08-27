package getch

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	CA_ENV_VAR = "GETCH_CACERT"
)

type Client struct {
	url        string
	httpClient *http.Client
}

func NewClient(url string) *Client {
	var httpClient *http.Client
	//prepare a TLS configured client when the url starts with https
	if strings.HasPrefix(url, "https") {
		//create an empty transport
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		}
		//get the CA cert from an environment variable
		cacert := os.Getenv(CA_ENV_VAR)
		if len(cacert) > 0 {
			//read the file
			pemData, err := ioutil.ReadFile(cacert)
			if err != nil {
				panic(err)
			}
			//append it to a new Certificate Pool
			pool := x509.NewCertPool()
			ok := pool.AppendCertsFromPEM(pemData)
			if ok {
				tr.TLSClientConfig.RootCAs = pool
			}
			tr.TLSClientConfig.InsecureSkipVerify = false
		}
		httpClient = &http.Client{Transport: tr}
	} else {
		httpClient = &http.Client{}
	}
	return &Client{
		url:        url,
		httpClient: httpClient,
	}
}

func mergeUrl(url, key string) string {
	if strings.HasSuffix(url, "/") {
		return fmt.Sprintf("%v%v", url, key)
	} else {
		return fmt.Sprintf("%v/%v", url, key)
	}
}
