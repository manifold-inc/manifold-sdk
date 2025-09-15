// Package utils
package utils

import (
	"errors"
	"fmt"
)

func Wrap(msg string, errs ...error) error {
	fullerr := errors.New(msg)
	for _, err := range errs {
		if err == nil {
			continue
		}
		fullerr = fmt.Errorf("%w: %w", fullerr, err)
	}
	return fullerr
}
