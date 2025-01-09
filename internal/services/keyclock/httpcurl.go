package keyclockservice

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type HttpCurl struct {
	Payload url.Values
}

func (httpcurl *HttpCurl) HttpPayload(payload url.Values) {
	fmt.Println(payload)
	httpcurl.Payload = payload
}
func (httpcurl *HttpCurl) BaseUrl() string {
	base, _ := os.LookupEnv("keyclock_base_url")
	realm, _ := os.LookupEnv("realms")
	keyclockUrl, _ := os.LookupEnv("keyclock_url")
	fullpath := fmt.Sprintf("%s%s", base, fmt.Sprintf(keyclockUrl, realm))
	return fullpath
}
func (httpcurl *HttpCurl) HttpPost(url string) (*http.Response, error) {

	resp, err := http.Post(httpcurl.BaseUrl()+url, "application/x-www-form-urlencoded", strings.NewReader(httpcurl.Payload.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("failed to get token. Status Code: %d", resp.StatusCode)
	}

	return resp, nil
}

func (httpcurl *HttpCurl) HttpGet(url string) (*http.Response, error) {
	resp, err := http.Get(httpcurl.BaseUrl() + url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("failed to get token. Status Code: %d", resp.StatusCode)
	}

	return resp, nil
}
