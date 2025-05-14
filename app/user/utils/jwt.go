package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtKey []byte

// Claims defines the JWT claims structure
type Claims struct {
	UserID int32 `json:"user_id"`
	jwt.RegisteredClaims
}

// InitJWTKey initializes the JWT secret key from an environment variable.
// Call this function at the startup of your service.
func InitJWTKey() {

	err := godotenv.Load("./.env")
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file in user service: %v", err))
	}
	key := os.Getenv("JWT_SECRET_KEY") // Must be the SAME key as user service
	if key == "" {
		panic("JWT_SECRET_KEY environment variable not set for user service")
	}
	jwtKey = []byte(key)
}

// GenerateJWT creates a new JWT token for a given userID.
func GenerateJWT(userID int32) (string, error) {
	if len(jwtKey) == 0 {
		return "", fmt.Errorf("JWT key not initialized. Call InitJWTKey first")
	}
	expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "gomall-user-service", // Identifies who issued the token
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return tokenString, nil
}

// ValidateJWT checks the validity of a token string.
// This function can be used by any service that needs to validate a token.
func ValidateJWT(tokenStr string) (*Claims, error) {
	if len(jwtKey) == 0 {
		return nil, fmt.Errorf("JWT key not initialized. Call InitJWTKey first")
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Check for expiration, though ParseWithClaims with RegisteredClaims usually handles this.
	if time.Now().After(claims.ExpiresAt.Time) {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}
