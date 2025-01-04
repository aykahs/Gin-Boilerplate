package keyclock

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type KeyclockAuthService struct{}

func (auth *KeyclockAuthService) login(username string, password string) {
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
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}
