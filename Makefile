unit-test:
	go test -tags unit -cover ./...

run:
	go run cmd/honeycombs/main.go

build:
	go build -o ./tmp/honeycombs cmd/honeycombs/main.go
