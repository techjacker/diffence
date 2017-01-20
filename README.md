# diffence

Checks a diff for offensive content.

-----------------------------------------------------------

## Install CLI tool

### Binary
[Download](../../releases) the latest stable release.

### Source
```
go get -u github.com/techjacker/diffence/cmd/diffence
```


-----------------------------------------------------------

## Usage
```
```


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
go build ./cmd/diffence && ./diffence -config ./test/fixtures/config.json
```


#### Check for race conditions

```shell
go run -race ./cmd/diffence/main.go -config ./test/fixtures/config.json
```

