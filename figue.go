/*
Configue is a dependency-free configuration library inspired by Go's flag
package from the standard library.
*/
package configue

import (
	"flag"
	"io"
	"os"

	"github.com/negrel/configue/option"
)

// Figue define the top level configuration loader.
type Figue struct {
	backends      []Backend
	name          string
	output        io.Writer
	Usage         func()
	errorHandling ErrorHandling
}

// New returns a new Fig instance. This function panics if 0 backend is
// provided. Backends orders matters as it defines priority. You should provides
// backends in ascending order.
func New(
	name string,
	errorHandling ErrorHandling,
	backends ...Backend,
) *Figue {
	if len(backends) < 1 {
		panic("you must provide at least one configue.Backend")
	}

	// Initialize backends.
	for _, b := range backends {
		b.Init(name)
	}

	f := &Figue{
		backends: backends,
		name:     name,
		output:   nil,
	}
	f.Usage = f.defaultUsage
	return f
}

// Var defines an option with the specified name and usage string.
// The type and value of the option are represented by the first argument, of
// type [Value], which typically holds a user-defined implementation of [Value].
// For instance, the caller could create an option that turns a comma-separated
// string into a slice of strings by giving the slice the methods of [Value];
// in particular, [Set] would decompose the comma-separated string into the
// slice.
func (f *Figue) Var(val option.Value, path string, usage string) {
	done := make(map[Backend]struct{})

	for _, b := range f.backends {
		if _, ok := done[b]; ok {
			continue
		}

		b.Var(val, path, usage)
		done[b] = struct{}{}
	}
}

// Parse parses and merges options from their sources. Must be called after all
// options in the Figue are defined and before options are accessed by the program.
func (f *Figue) Parse() error {
	for _, b := range f.backends {
		err := b.Parse()
		if err != nil {
			f.usage()

			switch f.errorHandling {
			case ContinueOnError:
				return err
			case ExitOnError:
				if err == flag.ErrHelp {
					os.Exit(0)
				}
				os.Exit(2)
			case PanicOnError:
				panic(err)
			}
		}
	}

	return nil
}

func (f *Figue) defaultUsage() {
	f.PrintDefaults()
}

func (f *Figue) usage() {
	if f.Usage == nil {
		f.defaultUsage()
	} else {
		f.Usage()
	}
}

// PrintDefaults prints, to standard error unless configured otherwise, the
// default values of all defined command-line options. To do so, it calls in
// reverse order [Backend.PrintDefaults] of all backends. Backend with higher
// priority are printed first.
func (f *Figue) PrintDefaults() {
	done := make(map[Backend]struct{})

	for i := len(f.backends) - 1; i >= 0; i-- {
		b := f.backends[i]

		if _, ok := done[b]; ok {
			continue
		}

		b.PrintDefaults()
		done[b] = struct{}{}
	}
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, [os.Stderr] is used.
func (f *Figue) SetOutput(w io.Writer) {
	f.output = w
	for _, b := range f.backends {
		b.SetOutput(w)
	}
}

// Output returns the destination for usage and error messages. [os.Stderr] is
// returned if output was not set or was set to nil.
func (f *Figue) Output() io.Writer {
	if f.output == nil {
		return os.Stderr
	}
	return f.output
}

// Parsed reports whether Figue.Parse has been called.
func (f *Figue) Parsed() bool {
	for _, b := range f.backends {
		if !b.Parsed() {
			return false
		}
	}
	return len(f.backends) > 0
}

// Set sets the value of the named command-line option.
func (f *Figue) Set(name, value string) error {
	return f.backends[0].Set(name, value)
}
