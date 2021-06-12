[![license](https://img.shields.io/github/license/nwillc/cryptoport.svg)](https://tldrlegal.com/license/-isc-license)
[![CI](https://github.com/nwillc/cryptoport/workflows/CI/badge.svg)](https://github.com/nwillc/cryptoport/actions?query=workflow%3CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/nwillc/cryptoport)](https://goreportcard.com/report/github.com/nwillc/cryptoport)
---

# cryptoport
cli crypto portfolio 

Cli that produces output like:
![screenshot](cryptoport.png)

## Build

```shell
go build .
```

## Use
### Help
```shell
$ cryptoport -h
A simple crypto portfolio status cli that reports the value of your portfolio

Usage:
  cryptoport [flags]
  cryptoport [command]

Available Commands:
  help        Help about any command
  setup       Setup your portfolio configuration.
  version     Print the version number

Flags:
  -h, --help   help for cryptoport
```

1. Have your Nomics API key and holdings info ready. Setup your configuration:
```shell
$ cryptoport setup
```
1. Run...
```shell
$ cryptoport
```

## TODO
- timestamp & deltas?
