package env

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/negrel/configue/option"
)

// EnvVar represents the state of an environment variable.
type EnvVar struct {
	Name     string
	Usage    string
	Value    option.Value
	DefValue string
}

// NewEnvSet returns a new, empty env var set with the specified name and error
// handling property. If the name is not empty, it will be printed in the
// default usage message and in error messages.
func NewEnvSet(name string, errorHandling ErrorHandling) *EnvSet {
	es := &EnvSet{
		name:          name,
		errorHandling: errorHandling,
	}
	es.Usage = es.defaultUsage
	return es
}

// EnvSet represents a set of defined environment variables. The zero value of a
// EnvSet has no name and has ContinueOnError error handling.
//
// EnvVar names must be unique within a EnvSet. An attempt to define an env var
// whose name is already in use will cause a panic.
type EnvSet struct {
	name          string
	parsed        bool
	formal        map[string]*EnvVar
	actual        map[string]*EnvVar
	undef         map[string]string
	output        io.Writer
	errorHandling ErrorHandling
	Usage         func()
}

// Init sets the name and error handling property for a env var set. By default,
// the zero EnvSet uses an empty name and the ContinueOnError error handling
// policy.
func (es *EnvSet) Init(name string, errorHandling ErrorHandling) {
	es.name = name
	es.errorHandling = errorHandling
}

// Var defines an environment variable with the specified name and usage string.
// The type and value of the env var are represented by the first argument, of
// type [Value], which typically holds a user-defined implementation of [Value].
// For instance, the caller could create a env var that turns a comma-separated
// string into a slice of strings by giving the slice the methods of [Value];
// in particular, [Set] would decompose the comma-separated string into the slice.
func (es *EnvSet) Var(val option.Value, name string, usage string) {
	if es.formal == nil {
		es.formal = make(map[string]*EnvVar)
	}

	e := &EnvVar{name, usage, val, val.String()}
	_, exists := es.formal[name]
	if exists {
		var msg string
		if es.name == "" {
			msg = es.sprintf("env var redefined: %s", name)
		} else {
			msg = es.sprintf("%s env var redefined: %s", es.name, name)
		}
		panic(msg)
	}
	if pos := es.undef[name]; pos != "" {
		panic(fmt.Sprintf("env var %s set at %s before being defined", name, pos))
	}

	es.formal[name] = e
}

// Set sets the value of the named command-line env var.
func (es *EnvSet) Set(name, value string) error {
	envVar, ok := es.formal[name]
	if !ok {
		// Remember that a env var that isn't defined is being set.
		// We return an error in this case, but in addition if
		// subsequently that env var is defined, we want to panic
		// at the definition point.
		// This is a problem which occurs if both the definition
		// and the Set call are in init code and for whatever
		// reason the init code changes evaluation order.
		// See issue 57411.
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "?"
			line = 0
		}
		if es.undef == nil {
			es.undef = map[string]string{}
		}
		es.undef[name] = fmt.Sprintf("%s:%d", file, line)

		return fmt.Errorf("no such env var %v", name)
	}
	err := envVar.Value.Set(value)
	if err != nil {
		return err
	}
	if es.actual == nil {
		es.actual = make(map[string]*EnvVar)
	}
	es.actual[name] = envVar
	return nil
}

// Parse parses env var definitions from the environment variable list. Must be
// called after all env vars in the EnvSet are defined and before env vars are
// accessed by the program.
func (es *EnvSet) Parse(envvars []string) error {
	es.parsed = true

	for _, envVar := range envvars {
		seen, err := es.parseOne(envVar)
		if seen {
			continue
		}
		if err == nil {
			break
		}

		es.usage()
		switch es.errorHandling {
		case ContinueOnError:
			return err
		case ExitOnError:
			os.Exit(2)
		case PanicOnError:
			panic(err)
		}
	}

	return nil
}

func (es *EnvSet) parseOne(envVar string) (bool, error) {
	splitted := strings.SplitN(envVar, "=", 2)
	if len(splitted) != 2 {
		return false, es.failf("bad env var syntax: %s", envVar)
	}

	key := splitted[0]
	val := splitted[1]

	// Lookup env var.
	env, ok := es.formal[key]
	if !ok {
		if IgnoreUndefined {
			return true, nil
		}
		return false, es.failf("env var provided but not defined: %s", key)
	}

	if fv, ok := env.Value.(interface{ IsBoolFlag() bool }); ok && fv.IsBoolFlag() {
		if val != "" {
			if err := env.Value.Set(val); err != nil {
				return false, es.failf("invalid boolean value %q for %s: %v", val, key, err)
			}
		} else {
			if err := env.Value.Set(strconv.FormatBool(EmptyBoolValue)); err != nil {
				return false, es.failf("invalid boolean env var %s: %v", key, err)
			}
		}
	} else {
		// Set env var.
		err := env.Value.Set(val)
		if err != nil {
			return false, es.failf("invalid value %q for env var %s: %v", val, key, err)
		}
	}

	// Mark env var as defined.
	if es.actual == nil {
		es.actual = make(map[string]*EnvVar)
	}
	es.actual[key] = env

	return true, nil
}

// Parsed reports whether EnvSet.Parse has been called.
func (es *EnvSet) Parsed() bool {
	return es.parsed
}

// failf prints to standard error a formatted error and usage message and
// returns the error.
func (es *EnvSet) failf(format string, a ...any) error {
	msg := es.sprintf(format, a...)
	es.usage()
	return errors.New(msg)
}

// sprintf formats the message, prints it to output, and returns it.
func (es *EnvSet) sprintf(format string, a ...any) string {
	msg := fmt.Sprintf(format, a...)
	_, _ = fmt.Fprintln(es.Output(), msg)
	return msg
}

// usage calls the Usage method for the env var set if one is specified,
// or the appropriate default usage function otherwise.
func (es *EnvSet) usage() {
	if es.Usage == nil {
		es.defaultUsage()
	} else {
		es.Usage()
	}
}

// defaultUsage is the default function to print a usage message.
func (es *EnvSet) defaultUsage() {
	if es.name == "" {
		_, _ = fmt.Fprintf(es.Output(), "Usage:\n")
	} else {
		_, _ = fmt.Fprintf(es.Output(), "Usage of %s:\n", es.name)
	}
	es.PrintDefaults()
}

// PrintDefaults prints, to standard error unless configured otherwise, the
// default values of all defined env vars in the set.
// For an integer valued env var X, the default output has the form
//
//	X int
//		usage-message-for-x (default 7)
//
// The usage message will appear on a separate line for anything but
// a bool env var with a one-byte name. For bool env vars, the type is
// omitted and if the env var name is one byte the usage message appears
// on the same line. The parenthetical default is omitted if the
// default is the zero value for the type. The listed type, here int,
// can be changed by placing a back-quoted name in the env var's usage
// string; the first such item in the message is taken to be a parameter
// name to show in the message and the back quotes are stripped from
// the message when displayed. For instance, given
//
//	env.String("I", "", "search `directory` for include files")
//
// the output will be
//
//	I directory
//		search directory for include files.
//
// Env vars with usage "HIDDEN" won't be printed:
//
//	env.String("EXPERIMENTAL_OPTION", "", "HIDDEN")
//
// To change the destination for env var messages, call [*EnvSet.SetOutput].
func (es *EnvSet) PrintDefaults() {
	var isZeroValueErrs []error
	es.VisitAll(func(envVar *EnvVar) {
		if envVar.Usage == "HIDDEN" {
			return
		}

		var b strings.Builder
		fmt.Fprintf(&b, "  %s", envVar.Name)
		name, usage := UnquoteUsage(envVar.Value, envVar.Usage)
		if len(name) > 0 {
			b.WriteString(" ")
			b.WriteString(name)
		}
		// Four spaces before the tab triggers good alignment
		// for both 4- and 8-space tab stops.
		b.WriteString("\n    \t")
		b.WriteString(strings.ReplaceAll(usage, "\n", "\n    \t"))

		// Print the default value only if it differs to the zero value
		// for this env var type.
		if isZero, err := isZeroValue(envVar, envVar.DefValue); err != nil {
			isZeroValueErrs = append(isZeroValueErrs, err)
		} else if !isZero {
			if _, ok := envVar.Value.(*option.String); ok {
				// put quotes on the value
				fmt.Fprintf(&b, " (default %q)", envVar.DefValue)
			} else {
				fmt.Fprintf(&b, " (default %v)", envVar.DefValue)
			}
		}
		_, _ = fmt.Fprint(es.Output(), b.String(), "\n")
	})
	// If calling String on any zero env.Values triggered a panic, print
	// the messages after the full set of defaults so that the programmer
	// knows to fix the panic.
	if errs := isZeroValueErrs; len(errs) > 0 {
		_, _ = fmt.Fprintln(es.Output())
		for _, err := range errs {
			_, _ = fmt.Fprintln(es.Output(), err)
		}
	}
}

// VisitAll visits the env vars in lexicographical order, calling fn for each.
// It visits all env vars, even those not set.
func (es *EnvSet) VisitAll(fn func(*EnvVar)) {
	for _, envVar := range sortEnvVars(es.formal) {
		fn(envVar)
	}
}

// Visit visits the command-line env vars in lexicographical order, calling fn
// for each. It visits only those env vars that have been set.
func (es *EnvSet) Visit(fn func(*EnvVar)) {
	for _, envVar := range sortEnvVars(es.actual) {
		fn(envVar)
	}
}

// Output returns the destination for usage and error messages. [os.Stderr] is
// returned if output was not set or was set to nil.
func (es *EnvSet) Output() io.Writer {
	if es.output == nil {
		return os.Stderr
	}
	return es.output
}

// Name returns the name of the env set.
func (es *EnvSet) Name() string {
	return es.name
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, [os.Stderr] is used.
func (es *EnvSet) SetOutput(output io.Writer) {
	es.output = output
}

// Lookup returns the [EnvVar] structure of the named env var, returning nil if
// none exists.
func (es *EnvSet) Lookup(name string) *EnvVar {
	return es.formal[name]
}
