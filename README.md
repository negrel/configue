<h1 align="center">
    <img height="250" alt="cute fig" src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fi.pinimg.com%2Foriginals%2Fb0%2Fb7%2F52%2Fb0b752d332e6e81e8dd7ed172aeefcd9.jpg&f=1&nofb=1&ipt=a52f6c53be292db57237cd5c7379525ee5b4146ff1765037c0bfea1cbb744b05">
</h1>

# `configue` - Configuration library for Go

[![Go doc](https://pkg.go.dev/badge/github.com/negrel/configue)](https://pkg.go.dev/github.com/negrel/configue)
[![go report card](https://goreportcard.com/badge/github.com/negrel/configue)](https://goreportcard.com/report/github.com/negrel/configue)
[![license card](https://img.shields.io/github/license/negrel/configue)](./LICENSE)
[![PRs welcome card](https://img.shields.io/badge/PRs-Welcome-brightgreen)](https://github.com/negrel/configue/pulls)
![Go version card](https://img.shields.io/github/go-mod/go-version/negrel/configue)

`configue` is a simple, extensible and dependency-free configuration library for
Go. It implements an API similar to `flag` package from standard library to load
and merge options from multiple sources.

Flags, environment variables and INI files are built-in but you can add custom
sources by implementing the `Backend` interface.

## Getting started

Here is a simple example:

```go
package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/negrel/configue"
)

type Option struct {
	Debug   bool
	MaxProc int
}

func main() {
	iniFile, err := os.Open("./config.ini")
	if err != nil {
		// handle error.
	}

	// Create backends.
	env := configue.NewEnv("MYAPP")
	flag := configue.NewFlag()

	figue := configue.New(
		"",                       // Subcommand name.
		configue.ContinueOnError, // Error handling strategy.
		configue.NewINI(iniFile), // INI file backend.
		env,                      // Environment variable backend with MYAPP_ prefix.
		flag,                     // Go's std `flag` backend.
	)

	// Custom usage.
	figue.Usage = func() {
		_, _ = fmt.Fprintln(figue.Output(), "myapp - a great app")
		_, _ = fmt.Fprintln(figue.Output())
		_, _ = fmt.Fprintln(figue.Output(), "Usage:")
		_, _ = fmt.Fprintln(figue.Output(), "  myapp [flags]")
		_, _ = fmt.Fprintln(figue.Output())
		flag.PrintDefaults()
		_, _ = fmt.Fprintln(figue.Output())
		env.PrintDefaults()
		// We don't print INI defaults.
	}

	// Define options.
	var options Option
	figue.BoolVar(&options.Debug, "debug", false, "enable debug logs")
	figue.IntVar(&options.MaxProc, "max.proc", runtime.NumCPU(), "maximum number of CPU that can be executed simultaneously")

	// Parse options.
	err = figue.Parse()
	if errors.Is(err, configue.ErrHelp) {
		return
	}
	if err != nil {
		// handle error
	}
}
```

```sh
$ myapp -h
myapp - a great app

Usage:
  myapp [flags]

Flags:
  -debug
        enable debug logs
  -max-proc value
        maximum number of CPU that can be executed simultaneously (default 16)

Environment variables:
  MYAPP_DEBUG
        enable debug logs
  MYAPP_MAX_PROC int
        maximum number of CPU that can be executed simultaneously (default 16)
```

For a real example, see [Prisme Analytics](https://github.com/prismelabs/analytics/blob/e6522e6502fef0ceb3f5df79f17a6a3b4b70ba02/cmd/prisme/main.go#L42-L98)
configuration loading.

## Contributing

If you want to contribute to `configue` to add a feature or improve the code contact
me at [alexandre@negrel.dev](mailto:alexandre@negrel.dev), open an
[issue](https://github.com/negrel/configue/issues) or make a
[pull request](https://github.com/negrel/configue/pulls).

## :stars: Show your support

Please give a :star: if this project helped you!

[![buy me a coffee](https://github.com/negrel/.github/blob/master/.github/images/bmc-button.png?raw=true)](https://www.buymeacoffee.com/negrel)

## :scroll: License

MIT © [Alexandre Negrel](https://www.negrel.dev/)
