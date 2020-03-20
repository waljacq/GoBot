default: run


build:
	go build ./cmd/main.go

run:
	go run ./cmd/main.go

# lint:
# 	go fmt

clean:
	go clean

test:
	go test