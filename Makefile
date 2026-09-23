format:
	gofmt -w ./src

build:
	go mod tidy && \
	go build -o dist/ src/main.go

run:
	./dist/main

test:
	go test ./src/...

static-check:
	go vet ./src/...
