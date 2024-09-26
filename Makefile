# Target to simply run the go run command in the backend directory

.PHONY: run
run:
	go run ./cmd/main.go

build:
	go build -o ./build/main -v ./cmd/main.go 

clean:
	go clean
	rm -rf build

.PHONY: test
test:
	go test -v ./...
