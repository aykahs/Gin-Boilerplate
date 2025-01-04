package utils

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/aykahs/Gin-Boilerplate/configs"
	keyclockservice "github.com/aykahs/Gin-Boilerplate/internal/services/keyclock"
	"github.com/golang-jwt/jwt/v5"

	JwtKC "github.com/aykahs/Gin-Boilerplate/internal/types"
)

type Claims struct {
	jwt.RegisteredClaims

	Uid      uint   `json:"uid"`
	Username string `json:"username"`
}

type KeyClockClaims struct {
	Jti                  string         `json:"jti"`                // JWT ID
	RealmAccess          RealmAccess    `json:"realm_access"`       // Realm-specific access roles
	ResourceAccess       ResourceAccess `json:"resource_access"`    // Resource-specific access roles
	EmailVerified        bool           `json:"email_verified"`     // Whether email is verified
	Name                 string         `json:"name"`               // Full name
	PreferredUsername    string         `json:"preferred_username"` // Preferred username
	Email                string         `json:"email"`              // User email
	jwt.RegisteredClaims                // Embedded registered claims like exp, iat, sub
}

// RealmAccess represents the roles in the realm (like offline_access, default-roles-core, etc.)
type RealmAccess struct {
	Roles []string `json:"roles"`
}

// ResourceAccess represents resource-specific access roles (e.g., account access roles)
type ResourceAccess struct {
	Account struct {
		Roles []string `json:"roles"`
	} `json:"account"`
}

var auth = new(keyclockservice.KeyClockAuthService)

// generate tokens used for auth
func GenerateToken(claims *Claims) string {
	Envconfig := configs.EnvConfig

	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 30)) // set expire time

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(Envconfig.Jwt.Secret))
	if err != nil {
		panic(err)
	}
	return token
}

// verify token
func JwtVerify(tokenStr string) (*Claims, error) {

	Envconfig := configs.EnvConfig
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Envconfig.Jwt.Secret), nil
	})

	if !token.Valid || err != nil {
		return nil, fmt.Errorf("token invalid")
	}
	claims, ok := token.Claims.(*Claims)

	if float64(claims.ExpiresAt.Unix()) < float64(time.Now().Unix()) {
		return nil, fmt.Errorf("token expired")
	}

	if !ok {
		return nil, err
	}
	fmt.Println(claims.Uid)
	return claims, err

}

func GetPublicKeyFromJWKS(jwks *JwtKC.JWKS, kid string) (*rsa.PublicKey, error) {
	for _, key := range jwks.Keys {
		if key.Kid == kid {
			// Decode the modulus (n)
			modulus, err := base64.RawURLEncoding.DecodeString(key.N)
			if err != nil {
				return nil, fmt.Errorf("error decoding modulus: %v", err)
			}

			// Decode the exponent (e)
			exponent, err := base64.RawURLEncoding.DecodeString(key.E)
			if err != nil {
				return nil, fmt.Errorf("error decoding exponent: %v", err)
			}

			// Convert exponent to integer
			eInt := int(new(big.Int).SetBytes(exponent).Int64())

			// Create the RSA public key
			pubKey := &rsa.PublicKey{
				N: new(big.Int).SetBytes(modulus),
				E: eInt,
			}
			return pubKey, nil
		}
	}

	return nil, errors.New("key ID not found in JWKS")
}
func JwtKeyClockVerify(tokenStr string) (*KeyClockClaims, error) {
	jwks, err := auth.FetchJWKS()
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokenStr, &KeyClockClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("missing kid in token header")
		}

		return GetPublicKeyFromJWKS(jwks, kid)
	})
	if err != nil {
		return nil, fmt.Errorf("token invalid")
	}
	claims, ok := token.Claims.(*KeyClockClaims)
	if float64(claims.ExpiresAt.Unix()) < float64(time.Now().Unix()) {
		return nil, fmt.Errorf("token expired")
	}

	if !ok {
		return nil, err
	}
	return claims, err

}
