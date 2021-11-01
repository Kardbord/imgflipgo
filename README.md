# imgflipgo

[![Unit Tests](https://github.com/TannerKvarfordt/imgflipgo/actions/workflows/unit-tests.yml/badge.svg)](https://github.com/TannerKvarfordt/imgflipgo/actions/workflows/unit-tests.yml)
[![CodeQL](https://github.com/TannerKvarfordt/imgflipgo/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/TannerKvarfordt/imgflipgo/actions/workflows/codeql-analysis.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/TannerKvarfordt/imgflipgo)](https://goreportcard.com/report/github.com/TannerKvarfordt/imgflipgo)

(Golang) Go bindings for the [Imgflip API](https://imgflip.com/api).

## Usage

To include this library in your Go project using modules, add the following import and run `go mod tidy`.

```Go
import "github.com/TannerKvarfordt/imgflipgo"
```

Otherwise, you can run the following command.

```sh
go get github.com/TannerKvarfordt/imgflipgo
```

The code is fairly self-documenting (said every developer too lazy to write real docs). There are only two [API](https://imgflip.com/api) endpoints.

- The `get_memes` endpoint (https://api.imgflip.com/get_memes) can be accessed via `imgflipgo.GetMemesWithResponse()` or `imgflipgo.GetMemes()`.
- The `caption_image` endpoint (https://api.imgflip.com/caption_image) can be accessed via `imgflipgo.CaptionImage(*CaptionRequest)`.

For a concrete example of how to use the library, check out [example.go](https://github.com/TannerKvarfordt/imgflipgo/blob/main/example/example.go).
