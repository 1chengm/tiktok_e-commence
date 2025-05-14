package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKeyFrontend []byte // Use a different variable name to avoid confusion if copied directly

// Claims defines the JWT claims structure (should match user service's Claims)
type Claims struct {
	UserID int32 `json:"user_id"`
	jwt.RegisteredClaims
}

// InitJWTKey initializes the JWT secret key from an environment variable for the frontend.
func InitJWTKey() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file in frontend service: %v", err))
	}
	key := os.Getenv("JWT_SECRET_KEY") // Must be the SAME key as user service
	if key == "" {
		panic("JWT_SECRET_KEY environment variable not set for frontend")
	}
	jwtKeyFrontend = []byte(key)
}

// ValidateJWT checks the validity of a token string.
func ValidateJWT(tokenStr string) (*Claims, error) {
	if len(jwtKeyFrontend) == 0 {
		return nil, fmt.Errorf("JWT key not initialized in frontend. Call InitJWTKey first")
	}


	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKeyFrontend, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		return nil, fmt.Errorf("token expired")
	}
	return claims, nil
}
