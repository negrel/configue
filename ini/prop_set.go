package ini

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"slices"
	"strconv"
	"strings"

	"github.com/negrel/configue/option"
)

// Property represents the state of an INI key value.
type Property struct {
	Name     string
	Usage    string
	Value    option.Value
	DefValue string
}

func NewPropSet(name string, errorHandling ErrorHandling) *PropSet {
	ps := &PropSet{
		name:          name,
		errorHandling: errorHandling,
	}
	ps.Usage = ps.defaultUsage
	return ps
}

// PropSet represents a set of defined options. The zero value of a PropSet
// has no name and has ContinueOnError error handling.
//
// PropSet names must be unique within a PropSet. An attempt to define an
// option whose name is already in use will cause a panic.
type PropSet struct {
	name          string
	parsed        bool
	formal        map[string]*Property
	actual        map[string]*Property
	undef         map[string]string
	output        io.Writer
	errorHandling ErrorHandling
	Usage         func()
}

// Init sets the name and error handling property for a property set. By default,
// the zero PropSet uses an empty name and the ContinueOnError error handling
// policy.
func (ps *PropSet) Init(name string, errorHandling ErrorHandling) {
	ps.name = name
	ps.errorHandling = errorHandling
}

// Parsed reports whether PropSet.Parse has been called.
func (ps *PropSet) Parsed() bool {
	return ps.parsed
}

// Name returns the name of the property set.
func (ps *PropSet) Name() string {
	return ps.name
}

// Parse parses INI properties from the provided io.Reader. Must be called after
// all properties in the PropSet are defined and before properties are accessed
// by the program.
func (ps *PropSet) Parse(r io.Reader) error {
	if r == nil {
		r = strings.NewReader("")
	}

	ps.parsed = true

	parser := newParser(r)

	for {
		seen, err := ps.parseOne(parser)
		if seen {
			continue
		}
		if err == nil {
			break
		}

		ps.usage()
		switch ps.errorHandling {
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

func (ps *PropSet) parseOne(parser *parser) (bool, error) {
	key, val, err := parser.parseNext()
	if key == "" && val == "" && err == nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	// Lookup property.
	prop, ok := ps.formal[key]
	if !ok {
		return false, ps.failf("property provided but not defined: %s", key)
	}

	if fv, ok := prop.Value.(interface{ IsBoolFlag() bool }); ok && fv.IsBoolFlag() {
		if val != "" {
			if err := prop.Value.Set(val); err != nil {
				return false, ps.failf("invalid boolean value %q for %s: %v", val, key, err)
			}
		}
	} else {
		// Set property.
		err := prop.Value.Set(val)
		if err != nil {
			return false, ps.failf("invalid value %q for property %s: %v", val, key, err)
		}
	}

	// Mark property as defined.
	if ps.actual == nil {
		ps.actual = make(map[string]*Property)
	}
	ps.actual[key] = prop

	return true, nil
}

// Lookup returns the [Property] structure of the named property, returning nil
// if none exists.
func (ps *PropSet) Lookup(name string) *Property {
	return ps.formal[name]
}

// Var defines a property with the specified name and usage string.
// The type and value of the option are represented by the first argument, of
// type [option.Value], which typically holds a user-defined implementation of
// [option.Value].
// For instance, the caller could create a property that turns a comma-separated
// string into a slice of strings by giving the slice the methods of
// [option.Value]; in particular, [Set] would decompose the comma-separated
// string into the slice.
func (ps *PropSet) Var(val option.Value, name string, usage string) {
	if ps.formal == nil {
		ps.formal = make(map[string]*Property)
	}

	if strings.Contains(name, "..") {
		panic(ps.sprintf("property %q has consecutive dots", name))
	} else if strings.Contains(name, "=") {
		panic(ps.sprintf("property %q contains =", name))
	}

	e := &Property{name, usage, val, val.String()}
	_, exists := ps.formal[name]
	if exists {
		var msg string
		if ps.name == "" {
			msg = ps.sprintf("property redefined: %s", name)
		} else {
			msg = ps.sprintf("%s property redefined: %s", ps.name, name)
		}
		panic(msg)
	}
	if pos := ps.undef[name]; pos != "" {
		panic(fmt.Sprintf("property %s set at %s before being defined", name, pos))
	}

	ps.formal[name] = e
}

func (ps *PropSet) Set(key, value string) error {
	prop, ok := ps.formal[key]
	if !ok {
		// Remember that a property that isn't defined is being set.
		// We return an error in this case, but in addition if
		// subsequently that property is defined, we want to panic
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
		if ps.undef == nil {
			ps.undef = map[string]string{}
		}
		ps.undef[key] = fmt.Sprintf("%s:%d", file, line)

		return fmt.Errorf("no such property %v", key)
	}
	err := prop.Value.Set(value)
	if err != nil {
		return err
	}
	if ps.actual == nil {
		ps.actual = make(map[string]*Property)
	}
	ps.actual[key] = prop
	return nil
}

// defaultUsage is the default function to print a usage message.
func (ps *PropSet) defaultUsage() {
	if ps.name == "" {
		_, _ = fmt.Fprintf(ps.Output(), "Usage:\n")
	} else {
		_, _ = fmt.Fprintf(ps.Output(), "Usage of %s:\n", ps.name)
	}
	ps.PrintDefaults()
}

// Output returns the destination for usage and error messages. [os.Stderr] is
// returned if output was not set or was set to nil.
func (ps *PropSet) Output() io.Writer {
	if ps.output == nil {
		return os.Stderr
	}
	return ps.output
}

// PrintDefaults prints, to standard error unless configured otherwise, the
// default values of all defined properties in the set.
// For an integer valued property X with a default value of 7, the output has
// the form
//
//	; usage-message-for-x
//	X = 7
//
// The usage message appear as an INI comment on a separate line.
func (ps *PropSet) PrintDefaults() {
	var isZeroValueErrs []error
	previousSection := ""

	ps.VisitAll(func(prop *Property) {
		var b strings.Builder

		path := strings.Split(prop.Name, ".")
		section := strings.Join(path[:len(path)-1], ".")
		key := path[len(path)-1]

		padding := ""

		// Section.
		if section != previousSection {
			_ = b.WriteByte('[')
			if after, ok := strings.CutPrefix(section, previousSection); ok {
				_, _ = b.WriteString(after)
			} else {
				_, _ = b.WriteString(section)
			}
			_ = b.WriteByte(']')
			_ = b.WriteByte('\n')
		}
		if section != "" {
			padding = "  "
		}

		// Usage + default value.
		name, usage := UnquoteUsage(prop.Value, prop.Usage)
		_, _ = b.WriteString("; ")
		if len(name) > 0 {
			_, _ = b.WriteString(usage)
			_ = b.WriteByte('\n')
			_, _ = b.WriteString("; ")
			_, _ = b.WriteString(key)
			_, _ = b.WriteString(" = ")
			_, _ = b.WriteString(name)
		} else {
			_, _ = b.WriteString(prop.Usage)
		}
		_ = b.WriteByte('\n')

		_, _ = b.WriteString(padding)
		_, _ = b.WriteString(key)
		_, _ = b.WriteString(" = ")
		if strings.ContainsAny(prop.DefValue, ";#\"`\\") ||
			strings.TrimSpace(prop.DefValue) != prop.DefValue {
			_, _ = b.WriteString(strconv.Quote(prop.DefValue))
		} else {
			_, _ = b.WriteString(prop.DefValue)
		}
		_, _ = b.WriteString("\n")

		_, _ = fmt.Fprint(ps.Output(), b.String(), "\n")

		previousSection = section
	})
	// If calling String on any zero option.Values triggered a panic, print
	// the messages after the full set of defaults so that the programmer
	// knows to fix the panic.
	if errs := isZeroValueErrs; len(errs) > 0 {
		_, _ = fmt.Fprintln(ps.Output())
		for _, err := range errs {
			_, _ = fmt.Fprintln(ps.Output(), err)
		}
	}
}

// VisitAll visits the options in lexicographical order, calling fn for each.
// It visits all options, even those not set.
func (ps *PropSet) VisitAll(fn func(*Property)) {
	for _, prop := range sortProperties(ps.formal) {
		fn(prop)
	}
}

// Visit visits the command-line properties in lexicographical order, calling fn
// for each. It visits only those properties that have been set.
func (ps *PropSet) Visit(fn func(*Property)) {
	for _, prop := range sortProperties(ps.actual) {
		fn(prop)
	}
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, [os.Stderr] is used.
func (ps *PropSet) SetOutput(output io.Writer) {
	ps.output = output
}

// usage calls the Usage method for the property set if one is specified,
// or the appropriate default usage function otherwise.
func (ps *PropSet) usage() {
	if ps.Usage == nil {
		ps.defaultUsage()
	} else {
		ps.Usage()
	}
}

// sprintf formats the message, prints it to output, and returns it.
func (ps *PropSet) sprintf(format string, a ...any) string {
	msg := fmt.Sprintf(format, a...)
	_, _ = fmt.Fprintln(ps.Output(), msg)
	return msg
}

// failf prints to standard error a formatted error and usage message and
// returns the error.
func (ps *PropSet) failf(format string, a ...any) error {
	msg := ps.sprintf(format, a...)
	ps.usage()
	return errors.New(msg)
}

func sortProperties(options map[string]*Property) []*Property {
	result := make([]*Property, len(options))
	i := 0
	for _, f := range options {
		result[i] = f
		i++
	}
	slices.SortFunc(result, func(a, b *Property) int {
		return strings.Compare(a.Name, b.Name)
	})
	return result
}
