package option

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestSome(t *testing.T) {
	opt := Some(42)
	if IsNone(opt) {
		t.Errorf("Some(42) should be Some")
	}
	if val := opt.Ok(); val != 42 {
		t.Errorf("Expected 42, got %v", val)
	}
}

func TestNone(t *testing.T) {
	customError := errors.New("custom error")
	opt := None[int](customError)
	if IsSome(opt) {
		t.Errorf("None should not be Some")
	}
	if cause := opt.Cause(); cause == nil || !errors.Is(cause, customError) {
		t.Errorf("Expected custom error, got %v", cause)
	}
}

func TestProcess(t *testing.T) {
	opt := Some(5)
	result := Process(opt, func(v int) (int, error) {
		if v == 0 {
			return 0, errors.New("zero value")
		}
		return 2 * v, nil
	})

	if result.Ok() != 10 {
		t.Errorf("Expected processed value 10, got %v", result.Ok())
	}

	optZero := Some(0)
	resultZero := Process(optZero, func(v int) (int, error) {
		if v == 0 {
			return 0, errors.New("zero value")
		}
		return 2 * v, nil
	})

	if !IsNone(resultZero) || resultZero.Cause() == nil || resultZero.Cause().Error() != "zero value" {
		t.Errorf("Expected zero value error, got %v", resultZero.Cause())
	}
}

func TestMap(t *testing.T) {
	opt := Some(5)
	newOpt := Map(opt, func(v int) int {
		return v + 3
	})
	if !IsSome(newOpt) {
		t.Errorf("Map should return Some")
	}
	if val := newOpt.Ok(); val != 8 {
		t.Errorf("Expected 8, got %v", val)
	}
}

func TestFlatten(t *testing.T) {
	opt := Some(Some(100))
	flatOpt := Flatten(opt)
	if !IsSome(flatOpt) {
		t.Errorf("Flatten should return Some")
	}
	if val := flatOpt.Ok(); val != 100 {
		t.Errorf("Expected 100, got %v", val)
	}
}

func TestNoneFlatten(t *testing.T) {
	opt := None[Option[int]]()
	flatOpt := Flatten(opt)
	if IsSome(flatOpt) {
		t.Errorf("Flatten of None should return None")
	}
	if !errors.Is(flatOpt.Cause(), Nil) {
		t.Errorf("")
	}
}

func TestOptionMarshalJSON(t *testing.T) {
	opt := Some(42)
	data, err := json.Marshal(opt)
	if err != nil {
		t.Errorf("MarshalJSON failed: %v", err)
	}
	if string(data) != "42" {
		t.Errorf("Expected 42, got %v", string(data))
	}

	opt = None[int](errors.New("custom error"))
	data, err = json.Marshal(opt)
	if err != nil {
		t.Errorf("MarshalJSON failed: %v", err)
	}
	if string(data) != "null" {
		t.Errorf("Expected null, got %v", string(data))
	}
}

func TestOptionUnmarshalJSON(t *testing.T) {
	opt := Some(0)
	err := json.Unmarshal([]byte("42"), opt)
	if err != nil {
		t.Errorf("UnmarshalJSON failed: %v", err)
	}
	if !IsSome(opt) {
		t.Errorf("Expected Some, got None")
	}
	if val := opt.Ok(); val != 42 {
		t.Errorf("Expected 42, got %v", val)
	}
	err = json.Unmarshal([]byte("null"), opt)
	if err != nil {
		t.Errorf("UnmarshalJSON failed: %v", err)
	}
	if IsSome(opt) {
		t.Errorf("Expected None, got Some")
	}
	if !errors.Is(opt.Cause(), Nil) {
		t.Errorf("Expected Nil, got %v", opt.Cause())
	}
}
