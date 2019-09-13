# Go Examples

Go Examples contains some practice programs for learning the basics of Go Language. 
It contains programs on topics, that will help you as a building block:

- Basics
    - Interface
    - Struct
    - Maps
    - Range
    - Defer
    - Panic
    - JSON

- Design Pattern
    - Closure
    - Recursion
    - Currying
    - Function Composition
    - Collection Functions (Index, Include, Any, All, Filter, Map)

- Concurrency
    - GoRoutine
    - Channels
    - Channels buffering
    - Channel Syncronization
    - Worker Pools
    - Wait Group

- Server
    - HttpServer
    - Environment Variable
    - Middleware
    - Authentication


## Basic

### Interface
Interfaces are named collection of method signature. To impliment an interface in Go, we just need to impliment all the methods in the interface

### Struct
Structs are type collection of fields. They are simalar to an object that contains only attributes.
Structs behaivour can be extended using methods.

### Maps
Maps are a key value based data structure also called as hash and dict in other language.The builtin len returns the number of key/value pairs when called on a map.

### Range
Range iterates over elements in a variety of data structure. Any collection of data can be iterated by using range and for loop

``` go
 nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
    //Output= sum:9
   
   for i, c := range "go" {
        fmt.Println(i, c)
    }
    //Output=
    //0 103
    //1 111
```

### Defer 
This is usualy used for cleanup purpose for runing the functions before exiting the context of function. defer is often used where e.g. ensure and finally would be used in other languages.