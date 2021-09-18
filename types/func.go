package types

import "fmt"

func OnlyCertainNumbers(n int64) (int64, error) {
	switch n {
	case 200, 300, 1505, 45450:
		return 0, fmt.Errorf("invalid number %d", n)
	}
	return n, nil
}
