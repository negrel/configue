package configue

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/negrel/configue/env"
	"github.com/negrel/configue/option"
)

// Backend define options registry and parser.
type Backend interface {
	Init(name string, errorHandling ErrorHandling)
	Var(val Value, name, usage string)
	Parse() error
	Parsed() bool
	Set(name, value string) error
	Visit(fn func(string, option.Value))
	PrintDefaults()
	SetOutput(io.Writer)
}

type Value interface {
	fmt.Stringer
	Set(string) error
}

var _ Backend = &envBackend{}
var _ Backend = &flagBackend{}

type envBackend struct {
	*env.EnvSet
	prefix  string
	nameMap map[string]string
}

// NewEnv returns a new environment variable based backend.
func NewEnv(prefix string) Backend {
	if prefix != "" && !strings.HasSuffix(prefix, "_") {
		prefix += "_"
	}

	eb := &envBackend{
		EnvSet:  env.NewEnvSet("", ContinueOnError),
		prefix:  prefix,
		nameMap: make(map[string]string),
	}
	eb.Usage = func() {
		if name := eb.Name(); name != "" {
			_, _ = fmt.Fprintf(eb.Output(), "Environment variables of %v:\n", name)
		} else {
			_, _ = fmt.Fprintln(eb.Output(), "Environment variables:")
		}
		eb.EnvSet.PrintDefaults()
	}
	return eb
}

func (ev *envBackend) envName(name string) string {
	// Convert "OPTION.path" to "OPTION_PATH" env var.
	path := strings.Split(name, ".")
	return strings.ToUpper(ev.prefix + strings.Join(path, "_"))
}

// Var implements Backend.
func (eb *envBackend) Var(val Value, name, usage string) {
	envName := eb.envName(name)
	eb.EnvSet.Var(val, envName, usage)
	eb.nameMap[envName] = name
}

// Set sets the value of the named command-line option.
func (eb *envBackend) Set(name, value string) error {
	envName := eb.envName(name)
	return eb.EnvSet.Set(envName, value)
}

// Parse implements Backend.
func (eb *envBackend) Parse() error {
	return eb.EnvSet.Parse(os.Environ())
}

// Visit implements Backend.
func (eb *envBackend) Visit(fn func(string, option.Value)) {
	eb.EnvSet.Visit(func(envVar *env.EnvVar) {
		fn(eb.nameMap[envVar.Name], envVar.Value)
	})
}

// PrintDefaults implements Backend.
func (eb *envBackend) PrintDefaults() {
	eb.Usage()
}

type flagBackend struct {
	*flag.FlagSet
	nameMap map[string]string
}

// NewFlag returns a new flag based backend.
func NewFlag() Backend {
	fb := &flagBackend{flag.NewFlagSet("", ContinueOnError), make(map[string]string)}
	fb.Usage = func() {
		if name := fb.Name(); name != "" {
			_, _ = fmt.Fprintf(fb.Output(), "Flags of %v:\n", name)
		} else {
			_, _ = fmt.Fprintln(fb.Output(), "Flags:")
		}
		fb.FlagSet.PrintDefaults()
	}
	return fb
}

func (fb *flagBackend) flagName(name string) string {
	// Convert "OPTION.path" to "option-path" flag.
	path := strings.Split(strings.ToLower(name), ".")
	return strings.Join(path, "-")
}

// Var implements Backend.
func (fb *flagBackend) Var(val Value, name, usage string) {
	flagName := fb.flagName(name)
	fb.FlagSet.Var(val, flagName, usage)
	fb.nameMap[flagName] = name
}

// Set sets the value of the named command-line option.
func (fb *flagBackend) Set(name, value string) error {
	flagName := fb.flagName(name)
	return fb.FlagSet.Set(flagName, value)
}

// Parse implements Backend.
func (fb *flagBackend) Parse() error {
	return fb.FlagSet.Parse(os.Args[1:])
}

// Visit implements Backend.
func (fb *flagBackend) Visit(fn func(string, option.Value)) {
	fb.FlagSet.Visit(func(flag *flag.Flag) {
		fn(fb.nameMap[flag.Name], flag.Value)
	})
}

// PrintDefaults implements Backend.
func (fb *flagBackend) PrintDefaults() {
	fb.Usage()
}
