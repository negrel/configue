/*
Package env implements command-line environment variables parsing.

# Usage

Define env vars using [env.String], [Bool], [Int], etc.

This declares an integer env variable, N, stored in the pointer nEnvVar, with type *int:

	import "github.com/negrel/configue/env"
	var nEnvVar = env.Int("N", 1234, "help message for env var N")

If you like, you can bind the env var to a variable using the Var() functions.

	var envVar int
	func init() {
		env.IntVar(&envVar, "ENVNAME", 1234, "help message for ENVNAME")
	}

Or you can create custom env vars that satisfy the Value interface (with
pointer receivers) and couple them to env var parsing by

	env.Var(&envVar, "ENVNAME", "help message for ENVNAME")

For such env vars, the default value is just the initial value of the variable.

After all env vars are defined, call

	env.Parse()

to parse the command line into the defined env vars.

Env vars may then be used directly. If you're using the env vars themselves,
they are all pointers; if you bind to variables, they're values.

	fmt.Println("ip has value ", *ip)
	fmt.Println("envVar has value ", envVar)

# Command line env var syntax

See [`option`](../option#pkg-overview) documentation for options syntax of basic
types.

The default set of command-line env vars is controlled by top-level functions.
The [EnvSet] type allows one to define independent sets of env vars, such as to
implement subcommands in a command-line interface. The methods of [EnvSet] are
analogous to the top-level functions for the command-line env var set.
*/
package env
