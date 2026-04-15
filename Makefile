format:
	gofmt -w src

build:
	go build -o dist/ src/main.go

run:
	./dist/main
