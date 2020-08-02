ci: vet test

vet:
	go vet ./...

test:
	go test ./...

example:
	go run example/example.go