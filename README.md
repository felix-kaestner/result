# Result

<p align="center">
    <span>Go Completion Implementation with support for Generics (requires Go v1.18+).</span>
    <br><br>
    <a href="https://github.com/felix-kaestner/result/issues">
        <img alt="Issues" src="https://img.shields.io/github/issues/felix-kaestner/result?color=29b6f6&style=flat-square">
    </a>
    <a href="https://github.com/felix-kaestner/result/stargazers">
        <img alt="Stars" src="https://img.shields.io/github/stars/felix-kaestner/result?color=29b6f6&style=flat-square">
    </a>
    <a href="https://github.com/felix-kaestner/result/blob/main/LICENSE">
        <img alt="License" src="https://img.shields.io/github/license/felix-kaestner/result?color=29b6f6&style=flat-square">
    </a>
    <a href="https://pkg.go.dev/github.com/felix-kaestner/result">
        <img alt="Stars" src="https://img.shields.io/badge/go-documentation-blue?color=29b6f6&style=flat-square">
    </a>
    <a href="https://goreportcard.com/report/github.com/felix-kaestner/result">
        <img alt="Issues" src="https://goreportcard.com/badge/github.com/felix-kaestner/result?style=flat-square">
    </a>
    <a href="https://codecov.io/gh/felix-kaestner/result">
        <img src="https://img.shields.io/codecov/c/github/felix-kaestner/result?style=flat-square&token=xfLSFrNzIU"/>
    </a>
    <a href="https://twitter.com/kaestner_felix">
        <img alt="Twitter" src="https://img.shields.io/badge/twitter-@kaestner_felix-29b6f6?style=flat-square">
    </a>
</p>

## Quickstart

```go
package main

import (
	"fmt"

	"github.com/felix-kaestner/result"
)

func main() {
	v := "Hello World"
	r := result.Success(v)
	if r.IsFailure() {
		panic(r.Error())
	}
	if r.IsSuccess() {
		fmt.Println(fmt.Printf("Success %v", r.Value()))
	}
}
```

##  Installation

Install with the `go get` command:

```
$ go get -u github.com/felix-kaestner/result
```

## Contribute

All contributions in any form are welcome! 🙌🏻  
Just use the [Issue](.github/ISSUE_TEMPLATE) and [Pull Request](.github/PULL_REQUEST_TEMPLATE) templates and I'll be happy to review your suggestions. 👍

---

Released under the [MIT License](LICENSE).
