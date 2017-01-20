run:
	@go build ./cmd/diffence && ./diffence

test:
	@go test -v ./...

.PHONY: test run
