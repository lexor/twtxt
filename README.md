# twtxt

[![Build Status](https://cloud.drone.io/api/badges/prologic/twtxt/status.svg)](https://cloud.drone.io/prologic/twtxt)
[![CodeCov](https://codecov.io/gh/prologic/twtxt/branch/master/graph/badge.svg)](https://codecov.io/gh/prologic/twtxt)
[![Go Report Card](https://goreportcard.com/badge/prologic/twtxt)](https://goreportcard.com/report/prologic/twtxt)
[![GoDoc](https://godoc.org/github.com/prologic/twtxt?status.svg)](https://godoc.org/github.com/prologic/twtxt) 
[![Sourcegraph](https://sourcegraph.com/github.com/prologic/twtxt/-/badge.svg)](https://sourcegraph.com/github.com/prologic/twtxt?badge)

twtxt is a [twtxt](https://twtxt.readthedocs.io/en/latest/) client in the form
of a web application and command-line client. It supports multiple users and
also hosts user feeds directly. It also  has a builtin registry and search.

There is also a publically (_free_) service online available at:

- https://twtxt.net/

![Screenshot](./screenshot.png)
![Screenshot 2](./screenshot2.png)

## Installation

### Source

```#!bash
$ go get -u github.com/prologic/twtxt/...
```

## Usage

### CLI

Run twt:

```#!bash
$ twt
```

### Web App

Run twtd:

```#!bash
$ twtd
```

Then visit: http://localhost:8000/

## License

twtwt is licensed under the terms of the [MIT License](/LICENSE)
