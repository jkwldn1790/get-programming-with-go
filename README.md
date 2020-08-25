# get-programming-with-go 
Capstone Projects from Get Programming with Go | Nathan Youngman &amp; Roger Peppe

## Capstone Solutions
* Ticket to Mars

## Notes : Lesson 6
* Go will infer types by the variable
* There are two types of floating point numbers in Go (double & single precision) or float64 & float32
* If you want to use float32 you have to declare the type
* Float64 uses 8 bytes of memory while float32 uses 4 bytes

## Notes: Lesson 7 - 8 
* You can check the type of variables by using the `%T` operator.
Example:
```
package main

import (
	"fmt"
)

func main() {
	year := 2020
	fmt.Printf("Type %T for %v\n", year, year)
}
```
* If you have the case where you are using the same variable twice. You can use `[1]%v` to indicate the first arg.
```
year := 2020
fmt.Printf("Type %T for %[1]v\n", year)
```
* You can use `fmt.Printf` with `%x` to print hexidecimal numbers.
* `%b` will show you the bits for an integer
* Calculations on constants and literals are done during compilation rather than while the program is running. This allows you to define numbers that are larger than 64-bit in a constant or literal.
* Untyped constants must be converted to typed variables when passed to functions

## Notes: Lesson 9
* Characters together in a string between quotes becomes a "literal string"
* If you declare a variable without a value it will be inferred as zero value and zero value for a string is an empty string `""`
* If you want to avoid substituting in a `\n` in a string literal, you can use a raw string literal using backticks.
* You can't modify a string in Go, but you can access it much like you would in Python. `message := "shalom" message[5] = "m"`
* There is a built-in `len` function for messages that can be used. `fmt.Println(len(message))`

> Signed vs. Unsigned integers ... unsigned integers cannot be negative and have a higher positive range. While signed integers can be negative and have a lower positive range as result.

## Notes: Lesson 10
* You can convert types using similar sytax to the following ...
```
age := 41
marsAge := float64(age)
```
* You can also convert a floating point to an integer `fmt.Println(int(earthDays))
* You can convert a rune or a byte to a string as well using like so ... 
```
var pi rune = 960
var alpha rune = 940
var omega rune = 969
var bang byte = 33

fmt.Println(string(pi), string(alpha), string(omega), string(bang))
```
*Convert Integers to Strings*
* You can convert integers to strings using `strconv.Itoa(10)`
* You can also use the `Sprintf` package that returns a string instead of displaying it.
* You should also probably use the `Sprintf` package for converting booleans to strings as well.
> You can use a string to evaluate a boolean like...
```
yesNo := "yes"

launch := (yesNo == "yes") // evaluates to true
```