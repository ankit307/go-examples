# Go Examples

[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://github.com/ankit307/go-examples/graphs/commit-activity) [![made-with-Go](https://img.shields.io/badge/Made%20with-Go-blue.svg)](https://golang.org/) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) [![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/) [![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php) [![HitCount](http://hits.dwyl.io/ankit307/go-examples.svg)](http://hits.dwyl.io/ankit307/go-examples)

### Show some :heart: and :star: the repo to support the project

Go Examples contains some practice programs for learning the basics of Go Language. 
It contains programs on topics, that will help you as a building block:

### Table of contents

- Basics
    - Interface
    - Struct
    - Methods
    - Maps
    - Range
    - Defer
    - Read File
    - Process CSV
    - Panic
    - JSON
    - Encrypt Decrypt
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
    - Channel Synchronization
    - Worker Pools
    - Wait Group
    - Atomic
- Server
    - HttpServer
    - Environment Variable
    - Middleware
    - Authentication


## Basic

### Interface
Interfaces are named collection of method signature. To implement an interface in Go, we just need to implement all the methods in the interface

### Struct
Structs are type collection of fields. They are similar to an object that contains only attributes.
Structs behavior can be extended using methods.

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
This is usually used for cleanup purpose for running the functions before exiting the context of function. defer is often used where e.g. ensure and finally would be used in other languages.


### [About the Author]()
Developer / Author: Ankit Garg
Let's Connect ! - | [LinkedIn](http://bit.ly/2lUHXQi) | [Twitter](http://bit.ly/2lO6WVJ)