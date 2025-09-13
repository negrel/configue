package option

import (
	"fmt"
	"math"
	"slices"
	"testing"
	"unsafe"
)

func TestSlice(t *testing.T) {
	t.Run("Uint64", func(t *testing.T) {
		t.Run("Ok", func(t *testing.T) {
			t.Run("Default", func(t *testing.T) {
				var u64s []uint64
				def := []uint64{1, 2, 3}
				val := NewSlice(def, &u64s)
				_ = val

				if unsafe.SliceData(u64s) == unsafe.SliceData(def) {
					t.FailNow()
				}
				if !slices.Equal(u64s, def) {
					t.FailNow()
				}
			})
			t.Run("MultipleSet", func(t *testing.T) {
				var u64s []uint64
				val := NewSlice([]uint64{1, 2, 3}, &u64s)

				err := val.Set("1,2,345,678")
				if err != nil {
					t.Fatal(err.Error())
				}
				err = val.Set("1")
				if err != nil {
					t.Fatal(err.Error())
				}

				if u64s[0] != 1 || u64s[1] != 2 || u64s[2] != 345 || u64s[3] != 678 ||
					u64s[4] != 1 {
					t.Fatal("value doesn't match expected", u64s)
				}
			})
		})
		t.Run("KO", func(t *testing.T) {
			t.Run("Negative", func(t *testing.T) {
				var u64s []uint64
				val := NewSlice(nil, &u64s)

				err := val.Set("1,2,345,-678")
				if err == nil {
					t.Fatal("error expected but got nil")
				}
			})
			t.Run("EmptyElement", func(t *testing.T) {
				var u64s []uint64
				val := NewSlice(nil, &u64s)

				err := val.Set("1,,345,678")
				if err == nil {
					t.Fatal("error expected but got nil")
				}
			})
			t.Run("TooBig", func(t *testing.T) {
				var u64s []uint64
				val := NewSlice(nil, &u64s)

				err := val.Set(fmt.Sprint(uint64(math.MaxUint64)) + "1")
				if err == nil {
					t.Fatal("error expected but got nil")
				}
			})
		})
	})
}
