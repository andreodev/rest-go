package tokens

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)

type Claims struct {
	Subject   string `json:"sub"`
	Email     string `json:"email"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
}

func GenerateJWT(userID, email string) (string, error) {
	now := time.Now()

	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}

	claims := Claims{
		Subject:   userID,
		Email:     email,
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(24 * time.Hour).Unix(),
	}

	headerSegment, err := encodeSegment(header)
	if err != nil {
		return "", err
	}

	claimsSegment, err := encodeSegment(claims)
	if err != nil {
		return "", err
	}

	unsignedToken := headerSegment + "." + claimsSegment
	signature := sign(unsignedToken)

	return unsignedToken + "." + signature, nil
}

func ValidateJWT(token string) (Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return Claims{}, ErrInvalidToken
	}

	unsignedToken := parts[0] + "." + parts[1]
	expectedSignature := sign(unsignedToken)
	if !hmac.Equal([]byte(parts[2]), []byte(expectedSignature)) {
		return Claims{}, ErrInvalidToken
	}

	var header struct {
		Algorithm string `json:"alg"`
		Type      string `json:"typ"`
	}
	if err := decodeSegment(parts[0], &header); err != nil {
		return Claims{}, ErrInvalidToken
	}

	if header.Algorithm != "HS256" || header.Type != "JWT" {
		return Claims{}, ErrInvalidToken
	}

	var claims Claims
	if err := decodeSegment(parts[1], &claims); err != nil {
		return Claims{}, ErrInvalidToken
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return Claims{}, ErrExpiredToken
	}

	return claims, nil
}

func encodeSegment(data any) (string, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(payload), nil
}

func decodeSegment(segment string, data any) error {
	payload, err := base64.RawURLEncoding.DecodeString(segment)
	if err != nil {
		return err
	}

	return json.Unmarshal(payload, data)
}

func sign(unsignedToken string) string {
	mac := hmac.New(sha256.New, getJWTSecret())
	mac.Write([]byte(unsignedToken))

	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "development-secret"
	}

	return []byte(secret)
}
