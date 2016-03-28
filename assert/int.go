package assert

import (
	"fmt"
)

func AssertInt(got, want ...int) error {
	for _, w := range want {
		if got == w {
			return nil
		}
	}
	return fmt.Errorf("got unexpected int val:%d", got)
}
