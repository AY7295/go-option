package option

import (
	"github.com/bytedance/sonic"
	"slices"
)

type JsonCoder interface {
	Marshal(any) ([]byte, error)
	Unmarshal([]byte, any) error
}

var (
	defaultCoder JsonCoder = sonic.ConfigStd
)

func SetJsonCoder(coder JsonCoder) {
	defaultCoder = coder
}

// MarshalJSON converts the option to JSON.
func (opt *option[T]) MarshalJSON() ([]byte, error) {
	if opt.cause != nil {
		return []byte(`null`), nil
	}

	return defaultCoder.Marshal(opt.value)
}

// UnmarshalJSON converts JSON to an option.
func (opt *option[T]) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || slices.Equal(data, []byte("null")) ||
		slices.Equal(data, []byte("{}")) || slices.Equal(data, []byte("[]")) {
		opt.cause = Nil
		return nil
	}

	return defaultCoder.Unmarshal(data, &opt.value)
}
