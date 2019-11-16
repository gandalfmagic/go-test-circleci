GOFILES = $(shell find . -name '*.go')

default: build

build: bin/pouch

build-native: $(GOFILES)
	go build -o bin/native-pouch$$(go env GOEXE) cmd/pouch/main.go

bin/pouch: $(GOFILES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/pouch cmd/pouch/main.go
