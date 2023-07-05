# Go Types

Provides Option[T], Result[T] types for go and supporting methods

## Option[T]

Option[T] represents Something or Nothing. 

### Some[T]

Some[T] is an Option[T] with a value

### None[T]

None[T] is an Option[T] with no value

## Result[T]

Result[T] represents a value Ok or an Error

## Ok[T]

Ok[T] is a successful Result[T] with a value

## Error[T]

Error[T] is a failed Result[T] with an error

## Usage

Result[T] can be used to avoid if err != nil repetition in code by passing in Result[T] objects instead of values

```go
package main

import (
    "github.com/patrickhuber/go-types"
    "github.com/patrickhuber/go-types/result"    
    "strconv"
)

func PossibleError(val int) Result[int]{
    return result.Ok(val)
}

func AcceptResult(res Result[int]) Result[string]{
    switch t := res.(type){
    
    case types.Ok[int]:
        return result.Ok(strconv.Itoa(t.Ok()))

    case types.Error[int]:
        return result.Error[string](v.Error())        
    }

    // no exhaustive matching so match default case and throw error
    return MatchError[string](res)
}

func MatchError[T any](val any) Result[T]{
    return result.Error(fmt.Errorf("unable to match type %T", val))
}

func main(){
    switch t := AcceptResult(PossibleError(1)).(type){

    case types.Ok[string]:
        fmt.Println(t.Ok())

    case types.Error[string]:
        fmt.Println(t.Error())
        panic(1)
    }
}
```

An idiomatic go program would look like the following

```go
package main

import (
    "strconv"
)

func PossibleError(val int) (int, error){
    return val, nil
}

func AcceptResult(val int) (string,error){
    return strconv.Itoa(val), nil
}

func main(){
    res, err := PossibleError(1)
    if err != nil{
        fmt.Println(err)
        panic(1)
    }

    accept, err := AcceptResult(res)
    if err != nil{
        fmt.Println(err)
        panic(1)
    }

    fmt.Println(accept)
}
```