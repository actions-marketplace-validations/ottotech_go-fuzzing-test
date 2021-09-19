// +build gofuzzbeta

package types

import (
	"testing"
)

func FuzzOnlyCertainNumbers(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int64) {
		t.Parallel()

		if n <= 0 {
			t.Skip() // only test positive numbers
		}

		if n == 1505 || n == 200 || n == 300 {
			t.Skip()
		}

		_, err := OnlyCertainNumbers(n)
		if err != nil {
			t.Fatal(err)
		}
	})
}
