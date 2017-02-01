USER = techjacker
REPO = systemdlogger
COMMIT_ID = 94c52865e3b449ca594d09995b99efc28a24c53d

RULES_DIR = test/fixtures/rules
RULES_URL = https://raw.githubusercontent.com/michenriksen/gitrob/master/signatures.json

install:
	@go install -race ./cmd/diffence

lint:
	@golint  -set_exit_status ./...
	@go vet ./...
	@interfacer $(go list ./... | grep -v /vendor/)


rules:
	@curl -s $(RULES_URL) > $(RULES_DIR)/gitrob.json

diff:
	@curl -s https://api.github.com/repos/$(USER)/$(REPO)/commits/$(COMMIT_ID) \
		-H "Accept: application/vnd.github.VERSION.diff"

run:
	@go build -race ./cmd/diffence && ./diffence

test:
	@go test ./...

test-cover:
	@go test -cover ./...

test-race:
	@go test -race ./...


.PHONY: test run
