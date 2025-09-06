package ini

import (
	"io"
	"math"
	"os"
	"strings"
	"testing"
	"time"
)

func TestPropSet(t *testing.T) {
	t.Run("Primitives", func(t *testing.T) {
		t.Run("Parse", func(t *testing.T) {
			t.Run("Default", func(t *testing.T) {
				var ps PropSet
				b := ps.Bool("bool", true, "a bool property")
				d := ps.Duration("dur", time.Second, "a duration property")
				f := ps.Float64("float", math.Pi, "a float property")
				i := ps.Int("int", -1234, "an int property")
				i64 := ps.Int64("int64", -12345, "an int64 property")
				s := ps.String("string", "foo;bar", "a string property")
				u := ps.Uint("uint", 1234, "an uint property")
				u64 := ps.Uint64("uint64", 12345, "an uint64 property")

				if ps.Parsed() {
					t.Fatal("PropSet.Parsed() should return false")
				}

				err := ps.Parse(nil)
				if err != nil {
					t.Fatal("unexpected parse error")
				}
				if !ps.Parsed() {
					t.Fatal("PropSet.Parsed() should return true")
				}

				if *b != true || *d != time.Second || *f != math.Pi ||
					*i != -1234 || *i64 != -12345 || *s != "foo;bar" ||
					*u != 1234 || *u64 != 12345 {
					t.Fatal("unexpected value")
				}
			})

			t.Run("Success", func(t *testing.T) {
				var ps PropSet
				b := ps.Bool("bool", true, "a bool property")
				d := ps.Duration("dur", time.Second, "a duration property")
				f := ps.Float64("float", math.Pi, "a float property")
				i := ps.Int("int", -1234, "an int property")
				i64 := ps.Int64("int64", -12345, "an int64 property")
				s := ps.String("string", "foo", "a string property")
				u := ps.Uint("uint", 1234, "an uint property")
				u64 := ps.Uint64("uint64", 12345, "an uint64 property")

				err := ps.Parse(strings.NewReader(`
	bool = f	
			bool = "false"
	dur = 0.5s  
float=  1.23 		
int=-1 	
int64=-2 	
string=     bar \
foo
uint	=	100   
uint64=101`))
				if err != nil {
					t.Fatal("unexpected parse error")
				}
				if !ps.Parsed() {
					t.FailNow()
				}

				if *b != false || *d != time.Second/2 || *f != 1.23 ||
					*i != -1 || *i64 != -2 || *s != "bar \nfoo" ||
					*u != 100 || *u64 != 101 {

					ps.VisitAll(func(p *Property) {
						println(p.Name, p.Value.String())
					})
					t.Fatal("unexpected value")
				}
			})

			t.Run("Error", func(t *testing.T) {
				t.Run("UndefinedProperty", func(t *testing.T) {
					var ps PropSet

					ps.SetOutput(io.Discard)

					// bool is not defined.
					err := ps.Parse(strings.NewReader(`bool=false`))
					if err == nil ||
						!strings.Contains(err.Error(),
							"property provided but not defined") {
						t.Fatal("error doesn't match expected:", err)
					}
				})

				t.Run("ParseError", func(t *testing.T) {
					var ps PropSet
					_ = ps.Bool("bool", true, "a bool property")
					_ = ps.Duration("dur", time.Second, "a duration property")
					_ = ps.Float64("float", math.Pi, "a float property")
					_ = ps.Int("int", -1234, "an int property")
					_ = ps.Int64("int64", -12345, "an int64 property")
					_ = ps.String("string", "foo", "a string property")
					_ = ps.Uint("uint", 1234, "an uint property")
					_ = ps.Uint64("uint64", 12345, "an uint64 property")

					ps.SetOutput(io.Discard)
					err := ps.Parse(strings.NewReader(`
	bool = t	
			bool = "true"
	dur = 0.5s  
float=  1.23 		
int=-1 	
int64=-2 	
string=     bar \
foo
uint	=	-100   ; negative
uint64=101`))
					if err == nil ||
						!strings.Contains(err.Error(),
							"invalid value \"-100\" for property uint: parse error") {
						t.Fatal("error doesn't match expected:", err)
					}
				})
			})

			t.Run("Name", func(t *testing.T) {
				var ps PropSet
				if ps.Name() != "" {
					t.Fatal("unexpected name")
				}
				ps.Init("foo", ContinueOnError)
				if ps.Name() != "foo" {
					t.Fatal("unexpected name")
				}
			})

			t.Run("Output", func(t *testing.T) {
				var ps PropSet
				if ps.Output() != os.Stderr {
					t.Fatal("unexpected output")
				}
				ps.SetOutput(io.Discard)
				if ps.Output() != io.Discard {
					t.Fatal("unexpected output")
				}
			})

			t.Run("PrintDefaults", func(t *testing.T) {
				var ps PropSet
				_ = ps.Bool("abc.bool", true, "a bool property")
				_ = ps.Duration("abc.def.dur", time.Second, "a duration property")
				_ = ps.Float64("abc.def.float", math.Pi, "a float property")
				_ = ps.Int("def.int", -1234, "an int property")
				_ = ps.Int64("int64", -12345, "an `big number` property")
				_ = ps.String("string", "foo", "a string property")
				_ = ps.Uint("uint", 1234, "an uint property")
				_ = ps.Uint64("uint64", 12345, "an uint64 property")

				b := strings.Builder{}
				ps.SetOutput(&b)
				ps.PrintDefaults()

				expectedOutput := `[abc]
; a bool property
  bool = true

[.def]
; a duration property
; dur = duration
  dur = 1s

; a float property
; float = float
  float = 3.141592653589793

[def]
; an int property
; int = int
  int = -1234

[]
; an big number property
; int64 = big number
int64 = -12345

; a string property
; string = string
string = foo

; an uint property
; uint = uint
uint = 1234

; an uint64 property
; uint64 = uint
uint64 = 12345

`
				if b.String() != expectedOutput {
					t.Log(b.String())
					t.Fatal("defaults doesn't match expected")
				}
			})

			t.Run("Lookup", func(t *testing.T) {
				var ps PropSet

				_ = ps.Bool("bool", true, "a bool property")
				d := ps.Duration("dur", time.Second, "a duration property")
				_ = ps.Float64("float", math.Pi, "a float property")

				dur := ps.Lookup("dur")
				if dur.Name != "dur" || dur.DefValue != "1s" ||
					dur.Usage != "a duration property" || dur.Value.String() != "1s" {
					t.Fatal("looked up value doesn't match expected")
				}

				*d = 2 * time.Second
				if dur.Value.String() != "2s" {
					t.Fatal("looked up value doesn't match expected")
				}
			})

			t.Run("TextVar", func(t *testing.T) {
				t.Run("Success", func(t *testing.T) {
					var ps PropSet
					var mock mockTextVar
					ps.TextVar(&mock, "text", &mock, "a text property")
					err := ps.Parse(strings.NewReader("text=foo"))
					if err != nil {
						t.Fatal("unexpected parse error")
					}
					if mock.str != "foo" {
						t.Fatal("unexpected text value")
					}
				})

				t.Run("ParseError", func(t *testing.T) {
					var ps PropSet
					var mock mockTextVar
					mock.err = io.EOF

					ps.Init("", ContinueOnError)
					ps.SetOutput(io.Discard)

					ps.TextVar(&mock, "text", &mock, "a text property")
					err := ps.Parse(strings.NewReader("text=foo"))
					if err == nil || !strings.Contains(err.Error(), mock.err.Error()) {
						t.Fatal("unexpected parse error")
					}
					if mock.str != "" {
						t.Fatal("unexpected text value")
					}
				})
			})
		})
	})
}

type mockTextVar struct {
	str string
	err error
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (m *mockTextVar) UnmarshalText(text []byte) error {
	if m.err != nil {
		return m.err
	}

	m.str = string(text)
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (m *mockTextVar) MarshalText() (text []byte, err error) {
	if m.err != nil {
		return nil, err
	}
	return []byte(m.str), nil
}
