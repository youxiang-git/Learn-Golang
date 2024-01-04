# Learn Golang 

[Main website](https://www.golang.org) 

## 1. Introduction

### Golang is:
1. Strong; and **(Variable type cannot change overtime i.e. int cannot hold bool)**
1. Statically typed **(Variables have to be defined at compile time)**
1. Excellent community
1. Key features
    * Simplicity
    * Fast compile times
    * Garbage collected
    * Built-in concurrency
    * Compile and standalone binaries **(Don't need to bundle additional dependencies for Golang)**

## 2. First GO file

If the package is called 'main', it tells GO to compile our file into an executable program at the end. All files that are part of the program will specify 'main' at the start of the program.

```golang
package main
```


There should only be one main function inside the application, and it serves as the entry point.

```golang
func main() { }    
```

## 3. Variables, Strings & Numbers
### Strings 
**NOTE** - For strings in Golang, we use double quotes, not single quotes

Refer to this [file](golang/main.go) for examples on variables, strings & numbers