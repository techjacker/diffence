USER = techjacker
REPO = systemdlogger
COMMIT_ID = a4e1c01c137c2218684fd59da54d769623ac567f

GITROB_RULES = https://raw.githubusercontent.com/michenriksen/gitrob/master/signatures.json

# https://github.com/michenriksen/gitrob/blob/master/signatures.json
rules:
	@curl -s $(GITROB_RULES) > ./test/fixtures/gitrob.json

diff:
	@curl -L https://api.github.com/repos/$(USER)/$(REPO)/commits/$(COMMIT_ID) \
		-H "Accept: application/vnd.github.VERSION.diff"

run:
	@go build ./cmd/diffence && ./diffence

test:
	@go test -v ./...

.PHONY: test run
