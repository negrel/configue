package env

import (
	"errors"
	"strconv"
)

// ErrorHandling defines how [EnvSet.Parse] behaves if the parse fails.
type ErrorHandling int

// These constants cause [EnvSet.Parse] to behave as described if the parse
// fails.
const (
	ContinueOnError ErrorHandling = iota // Return a descriptive error.
	ExitOnError                          // Call os.Exit(2) or for -h/-help Exit(0).
	PanicOnError                         // Call panic with a descriptive error.
)

// errParse is returned by Set if a Value fails to parse, such as with
// an invalid integer for Int.
var errParse = errors.New("parse error")

// errRange is returned by Set if an env var's value is out of range.
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
