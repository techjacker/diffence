# diffence

Checks a git diff for offensive content.
Golang 1.7+

-----------------------------------------------------------
## Install

### Binary
[Download](../../releases) the latest stable release.

### CLI
```
go get -u github.com/techjacker/diffence/cmd/diffence
```

### Library
```
go get -u github.com/techjacker/diffence
```

-----------------------------------------------------------
## CLI tool

### Example Usage
```
touch key.pem
git add -N key.pem
git diff HEAD |diffence
```

-----------------------------------------------------------
## Rules
- currently uses [gitrob rules](https://github.com/michenriksen/gitrob#signature-keys)
- file body rules coming soon (gitrob rules only check the filename for violations)
- option to input your own rules coming soon

-----------------------------------------------------------
## Tests
```
go test ./...
```

-----------------------------------------------------------
## Release
Update the vars in ```env.sh```.

```shell
$ release
```

-----------------------------------------------------------
## Local Development

#### Build & Run Locally
```shell
go install -race ./cmd/diffence
```
OR
```shell
go build -race ./cmd/diffence
```

#### Check for race conditions
```shell
go run -race ./cmd/diffence/main.go
```

