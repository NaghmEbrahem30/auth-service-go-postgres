package main

import (
	"log"
	"net/http"
	"os"

	"auth-service-go-postgres/internal/api"
	"auth-service-go-postgres/internal/repository"
	"auth-service-go-postgres/internal/service"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	userRepo := repository.NewInMemoryUserRepository()
	authService := service.NewAuthService(userRepo, []byte("dev-secret-key-change-me"))
	handler := api.NewHandler(authService)

	log.Printf("auth service listening on :%s", port)
	if err := http.ListenAndServe(":"+port, handler.Routes()); err != nil {
		log.Fatal(err)
	}
}

