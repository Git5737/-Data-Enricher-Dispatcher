package app

import (
	"context"
	"data_enricher_dispatcher/internal/config"
	"data_enricher_dispatcher/internal/service"
	"log"
	"time"
)

const timeout = time.Second * 30

func Run() {
	cfg := config.LoadConfig()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := service.Process(ctx, cfg); err != nil {
		log.Fatalf("service process failed: %v", err)
	}
}
