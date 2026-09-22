format:
	gofmt -w src

build:
	go mod tidy && \
	go build -o dist/ src/main.go

run:
	./dist/main

test:
	go test \
		./src/controllers/transaction.go ./src/controllers/transaction_test.go

lint:
	go vet ./src
