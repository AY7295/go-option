package option

import "fmt"

func (opt *option[T]) String() string {
	if IsNone[T](opt) {
		return "none"
	}

	return fmt.Sprint(opt.value)
}
