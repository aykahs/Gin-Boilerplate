package utils

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	keyclockservice "github.com/aykahs/Gin-Boilerplate/internal/services/keyclock"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	JwtKC "github.com/aykahs/Gin-Boilerplate/internal/types"
)

var auth = &keyclockservice.KeyClockAuthService{
	HttpCurl: &keyclockservice.HttpCurl{},
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
func JwtKeyClockVerify(tokenStr string) (*JwtKC.KeyClockClaims, error) {
	jwks, err := auth.FetchJWKS()
	fmt.Print(err)
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokenStr, &JwtKC.KeyClockClaims{}, func(token *jwt.Token) (interface{}, error) {
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
	claims, ok := token.Claims.(*JwtKC.KeyClockClaims)
	if float64(claims.ExpiresAt.Unix()) < float64(time.Now().Unix()) {
		return nil, fmt.Errorf("token expired")
	}

	if !ok {
		return nil, err
	}
	return claims, err

}

func GetToken(ctx *gin.Context) (string, error) {
	tokenStr := ctx.GetHeader("Authorization")
	if tokenStr == "" {
		ctx.Abort()
		return "", errors.New("authorization header is missing")
	}
	parts := strings.Split(tokenStr, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("authorization header is missing")
	}

	return parts[1], nil
}
