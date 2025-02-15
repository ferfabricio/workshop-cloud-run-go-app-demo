package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("APP_NAME")
	if name == "" {
		name = "Cloud Run App"
	}
	fmt.Fprintf(w, "Olá! Bem-vindo ao %s 🚀", name)
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handler)
	logger.Info("servidor inicializado", "porta", port, "app-name", os.Getenv("APP_NAME"))

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		logger.Error("falha ao iniciar o servidor", "erro", err)
	}
}
