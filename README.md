# loaderiotoken-go

[![Build Status](https://travis-ci.org/TV4/loaderiotoken-go.svg?branch=master)](https://travis-ci.org/TV4/loaderiotoken-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/TV4/loaderiotoken-go)](https://goreportcard.com/report/github.com/TV4/loaderiotoken-go)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/TV4/loaderiotoken-go)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/TV4/loaderiotoken-go#license)

loaderiotoken-go is a Go package for setting up a verification endpoint for the
loader.io load test service.

## Usage
```go
package main

import (
	"net/http"

	loaderiotoken "github.com/TV4/loaderiotoken-go"
)

func main() {
	mux := http.NewServeMux()

	loaderiotoken.Register(mux, "97bd38305a81f2d89b5f3aa44500ec96")

	http.ListenAndServe(":8080", mux)
}
```

```
$ http get :8080/loaderio-97bd38305a81f2d89b5f3aa44500ec96.txt
HTTP/1.1 200 OK
Cache-Control: public, max-age=10
Content-Length: 42
Content-Type: text/plain; charset=utf-8
Date: Wed, 08 Nov 2017 22:22:05 GMT

loaderio-97bd38305a81f2d89b5f3aa44500ec96
```

## License
Copyright (c) 2017 Bonnier Broadcasting

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
