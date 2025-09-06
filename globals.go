package configue

import (
	"flag"
	"fmt"
	"os"
	"path"
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
	// CommandLine is the default set of command-line options, parsed from
	// environments variable and flags. The top-level functions such as BoolVar,
	// and so on are wrappers for the methods of CommandLine.
	CommandLine = New("", ExitOnError, NewINI(File("./", "config.ini")), NewEnv(""), NewFlag())
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

	// ErrHelp is the error returned if the -help or -h flag is invoked but no
	// such flag is defined.
	ErrHelp = flag.ErrHelp
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
	return CommandLine.Parse()
}

// UserDir returns user configuration directory or provided default on error
// (e.g. when os.UserConfigDir fails).
func UserDir(def string) string {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return def
	}
	return cfgDir
}

// AppDir returns application configuration directory which is [UserDir] joined
// with os.Args[0].
func AppDir(defaultUserDir string) string {
	return path.Join(UserDir(defaultUserDir), path.Base(os.Args[0]))
}

// File returns path to configuration file by joining [AppDir] with provided
// `fname`.
func File(defaultUserDir, fname string) string {
	return path.Join(AppDir(defaultUserDir), fname)
}
