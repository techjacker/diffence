# diffence

Checks a git diff for offensive content.

Golang 1.7+

-----------------------------------------------------------

## Install CLI tool

### Binary
[Download](../../releases) the latest stable release.

### Source

##### CLI tool
```
go get -u github.com/techjacker/diffence/cmd/diffence
```

##### Library
```
go get -u github.com/techjacker/diffence
```

-----------------------------------------------------------

## Usage
```
```

-----------------------------------------------------------
## [Gitrob Rules - Signature Keys](https://github.com/michenriksen/gitrob#signature-keys)





-----------------------------------------------------------
## JSON Configuration

By default the CLI tool looks for config.json in $PWD. You can specify a custom location with the `config` flag, eg:

```Shell
diffence -config ./test/fixtures/config.json
```


#### Example JSON Config
```json
{
	"rules": {
		"jsonPath": "./test/fixtures/rules.json"
	}
}
```

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
go install ./cmd/diffence && diffence -config ./test/fixtures/config.json
```

OR

```shell
go build -race ./cmd/diffence && ./diffence -config ./test/fixtures/config.json
```


#### Check for race conditions

```shell
go run -race ./cmd/diffence/main.go -config ./test/fixtures/config.json
```

