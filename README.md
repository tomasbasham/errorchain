# errors ![test](https://github.com/crumbandbase/errors/workflows/test/badge.svg?event=push)

It is commonly seen within the Go community `error` values being wrapped within
other `error` values before being returned from function and methods, adding
context at each call site. Before Go 1.13 the typical way achieve this was to
create custom types that implement the `error` interface and storing reference
to the parent `error` value. When an application includes many custom error
types this can litter the codebase with a lot of boilerplate.

Go 1.13 introduced the `%w` verb to be used with the `fmt.Errorf` method,
taking any `error` value as an operand and returning a new `error` value with
the
[`Unwrap`](https://github.com/golang/go/blob/62f5e8156ef56fa61e6af56f4ccc633bde1a9120/src/errors/wrap.go#L38)
method already implemented. However multiple custom `error` values cannot be
wrapped using this method without losing type information.

This package tackles this problem, exposing a new `Wrap` method that may be
used to wrap as many custom `error` values as necessary, and returning a new
`error` value that implements the `Unwrap` method. Unlike the value returned
from `fmt.Errorf`, the value returned from `Wrap` does not lose the type
information of any error in the hierarchy.

## Prerequisites

You will need the following things properly installed on your computer.

* [Go](https://golang.org/): any one of the **three latest major**
  [releases](https://golang.org/doc/devel/release.html)

## Installation

With [Go module](https://github.com/golang/go/wiki/Modules) support (Go 1.11+),
simply add the following import

```go
import "github.com/crumbandbase/errors"
```

to your code, and then `go [build|run|test]` will automatically fetch the
necessary dependencies.

Otherwise, to install the `expect` package, run the following command:

```bash
$ go get -u github.com/crumbandbase/errors
```

## Usage

To use this package to wrap errors simply invoke the `Wrap` function taking two
types conforming to the `error` interface; where the first argument is the
error to be wrapped within the error passed as the second argument.

```go
package main

import (
	"fmt"

	"github.com/crumbandbase/errors"
)

const (
	grandparentError = customError("grandparent")
	parentError      = customError("parent")
	childError       = customError("child")
)

type customError string

func (e customError) Error() string {
	return string(e)
}

func main() {
	err := errors.NewWithMessage(grandparentError, "top-most error")
	err = errors.Wrap(parentError, err)
	err = errors.Wrap(childError, err)

	if errors.Is(err, grandparentError) {
		fmt.Println(grandparentError)
	}

	if errors.Is(err, parentError) {
		fmt.Println(parentError)
	}

	if errors.Is(err, childError) {
		fmt.Println(childError)
	}
}
```

## License

This project is licensed under the [MIT License](LICENSE.md).
