# Option Type for Go

The `Option` package provides an `Option` type that represents encapsulation of an optional value; a value that may either contain a value (`Some`) or not contain a value (`None`), similar to `Option` in Rust.

## Features

- **Some**: Create an `Option` that contains a value.
- **None**: Create an `Option` that represents no value.
- **Ok**: Return the value if it exists.
- **Cause**: Return the cause of `None` if it is due to an error.
- **Process**: Transform the `Option` value with a function that might fail.
- **Map**: Transform the `Option` value with a function that cannot fail.
- **Flatten**: Nested `Option` values into a single `Option`.
- **Wrap**: Wraps a value and an error into an `Option`.
- **WrapFn**: Wraps a function that returns a value and an error into a function that returns an `Option`.
## Installation

To use the `Option` package in your Go project, place the package in your project directory or a suitable location in your GOPATH. Import it using:

```shell
go get -u github.com/AY7295/go-option
```


## Usage Examples
Here are some basic examples of how to use the Option type:

### Creating Some and None
```go
opt1 := option.Some(10)
opt2 := option.None[int]()
```

### Processing an Option

```go
opt := option.Some(5)
result := option.Process(opt, func(v int) (int, error) {
    if v == 0 {
        return 0, errors.New("zero value")
    }
    return 2 * v, nil
})
```

### Mapping over an Option

```go
opt := option.Some(10)
mapped := option.Map(opt, func(v int) int {
    return v + 1
})
```

### Flattening an Option

```go
doubleWrapped := option.Some(option.Some(20))
flattened := option.Flatten(doubleWrapped)
```

### Testing
To run the tests for the Option package, navigate to the package directory and execute:

```shell
go test
```
This will run all the tests defined in option_test.go.


