package option

import (
	"encoding"
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// Option define a configuration option.
type Option struct {
	Name     string
	Usage    string
	Value    Value
	DefValue string
}

// Value is the interface to the dynamic value stored in an option. (The default
// value is represented as a string.)
type Value = flag.Value

// Getter is an interface that allows the contents of a Value to be retrieved.
// It wraps the Value interface, rather than being part of it, because it
// appeared after Go 1 and its compatibility rules. All Value types provided by
// this package satisfy the Getter interface, except the type used by Func.
type Getter = flag.Getter

type Bool bool

// NewBool creates a new boolean value.
func NewBool(val bool, p *bool) *Bool {
	*p = val
	return (*Bool)(p)
}

// Set implements Value.
func (b *Bool) Set(s string) error {
	v, err := strconv.ParseBool(s)
	if err != nil {
		err = errParse
	}
	*b = Bool(v)
	return err
}

// Get implements Getter.
func (b *Bool) Get() any { return bool(*b) }

// String implements Value.
func (b *Bool) String() string { return strconv.FormatBool(bool(*b)) }

// IsBoolFlag always return true.
func (b *Bool) IsBoolFlag() bool { return true }

type Duration time.Duration

// NewDuration creates a new boolean value.
func NewDuration(val time.Duration, p *time.Duration) *Duration {
	*p = val
	return (*Duration)(p)
}

func (d *Duration) Set(s string) error {
	v, err := time.ParseDuration(s)
	if err != nil {
		err = errParse
	}
	*d = Duration(v)
	return err
}

// Get implements Getter.
func (d *Duration) Get() any { return time.Duration(*d) }

// String implements Value.
func (d *Duration) String() string { return (*time.Duration)(d).String() }

type Float64 float64

// NewFloat64 creates a new float64 value.
func NewFloat64(val float64, p *float64) *Float64 {
	*p = val
	return (*Float64)(p)
}

// Set implements Value.
func (f *Float64) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		err = numError(err)
	}
	*f = Float64(v)
	return err
}

// Get implements Getter.
func (f *Float64) Get() any { return float64(*f) }

// String implements Value.
func (f *Float64) String() string { return strconv.FormatFloat(float64(*f), 'g', -1, 64) }

type Int int

// NewInt creates a new int value.
func NewInt(val int, p *int) *Int {
	*p = val
	return (*Int)(p)
}

// Set implements Value.
func (i *Int) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		err = numError(err)
	}
	*i = Int(v)
	return err
}

// Get implements Getter.
func (i *Int) Get() any { return int(*i) }

// String implements Value.
func (i *Int) String() string { return strconv.Itoa(int(*i)) }

type Int64 int64

// NewInt64 creates a new int64 value.
func NewInt64(val int64, p *int64) *Int64 {
	*p = val
	return (*Int64)(p)
}

// Set implements Value.
func (i *Int64) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		err = numError(err)
	}
	*i = Int64(v)
	return err
}

// Get implements Getter.
func (i *Int64) Get() any { return int64(*i) }

// String implements Value.
func (i *Int64) String() string { return strconv.FormatInt(int64(*i), 10) }

type String string

// NewString creates a new string value.
func NewString(val string, p *string) *String {
	*p = val
	return (*String)(p)
}

// Set implements Value.
func (s *String) Set(val string) error {
	*s = String(val)
	return nil
}

// Get implements Getter.
func (s *String) Get() any { return string(*s) }

// String implements Value.
func (s *String) String() string { return string(*s) }

type Text struct{ p encoding.TextUnmarshaler }

// NewText creates a new encoding.TextMarshaler/encoding.TextUnmarshaler value.
func NewText(val encoding.TextMarshaler, p encoding.TextUnmarshaler) Text {
	ptrVal := reflect.ValueOf(p)
	if ptrVal.Kind() != reflect.Ptr {
		panic("variable value type must be a pointer")
	}
	defVal := reflect.ValueOf(val)
	if defVal.Kind() == reflect.Ptr {
		defVal = defVal.Elem()
	}
	if defVal.Type() != ptrVal.Type().Elem() {
		panic(fmt.Sprintf("default type does not match variable type: %v != %v", defVal.Type(), ptrVal.Type().Elem()))
	}
	// Set implements Value.
	ptrVal.Elem().Set(defVal)
	return Text{p}
}

// Set implements Value.
func (v Text) Set(s string) error {
	return v.p.UnmarshalText([]byte(s))
}

// Get implements Getter.
func (v Text) Get() any {
	return v.p
}

// String implements Value.
func (v Text) String() string {
	if m, ok := v.p.(encoding.TextMarshaler); ok {
		if b, err := m.MarshalText(); err == nil {
			return string(b)
		}
	}
	return ""
}

type Uint uint

// NewUint creates a new uint value.
func NewUint(val uint, p *uint) *Uint {
	*p = val
	return (*Uint)(p)
}

// Set implements Value.
func (u *Uint) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	if err != nil {
		err = numError(err)
	}
	*u = Uint(v)
	return err
}

// Get implements Getter.
func (u *Uint) Get() any { return uint(*u) }

// String implements Value.
func (u *Uint) String() string { return strconv.FormatUint(uint64(*u), 10) }

type Uint64 uint64

// NewUint64 creates a new uint64 value.
func NewUint64(val uint64, p *uint64) *Uint64 {
	*p = val
	return (*Uint64)(p)
}

// Set implements Value.
func (i *Uint64) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		err = numError(err)
	}
	*i = Uint64(v)
	return err
}

// Get implements Getter.
func (i *Uint64) Get() any { return uint64(*i) }

// String implements Value.
func (i *Uint64) String() string { return strconv.FormatUint(uint64(*i), 10) }

// Func implements Value.
type Func func(string) error

func (f Func) Set(s string) error { return f(s) }

func (f Func) String() string { return "" }
