package service

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"auth-service-go-postgres/internal/domain"
	"auth-service-go-postgres/internal/repository"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidInput       = errors.New("invalid input")
)

type AuthService struct {
	users  repository.UserRepository
	secret []byte
}

func NewAuthService(users repository.UserRepository, secret []byte) *AuthService {
	return &AuthService{users: users, secret: secret}
}

func (s *AuthService) Register(email, password string) error {
	if !strings.Contains(email, "@") || len(password) < 8 {
		return ErrInvalidInput
	}
	user := domain.User{
		ID:       fmt.Sprintf("u_%d", time.Now().UnixNano()),
		Email:    email,
		Password: hashPassword(password),
	}
	return s.users.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, ok := s.users.FindByEmail(email)
	if !ok || user.Password != hashPassword(password) {
		return "", ErrInvalidCredentials
	}
	return s.issueToken(user.Email)
}

func (s *AuthService) ValidateToken(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return "", ErrInvalidCredentials
	}
	payloadRaw, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return "", ErrInvalidCredentials
	}
	signatureRaw, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", ErrInvalidCredentials
	}
	if string(sign(payloadRaw, s.secret)) != string(signatureRaw) {
		return "", ErrInvalidCredentials
	}
	var payload map[string]string
	if err := json.Unmarshal(payloadRaw, &payload); err != nil {
		return "", ErrInvalidCredentials
	}
	return payload["sub"], nil
}

func (s *AuthService) issueToken(subject string) (string, error) {
	payload := map[string]string{"sub": subject}
	raw, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	sig := sign(raw, s.secret)
	return base64.RawURLEncoding.EncodeToString(raw) + "." + base64.RawURLEncoding.EncodeToString(sig), nil
}

func sign(msg, secret []byte) []byte {
	sum := sha256.Sum256(append(msg, secret...))
	return sum[:]
}

func hashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", sum[:])
}
