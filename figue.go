/*
Configue is a dependency-free configuration library inspired by Go's flag
package from the standard library.
*/
package configue

import (
	"fmt"
	"io"
	"maps"
	"os"
	"slices"

	"github.com/negrel/configue/option"
)

// Figue define the top level configuration loader.
type Figue struct {
	backends []Backend
	options  map[string]struct{}
	name     string
	output   io.Writer
	Usage    func()
}

// New returns a new Fig instance. This function panics if 0 backend is provided.
func New(name string, errorHandling ErrorHandling, newBackends ...func(string, ErrorHandling) Backend) *Figue {
	if len(newBackends) < 1 {
		panic("you must provide at least one configue.Backend")
	}

	var backends []Backend
	for _, b := range newBackends {
		backends = append(backends, b(name, errorHandling))
	}

	f := &Figue{
		backends: backends,
		options:  make(map[string]struct{}),
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
	for _, b := range f.backends {
		b.Var(val, path, usage)
	}
	f.options[path] = struct{}{}
}

// Parse parses and merges options from their sources. Must be called after all
// options in the Figue are defined and before options are accessed by the program.
func (f *Figue) Parse() error {
	undefinedOptions := maps.Clone(f.options)

	for _, b := range f.backends {
		err := b.Parse()
		if err != nil {
			return err
		}

		// Remove defined options.
		b.Visit(func(name string, _ option.Value) {
			delete(undefinedOptions, name)
		})
	}

	undefined := slices.Collect(maps.Keys(undefinedOptions))
	_ = undefined

	return nil
}

func (f *Figue) defaultUsage() {
	if f.name != "" {
		_, _ = fmt.Fprintln(f.Output(), "Usage of", f.name)
	} else {
		_, _ = fmt.Fprintln(f.Output(), "Usage:")
	}
	f.PrintDefaults()
}

// PrintDefaults prints, to standard error unless configured otherwise, the
// default values of all defined command-line options. See the
// documentation for the global function PrintDefaults for more information.
func (f *Figue) PrintDefaults() {
	for _, b := range f.backends {
		b.PrintDefaults()
		_, _ = fmt.Fprintln(f.Output())
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
		if b.Parsed() {
			return true
		}
	}
	return false
}

// Set sets the value of the named command-line option.
func (f *Figue) Set(name, value string) error {
	return f.backends[0].Set(name, value)
}
