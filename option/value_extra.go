package option

import (
	"encoding"
	"encoding/csv"
	"strings"
	"time"
)

// This file contains extra flag.Value that are not part of standard library.

// NewSlice creates a new slice value.
func NewSlice[T any](val []T, p *[]T) *Slice[T] {
	*p = val
	return (*Slice[T])(p)
}

// Slice is a wrapper around []T that implements Value. *T MUST implement Value
// otherwise Slice.Set / Slice.String will panic at runtime.
type Slice[T any] []T

func (s *Slice[T]) Set(str string) error {
	r := csv.NewReader(strings.NewReader(str))
	record, err := r.Read()
	if err != nil {
		return err
	}

	for _, str := range record {
		var (
			err  error
			t    T
			tAny any = &t
		)
		switch val := tAny.(type) {
		case *bool:
			err = (*Bool)(val).Set(str)
		case *time.Duration:
			err = (*Duration)(val).Set(str)
		case *float64:
			err = (*Float64)(val).Set(str)
		case *int:
			err = (*Int)(val).Set(str)
		case *int64:
			err = (*Int64)(val).Set(str)
		case *string:
			err = (*String)(val).Set(str)
		case *uint:
			err = (*Uint)(val).Set(str)
		case *uint64:
			err = (*Uint64)(val).Set(str)
		case *func(str string) error:
			err = (*Func)(val).Set(str)
		default:
			if unmarshaler, ok := tAny.(encoding.TextUnmarshaler); ok {
				err = Text{p: unmarshaler}.Set(str)
			} else if val, ok := tAny.(Value); ok {
				err = val.Set(str)
			} else {
				panic("T doesn't implement Value")
			}
		}
		if err != nil {
			return err
		}

		*s = append(*s, t)
	}

	return nil
}

func (s *Slice[T]) String() string {
	if s == nil || *s == nil {
		return ""
	}

	var b strings.Builder
	w := csv.NewWriter(&b)

	strs := make([]string, len(*s))
	for _, t := range *s {
		var (
			str  string
			tAny any = &t
		)
		switch val := tAny.(type) {
		case *bool:
			str = (*Bool)(val).String()
		case *time.Duration:
			str = (*Duration)(val).String()
		case *float64:
			str = (*Float64)(val).String()
		case *int:
			str = (*Int)(val).String()
		case *int64:
			str = (*Int64)(val).String()
		case *string:
			str = (*String)(val).String()
		case *uint:
			str = (*Uint)(val).String()
		case *uint64:
			str = (*Uint64)(val).String()
		case *func(str string) error:
			str = (*Func)(val).String()
		default:
			if unmarshaler, ok := tAny.(encoding.TextUnmarshaler); ok {
				str = Text{p: unmarshaler}.String()
			} else if val, ok := tAny.(Value); ok {
				str = val.String()
			} else {
				panic("T doesn't implement Value")
			}
		}

		strs = append(strs, str)
	}
	_ = w.Write(strs)

	return b.String()
}
