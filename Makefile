REVISION := $(shell git rev-parse --short HEAD)

default:
	go build -o bin/meshi-api -ldflags "-X 'main.revision=$(REVISION)'" main.go

gen-schema:
	gqlgen generate
