package types

import "github.com/golang-jwt/jwt/v5"

type JWKS struct {
	Keys []JWK `json:"keys"`
}

// JWK represents a single JSON Web Key
type JWK struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	N   string `json:"n"` // Modulus
	E   string `json:"e"` // Exponent
}
type KeyClockClaims struct {
	Jti                  string         `json:"jti"`                // JWT ID
	RealmAccess          RealmAccess    `json:"realm_access"`       // Realm-specific access roles
	ResourceAccess       ResourceAccess `json:"resource_access"`    // Resource-specific access roles
	EmailVerified        bool           `json:"email_verified"`     // Whether email is verified
	Name                 string         `json:"name"`               // Full name
	PreferredUsername    string         `json:"preferred_username"` // Preferred username
	Email                string         `json:"email"`              // User email
	ExpiresIn            int            `json:"expires_in"`
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
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type UserInfo struct {
	RealmAccess    RealmAccess    `json:"realm_access"`
	ResourceAccess ResourceAccess `json:"resource_access"`
	EmailVerified  bool           `json:"email_verified"`
	Name           string         `json:"name"`
	Email          string         `json:"email"`
	Username       string         `json:"username"`
	Active         bool           `json:"active"`
}
