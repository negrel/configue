/*
Package configue facilitate parsing command-line options from multiple sources
(environments variables, flag, INI files, etc).

# Usage

Define options using [configue.String], [Bool], [Int], etc.

This declares an integer option, n, stored in the pointer n, with type *int:

	import "github.com/negrel/configue"
	var n = configue.Int("n", 1234, "help message for option n")

If you like, you can bind the option to a variable using the Var() functions.

	var n int
	func init() {
		configue.IntVar(&n, "n", 1234, "help message for n")
	}

Or you can create custom options that satisfy the Value interface (with
pointer receivers) and couple them to options parsing by

	configue.Var(&n, "n", "help message for n")

For such options, the default value is just the initial value of the variable.

After all options are defined, call

	configue.Parse()

to parse the command line into the defined options.

Options may then be used directly. If you're using the options themselves,
they are all pointers; if you bind to variables, they're values.

	fmt.Println("ip has value ", *ip)
	fmt.Println("n has value ", n)

# Command line option syntax

Options are loaded/parsed by [Backend]. Built-in flag and environment variable
based backends are provided by [NewFlag] and [NewEnv] respectively. They parse
options value the same way. See [`option`](./option#pkg-overview) documentation
for more information.
*/
package configue
