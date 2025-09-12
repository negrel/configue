package env

import (
	"io"
	"math"
	"os"
	"strings"
	"testing"
	"time"
)

func TestEnvSet(t *testing.T) {
	t.Run("Primitives", func(t *testing.T) {
		t.Run("Parse", func(t *testing.T) {
			t.Run("Default", func(t *testing.T) {
				var es EnvSet
				b := es.Bool("bool", true, "a bool env var")
				d := es.Duration("dur", time.Second, "a duration env var")
				f := es.Float64("float", math.Pi, "a float env var")
				i := es.Int("int", -1234, "an int env var")
				i64 := es.Int64("int64", -12345, "an int64 env var")
				s := es.String("string", "foo", "a string env var")
				u := es.Uint("uint", 1234, "an uint env var")
				u64 := es.Uint64("uint64", 12345, "an uint64 env var")
				u64s := es.Uint64Slice("uint64.slice", []uint64{12345}, "an uint64 slice")

				if es.Parsed() {
					t.Fatal("EnvSet.Parsed() should return false")
				}

				err := es.Parse(nil)
				if err != nil {
					t.Fatal("unexpected parse error")
				}
				if !es.Parsed() {
					t.Fatal("EnvSet.Parsed() should return true")
				}

				if *b != true || *d != time.Second || *f != math.Pi ||
					*i != -1234 || *i64 != -12345 || *s != "foo" ||
					*u != 1234 || *u64 != 12345 || (*u64s)[0] != 12345 {
					t.Fatal("unexpected value")
				}
			})

			t.Run("Success", func(t *testing.T) {
				var es EnvSet
				b := es.Bool("bool", true, "a bool env var")
				d := es.Duration("dur", time.Second, "a duration env var")
				f := es.Float64("float", math.Pi, "a float env var")
				i := es.Int("int", -1234, "an int env var")
				i64 := es.Int64("int64", -12345, "an int64 env var")
				s := es.String("string", "foo", "a string env var")
				u := es.Uint("uint", 1234, "an uint env var")
				u64 := es.Uint64("uint64", 12345, "an uint64 env var")
				u64s := es.Uint64Slice("uint64_slice", []uint64{12345}, "an uint64 slice")

				err := es.Parse([]string{
					"bool=f",
					"dur=0.5s",
					"float=1.23",
					"int=-1",
					"int64=-2",
					"string=bar",
					"uint=100",
					"uint64=101",
					"uint64_slice=1,2,56,239232",
				})
				if err != nil {
					t.Fatal("unexpected parse error")
				}
				if !es.Parsed() {
					t.FailNow()
				}

				if *b != false || *d != time.Second/2 || *f != 1.23 ||
					*i != -1 || *i64 != -2 || *s != "bar" ||
					*u != 100 || *u64 != 101 ||
					(*u64s)[0] != 1 || (*u64s)[1] != 2 || (*u64s)[2] != 56 ||
					(*u64s)[3] != 239232 {
					t.Fatal("unexpected value")
				}
			})

			t.Run("Error", func(t *testing.T) {
				t.Run("EmptyStringEnvVarName", func(t *testing.T) {
					var es EnvSet

					es.Bool("", false, "")
				})

				t.Run("UndefinedEnvVar", func(t *testing.T) {
					var es EnvSet

					es.SetOutput(io.Discard)

					// bool is not defined.
					IgnoreUndefined = false
					err := es.Parse([]string{"bool=f"})
					if err == nil ||
						!strings.Contains(err.Error(),
							"env var provided but not defined") {
						t.Fatal("error doesn't match expected:", err)
					}
					IgnoreUndefined = true

					err = es.Parse([]string{"bool=f"})
					if err != nil {
						t.Fatal("unexpected parse error", err)

					}
				})

				t.Run("ParseError", func(t *testing.T) {
					var es EnvSet
					_ = es.Bool("bool", true, "a bool env var")
					_ = es.Duration("dur", time.Second, "a duration env var")
					_ = es.Float64("float", math.Pi, "a float env var")
					_ = es.Int("int", -1234, "an int env var")
					_ = es.Int64("int64", -12345, "an int64 env var")
					_ = es.String("string", "foo", "a string env var")
					_ = es.Uint("uint", 1234, "an uint env var")
					_ = es.Uint64("uint64", 12345, "an uint64 env var")

					es.SetOutput(io.Discard)
					err := es.Parse([]string{
						"bool=f",
						"dur=0.5s",
						"float=1.23",
						"int=-1",
						"int64=-2",
						"string=bar",
						"uint=-100", // Negative uint
						"uint64=101",
					})

					if err == nil ||
						!strings.Contains(err.Error(),
							"invalid value \"-100\" for env var uint: parse error") {
						t.Fatal("error doesn't match expected:", err)
					}
				})
			})
		})

		t.Run("Name", func(t *testing.T) {
			var es EnvSet
			if es.Name() != "" {
				t.Fatal("unexpected name")
			}
			es.Init("foo", ContinueOnError)
			if es.Name() != "foo" {
				t.Fatal("unexpected name")
			}
		})

		t.Run("Output", func(t *testing.T) {
			var es EnvSet
			if es.Output() != os.Stderr {
				t.Fatal("unexpected output")
			}
			es.SetOutput(io.Discard)
			if es.Output() != io.Discard {
				t.Fatal("unexpected output")
			}
		})

		t.Run("PrintDefaults", func(t *testing.T) {
			var es EnvSet
			_ = es.Bool("bool", true, "a bool env var")
			_ = es.Duration("dur", time.Second, "a duration env var")
			_ = es.Float64("float", math.Pi, "a float env var")
			_ = es.Int("int", -1234, "an int env var")
			_ = es.Int64("int64", -12345, "an `big number` env var")
			_ = es.String("string", "foo", "a string env var")
			_ = es.Uint("uint", 1234, "an uint env var")
			_ = es.Uint64("uint64", 12345, "an uint64 env var")

			b := strings.Builder{}
			es.SetOutput(&b)
			es.PrintDefaults()

			expectedOutput := `  bool
    	a bool env var (default true)
  dur duration
    	a duration env var (default 1s)
  float float
    	a float env var (default 3.141592653589793)
  int int
    	an int env var (default -1234)
  int64 big number
    	an big number env var (default -12345)
  string string
    	a string env var (default "foo")
  uint uint
    	an uint env var (default 1234)
  uint64 uint
    	an uint64 env var (default 12345)
`

			if b.String() != expectedOutput {
				t.Fatal("defaults doesn't match expected")
			}
		})

		t.Run("Lookup", func(t *testing.T) {
			var es EnvSet

			_ = es.Bool("bool", true, "a bool env var")
			d := es.Duration("dur", time.Second, "a duration env var")
			_ = es.Float64("float", math.Pi, "a float env var")

			dur := es.Lookup("dur")
			if dur.Name != "dur" || dur.DefValue != "1s" ||
				dur.Usage != "a duration env var" || dur.Value.String() != "1s" {
				t.Fatal("looked up value doesn't match expected")
			}

			*d = 2 * time.Second
			if dur.Value.String() != "2s" {
				t.Fatal("looked up value doesn't match expected")
			}
		})

		t.Run("TextVar", func(t *testing.T) {
			t.Run("Success", func(t *testing.T) {
				var es EnvSet
				var mock mockTextVar
				es.TextVar(&mock, "text", &mock, "a text env var")
				err := es.Parse([]string{"text=foo"})
				if err != nil {
					t.Fatal("unexpected parse error")
				}
				if mock.str != "foo" {
					t.Fatal("unexpected text value")
				}
			})

			t.Run("ParseError", func(t *testing.T) {
				var es EnvSet
				var mock mockTextVar
				mock.err = io.EOF

				es.Init("", ContinueOnError)
				es.SetOutput(io.Discard)

				es.TextVar(&mock, "text", &mock, "a text env var")
				err := es.Parse([]string{"text=foo"})
				if err == nil || !strings.Contains(err.Error(), mock.err.Error()) {
					t.Fatal("unexpected parse error")
				}
				if mock.str != "" {
					t.Fatal("unexpected text value")
				}
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
