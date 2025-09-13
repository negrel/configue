package env

// This file contains generated code, do not edit.

import (
	"encoding"
	"time"

	"github.com/negrel/configue/option"
)

// Bool defines a bool option with specified
// name, default value, and usage string. The return value is the address of
// a bool variable that stores the value of the option.
func (es *EnvSet) Bool(
	name string,
	value bool,
	usage string,
) *bool {
	t := new(bool)
	es.BoolVar(t, name, value, usage)
	return t
}

// Bool defines a bool option with specified
// name, default value, and usage string. The argument p points to a bool
// variable in which to store the value of the option.
func (es *EnvSet) BoolVar(
	p *bool,
	name string,
	value bool,
	usage string,
) {
	es.Var(option.NewBool(value, p), name, usage)
}

// Bool defines a slice of bool option with specified
// name, default value, and usage string. The return value is the address of
// a slice of bool variable that stores the value of the option.
func (es *EnvSet) BoolSlice(
	name string,
	value []bool,
	usage string,
) *[]bool {
	t := new([]bool)
	es.BoolSliceVar(t, name, value, usage)
	return t
}

// BoolSliceVar defines a slice of bool option with specified
// name, default value, and usage string. The argument p points to a
// slice of bool variable in which to store the value of the option.
func (es *EnvSet) BoolSliceVar(
	p *[]bool,
	name string,
	value []bool,
	usage string,
) {
	es.Var(option.NewSlice(value, p), name, usage)
}

// Bool defines a bool env var with
// specified name, default value, and usage string. The return value is the
// address of a bool variable that stores the value of
// the env var.
func Bool(name string, value bool, usage string) *bool {
	return CommandLine.Bool(name, value, usage)
}

// BoolVar defines a bool env var with
// specified name, default value, and usage string. The argument p points to
// a bool variable in which to store the value of the
// env var.
func BoolVar(p *bool, name string, value bool, usage string) {
	CommandLine.Var(option.NewBool(value, p), name, usage)
}

// BoolSlice defines a slice of bool env var with
// specified name, default value, and usage string. The return value is the
// address of a slice of bool variable that stores the value of
// the env var.
func BoolSlice(name string, value []bool, usage string) *[]bool {
	return CommandLine.BoolSlice(name, value, usage)
}

// BoolSliceVar defines a slice of bool env var with
// specified name, default value, and usage string. The argument p points to
// a slice of bool variable in which to store the value of the
// env var.
func BoolSliceVar(p *[]bool, name string, value []bool, usage string) {
	CommandLine.Var(option.NewSlice(value, p), name, usage)
}

// Float64 defines a float64 option with specified
// name, default value, and usage string. The return value is the address of
// a float64 variable that stores the value of the option.
func (es *EnvSet) Float64(
	name string,
	value float64,
	usage string,
) *float64 {
	t := new(float64)
	es.Float64Var(t, name, value, usage)
	return t
}

// Float64 defines a float64 option with specified
// name, default value, and usage string. The argument p points to a float64
// variable in which to store the value of the option.
func (es *EnvSet) Float64Var(
	p *float64,
	name string,
	value float64,
	usage string,
) {
	es.Var(option.NewFloat64(value, p), name, usage)
}

// Float64 defines a slice of float64 option with specified
// name, default value, and usage string. The return value is the address of
// a slice of float64 variable that stores the value of the option.
func (es *EnvSet) Float64Slice(
	name string,
	value []float64,
	usage string,
) *[]float64 {
	t := new([]float64)
	es.Float64SliceVar(t, name, value, usage)
	return t
}

// Float64SliceVar defines a slice of float64 option with specified
// name, default value, and usage string. The argument p points to a
// slice of float64 variable in which to store the value of the option.
func (es *EnvSet) Float64SliceVar(
	p *[]float64,
	name string,
	value []float64,
	usage string,
) {
	es.Var(option.NewSlice(value, p), name, usage)
}

// Float64 defines a float64 env var with
// specified name, default value, and usage string. The return value is the
// address of a float64 variable that stores the value of
// the env var.
func Float64(name string, value float64, usage string) *float64 {
	return CommandLine.Float64(name, value, usage)
}

// Float64Var defines a float64 env var with
// specified name, default value, and usage string. The argument p points to
// a float64 variable in which to store the value of the
// env var.
func Float64Var(p *float64, name string, value float64, usage string) {
	CommandLine.Var(option.NewFloat64(value, p), name, usage)
}

// Float64Slice defines a slice of float64 env var with
// specified name, default value, and usage string. The return value is the
// address of a slice of float64 variable that stores the value of
// the env var.
func Float64Slice(name string, value []float64, usage string) *[]float64 {
	return CommandLine.Float64Slice(name, value, usage)
}

// Float64SliceVar defines a slice of float64 env var with
// specified name, default value, and usage string. The argument p points to
// a slice of float64 variable in which to store the value of the
// env var.
func Float64SliceVar(p *[]float64, name string, value []float64, usage string) {
	CommandLine.Var(option.NewSlice(value, p), name, usage)
}

// String defines a string option with specified
// name, default value, and usage string. The return value is the address of
// a string variable that stores the value of the option.
func (es *EnvSet) String(
	name string,
	value string,
	usage string,
) *string {
	t := new(string)
	es.StringVar(t, name, value, usage)
	return t
}

// String defines a string option with specified
// name, default value, and usage string. The argument p points to a string
// variable in which to store the value of the option.
func (es *EnvSet) StringVar(
	p *string,
	name string,
	value string,
	usage string,
) {
	es.Var(option.NewString(value, p), name, usage)
}

// String defines a slice of string option with specified
// name, default value, and usage string. The return value is the address of
// a slice of string variable that stores the value of the option.
func (es *EnvSet) StringSlice(
	name string,
	value []string,
	usage string,
) *[]string {
	t := new([]string)
	es.StringSliceVar(t, name, value, usage)
	return t
}

// StringSliceVar defines a slice of string option with specified
// name, default value, and usage string. The argument p points to a
// slice of string variable in which to store the value of the option.
func (es *EnvSet) StringSliceVar(
	p *[]string,
	name string,
	value []string,
	usage string,
) {
	es.Var(option.NewSlice(value, p), name, usage)
}

// String defines a string env var with
// specified name, default value, and usage string. The return value is the
// address of a string variable that stores the value of
// the env var.
func String(name string, value string, usage string) *string {
	return CommandLine.String(name, value, usage)
}

// StringVar defines a string env var with
// specified name, default value, and usage string. The argument p points to
// a string variable in which to store the value of the
// env var.
func StringVar(p *string, name string, value string, usage string) {
	CommandLine.Var(option.NewString(value, p), name, usage)
}

// StringSlice defines a slice of string env var with
// specified name, default value, and usage string. The return value is the
// address of a slice of string variable that stores the value of
// the env var.
func StringSlice(name string, value []string, usage string) *[]string {
	return CommandLine.StringSlice(name, value, usage)
}

// StringSliceVar defines a slice of string env var with
// specified name, default value, and usage string. The argument p points to
// a slice of string variable in which to store the value of the
// env var.
func StringSliceVar(p *[]string, name string, value []string, usage string) {
	CommandLine.Var(option.NewSlice(value, p), name, usage)
}

// Uint defines an uint option with specified
// name, default value, and usage string. The return value is the address of
// an uint variable that stores the value of the option.
func (es *EnvSet) Uint(
	name string,
	value uint,
	usage string,
) *uint {
	t := new(uint)
	es.UintVar(t, name, value, usage)
	return t
}

// Uint defines an uint option with specified
// name, default value, and usage string. The argument p points to a uint
// variable in which to store the value of the option.
func (es *EnvSet) UintVar(
	p *uint,
	name string,
	value uint,
	usage string,
) {
	es.Var(option.NewUint(value, p), name, usage)
}

// Uint defines a slice of uint option with specified
// name, default value, and usage string. The return value is the address of
// a slice of uint variable that stores the value of the option.
func (es *EnvSet) UintSlice(
	name string,
	value []uint,
	usage string,
) *[]uint {
	t := new([]uint)
	es.UintSliceVar(t, name, value, usage)
	return t
}

// UintSliceVar defines a slice of uint option with specified
// name, default value, and usage string. The argument p points to a
// slice of uint variable in which to store the value of the option.
func (es *EnvSet) UintSliceVar(
	p *[]uint,
	name string,
	value []uint,
	usage string,
) {
	es.Var(option.NewSlice(value, p), name, usage)
}

// Uint defines an uint env var with
// specified name, default value, and usage string. The return value is the
// address of an uint variable that stores the value of
// the env var.
func Uint(name string, value uint, usage string) *uint {
	return CommandLine.Uint(name, value, usage)
}

// UintVar defines an uint env var with
// specified name, default value, and usage string. The argument p points to
// an uint variable in which to store the value of the
// env var.
func UintVar(p *uint, name string, value uint, usage string) {
	CommandLine.Var(option.NewUint(value, p), name, usage)
}

// UintSlice defines a slice of uint env var with
// specified name, default value, and usage string. The return value is the
// address of a slice of uint variable that stores the value of
// the env var.
func UintSlice(name string, value []uint, usage string) *[]uint {
	return CommandLine.UintSlice(name, value, usage)
}

// UintSliceVar defines a slice of uint env var with
// specified name, default value, and usage string. The argument p points to
// a slice of uint variable in which to store the value of the
// env var.
func UintSliceVar(p *[]uint, name string, value []uint, usage string) {
	CommandLine.Var(option.NewSlice(value, p), name, usage)
}

// Uint64 defines an uint64 option with specified
// name, default value, and usage string. The return value is the address of
// an uint64 variable that stores the value of the option.
func (es *EnvSet) Uint64(
	name string,
	value uint64,
	usage string,
) *uint64 {
	t := new(uint64)
	es.Uint64Var(t, name, value, usage)
	return t
}

// Uint64 defines an uint64 option with specified
// name, default value, and usage string. The argument p points to a uint64
// variable in which to store the value of the option.
func (es *EnvSet) Uint64Var(
	p *uint64,
	name string,
	value uint64,
	usage string,
) {
	es.Var(option.NewUint64(value, p), name, usage)
}

// Uint64 defines a slice of uint64 option with specified
// name, default value, and usage string. The return value is the address of
// a slice of uint64 variable that stores the value of the option.
func (es *EnvSet) Uint64Slice(
	name string,
	value []uint64,
	usage string,
) *[]uint64 {
	t := new([]uint64)
	es.Uint64SliceVar(t, name, value, usage)
	return t
}

// Uint64SliceVar defines a slice of uint64 option with specified
// name, default value, and usage string. The argument p points to a
// slice of uint64 variable in which to store the value of the option.
func (es *EnvSet) Uint64SliceVar(
	p *[]uint64,
	name string,
	value []uint64,
	usage string,
) {
	es.Var(option.NewSlice(value, p), name, usage)
}

// Uint64 defines an uint64 env var with
// specified name, default value, and usage string. The return value is the
// address of an uint64 variable that stores the value of
// the env var.
func Uint64(name string, value uint64, usage string) *uint64 {
	return CommandLine.Uint64(name, value, usage)
}

// Uint64Var defines an uint64 env var with
// specified name, default value, and usage string. The argument p points to
// an uint64 variable in which to store the value of the
// env var.
func Uint64Var(p *uint64, name string, value uint64, usage string) {
	CommandLine.Var(option.NewUint64(value, p), name, usage)
}

// Uint64Slice defines a slice of uint64 env var with
// specified name, default value, and usage string. The return value is the
// address of a slice of uint64 variable that stores the value of
// the env var.
func Uint64Slice(name string, value []uint64, usage string) *[]uint64 {
	return CommandLine.Uint64Slice(name, value, usage)
}

// Uint64SliceVar defines a slice of uint64 env var with
// specified name, default value, and usage string. The argument p points to
// a slice of uint64 variable in which to store the value of the
// env var.
func Uint64SliceVar(p *[]uint64, name string, value []uint64, usage string) {
	CommandLine.Var(option.NewSlice(value, p), name, usage)
}

// Int defines an int option with specified
// name, default value, and usage string. The return value is the address of
// an int variable that stores the value of the option.
func (es *EnvSet) Int(
	name string,
	value int,
	usage string,
) *int {
	t := new(int)
	es.IntVar(t, name, value, usage)
	return t
}

// Int defines an int option with specified
// name, default value, and usage string. The argument p points to a int
// variable in which to store the value of the option.
func (es *EnvSet) IntVar(
	p *int,
	name string,
	value int,
	usage string,
) {
	es.Var(option.NewInt(value, p), name, usage)
}

// Int defines a slice of int option with specified
// name, default value, and usage string. The return value is the address of
// a slice of int variable that stores the value of the option.
func (es *EnvSet) IntSlice(
	name string,
	value []int,
	usage string,
) *[]int {
	t := new([]int)
	es.IntSliceVar(t, name, value, usage)
	return t
}

// IntSliceVar defines a slice of int option with specified
// name, default value, and usage string. The argument p points to a
// slice of int variable in which to store the value of the option.
func (es *EnvSet) IntSliceVar(
	p *[]int,
	name string,
	value []int,
	usage string,
) {
	es.Var(option.NewSlice(value, p), name, usage)
}

// Int defines an int env var with
// specified name, default value, and usage string. The return value is the
// address of an int variable that stores the value of
// the env var.
func Int(name string, value int, usage string) *int {
	return CommandLine.Int(name, value, usage)
}

// IntVar defines an int env var with
// specified name, default value, and usage string. The argument p points to
// an int variable in which to store the value of the
// env var.
func IntVar(p *int, name string, value int, usage string) {
	CommandLine.Var(option.NewInt(value, p), name, usage)
}

// IntSlice defines a slice of int env var with
// specified name, default value, and usage string. The return value is the
// address of a slice of int variable that stores the value of
// the env var.
func IntSlice(name string, value []int, usage string) *[]int {
	return CommandLine.IntSlice(name, value, usage)
}

// IntSliceVar defines a slice of int env var with
// specified name, default value, and usage string. The argument p points to
// a slice of int variable in which to store the value of the
// env var.
func IntSliceVar(p *[]int, name string, value []int, usage string) {
	CommandLine.Var(option.NewSlice(value, p), name, usage)
}

// Int64 defines an int64 option with specified
// name, default value, and usage string. The return value is the address of
// an int64 variable that stores the value of the option.
func (es *EnvSet) Int64(
	name string,
	value int64,
	usage string,
) *int64 {
	t := new(int64)
	es.Int64Var(t, name, value, usage)
	return t
}

// Int64 defines an int64 option with specified
// name, default value, and usage string. The argument p points to a int64
// variable in which to store the value of the option.
func (es *EnvSet) Int64Var(
	p *int64,
	name string,
	value int64,
	usage string,
) {
	es.Var(option.NewInt64(value, p), name, usage)
}

// Int64 defines a slice of int64 option with specified
// name, default value, and usage string. The return value is the address of
// a slice of int64 variable that stores the value of the option.
func (es *EnvSet) Int64Slice(
	name string,
	value []int64,
	usage string,
) *[]int64 {
	t := new([]int64)
	es.Int64SliceVar(t, name, value, usage)
	return t
}

// Int64SliceVar defines a slice of int64 option with specified
// name, default value, and usage string. The argument p points to a
// slice of int64 variable in which to store the value of the option.
func (es *EnvSet) Int64SliceVar(
	p *[]int64,
	name string,
	value []int64,
	usage string,
) {
	es.Var(option.NewSlice(value, p), name, usage)
}

// Int64 defines an int64 env var with
// specified name, default value, and usage string. The return value is the
// address of an int64 variable that stores the value of
// the env var.
func Int64(name string, value int64, usage string) *int64 {
	return CommandLine.Int64(name, value, usage)
}

// Int64Var defines an int64 env var with
// specified name, default value, and usage string. The argument p points to
// an int64 variable in which to store the value of the
// env var.
func Int64Var(p *int64, name string, value int64, usage string) {
	CommandLine.Var(option.NewInt64(value, p), name, usage)
}

// Int64Slice defines a slice of int64 env var with
// specified name, default value, and usage string. The return value is the
// address of a slice of int64 variable that stores the value of
// the env var.
func Int64Slice(name string, value []int64, usage string) *[]int64 {
	return CommandLine.Int64Slice(name, value, usage)
}

// Int64SliceVar defines a slice of int64 env var with
// specified name, default value, and usage string. The argument p points to
// a slice of int64 variable in which to store the value of the
// env var.
func Int64SliceVar(p *[]int64, name string, value []int64, usage string) {
	CommandLine.Var(option.NewSlice(value, p), name, usage)
}

// Duration defines a time.Duration option with specified
// name, default value, and usage string. The return value is the address of
// a time.Duration variable that stores the value of the option.
func (es *EnvSet) Duration(
	name string,
	value time.Duration,
	usage string,
) *time.Duration {
	t := new(time.Duration)
	es.DurationVar(t, name, value, usage)
	return t
}

// Duration defines a time.Duration option with specified
// name, default value, and usage string. The argument p points to a time.Duration
// variable in which to store the value of the option.
func (es *EnvSet) DurationVar(
	p *time.Duration,
	name string,
	value time.Duration,
	usage string,
) {
	es.Var(option.NewDuration(value, p), name, usage)
}

// Duration defines a slice of time.Duration option with specified
// name, default value, and usage string. The return value is the address of
// a slice of time.Duration variable that stores the value of the option.
func (es *EnvSet) DurationSlice(
	name string,
	value []time.Duration,
	usage string,
) *[]time.Duration {
	t := new([]time.Duration)
	es.DurationSliceVar(t, name, value, usage)
	return t
}

// DurationSliceVar defines a slice of time.Duration option with specified
// name, default value, and usage string. The argument p points to a
// slice of time.Duration variable in which to store the value of the option.
func (es *EnvSet) DurationSliceVar(
	p *[]time.Duration,
	name string,
	value []time.Duration,
	usage string,
) {
	es.Var(option.NewSlice(value, p), name, usage)
}

// Duration defines a time.Duration env var with
// specified name, default value, and usage string. The return value is the
// address of a time.Duration variable that stores the value of
// the env var.
func Duration(name string, value time.Duration, usage string) *time.Duration {
	return CommandLine.Duration(name, value, usage)
}

// DurationVar defines a time.Duration env var with
// specified name, default value, and usage string. The argument p points to
// a time.Duration variable in which to store the value of the
// env var.
func DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	CommandLine.Var(option.NewDuration(value, p), name, usage)
}

// DurationSlice defines a slice of time.Duration env var with
// specified name, default value, and usage string. The return value is the
// address of a slice of time.Duration variable that stores the value of
// the env var.
func DurationSlice(name string, value []time.Duration, usage string) *[]time.Duration {
	return CommandLine.DurationSlice(name, value, usage)
}

// DurationSliceVar defines a slice of time.Duration env var with
// specified name, default value, and usage string. The argument p points to
// a slice of time.Duration variable in which to store the value of the
// env var.
func DurationSliceVar(p *[]time.Duration, name string, value []time.Duration, usage string) {
	CommandLine.Var(option.NewSlice(value, p), name, usage)
}

// PrintDefaults prints, to standard error unless configured otherwise,
// a usage message showing the default settings of all defined
// env vars.
//
// See the documentation [*EnvSet.PrintDefaults] for more information.
//
// To change the destination for env var messages, call [CommandLine].SetOutput.
func PrintDefaults() {
	CommandLine.PrintDefaults()
}

// Parsed reports whether the command-line env vars have been parsed.
func Parsed() bool {
	return CommandLine.Parsed()
}

// Set sets the value of the named command-line env var.
func Set(name, value string) error {
	return CommandLine.Set(name, value)
}

// UnquoteUsage extracts a back-quoted name from the usage
// string for an env var and returns it and the un-quoted usage.
// Given "a `name` to show" it returns ("name", "a name to show").
// If there are no back quotes, the name is an educated guess of the
// type of the env var's value, or the empty string if the env var is boolean.
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

// TextVar defines an option with a specified name, default value, and usage
// string. The argument p must be a pointer to a variable that will hold the
// value of the option, and p must implement encoding.TextUnmarshal. If the
// option is used, the option value will be passed to p's UnmarshalText
// method. The type of the default value must be the same as the type of p.
func (es *EnvSet) TextVar(
	p encoding.TextUnmarshaler,
	name string,
	value encoding.TextMarshaler,
	usage string,
) {
	es.Var(option.NewText(value, p), name, usage)
}

// TextVar defines an env var with a specified name, default value, and usage string.
// The argument p must be a pointer to a variable that will hold the value
// of the env var, and p must implement encoding.TextUnmarshaler.
// If the env var is used, the env var value will be passed to p's UnmarshalText method.
// The type of the default value must be the same as the type of p.
func TextVar(
	p encoding.TextUnmarshaler,
	name string,
	value encoding.TextMarshaler,
	usage string,
) {
	CommandLine.Var(option.NewText(value, p), name, usage)
}

// Var defines an env var with the specified name and usage string. The type and
// value of the env var are represented by the first argument, of type [Value], which
// typically holds a user-defined implementation of [Value]. For instance, the
// caller could create an env var that turns a comma-separated string into a slice
// of strings by giving the slice the methods of [Value]; in particular, [Set] would
// decompose the comma-separated string into the slice.
func Var(value option.Value, name string, usage string) {
	CommandLine.Var(value, name, usage)
}

// Func defines an env var with the specified name and
// usage string. Each time the env var is seen, fn is called with the
// value of the env var. If fn returns a non-nil error, it will be
// treated as a value parsing error.
func Func(name, usage string, fn func(string) error) {
	CommandLine.Func(name, usage, fn)
}

// Func defines an env var with the specified name and
// usage string. Each time the env var is seen, fn is called with the
// value of the env var. If fn returns a non-nil error, it will be
// treated as a value parsing error.
func (es *EnvSet) Func(name, usage string, fn func(string) error) {
	es.Var(option.Func(fn), name, usage)
}
