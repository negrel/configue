package ini

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// ErrorHandling defines how [PropSet.Parse] behaves if the parse fails.
type ErrorHandling = flag.ErrorHandling

// These constants cause [PropSet.Parse] to behave as described if the parse
// fails.
const (
	ContinueOnError ErrorHandling = flag.ContinueOnError // Return a descriptive error.
	ExitOnError                   = flag.ExitOnError     // Call os.Exit(2).
	PanicOnError                  = flag.PanicOnError    // Call panic with a descriptive error.
)

var (
	// CommandLine is the default set of command-line properties, parsed from
	// provided io.Reader when calling [Parse]. The top-level functions such as
	// BoolVar, and so on are wrappers for the methods of CommandLine.
	CommandLine = NewPropSet("", ExitOnError)
	// Usage prints a usage message documenting all defined command-line
	// properties to CommandLine's output, which by default is os.Stderr. It is
	// called when an error occurs while parsing INI properties. The function is a
	// variable that may be changed to point to a custom function. By default it
	// prints a simple header and calls PrintDefaults; for details about the
	// format of the Output and how to control it, see the documentation for
	// PrintDefaults. Custom usage functions may choose to exit the program; by
	// default exiting happens anyway as the command line's error handling
	// strategy is set to ExitOnError.
	Usage = func() {
		_, _ = fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		PrintDefaults()
	}
)

func init() {
	CommandLine.Usage = cliUsage
}

func cliUsage() {
	Usage()
}

// Parse parses the command-line properties from provided io.Reader. Must be
// called after all properties are defined and before properties are accessed by
// the program.
func Parse(r io.Reader) error {
	return CommandLine.Parse(r)
}

// Visit visits the command-line properties in lexicographical order, calling fn
// for each. It visits only those properties that have been set.
func Visit(fn func(*Property)) {
	CommandLine.Visit(fn)
}

// VisitAll visits the command-line properties in lexicographical order, calling
// fn for each. It visits all properties, even those not set.
func VisitAll(fn func(*Property)) {
	CommandLine.VisitAll(fn)
}
