package env

import (
	"flag"
	"fmt"
	"os"
)

// ErrorHandling defines how [EnvSet.Parse] behaves if the parse fails.
type ErrorHandling = flag.ErrorHandling

// These constants cause [EnvSet.Parse] to behave as described if the parse
// fails.
const (
	ContinueOnError ErrorHandling = flag.ContinueOnError // Return a descriptive error.
	ExitOnError                   = flag.ExitOnError     // Call os.Exit(2).
	PanicOnError                  = flag.PanicOnError    // Call panic with a descriptive error.
)

var (
	// CommandLine is the default set of command-line env vars, parsed from
	// os.Environ. The top-level functions such as BoolVar, and so on are
	// wrappers for the methods of CommandLine.
	CommandLine = NewEnvSet("", ExitOnError)
	// Usage prints a usage message documenting all defined command-line env vars
	//  to CommandLine's output, which by default is os.Stderr. It is called when
	// an error occurs while parsing env vars. The function is a variable that may
	// be changed to point to a custom function. By default it prints a simple
	// header and calls PrintDefaults; for details about the format of the Output
	// and how to control it, see the documentation for PrintDefaults. Custom
	// usage functions may choose to exit the program; by default exiting happens
	// anyway as the command line's error handling strategy is set to ExitOnError.
	Usage = func() {
		_, _ = fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		PrintDefaults()
	}
	// Value to use when an env var of type bool is an empty string.
	EmptyBoolValue = true
	// Ignore undefined variables instead of returning error.
	IgnoreUndefined = true
)

func init() {
	CommandLine.Usage = cliUsage
}

func cliUsage() {
	Usage()
}

// Parse parses the command-line env vars from os.Environ(). Must be called
// after all env vars are defined and before env vars are accessed by the
// program.
func Parse() error {
	return CommandLine.Parse(os.Environ())
}

// Visit visits the command-line env vars in lexicographical order, calling fn
// for each. It visits only those env vars that have been set.
func Visit(fn func(*EnvVar)) {
	CommandLine.Visit(fn)
}

// VisitAll visits the command-line env vars in lexicographical order, calling
// fn for each. It visits all env vars, even those not set.
func VisitAll(fn func(*EnvVar)) {
	CommandLine.VisitAll(fn)
}
