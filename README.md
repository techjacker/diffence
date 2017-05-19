[![Build Status](https://travis-ci.org/techjacker/diffence.svg?branch=master)](https://travis-ci.org/techjacker/diffence)
[![Go Report Card](https://goreportcard.com/badge/github.com/techjacker/diffence)](https://goreportcard.com/report/github.com/techjacker/diffence)

# diffence
- Checks a git diff for passwords/secret keys accidentally committed
- Golang 1.7+

-----------------------------------------------------------
### Check the entire history of current branch for passwords/keys committed


```$ git log -p --full-diff | diffence```
```$ git log -p | diffence```


### Git Diff Formats
```
Path names in extended headers do not include the a/ and b/ prefixes.
Only supports regular diff formats.
```

### bufio.NewScanner Limitations
```
// Programs that need more control over error handling or large tokens,
// or must run sequential scans on a reader, should use bufio.Reader instead.
```
-----------------------------------------------------------
### Add false positives to `.secignore`

```
$ cat .secignore
file/that/is/not/really/a/secret/but/looks/like/one/to/diffence
these/pems/are/ok/*.pem
```

[See example in this repo](./.secignore).

-----------------------------------------------------------
## Install

### Binary
[Download](../../releases) the latest stable release.

### CLI
```
$ go get -u github.com/techjacker/diffence/cmd/diffence
```

### Library
```
$ go get -u github.com/techjacker/diffence
```

-----------------------------------------------------------
## CLI tool

### Example Usage
```
$ touch key.pem

$ git add -N key.pem

$ git diff --stat HEAD
gds HEAD
 key.pem | 0
 1 file changed, 0 insertions(+), 0 deletions(-)

$ git diff HEAD |diffence
File key.pem violates 1 rules:

Caption: Potential cryptographic private key
Description: <nil>
Part: extension
Pattern: pem
Type: match


```

-----------------------------------------------------------
## Rules
- [x] Analyse fPaths with [gitrob rules](https://github.com/michenriksen/gitrob#signature-keys)
- [ ] Analyse added lines - need to find/create ruleset that can analyse file contents
- [ ] Add option to use your own rules again file path/contents


-----------------------------------------------------------
## Tests
```
$ go test ./...
```

-----------------------------------------------------------
## Local Development

#### Build & Run Locally
```shell
$ go install -race ./cmd/diffence
```
OR
```shell
$ go build -race ./cmd/diffence
```

#### Check for race conditions
```shell
$ go run -race ./cmd/diffence/main.go
```

