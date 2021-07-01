REVISION := $(shell git rev-parse --short HEAD)

default:
	go build -o bin/meshi-api -ldflags "-X 'main.revision=$(REVISION)'" main.go

install-devdeps: install-golangcilint install-gotests install-ent install-gqlgen

install-golangcilint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.1

install-gotests:
	go install github.com/cweill/gotests/...@v1.6.0

install-ent:
	go install entgo.io/ent/cmd/ent@v0.8.0

install-gqlgen:
	go install github.com/99designs/gqlgen@v0.13.0

lint:
	golangci-lint run --timeout 5m
	
gen-schema:
	gqlgen generate

gen-meshi-db-schema:
	go generate ./ent
