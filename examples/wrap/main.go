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
	err := generateChildError()
	fmt.Println(err)

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

func generateChildError() error {
	if err := generateParentError(); err != nil {
		return errors.Wrap(childError, err)
	}

	return nil
}

func generateParentError() error {
	if err := generateGrandparentError(); err != nil {
		return errors.Wrap(parentError, err)
	}

	return nil
}

func generateGrandparentError() error {
	return errors.NewWithMessage(grandparentError, "top-most error")
}
