package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bamboooo-dev/meshi-api/app/interface/handler"
	"github.com/bamboooo-dev/meshi-api/app/registry"

	"go.uber.org/zap"
)

var revision = "undefined"

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic '%v' captured\n", err)
		}
	}()

	fmt.Printf("Version is %s\n", revision)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, _ := zap.NewProduction()
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error(err.Error())
			return
		}
	}()

	sugar := logger.Sugar()

	registry := registry.NewRegistry(sugar)

	mux := http.NewServeMux()

	pingHandler := handler.NewPingHandler(sugar, registry)

	mux.HandleFunc("/ping", pingHandler.Handle(ctx))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		sugar.Error(ctx, err)
	}
}
