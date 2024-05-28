## Packages
Every program contains packages. 
Programs start running in package `main`
The package name is the same as last element of the import path. i.e `'math/rand'` package contains files that begin with statement `package rand`
- Each of the imported files begin with and belong to `package rand`

## Imports
It is best practice to use imports in "factored" statement practice i.e 
```go
import (
	"fmt"
	"math"
)```
is better than
```go
import "fmt"
import "math"
```

## Exported Names
A name is exported if it begins with a capital letter.
- like `Pizza` or `Pi` from `math` package.
A name is **not** exported if it does not begin with a capital letter
- like `pizza` or `pi`

When importing  packages, you can **only** refer to its exported names.
```go
fmt.Println(math.pi)
```
throws an undefined exception, but
```go
fmt.Println(math.Pi)
```
runs successfully!

## Functions
A function can take zero or more arguments.
```go
func add(x int, y int) int {
	return x + y
}
```
add() takes two parameters, x and y, both of type `int`
- **notice how types come AFTER the variable name**
The return type of add() is `int`

When two or more consective function params are of the same type, you can omit the type from all but the last. We can shorten add() to:
```go
func add(x, y int) int {
	return x + y
}
```

### Detour - [Go's Declaration Syntax](https://go.dev/blog/declaration-syntax)
C syntax follows the approach where an expression involves:
- the type of the item being declared
- the name of the item being declared

```C
int x;
```
declares x to be an int
```C
int *p
```
declares p is a pointer to int
```C
int a[3]
```
declares a is an array of ints because a[3] has type int.

This is a clever syntactic idea that works well for simple types but can get confusing fast. The famous example is declaring a function pointer. Follow the rules and you get this:
```C
int (*fp)(int a, int b);
```
and it just gets even more confusing.
```C
int (*fp)(int (*ff)(int x, int y), int b)
```
and more confusing
```C
int (*(*fp)(int (*)(int, int), int))(int, int)
```

Go syntax follows something a little more descriptive, where you:
- declare the name of the object
- declare the type of the object
In a fictional language:
```
x: int
p: pointer to int
a: array[3] of int
```
but in Go:
```go
x int
p *int
a [3]int
```

In summary, Go's declarations read left to right, not right to left.

## Multiple Returns
A function can return any number of results.
```go
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

```
returns two strings, prints "world hello"

## Named return values
In Go, you can name return values.
- if done, they are treated as variables defined at the 'top' of the function
```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```
split() returns val of x = 7 and val of y = 4.

A return statement without args returns the **named** return values.
- in the example, that would be x and y
This is known as a naked return.
- they should only be used in short functions, as they can harm readability.

## Variables
Static typed lang
The `var` statement declares a list of variables. The type is last (as mentioned in detour).

A `var` statement can be at the package or function level, i.e
```go
var c, python, java bool //defined at package lvl

func main() {
	var i int // defined at function lvl
	fmt.Println(i, c, python, java)
}
```

## Variables with initializers
A `var` declaration can include initializers, one per variable. 
Go also has **type inference** - if a value is assigned, the variable will take the type of the value.
- similar to C#'s `var` statement in this way

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```
Type inference in use here at the function level. No type in sight

## Short variable declarations
Inside a function, the `:=` statement can be used instead of a `var` declaration with implicit type inferencing.
- only works in the function level
- at the package level, every statement begins with a keyword i.e `func` or `var`
```go
var i, j int = 1, 2
k := 3
```
all 3 vars evaluate to type int

## Basic types
- `bool`
- `string
- `int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr`
- `byte` alias for uint8
- `rune` alis for int32, represents a Unicode code point
- `float32 float64`
- `complex64 complex128`

When you need an integer value you should use `int` unless you have a specific reason to use a sized or unsigned integer type.

```go
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```
prints
```
Type: bool Value: false
Type: uint64 Value: 18446744073709551615
Type: complex128 Value: (2+3i)
```

**Note** - variable groups may be factored into blocks, like import statements

## Zero values
Variables declared without explicit values are given default `zero values`
- Numeric: 0
- Boolean: false
- String: empty string ""

## Type conversions
The expression `T(v)` converts the value `v` to the type `T`.
```go
i := 42 
f := float64(i) //42 -> 42.0
u := uint(f) //42.0 -> 42
```
In Go, assignment between items of different types **requires** an explicit cast

## Type inference
She exists
```go
var i int
j := i 
```
Here, j is inferred to be an int

For an untyped numeric constant, the type is inferred based on the precision of the constant:
```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

## Constants
Declared like variables, but with the `const` reserved keyword
They can be string, char, bool or numeric values, but cannot be declared using `:=` syntax :(
- technically, no syntax/compilation error from doing so. but it becomes a variable, not a constant.

**Numeric Constants**
high-precision values, takes the type needed by its context