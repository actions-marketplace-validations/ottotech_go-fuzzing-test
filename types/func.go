package types

import "fmt"

func OnlyCertainNumbers(n int64) (int64, error) {
	switch n {
	case 300, 1505, 15000:
		return 0, fmt.Errorf("invalid number %d", n)
	case 200:
		panic("BOOM")
	}
	return n, nil
}
