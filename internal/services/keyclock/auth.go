package keyclockservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	JwtKC "github.com/aykahs/Gin-Boilerplate/internal/types"
)

type KeyClockAuthService struct{}
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}
type RealmAccess struct {
	Roles []string `json:"roles"`
}

// Struct for Resource Access
type ResourceAccess struct {
	Account struct {
		Roles []string `json:"roles"`
	} `json:"account"`
}

// Main struct representing the user's information
type UserInfo struct {
	RealmAccess    RealmAccess    `json:"realm_access"`
	ResourceAccess ResourceAccess `json:"resource_access"`
	EmailVerified  bool           `json:"email_verified"`
	Name           string         `json:"name"`
	Email          string         `json:"email"`
	Username       string         `json:"username"`
	Active         bool           `json:"active"`
}

func (auth *KeyClockAuthService) Login(username string, password string) (*TokenResponse, error) {
	client_id := os.Getenv("client_id")
	client_secret := os.Getenv("client_secret")
	keyclock_url := os.Getenv("keyclock_url")
	formData := url.Values{}
	formData.Set("username", username)
	formData.Set("password", password)
	formData.Set("client_id", client_id)
	formData.Set("client_secret", client_secret)
	formData.Set("grant_type", "password")
	resp, err := http.Post(keyclock_url+"/token", "application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get token. Status Code: %d", resp.StatusCode)
	}
	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return &tokenResponse, nil
}

func (auth *KeyClockAuthService) Me(token string) (*UserInfo, error) {
	client_id := os.Getenv("client_id")
	client_secret := os.Getenv("client_secret")
	keyclock_url := os.Getenv("keyclock_url")
	formData := url.Values{}
	formData.Set("client_id", client_id)
	formData.Set("client_secret", client_secret)
	formData.Set("token", token)
	resp, err := http.Post(keyclock_url+"/token/introspect", "application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get token. Status Code: %d", resp.StatusCode)
	}
	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return &userInfo, nil
}

func (auth *KeyClockAuthService) FetchJWKS() (*JwtKC.JWKS, error) {
	jwksURL := os.Getenv("keyclock_url")
	resp, err := http.Get(jwksURL + "/certs")
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
