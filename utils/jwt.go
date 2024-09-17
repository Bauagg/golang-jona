package utils

import (
	"backend-jona-golang/config"
	"errors"

	"github.com/golang-jwt/jwt"
)

// Define the secret key for signing the JWT
var jwtSecret = []byte(config.SECRET_KEY)

// Claims struct for JWT payload
type Claims struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

// SignToken generates a JWT token with the given ID, email, and role
func SignToken(id uint64, email, role string) (string, error) {
	claims := Claims{
		ID:    id,
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			Issuer: "JonaNusantara", // Issuer of the token
		},
	}

	// Use HS256 signing method for HMAC with secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// VerifyToken verifies the given JWT token string and returns the claims if valid
func VerifyToken(tokenStr string) (*Claims, error) {
	// Parse the token with claims
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC (HS256)
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Validate the token and its claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// Additional validation (check the issuer)
		if claims.Issuer != "JonaNusantara" {
			return nil, jwt.NewValidationError("Invalid issuer", jwt.ValidationErrorIssuer)
		}
		return claims, nil
	}

	// If token is invalid or signature is invalid
	return nil, jwt.NewValidationError("Invalid token", jwt.ValidationErrorSignatureInvalid)
}
