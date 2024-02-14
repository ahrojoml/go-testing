.PHONY: test
test:
	go test ./...

.PHONY: cov
cov:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
