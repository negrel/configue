/*
Package ini implements INI properties parsing.

# Usage

Define properties using [ini.String], [Bool], [Int], etc.

This declares an integer property, N, stored in the pointer nProp, with type *int:

	import "github.com/negrel/configue/ini"
	var nProp = ini.Int("N", 1234, "help message for INI property N")

If you like, you can bind the property to a variable using the Var() functions.

	var prop int
	func init() {
		ini.IntVar(&prop, "PROPNAME", 1234, "help message for PROPNAME")
	}

Or you can create custom properties that satisfy the Value interface (with
pointer receivers) and couple them to property parsing by

	ini.Var(&prop, "PROPNAME", "help message for PROPNAME")

For such properties, the default value is just the initial value of the variable.

After all properties are defined, call

	ini.Parse(r) // r is an io.Reader.

to parse the command line into the defined properties.

Properties may then be used directly. If you're using the properties themselves,
they are all pointers; if you bind to variables, they're values.

	fmt.Println("ip has value ", *ip)
	fmt.Println("property has value ", prop)

# Command line property syntax

See [`option`](../option#pkg-overview) documentation for options syntax of basic
types.

The default set of command-line properties is controlled by top-level functions.
The [PropSet] type allows one to define independent sets of properties.
*/
package ini
