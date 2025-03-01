# Examples

### Table of Contents
- [Either](#either)
- [Option](#option)
- [Functions](#functions)
    - [Map](#map)
    - [Filter](#filter)
- [XError](#xerror)

## Either

```go
package main

import (
    "fmt"
    "github.com/livestock-market/go-fp-slim/either"
)

func failIfNegative(input int) either.Either[int] {
    if input < 0 {
        return either.Left[int]("input is negative")
    }
    return either.Right(input)
}

func main() {
    input := -3
    result := failIfNegative(input)
    if result.IsLeft() {
        fmt.Println(result.IsRight()) // false
        fmt.Println(result.Left()) // input is negative
    } else {
        fmt.Println(result.Right()) // won't be printed
    }
}
```
---

## Option

```go
package main

import (
    "fmt"
    "github.com/livestock-market/go-fp-slim/option"
)

func divide(a, b int) option.Option[int] {
    if b == 0 {
        return option.None[int]()
    }
    return option.Some(a / b)
}

type Payload struct {
    ID int
    Name string
    Description Option[string] // Optional field
}

func main() {
    a, b := 10, 0
    result := divide(a, b)
    if result.IsNone() {
        fmt.Println(result.IsSome()) // false
        fmt.Println(result.None()) // won't be printed
    } else {
        fmt.Println(result.Some()) // won't be printed
    }

    payload := Payload{
        ID: 1,
        Name: "John Doe",
        Description: option.None[string](),
    }

    if payload.Description.IsNone() {
        fmt.Println("Description is not provided")
    } else {
        fmt.Println(payload.Description.Some())
    }

    // encode payload to JSON
    jsonResult, _ := json.Marshal(payload)
    fmt.Println(string(jsonResult)) // {"ID":1,"Name":"John Doe","Description":null}
}
```
---

## Functions 
### Map

```go

package main

import (
    "fmt"
    "github.com/livestock-market/go-fp-slim/functions"
)

func square(a) int {
    return a * a
}

func main() {
    // slice of integers
    numbers := []int{1, 2, 3, 4, 5}
    // map the slice with square function
    result := functions.Map(numbers, square)
    fmt.Println(result) // [1 4 9 16 25]
}
```
---

### Filter

```go
package main

import (
    "fmt"
    "github.com/livestock-market/go-fp-slim/functions"
)

func isEven(a int) bool {
    return a % 2 == 0
}

func main() {
    // slice of integers
    numbers := []int{1, 2, 3, 4, 5}
    // filter the slice with isEven function
    result := functions.Filter(numbers, isEven)
    fmt.Println(result) // [2 4]
}
```
---

## XError 

```go

package main

import (
    "fmt"
    gfp "github.com/livestock-market/go-fp-slim"
)



func main() {
    actualErr := fmt.Errorf("Something went wrong")
    xerr := NewXError("Something went horribly wrong", actualErr, http.StatusInternalServerError)
    fmt.Println(xerr.Message) // Something went horribly wrong
}
```


