package internal

import (
	"testing"

	"auth-service-go-postgres/internal/repository"
	"auth-service-go-postgres/internal/service"
)

func TestRegisterAndLogin(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	svc := service.NewAuthService(repo, []byte("secret"))

	if err := svc.Register("demo@example.com", "password123"); err != nil {
		t.Fatalf("register failed: %v", err)
	}
	token, err := svc.Login("demo@example.com", "password123")
	if err != nil || token == "" {
		t.Fatalf("login failed: %v", err)
	}
	sub, err := svc.ValidateToken(token)
	if err != nil || sub != "demo@example.com" {
		t.Fatalf("token validation failed, sub=%s err=%v", sub, err)
	}
}

