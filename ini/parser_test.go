package ini

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	type testCase struct {
		name   string
		input  string
		output [][2]string
		err    string
	}

	testCases := []testCase{
		{
			name:   "EmptyLine",
			input:  "",
			output: [][2]string{},
		},
		{
			name:  "NoValue",
			input: "foo=",
			output: [][2]string{
				{"foo", ""},
			},
		},
		{
			name:  "Compact",
			input: "foo=bar",
			output: [][2]string{
				{"foo", "bar"},
			},
		},
		{
			name:  "SingleSpace",
			input: "     \t   foo  = bar",
			output: [][2]string{
				{"foo", "bar"},
			},
		},
		{
			name: "StringUnquoted",
			input: `str = a ; \
`,
			output: [][2]string{
				{"str", "a"},
			},
		},
		{
			name:  "MultipleWhiteSpaces",
			input: "     \t   foo   =    \tbar",
			output: [][2]string{
				{"foo", "bar"},
			},
		},
		{
			name:  "MultipleWhiteSpacesQuotedValue",
			input: "     \t   foo   =    \t \"bar\"",
			output: [][2]string{
				{"foo", "bar"},
			},
		},
		{
			name: "MultiLineStringUnquoted",
			input: `str = a \
b ; foo`,
			output: [][2]string{
				{"str", "a \nb"},
			},
		},
		{
			name:  "MultiLineStringDoubleQuote",
			input: `str = "1\n2"`,
			output: [][2]string{
				{"str", "1\n2"},
			},
		},
		{
			name:  "MultiLineStringBackQuote",
			input: "str = `1\n2\n`",
			output: [][2]string{
				{"str", "1\n2\n"},
			},
		},
		{
			name:  "MultipleWhiteSpacesQuotedValueInlineComment",
			input: "     \t   foo   =    \t \"bar\" ;foo",
			output: [][2]string{
				{"foo", "bar"},
			},
		},
		{
			name:  "EscapedQuote",
			input: `foo=\"`,
			output: [][2]string{
				{"foo", `\"`},
			},
		},
		{
			name:  "MutlipleSections",
			input: "[section]\n  foo=bar\n2.foo=baz\n[.inner]\nbaz=qux\n[]\nroot=true",
			output: [][2]string{
				{"section.foo", "bar"},
				{"section.2.foo", "baz"},
				{"section.inner.baz", "qux"},
				{"root", "true"},
			},
		},
		{
			name:  "InvalidSection",
			input: "[.]",
			output: [][2]string{
				{},
			},
			err: "invalid section at 1:4",
		},
		{
			name:  "InvalidString",
			input: "foo = `baz` value after end of string",
			output: [][2]string{
				{},
			},
			err: "invalid content after value at 1:13",
		},
	}

	for _, tcase := range testCases {
		t.Run(tcase.name, func(t *testing.T) {
			var (
				k, v string
				err  error
			)

			p := newParser(strings.NewReader(tcase.input))
			for _, out := range tcase.output {
				k, v, err = p.parseNext()
				if tcase.err != "" {
					if err == nil {
						t.Fatalf("error nil doesn't match expected error %q", tcase.err)
					}

					if err.Error() == tcase.err {
						return
					} else {
						t.Fatalf("error %q doesn't match expected %q", err, tcase.err)
					}
				} else {
					if err != nil {
						t.Fatalf("unexpected error %q", err)
					}

					if k != out[0] {
						t.Fatalf("key %q doesn't match expected %q", k, out[0])
					}
					if v != out[1] {
						t.Fatalf("value %q doesn't match expected %q", v, out[1])
					}
				}

			}

			k, v, err = p.parseNext()
			if k != "" || v != "" || err != nil {
				t.Fatalf("parser continue to return data: k=%q v=%q err=%q", k, v, err)
			}
		})
	}
}

func FuzzParser(f *testing.F) {
	f.Add("[section]\nkey=value\n")
	f.Add("[section]\nkey1=value1\nkey2=value2\n")

	f.Fuzz(func(t *testing.T, val string) {
		p := newParser(strings.NewReader("foo=" + val))
		_, _, _ = p.parseNext()
	})
}
