package ini

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	sectionRegex = regexp.MustCompile(`\.?\w+(.\w+)?`)
)

// parser defines a parser for the INI format.
type parser struct {
	scanner   *bufio.Scanner
	section   string
	line, col int
	buf       []byte
}

func newParser(r io.Reader) *parser {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	return &parser{scanner: scanner, section: "", line: 0}
}

func (p *parser) nextLine() bool {
	if p.scanner.Scan() {
		p.line++
		p.col = 0
		p.buf = p.scanner.Bytes()
	} else {
		p.buf = nil
		p.col = 0
	}

	return p.buf != nil
}

func (p *parser) trimSpace() {
	buf := p.bytes()
	trimmed := bytes.TrimLeftFunc(buf, unicode.IsSpace)
	p.skip(len(buf) - len(trimmed))
}

func (p *parser) bytes() []byte {
	return p.buf[p.col:]
}

func (p *parser) peek() byte {
	return p.buf[p.col]
}

func (p *parser) empty() bool {
	return p.col >= len(p.buf)
}

// Move cursor by n bytes. If n > len(p.bytes()) cursor is move to end.
func (p *parser) skip(n int) {
	p.col += min(n, len(p.bytes()))
}

// Remove trailing comment from current buffer.
func (p *parser) trimComment() {
	if i := bytes.IndexAny(p.bytes(), ";#"); i != -1 {
		p.buf = p.buf[:p.col+i]
	}
}

// Returns a slice of current buffer if any Unicode code point from `str` is
// found.
func (p *parser) sliceAny(str string) []byte {
	i := bytes.IndexAny(p.bytes(), str)
	if i == -1 {
		return nil
	}
	return p.bytes()[:i]
}

func (p *parser) parseNext() (string, string, error) {
	if p.nextLine() {
		p.trimSpace()
		if p.empty() {
			return p.parseNext()
		}

		// Parse section.
		if p.peek() == '[' {
			p.trimComment()

			// Skip '['
			p.skip(1)

			// Extract section.
			section := p.sliceAny("]")
			if section == nil {
				return "", "", p.error("invalid section")
			}
			p.skip(len(section) + 1)
			p.trimSpace()

			if !p.empty() {
				return "", "", p.error("invalid content after section")
			}

			if len(section) > 0 && !sectionRegex.Match(section) {
				return "", "", p.error("invalid section")
			}

			if len(section) == 0 {
				p.section = ""
			} else if section[0] == '.' {
				p.section += string(section[1:]) + "."
			} else {
				p.section = string(section) + "."
			}

			return p.parseNext()
		}

		// Skip comments.
		if p.peek() == ';' || p.peek() == '#' {
			return p.parseNext()
		}

		// Parse key = val
		{
			key := p.sliceAny("=:")
			if key == nil {
				return "", "", p.error("invalid option, separators '=' or ':' are missing")
			}
			p.skip(len(key) + 1)

			key = bytes.TrimSpace(key)

			value, err := p.parseValue()
			if err != nil {
				return "", "", err
			}

			return p.section + string(key), value, nil
		}
	} else if err := p.scanner.Err(); err != nil {
		return "", "", err
	}

	return "", "", nil
}

func (p *parser) parseValue() (string, error) {
	p.trimSpace()
	if p.empty() {
		return "", nil
	}

	if b := p.peek(); b == '"' || b == '\'' || b == '`' {
		return p.parseString()
	}

	line := p.bytes()
	if len(line) == 0 || line[len(line)-1] != '\\' {
		// Single line unquoted string.
		return string(line), nil
	}

	// Multi line unquoted string
	var b strings.Builder
	_, _ = b.Write(line[:len(line)-1])
	_ = b.WriteByte('\n')
	for p.nextLine() {
		line = p.bytes()
		if len(line) == 0 || line[len(line)-1] != '\\' {
			_, _ = b.Write(line)
			break
		}
		_, _ = b.Write(line[:len(line)-1])
		_ = b.WriteByte('\n')
	}
	if p.scanner.Err() != nil {
		return "", p.scanner.Err()
	}

	return b.String(), nil
}

func (p *parser) parseString() (string, error) {
	var b strings.Builder

	if len(p.bytes()) < 2 {
		return "", p.error("invalid string")
	}

	quote := p.peek()
	_ = b.WriteByte(quote)
	p.skip(1)

	inner := p.sliceAny(string(quote))
	if inner == nil {
		if quote != '`' {
			return "", p.error("unclosed string")
		} else {
			_, _ = b.Write(p.bytes())
			_ = b.WriteByte('\n')

			// Read until we find end of string.
			var l []byte
			for p.nextLine() {
				l = p.sliceAny("`")
				if l != nil {
					_, _ = b.Write(l)
					_ = b.WriteByte(quote)
					p.skip(len(l) + 1)
					break
				}

				_, _ = b.Write(p.bytes())
				_ = b.WriteByte('\n')
			}
			if p.scanner.Err() != nil {
				return "", p.scanner.Err()
			}
			if str := b.String(); len(str) > 0 && str[len(str)-1] != quote {
				return "", p.error("unclosed multiline string")
			}
		}
	} else {
		_, _ = b.Write(inner)
		_ = b.WriteByte(quote)
		p.skip(len(inner) + 1)
	}

	p.trimComment()
	p.trimSpace()
	if !p.empty() {
		return "", p.error("invalid content after value")
	}

	return strconv.Unquote(b.String())
}

func (p *parser) error(msg string) error {
	return fmt.Errorf("%v at %v:%v", msg, p.line, p.col+1)
}
