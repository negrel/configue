/*
Package option provides flag.Value implementation for basic types.

# Option syntax

Integer options accept 1234, 0664, 0x1234 and may be negative.
Boolean options may be:

	1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False

Duration options accept any input valid for time.ParseDuration.
*/
package option
