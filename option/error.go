package option

import (
	"errors"
	"strconv"
)

// errParse is returned by Set if a Value fails to parse, such as with
// an invalid integer for Int.
var errParse = errors.New("parse error")

// errRange is returned by Set if an option's value is out of range.
var errRange = errors.New("value out of range")

func numError(err error) error {
	ne, ok := err.(*strconv.NumError)
	if !ok {
		return err
	}
	if ne.Err == strconv.ErrSyntax {
		return errParse
	}
	if ne.Err == strconv.ErrRange {
		return errRange
	}
	return err
}
