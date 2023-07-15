# Go Types

Provides Option, Result and Tuple generic types for go and supporting methods

## Tuple

Tuples represent 2 or more values as a single type. They are very helpful when creating generic go code that utilizes function signatures.

For full examples of the tuple types see [the tuple tests](tuple_test.go)

### Tuple[T1, T2 any]

Represents 2 values in a single type.

The easiest way to create a Tuple[T1, T2 any] is to use the `tuple` package. 

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

[Go Playground](https://go.dev/play/p/-a2MXvUdCSP)

### Tuple3[T1, T2, T3 any]

Represents 3 values in a single type.

The easiest way to create a Tuple3[T1, T2, T3 any] is to use the `tuple` package

```golang
package main

import(
    "fmt"
    "github.com/patrickhuber/go-types/tuple"
)
func main(){
    t := tuple.New3(1, "hello world", 1.234)
    i, s, f := t.Deconstruct()
    fmt.Printf("%d %s %f", i, s, f)
}
```

[Go Playground](https://go.dev/play/p/0MdilHXDZaF)

### Tuple4[T1, T2, T3, T4 any]

Represents 4 values in a single type.

The easiest way to create a Tuple4[T1, T2, T3, T4 any] is to use the `tuple` package

```golang
package main

import (
	"fmt"

	"github.com/patrickhuber/go-types/tuple"
)

func main() {
	t := tuple.New4(1, "hello world", 1.234, true)
	i, s, f, b := t.Deconstruct()
	fmt.Printf("%d %s %f %t", i, s, f, b)
}
```

[Go Playground](https://go.dev/play/p/E3CPO3DWm5O)

## Option[T any]

Option[T any] represents Something or Nothing. 

The core types for the option type are in the `types` pacakge. Helper functions for creating options are located in the `option` package.

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
        fmt.Println("some", o.Value())
    case types.None[int]:
        fmt.Println("none")
    }    
}
```

[Go Playground](https://go.dev/play/p/a5kpg-AMw9a)

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
        fmt.Println("some", o.Value())
    case types.None[int]:
        fmt.Println("none")
    }    
}
```

[Go Playground](https://go.dev/play/p/MKYSi4NyGhi)

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
        fmt.Println("value", r.Ok())
    case types.Error[int]:
        fmt.Println(r.Error())
    }    
}
```

[Go Playground](https://go.dev/play/p/SybyvlOyH80)

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
        fmt.Println("value", r.Ok())
    case types.Error[int]:
        fmt.Println(r.Error())
    }    
}
```

[Go Playground](https://go.dev/play/p/13IyH_tP7qt)

## Error Handling 

Result[T any] can be used to avoid if err != nil repetition in code by passing in Result[T any] objects instead of values

### Baseline

As a baseline look at the [idiomatic go example](examples/idiomatic_test.go)

### Assert

To utilize the library without modifying dependencies you can use `defer handle.Error(&res)` along with the `assert` package to remove `if err != nil { return err}` checks. The [assert example](examples/assert_test.go) shows how.

### Wrapping

To use results in return types by wrapping existing functions that return errors see the [wrap example](examples/wrap_test.go)

The `result.Unwrap()` syntax is very similar to rust's `?` syntax and dramatically reduces the verticle size of the code while still handing errors.