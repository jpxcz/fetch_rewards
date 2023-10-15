package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/jpxcz/fetch_rewards/internal/routes"
)

func main() {
	r := routes.NewRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		slog.Error("problem starting up application", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
