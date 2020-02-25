.PHONY: build test clean

build:
	cfn generate
	env GOOS=linux go build -ldflags="-s -w" -tags="logging callback scheduler" -o bin/handler cmd/main.go

test:
	go test ./...

clean:
	rm -rf bin

deploy: build
	cfn submit --set-default

lint:
	golangci-lint run ./...