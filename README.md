# kotrcli

[![GoDoc](https://godoc.org/github.com/inabajunmr/kotrcli/github?status.svg)](https://godoc.org/github.com/inabajunmr/kotrcli)

kotrcli is api client for KING OF TIME My Recorder.

## Install
```
go get github.com/inabajunmr/kotrcli/cli/kot
```

## Config
You need `~/.kot/config/config` and put data as follow format. You need get token from Web UI.

```
[default]
user_token = xxxxxx
token = xxxxxx
```

## Usage
```
$ kot syukkin
$ kot taikin
```

## Caution
This tool doesn't use official API.
If official behavior changed, this will doesn't work so you must confirm result of syukkin/taikin.
