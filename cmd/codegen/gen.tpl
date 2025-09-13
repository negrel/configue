package {{ .Package }}

// This file contains generated code, do not edit.

import (
	"encoding"
	"time"

	"github.com/negrel/configue/option"
)

{{ define "option" }}
// {{ .Type }} defines {{ .typeArticle }} {{ .type }} option with specified
// name, default value, and usage string. The return value is the address of
// {{ .typeArticle }} {{ .type }} variable that stores the value of the option.
func ({{ .params.MethodReceiver }} {{ .params.MethodType }}) {{ .Type }}(
	name string,
	value {{ .type }},
	usage string,
) *{{ .type }} {
	t := new({{ .type }})
	{{ .params.MethodReceiver }}.{{ .Type }}Var(t, name, value, usage)
	return t
}

// {{ .Type }} defines {{ .typeArticle }} {{ .type }} option with specified
// name, default value, and usage string. The argument p points to a {{ .type }}
// variable in which to store the value of the option.
func ({{ .params.MethodReceiver }} {{ .params.MethodType }}) {{ .Type }}Var(
	p *{{ .type }},
	name string,
	value {{ .type }},
	usage string,
) {
	{{ .params.MethodReceiver }}.Var(option.New{{ .Type }}(value, p), name, usage)
}

// {{ .Type }} defines a slice of {{ .type }} option with specified
// name, default value, and usage string. The return value is the address of
// a slice of {{ .type }} variable that stores the value of the option.
func ({{ .params.MethodReceiver }} {{ .params.MethodType }}) {{ .Type }}Slice(
	name string,
	value []{{ .type }},
	usage string,
) *[]{{ .type }} {
	t := new([]{{ .type }})
	{{ .params.MethodReceiver }}.{{ .Type }}SliceVar(t, name, value, usage)
	return t
}

// {{ .Type }}SliceVar defines a slice of {{ .type }} option with specified
// name, default value, and usage string. The argument p points to a
// slice of {{ .type }} variable in which to store the value of the option.
func ({{ .params.MethodReceiver }} {{ .params.MethodType }}) {{ .Type }}SliceVar(
	p *[]{{ .type }},
	name string,
	value []{{ .type }},
	usage string,
) {
	{{ .params.MethodReceiver }}.Var(option.NewSlice(value, p), name, usage)
}

// {{ .Type }} defines {{ .typeArticle }} {{ .type }} {{.params.OptionName }} with
// specified name, default value, and usage string. The return value is the
// address of {{ .typeArticle }} {{ .type }} variable that stores the value of
// the {{.params.OptionName }}.
func {{ .Type }}(name string, value {{ .type }}, usage string) *{{ .type }} {
	return CommandLine.{{ .Type }}(name, value, usage)
}

// {{ .Type }}Var defines {{ .typeArticle }} {{ .type }} {{.params.OptionName }} with
// specified name, default value, and usage string. The argument p points to
// {{ .typeArticle }} {{ .type }} variable in which to store the value of the
// {{.params.OptionName }}.
func {{ .Type }}Var(p *{{ .type }}, name string, value {{ .type }}, usage string) {
	CommandLine.Var(option.New{{ .Type }}(value, p), name, usage)
}

// {{ .Type }}Slice defines a slice of {{ .type }} {{.params.OptionName }} with
// specified name, default value, and usage string. The return value is the
// address of a slice of {{ .type }} variable that stores the value of
// the {{.params.OptionName }}.
func {{ .Type }}Slice(name string, value []{{ .type }}, usage string) *[]{{ .type }} {
	return CommandLine.{{ .Type }}Slice(name, value, usage)
}

// {{ .Type }}SliceVar defines a slice of {{ .type }} {{.params.OptionName }} with
// specified name, default value, and usage string. The argument p points to
// a slice of {{ .type }} variable in which to store the value of the
// {{.params.OptionName }}.
func {{ .Type }}SliceVar(p *[]{{ .type }}, name string, value []{{ .type }}, usage string) {
	CommandLine.Var(option.NewSlice(value, p), name, usage)
}
{{ end }}

{{ $params := . }}
{{ range $type := (list "bool" "float64" "string") }}
{{ template "option" (dict "typeArticle" "a" "type" $type "Type" (title $type) "params" $params) }}
{{ end }}

{{ $params := . }}
{{ range $type := (list "uint" "uint64" "int" "int64") }}
{{ template "option" (dict "typeArticle" "an" "type" $type "Type" (title $type) "params" $params) }}
{{ end }}

{{ template "option" (dict "typeArticle" "a" "type" "time.Duration" "Type" "Duration" "params" $params) }}

// PrintDefaults prints, to standard error unless configured otherwise,
// a usage message showing the default settings of all defined
// {{ .OptionName }}s.
//
// See the documentation [{{ .MethodType }}.PrintDefaults] for more information.
//
// To change the destination for {{ .OptionName }} messages, call [CommandLine].SetOutput.
func PrintDefaults() {
	CommandLine.PrintDefaults()
}

// Parsed reports whether the command-line {{ .OptionName }}s have been parsed.
func Parsed() bool {
	return CommandLine.Parsed()
}

// Set sets the value of the named command-line {{ .OptionName }}.
func Set(name, value string) error {
	return CommandLine.Set(name, value)
}

// UnquoteUsage extracts a back-quoted name from the usage
// string for an {{ .OptionName }} and returns it and the un-quoted usage.
// Given "a `name` to show" it returns ("name", "a name to show").
// If there are no back quotes, the name is an educated guess of the
// type of the {{ .OptionName }}'s value, or the empty string if the {{ .OptionName }} is boolean.
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
func ({{ .MethodReceiver }} {{ .MethodType }}) TextVar(
	p encoding.TextUnmarshaler,
	name string,
	value encoding.TextMarshaler,
	usage string,
) {
	{{ .MethodReceiver }}.Var(option.NewText(value, p), name, usage)
}

// TextVar defines {{ .OptionArticle }} {{ .OptionName }} with a specified name, default value, and usage string.
// The argument p must be a pointer to a variable that will hold the value
// of the {{ .OptionName }}, and p must implement encoding.TextUnmarshaler.
// If the {{ .OptionName }} is used, the {{ .OptionName }} value will be passed to p's UnmarshalText method.
// The type of the default value must be the same as the type of p.
func TextVar(
	p encoding.TextUnmarshaler,
	name string,
	value encoding.TextMarshaler,
	usage string,
) {
	CommandLine.Var(option.NewText(value, p), name, usage)
}

// Var defines an {{ .OptionName }} with the specified name and usage string. The type and
// value of the {{ .OptionName }} are represented by the first argument, of type [Value], which
// typically holds a user-defined implementation of [Value]. For instance, the
// caller could create {{ .OptionArticle }} {{ .OptionName }} that turns a comma-separated string into a slice
// of strings by giving the slice the methods of [Value]; in particular, [Set] would
// decompose the comma-separated string into the slice.
func Var(value option.Value, name string, usage string) {
	CommandLine.Var(value, name, usage)
}

// Func defines {{ .OptionArticle }} {{ .OptionName }} with the specified name and
// usage string. Each time the {{ .OptionName }} is seen, fn is called with the
// value of the {{ .OptionName }}. If fn returns a non-nil error, it will be
// treated as a value parsing error.
func Func(name, usage string, fn func(string) error) {
	CommandLine.Func(name, usage, fn)
}

// Func defines {{ .OptionArticle }} {{ .OptionName }} with the specified name and
// usage string. Each time the {{ .OptionName }} is seen, fn is called with the
// value of the {{ .OptionName }}. If fn returns a non-nil error, it will be
// treated as a value parsing error.
func ({{ .MethodReceiver }} {{ .MethodType }}) Func(name, usage string, fn func(string) error) {
	{{ .MethodReceiver }}.Var(option.Func(fn), name, usage)
}


