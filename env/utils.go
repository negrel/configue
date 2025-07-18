package env

import (
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/negrel/configue/option"
)

// isZeroValue determines whether the string represents the zero
// value for an env var.
func isZeroValue(envvar *EnvVar, value string) (ok bool, err error) {
	// Build a zero value of the env var's Value type, and see if the
	// result of calling its String method equals the value passed in.
	// This works unless the Value type is itself an interface type.
	typ := reflect.TypeOf(envvar.Value)
	var z reflect.Value
	if typ.Kind() == reflect.Pointer {
		z = reflect.New(typ.Elem())
	} else {
		z = reflect.Zero(typ)
	}
	// Catch panics calling the String method, which shouldn't prevent the
	// usage message from being printed, but that we should report to the
	// user so that they know to fix their code.
	defer func() {
		if e := recover(); e != nil {
			if typ.Kind() == reflect.Pointer {
				typ = typ.Elem()
			}
			err = fmt.Errorf("panic calling String method on zero %v for env var %s: %v", typ, envvar.Name, e)
		}
	}()
	return value == z.Interface().(option.Value).String(), nil
}

func sortEnvVars(envVars map[string]*EnvVar) []*EnvVar {
	result := make([]*EnvVar, len(envVars))
	i := 0
	for _, f := range envVars {
		result[i] = f
		i++
	}
	slices.SortFunc(result, func(a, b *EnvVar) int {
		return strings.Compare(a.Name, b.Name)
	})
	return result
}
