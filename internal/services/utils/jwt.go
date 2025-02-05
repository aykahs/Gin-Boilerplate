package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	JwtKC "github.com/aykahs/Gin-Boilerplate/internal/types"
)

func GetPublicKeyFromJWKS() (*rsa.PublicKey, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting current directory: %w", err)
	}
	filePath := wd + "/internal/services/utils/public.pem"
	pemData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading public key file: %w", err)
	}

	// Decode the PEM block
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	// Parse the key
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing public key: %w", err)
	}

	// Assert the type
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}

	return rsaPub, nil
}

func JwtKeyClockVerify(tokenStr string) (*JwtKC.KeyClockClaims, error) {
	// Fetch the RSA public key
	jwks, err := GetPublicKeyFromJWKS()
	if err != nil {
		return nil, fmt.Errorf("error fetching public key: %w", err)
	}

	// Prepare claims
	claims := &JwtKC.KeyClockClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is RSA
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwks, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}
	// Validate token claims
	claims, ok := token.Claims.(*JwtKC.KeyClockClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// if time.Now().After(time.Unix(claims.ExpiresAt, 0)) {
	// 	return nil, fmt.Errorf("token expired")
	// }

	return claims, nil
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
