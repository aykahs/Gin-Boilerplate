package keyclockservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	JwtKC "github.com/aykahs/Gin-Boilerplate/internal/types"
)

type KeyClockAuthService struct {
	*HttpCurl
}

func (auth *KeyClockAuthService) Login(username string, password string) (*JwtKC.TokenResponse, error) {
	if auth == nil {
		return nil, fmt.Errorf("HttpCurl is not initialized")
	}
	client_id := os.Getenv("client_id")
	client_secret := os.Getenv("client_secret")
	formData := url.Values{}
	formData.Set("username", username)
	formData.Set("password", password)
	formData.Set("client_id", client_id)
	formData.Set("client_secret", client_secret)
	formData.Set("grant_type", "password")
	auth.HttpPayload(formData)

	resp, err := auth.HttpPost("token")
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get token. Status Code: %d", resp.StatusCode)
	}
	var tokenResponse JwtKC.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return &tokenResponse, nil
}

func (auth *KeyClockAuthService) Refresh(refresh_token string) (*JwtKC.TokenResponse, error) {
	client_id := os.Getenv("client_id")
	client_secret := os.Getenv("client_secret")
	formData := url.Values{}
	formData.Set("client_id", client_id)
	formData.Set("client_secret", client_secret)
	formData.Set("grant_type", "refresh_token")
	formData.Set("refresh_token", refresh_token)
	auth.HttpPayload(formData)
	resp, err := auth.HttpPost("token")
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get token. Status Code: %d", resp.StatusCode)
	}
	var tokenResponse JwtKC.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return &tokenResponse, nil
}

func (auth *KeyClockAuthService) Me(token string) (*JwtKC.UserInfo, error) {
	client_id := os.Getenv("client_id")
	client_secret := os.Getenv("client_secret")
	formData := url.Values{}
	formData.Set("client_id", client_id)
	formData.Set("client_secret", client_secret)
	formData.Set("token", token)
	auth.HttpPayload(formData)
	resp, err := auth.HttpPost("token/introspect")
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get token. Status Code: %d", resp.StatusCode)
	}
	var userInfo JwtKC.UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return &userInfo, nil
}

func (auth *KeyClockAuthService) FetchJWKS() (*JwtKC.JWKS, error) {
	resp, err := auth.HttpGet("certs")
	if err != nil {
		return nil, fmt.Errorf("error fetching JWKS: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch JWKS: status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading JWKS response: %v", err)
	}

	var jwks JwtKC.JWKS
	if err := json.Unmarshal(body, &jwks); err != nil {
		return nil, fmt.Errorf("error unmarshalling JWKS: %v", err)
	}

	return &jwks, nil
}
