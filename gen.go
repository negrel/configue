package configue

import (
	"encoding"
	"time"

	"github.com/negrel/configue/option"
)

// Bool defines a bool option with specified name, default value, and usage
// string. The return value is the address of a bool variable that stores the
// value of the option.
func (f *Figue) Bool(name string, value bool, usage string) *bool {
	b := new(bool)
	f.BoolVar(b, name, value, usage)
	return b
}

// BoolVar defines a bool option with specified name, default value, and usage
// string.
// The argument p points to a bool variable in which to store the value of the
// option.
func (f *Figue) BoolVar(p *bool, name string, value bool, usage string) {
	f.Var(option.NewBool(value, p), name, usage)
}

// Duration defines a time.Duration option with specified name, default value,
// and usage string. The return value is the address of a time.Duration variable
// that stores the value of the option. The option accepts a value acceptable to
// time.ParseDuration.
func (f *Figue) Duration(name string, value time.Duration, usage string) *time.Duration {
	d := new(time.Duration)
	f.DurationVar(d, name, value, usage)
	return d
}

// DurationVar defines a time.Duration option with specified name, default value,
// and usage string. The argument p points to a time.Duration variable in which
// to store the value of the option. The option accepts a value acceptable to
// time.ParseDuration.
func (f *Figue) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.Var(option.NewDuration(value, p), name, usage)
}

// Float64 defines a float64 option with specified name, default value, and usage
// string. The return value is the address of a float64 variable that stores the
// value of the option.
func (f *Figue) Float64(name string, value float64, usage string) *float64 {
	f64 := new(float64)
	f.Float64Var(f64, name, value, usage)
	return f64
}

// Float64Var defines a float64 option with specified name, default value, and
// usage string. The argument p points to a float64 variable in which to store
// the value of the option.
func (f *Figue) Float64Var(p *float64, name string, value float64, usage string) {
	f.Var(option.NewFloat64(value, p), name, usage)
}

// Int defines an int option with specified name, default value, and usage
// string. The return value is the address of an int variable that stores the
// value of the option.
func (f *Figue) Int(name string, value int, usage string) *int {
	i := new(int)
	f.IntVar(i, name, value, usage)
	return i
}

// IntVar defines an int option with specified name, default value, and usage
// string. The argument p points to an int variable in which to store the value
// of the option.
func (f *Figue) IntVar(p *int, name string, value int, usage string) {
	f.Var(option.NewInt(value, p), name, usage)
}

// Int64 defines an int64 option with specified name, default value, and usage
// string. The return value is the address of an int64 variable that stores the
// value of the option.
func (f *Figue) Int64(name string, value int64, usage string) *int64 {
	i := new(int64)
	f.Int64Var(i, name, value, usage)
	return i
}

// Int64Var defines an int64 option with specified name, default value, and
// usage string. The argument p points to an int64 variable in which to store
// the value of the option.
func (f *Figue) Int64Var(p *int64, name string, value int64, usage string) {
	f.Var(option.NewInt64(value, p), name, usage)
}

// String defines a string option with specified name, default value, and usage
// string. The return value is the address of a string variable that stores the
// value of the option.
func (f *Figue) String(name string, value string, usage string) *string {
	i := new(string)
	f.StringVar(i, name, value, usage)
	return i
}

// StringVar defines a string option with specified name, default value, and
// usage string. The argument p points to a string variable in which to store
// the value of the option.
func (f *Figue) StringVar(p *string, name string, value string, usage string) {
	f.Var(option.NewString(value, p), name, usage)
}

// TextVar defines an option with a specified name, default value, and usage
// string. The argument p must be a pointer to a variable that will hold the
// value of the option, and p must implement encoding.TextUnmarshalf. If the
// option is used, the option value will be passed to p's UnmarshalText
// method. The type of the default value must be the same as the type of p.
func (f *Figue) TextVar(p encoding.TextUnmarshaler, name string, value encoding.TextMarshaler, usage string) {
	f.Var(option.NewText(value, p), name, usage)
}

// Uint defines an uint option with specified name, default value, and usage
// string. The return value is the address of an uint variable that stores the
// value of the option.
func (f *Figue) Uint(name string, value uint, usage string) *uint {
	u := new(uint)
	f.UintVar(u, name, value, usage)
	return u
}

// UintVar defines an uint option with specified name, default value, and usage
// string. The argument p points to an uint variable in which to store the value of the option.
func (f *Figue) UintVar(p *uint, name string, value uint, usage string) {
	f.Var(option.NewUint(value, p), name, usage)
}

// Uint64 defines an uint64 option with specified name, default value, and usage
// string. The return value is the address of an uint64 variable that stores the
// value of the option.
func (f *Figue) Uint64(name string, value uint64, usage string) *uint64 {
	u := new(uint64)
	f.Uint64Var(u, name, value, usage)
	return u
}

// Uint64Var defines an uint64 option with specified name, default value, and
// usage string. The argument p points to an uint64 variable in which to store
// the value of the option.
func (f *Figue) Uint64Var(p *uint64, name string, value uint64, usage string) {
	f.Var(option.NewUint64(value, p), name, usage)
}

// Func defines an option with the specified name and
// usage string. Each time the option is seen, fn is called with the
// value of the option. If fn returns a non-nil error, it will be
// treated as a value parsing error.
func (f *Figue) Func(name, usage string, fn func(string) error) {
	f.Var(option.Func(fn), name, usage)
}

// PrintDefaults prints, to standard error unless configured otherwise,
// a usage message showing the default settings of all defined
// options.
//
// See the documentation [*Figue.PrintDefaults] for more information.
//
// To change the destination for option messages, call [CommandLine].SetOutput.
func PrintDefaults() {
	CommandLine.PrintDefaults()
}

// Parsed reports whether the command-line options have been parsed.
func Parsed() bool {
	return CommandLine.Parsed()
}

// Set sets the value of the named command-line option.
func Set(name, value string) error {
	return CommandLine.Set(name, value)
}

// UnquoteUsage extracts a back-quoted name from the usage
// string for an option and returns it and the un-quoted usage.
// Given "a `name` to show" it returns ("name", "a name to show").
// If there are no back quotes, the name is an educated guess of the
// type of the option's value, or the empty string if the option is boolean.
func UnquoteUsage(optValue option.Value, optUsage string) (name string, usage string) {
	// Look for a back-quoted name, but avoid the strings package.
	usage = optUsage
	for i := 0; i < len(usage); i++ {
		if usage[i] == '`' {
			for j := i + 1; j < len(usage); j++ {
				if usage[j] == '`' {
					name = usage[i+1 : j]
					usage = usage[:i] + name + usage[j+1:]
					return name, usage
				}
			}
			break // Only one back quote; use type name.
		}
	}
	// No explicit name, so use type if we can find one.
	name = "value"
	switch optValue.(type) {
	case *option.Bool:
		name = ""
	case *option.Duration:
		name = "duration"
	case *option.Float64:
		name = "float"
	case *option.Int, *option.Int64:
		name = "int"
	case *option.String:
		name = "string"
	case *option.Uint, *option.Uint64:
		name = "uint"
	}
	return
}

// Bool defines a bool option with specified name, default value, and usage
// string. The return value is the address of a bool variable that stores the
// value of the option.
func Bool(name string, value bool, usage string) *bool {
	return CommandLine.Bool(name, value, usage)
}

// BoolVar defines a bool option with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the option.
func BoolVar(p *bool, name string, value bool, usage string) {
	CommandLine.Var(option.NewBool(value, p), name, usage)
}

// Int defines an int option with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the option.
func Int(name string, value int, usage string) *int {
	return CommandLine.Int(name, value, usage)
}

// IntVar defines an int option with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the option.
func IntVar(p *int, name string, value int, usage string) {
	CommandLine.Var(option.NewInt(value, p), name, usage)
}

// Int64 defines an int64 option with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the option.
func Int64(name string, value int64, usage string) *int64 {
	return CommandLine.Int64(name, value, usage)
}

// Int64Var defines an int64 option with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the option.
func Int64Var(p *int64, name string, value int64, usage string) {
	CommandLine.Var(option.NewInt64(value, p), name, usage)
}

// Uint defines an uint option with specified name, default value, and usage string.
// The return value is the address of an uint variable that stores the value of the option.
func Uint(name string, value uint, usage string) *uint {
	return CommandLine.Uint(name, value, usage)
}

// UintVar defines an uint option with specified name, default value, and usage string.
// The argument p points to an uint variable in which to store the value of the option.
func UintVar(p *uint, name string, value uint, usage string) {
	CommandLine.Var(option.NewUint(value, p), name, usage)
}

// Uint64 defines an uint64 option with specified name, default value, and usage string.
// The return value is the address of an uint64 variable that stores the value of the option.
func Uint64(name string, value uint64, usage string) *uint64 {
	return CommandLine.Uint64(name, value, usage)
}

// Uint64Var defines an uint64 option with specified name, default value, and usage string.
// The argument p points to an uint64 variable in which to store the value of the option.
func Uint64Var(p *uint64, name string, value uint64, usage string) {
	CommandLine.Var(option.NewUint64(value, p), name, usage)
}

// String defines a string option with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the option.
func String(name string, value string, usage string) *string {
	return CommandLine.String(name, value, usage)
}

// StringVar defines a string option with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the option.
func StringVar(p *string, name string, value string, usage string) {
	CommandLine.Var(option.NewString(value, p), name, usage)
}

// Float64 defines a float64 option with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the option.
func Float64(name string, value float64, usage string) *float64 {
	return CommandLine.Float64(name, value, usage)
}

// Float64Var defines a float64 option with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the option.
func Float64Var(p *float64, name string, value float64, usage string) {
	CommandLine.Var(option.NewFloat64(value, p), name, usage)
}

// Duration defines a time.Duration option with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the option.
// The option accepts a value acceptable to time.ParseDuration.
func Duration(name string, value time.Duration, usage string) *time.Duration {
	return CommandLine.Duration(name, value, usage)
}

// DurationVar defines a time.Duration option with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the option.
// The option accepts a value acceptable to time.ParseDuration.
func DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	CommandLine.Var(option.NewDuration(value, p), name, usage)
}

// TextVar defines an option with a specified name, default value, and usage string.
// The argument p must be a pointer to a variable that will hold the value
// of the option, and p must implement encoding.TextUnmarshaler.
// If the option is used, the option value will be passed to p's UnmarshalText method.
// The type of the default value must be the same as the type of p.
func TextVar(p encoding.TextUnmarshaler, name string, value encoding.TextMarshaler, usage string) {
	CommandLine.Var(option.NewText(value, p), name, usage)
}

// Var defines an option with the specified name and usage string. The type and
// value of the option are represented by the first argument, of type [Value], which
// typically holds a user-defined implementation of [Value]. For instance, the
// caller could create an option that turns a comma-separated string into a slice
// of strings by giving the slice the methods of [Value]; in particular, [Set] would
// decompose the comma-separated string into the slice.
func Var(value option.Value, name string, usage string) {
	CommandLine.Var(value, name, usage)
}

// Func defines an option with the specified name and
// usage string. Each time the option is seen, fn is called with the
// value of the option. If fn returns a non-nil error, it will be
// treated as a value parsing error.
func Func(name, usage string, fn func(string) error) {
	CommandLine.Func(name, usage, fn)
}
