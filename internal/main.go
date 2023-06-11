package main

import (
	"context"
	"log"

	"github.com/delaneyj/gomponents-iconify/internal/iconify"
)

func main() {
	ctx := context.Background()
	if err := iconify.GenerateIconify(ctx, "gentmp", "."); err != nil {
		log.Fatalf("could not generate iconify: %v", err)
	}
}
