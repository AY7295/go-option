package option

import "fmt"

func (opt *option[T]) String() string {
	if opt.cause != nil {
		return fmt.Sprintf("None Option: %v", opt.cause)
	}

	return fmt.Sprint(opt.value)
}
