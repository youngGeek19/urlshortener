package main

import (
	"fmt"
	"urlshortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Print(cfg)

	// TODO: init logger: slog

	// TODO: init storage: sqlite3

	// TODO: init router: chi

	// TODO: run server:
}
