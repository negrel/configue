package configue

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/negrel/configue/env"
	"github.com/negrel/configue/ini"
	"github.com/negrel/configue/option"
)

// Backend define options registry and parser.
type Backend interface {
	Init(name string)
	Var(val Value, name, usage string) string
	Parse() error
	Parsed() bool
	Set(name, value string) error
	PrintDefaults()
	SetOutput(io.Writer)
}

type Value interface {
	fmt.Stringer
	Set(string) error
}

var _ Backend = &Env{}
var _ Backend = &Flag{}
var _ Backend = &Ini{}

// Env defines an environment variables based backend.
type Env struct {
	*env.EnvSet
	prefix  string
	nameMap map[string]string
}

// NewEnv returns a new environment variable based Backend implementation.
func NewEnv(prefix string) *Env {
	if prefix != "" && !strings.HasSuffix(prefix, "_") {
		prefix += "_"
	}

	eb := &Env{
		EnvSet:  env.NewEnvSet("", ContinueOnError),
		prefix:  prefix,
		nameMap: make(map[string]string),
	}
	eb.Usage = func() {}
	return eb
}

func (env *Env) envName(name string) string {
	// Convert "OPTION.path" to "OPTION_PATH" env var.
	path := strings.Split(name, ".")
	return strings.ToUpper(env.prefix + strings.Join(path, "_"))
}

// Init implements Backend.
func (env *Env) Init(name string) {
	env.EnvSet.Init(name, flag.ContinueOnError)
}

// Var implements Backend.
func (env *Env) Var(val Value, name, usage string) string {
	envName := env.envName(name)
	env.EnvSet.Var(val, envName, usage)
	env.nameMap[envName] = name
	return envName
}

// Set sets the value of the named command-line option.
func (env *Env) Set(name, value string) error {
	envName := env.envName(name)
	return env.EnvSet.Set(envName, value)
}

// Parse implements Backend by parsing os.Environ().
func (env *Env) Parse() error {
	return env.EnvSet.Parse(os.Environ())
}

// Visit implements Backend.
func (eb *Env) Visit(fn func(option.Option)) {
	eb.EnvSet.Visit(func(envVar *env.EnvVar) {
		opt := *envVar
		opt.Name = eb.nameMap[envVar.Name]
		fn(opt)
	})
}

// PrintDefaults implements Backend.
func (env *Env) PrintDefaults() {
	if name := env.Name(); name != "" {
		_, _ = fmt.Fprintf(env.Output(), "Environment variables of %v:\n", name)
	} else {
		_, _ = fmt.Fprintln(env.Output(), "Environment variables:")
	}
	env.EnvSet.PrintDefaults()
}

// Flag defines a flag based Backend implementation.
type Flag struct {
	*flag.FlagSet
	nameMap map[string]string
}

// NewFlag returns a new flag based backend.
func NewFlag() *Flag {
	fb := &Flag{flag.NewFlagSet("", ContinueOnError), make(map[string]string)}
	fb.Usage = func() {}
	return fb
}

func (flag *Flag) flagName(name string) string {
	// Convert "OPTION.path" to "option-path" flag.
	path := strings.Split(strings.ToLower(name), ".")
	return strings.Join(path, "-")
}

// Init implements Backend.
func (flag *Flag) Init(name string) {
	flag.FlagSet.Init(name, ContinueOnError)
}

// Var implements Backend.
func (flag *Flag) Var(val Value, name, usage string) string {
	flagName := flag.flagName(name)
	flag.FlagSet.Var(val, flagName, usage)
	flag.nameMap[flagName] = name
	return flagName
}

// Set sets the value of the named command-line option.
func (flag *Flag) Set(name, value string) error {
	flagName := flag.flagName(name)
	return flag.FlagSet.Set(flagName, value)
}

// Parse implements Backend by parsing flags from os.Args[1:].
func (flag *Flag) Parse() error {
	return flag.FlagSet.Parse(os.Args[1:])
}

// Visit visits the flags in lexicographical order, calling fn for each. It
// visits only those flags that have been set.
func (fb *Flag) Visit(fn func(option.Option)) {
	fb.FlagSet.Visit(func(flag *flag.Flag) {
		opt := option.Option(*flag)
		opt.Name = fb.nameMap[flag.Name]
		fn(opt)
	})
}

// VisitAll visits the flags in lexicographical order, calling fn for each. It
// visits all flags, even those not set.
func (fb *Flag) VisitAll(fn func(option.Option)) {
	fb.FlagSet.VisitAll(func(flag *flag.Flag) {
		opt := option.Option(*flag)
		opt.Name = fb.nameMap[flag.Name]
		fn(opt)
	})
}

// PrintDefaults implements Backend.
func (flag *Flag) PrintDefaults() {
	if name := flag.Name(); name != "" {
		_, _ = fmt.Fprintf(flag.Output(), "Flags of %v:\n", name)
	} else {
		_, _ = fmt.Fprintln(flag.Output(), "Flags:")
	}
	flag.FlagSet.PrintDefaults()
}

// Ini defines an INI file based Backend implementation.
type Ini struct {
	*ini.PropSet
	FilePath string
}

// NewINI returns a new INI based backend that will parse data from provided
// filepath. If the file doesn't exist, this backend will parse nothing.
func NewINI(fpath string) *Ini {
	ib := &Ini{ini.NewPropSet("", ContinueOnError), fpath}
	ib.Usage = func() {}
	return ib
}

// Init implements Backend.
func (ini *Ini) Init(name string) {
	ini.PropSet.Init(name, ContinueOnError)
}

// Var implements Backend.
func (ini *Ini) Var(val Value, name, usage string) string {
	ini.PropSet.Var(val, name, usage)
	return name
}

// Set sets the value of the named command-line option.
func (ini *Ini) Set(name, value string) error {
	return ini.PropSet.Set(name, value)
}

// Parse implements Backend.
func (ini *Ini) Parse() error {
	f, err := os.Open(ini.FilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	return errors.Join(ini.PropSet.Parse(f), f.Close())
}

// Visit implements Backend.
func (ib *Ini) Visit(fn func(option.Option)) {
	ib.PropSet.Visit(func(prop *ini.Property) {
		fn(*prop)
	})
}

// PrintDefaults implements Backend. Unlike other backends, we only print path
// to config file here.
func (ini *Ini) PrintDefaults() {
	if ini.FilePath == "" {
		return
	}

	if name := ini.Name(); name != "" {
		_, _ = fmt.Fprintf(ini.Output(), "Configuration file of %v is located at %v\n", name, ini.FilePath)
	} else {
		_, _ = fmt.Fprintf(ini.Output(), "Configuration file is located at %v\n", ini.FilePath)
	}
}
