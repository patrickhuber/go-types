# Go Types

Provides Option, Result and Tuple generic types for go and supporting methods

## Tuple

Tuples represent 2 or more values as a single type. They are very helpful when creating generic go code that utilizes function signatures.

For full examples of the tuple types see [the tuple tests](tuple_test.go)

### Tuple2[T1, T2 any]

Represents 2 values in a single type.

The easiest way to create a Tuple2[T1, T2 any] is to use the `tuple` package. 

```golang
package main

import(
    "fmt"
    "github.com/patrickhuber/go-types/tuple"
)
func main(){
    t := tuple.New2(1, "hello world")
    i, s := t.Deconstruct()
    fmt.Printf("%d %s", i, s)
}
```

[Go Playground](https://go.dev/play/p/FK_b8UvaiGL)

### TupleN[T1 .. TN]

There are a total of 16 Tuple types in this library. They each have similar methods and support functions generated by code gen. 

## Option[T any]

Option[T any] represents Something or Nothing. 

The core types for the option type are in the `types` package. Helper functions for creating options are located in the `option` package.

For full examples of the option types see [the option tests](option_test.go)

### Some[T any]

Some[T any] is an Option[T any] with a value. 

The easiest way to create a Some type is to use the `option` package

```golang
package main

import (
    "github.com/patrickhuber/go-types/option"
    "github.com/patrickhuber/go-types"
    "fmt"
)
func main(){
    // creates a types.Some[int] using type inference
    op := option.Some(1) 
    // because Some and None are types, a types switch can be used
    switch o := op.(type){
    case types.Some[int]:
        fmt.Println("some", o.Value)
    case types.None[int]:
        fmt.Println("none")
    }    
}
```

[Go Playground](https://go.dev/play/p/FK_b8UvaiGL)

### None[T any]

None[T any] is an Option[T any] with no value

The easiest way to create a None type is to use the `option` package

```golang
package main

import (
    "github.com/patrickhuber/go-types/option"
    "github.com/patrickhuber/go-types"
    "fmt"
)
func main(){
    // creates a types.None[int]
    op := option.None[int]() 
    // because Some and None are types, a types switch can be used
    switch o := op.(type){
    case types.Some[int]:
        fmt.Println("some", o.Value)
    case types.None[int]:
        fmt.Println("none")
    }    
}
```

[Go Playground](https://go.dev/play/p/SFAmmKi6wTi)

## Result[T any]

Result[T any] represents a value Ok or an Error

The core types for result type are in the `types` package. Helper functions for creating results are located in the `result` package.

For full examples of the result types see [the result tests](result_test.go)

### Ok[T any]

Ok[T any] is a successful Result[T any] with a value

The easiest way to create a Ok type is to use the `result` package

```golang
package main

import (
    "github.com/patrickhuber/go-types/result"
    "github.com/patrickhuber/go-types"
    "fmt"
)
func main(){
    // creates a Ok[int] result using type inference
    res := result.Ok(1) 
    // because Ok and Error are types, a types switch can be used
    switch r := res.(type){
    case types.Ok[int]:
        fmt.Println("value", r.Value)
    case types.Error[int]:
        fmt.Println(r.Value)
    }    
}
```

[Go Playground](https://go.dev/play/p/cilRscVSBtG)

### Error[T any]

Error[T any] is a failed Result[T any] with an error

The easiest way to create an Error type is to use the `result` package

```golang
package main

import (
    "github.com/patrickhuber/go-types/result"
    "github.com/patrickhuber/go-types"
    "fmt"
)
func main(){
    // creates a Error[int] result
    res := result.Errorf[int]("error")

    // because Ok and Error are types, a types switch can be used
    switch r := res.(type){
    case types.Ok[int]:
        fmt.Println("value", r.Value)
    case types.Error[int]:
        fmt.Println(r.Value)
    }    
}
```

[Go Playground](https://go.dev/play/p/6Xq2BehmNJL)

## Error Handling 

Result[T any] can be used to avoid if err != nil repetition in code by passing in Result[T any] objects instead of values.

Result[T any] must be called with defer or panics will not be recovered.

```golang
func MyFunction() (res Result[any]){
    defer handle.Error(&res)
    assert.IsNotNil(nil)
}
```

System panics will not be trapped by handle.Error because they will not wrap ErrRecoverable. Any function in this module that panics, will panic with ErrRecoverable. 

Use the method `Throw` if you want to panic with a recoverable error as it performs the error wrapping of ErrRecoverable using fmt.Errorf.

### Baseline

As a baseline look at the [idiomatic go example](examples/idiomatic_test.go)

### Assert

To utilize the library without modifying dependencies you can use `defer handle.Error(&res)` along with the `assert` package to remove `if err != nil { return err}` checks. The [assert example](examples/assert_test.go) shows how.

### Wrapping

To use results in return types by wrapping existing functions that return errors see the [wrap example](examples/wrap_test.go)

The `result.Unwrap()` syntax is very similar to rust's `?` syntax and dramatically reduces the verticle size of the code while still handing errors.
