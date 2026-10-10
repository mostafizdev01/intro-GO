# Go Programming Tutorial

---

## Table of Contents

- [1. Introduction to Go Programming](#1-introduction-to-go-programming)
  - [What is Go?](#what-is-go)
  - [Why Go?](#why-go)
  - [Prerequisites](#prerequisites)
- [2. Go Program Structure](#2-go-program-structure)
- [3. Environment Setup](#3-environment-setup)
- [4. First Go Program](#4-first-go-program)
  - [A Simple Go Program](#a-simple-go-program)
  - [Running the Program](#running-the-program)
- [5. Tokens - Escape Sequences & Comments](#5-tokens---escape-sequences--comments)
  - [Tokens of Any Programming Language](#tokens-of-any-programming-language)
  - [Escape Sequences Example](#escape-sequences-example)
  - [Comments in Go](#comments-in-go)
- [6. Keywords & Data Types](#6-keywords--data-types)
  - [Keywords](#keywords)
  - [Data Types](#data-types)
  - [Example of Data Types](#example-of-data-types)
  - [More on int, float](#more-on-int-float)
- [7. Variables](#7-variables)
  - [Syntax of Variable Declaration](#syntax-of-variable-declaration)
  - [Variable Naming Convention](#variable-naming-convention)
  - [Example](#example)
- [8. Formatting Output](#8-formatting-output)
- [9. Constants && Getting User Input](#9-constants--getting-user-input)
- [10. Number Conversion Calculator & Number Formatting](#10-number-conversion-calculator--number-formatting)
- [11. Operators](#11-operators)
  - [Types of Operators](#types-of-operators)
  - [Example of Arithmetic Operator](#example-of-arithmetic-operator)
- [12. Find the Area of Triangle and Circle](#12-find-the-area-of-triangle-and-circle)
- [13. Assignment Operator](#13-assignment-operator)
- [14. Unary Operator](#14-unary-operator)
- [15. Relational Operator](#15-relational-operator)
- [16. Logical Operator](#16-logical-operator)
- [17. Bitwise Operator](#17-bitwise-operator)
- [18. if, else if, else Control Statement](#18-if-else-if-else-control-statement)
- [19. A Go Program to Determine Whether a Number is Even/Odd](#19-a-go-program-to-determine-whether-a-number-is-evenodd)
- [20. A Go Program to Find Largest Number Among 3 Numbers Using Relational Operators](#20-a-go-program-to-find-largest-number-among-3-numbers-using-relational-operators)
- [21. A Go Program to Find Largest Number Among 3 Numbers Using Logical Operators](#21-a-go-program-to-find-largest-number-among-3-numbers-using-logical-operators)
- [22. A Go Program to Spell Digits](#22-a-go-program-to-spell-digits)
- [23. switch Control Statement](#23-switch-control-statement)
- [24. Loop Control Statement](#24-loop-control-statement)
- [25. Series Related Programs](#25-series-related-programs)
- [26. break and continue Statement](#26-break-and-continue-statement)
- [27. Functions](#27-functions)
- [28. Create a Function for Calculator](#28-create-a-function-for-calculator)
- [29. Pointers](#29-pointers)
- [30. Structure](#30-structure)
- [31. Arrays & Slices Data Types](#31-arrays--slices-data-types)

---

## 1. Introduction to Go Programming

### What is Go?

- Simple but fast, popular programming language for cloud engineering
- Statically typed, compiled programming language
- Developed by Google in 2007
- Open sourced in 2009
- Often known as Golang
- Syntax has similarity with C, C++
- Mainly used as a server side or backend language
- Supports cross-OS

### Why Go?

- For making simple, reliable and efficient computer program/software
- Helps to create high performance applications but use less resources
- It helps to write concurrent multithread programs

### Prerequisites

- Nothing, but knowing C/C++ or any programming language is an advantage

---

## 2. Go Program Structure

- **package declaration** - every Go program consists of packages. It is a mandatory thing. It has to be the first statement of a Go program.
- **import packages**
- **function**
- **statement**

```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("My name is Anisul Islam")
}
```

**Output:**
```text
My name is Anisul Islam
```

---

## 3. Environment Setup

- **IDE / Text Editor** -> Any text editor: Notepad / Notepad++ / VSCode (free), Vim (free) / GoLand (paid)
- **Go compiler** -> Compiler translates the source code into machine/executable code so that machine can understand the instructions
- **Go extension for VSCode**

---

## 4. First Go Program

### A Simple Go Program

```go
// filename: main.go
// a program for printing student details
package main

import "fmt"

func main() {
    fmt.Println("Anisul Islam is a student")
    fmt.Println("Anisul Islam is 32 years old")
    fmt.Println("Anisul Islam has got GPA 3.92/5 in SSC")
    fmt.Print("Anisul Islam originally from Bangladesh")
}
```

### Running the Program

```bash
go mod init project-name
go run fileName.go
```

---

## 5. Tokens - Escape Sequences & Comments

### Tokens of Any Programming Language

- keywords
- data types
- variables
- escape sequences
- operators

### Escape Sequences Example

```go
package main

import "fmt"

func main() {
    fmt.Print("Name \t Age\n")
    fmt.Print("----- \t ------\n")
    fmt.Print("Anis \t 32\n")
    fmt.Print("Rakib \t 29\n")
}
```

### Comments in Go

There are 2 types of comment available in Go -> single line, multiple line comment. Comments are ignored by the compiler.

```go
// single line comment

/*
  multiple
  line
  comment
*/
```

---

## 6. Keywords & Data Types

### Keywords

Keywords are reserved words.

### Data Types

- **boolean**
- **string**
- **numeric**: integer, floating
- **derived types**: pointer, array, structure, slice, map, interface etc.

### Example of Data Types

- `"anisul islam"` -> string type
- `21` -> int type
- `21.5` -> floating type
- `true/false` -> bool

### More on int, float

- **variation of integer** -> `int8` (8 bits), `int16`, `int32`, `int64`, `uint8` (only positive 0 to 255), `uint6`, `uint32`, `uint64`
- **variation of float** -> `float32`, `float64`

---

## 7. Variables

### Syntax of Variable Declaration

```text
var variable1, variable2, ... variableN dataType
```

### Variable Naming Convention

- camelCasing

### Example

```go
package main

import "fmt"

func main() {
    // static variable declaration
    var name, country string
    var age int
    var gpa float32

    // variable initialization
    name = "anisul islam"
    country = "Bangladesh"
    age = 32
    gpa = 3.92

    // dynamic variable declaration + initialization
    // var name = "Anisul Islam"
    // var country = "Bangladesh"
    // var age = 32
    // var gpa = 3.92

    // var name, country, age, gpa = "anisul islam", "Bangladesh", 32, 3.92

    // shortcuts
    // name := "Anisul Islam"
    // country := "Bangladesh"
    // age := 32
    // gpa := 3.92

    fmt.Println(name)
    fmt.Println(name, " is a student")
    fmt.Println(name, "is", age, "years old")
    fmt.Println(name, "has got", gpa, "/5 in SSC")
    fmt.Print(name, "originally from ", country)
}
```

---

## 8. Formatting Output

```go
package main

import "fmt"

func main() {
    name := "Anisul Islam"
    country := "Bangladesh"
    age := 32
    gpa := 3.92

    fmt.Printf("%v is a student\n", name)
    fmt.Printf("%v is %v years old\n", name, age)
    fmt.Printf("%v has got GPA %v/5 in SSC\n", name, gpa)
    fmt.Printf("%v originally from %v\n", name, country)
}
```

---

## 9. Constants && Getting User Input

Constants indicate to fix values.

```go
package main

import "fmt"

func main() {
    var name string
    const COUNTRY = "Bangladesh"
    var age, num1, num2 int
    var gpa float32

    fmt.Printf("Enter your name: ")
    fmt.Scanf("%v", &name)

    fmt.Printf("Enter your age: ")
    fmt.Scan(&age)

    fmt.Printf("Enter your SSC gpa: ")
    fmt.Scanln(&gpa)

    fmt.Printf("Enter 2 numbers: ")
    fmt.Scan(&num1, &num2)

    fmt.Printf("%v is a student\n", name)
    fmt.Printf("%v is %v years old\n", name, age)
    fmt.Printf("%v has got GPA %v/5 in SSC\n", name, gpa)
    fmt.Printf("%v originally from %v\n", name, COUNTRY)
    fmt.Printf("num1 = %v, num2 = %v\n", num1, num2)
}
```

---

## 10. Number Conversion Calculator & Number Formatting

```go
package main

import "fmt"

func main() {
    // string formatting
    // var name = "Anisul Islam"
    // fmt.Printf("%s\n", name)
    // fmt.Printf("%q\n", name)

    // floating number formatting
    // var number = 3.1416
    // fmt.Printf("%.2f\n", number)

    var decimalNumber int

    fmt.Printf("decimal number= ")
    fmt.Scanf("%v", &decimalNumber)

    fmt.Printf("Binary Number = %b\n", decimalNumber)
    fmt.Printf("Octal Number = %o\n", decimalNumber)
    fmt.Printf("Hexa Number = %x\n", decimalNumber)
}
```

---

## 11. Operators

Operator - symbol for doing calculation.

### Types of Operators

- **Arithmetic Operators** -> `+`, `-`, `*`, `/`, `%`
- **Assignment Operators** -> `=`, `+=`, `-=`, `*=`, `/=`, `%=`
- **Unary Operators** -> `++`, `--`
- **Relational Operators**
- **Logical Operators** -> `&&`, `||`, `!`
- **Bitwise Operators** -> `&`, `|`, `^`
- **Others** -> `,` (comma), `&` (returns address), `*` (pointer)

### Example of Arithmetic Operator

```go
package main

import "fmt"

func main() {
    // arithmetic operator -> +, -, *, /, %
    var num1, num2 int
    fmt.Printf("num1= ")
    fmt.Scan(&num1)
    fmt.Printf("num2= ")
    fmt.Scan(&num2)

    result := num1 + num2
    fmt.Printf("%v+%v=%v\n", num1, num2, result)

    result = num1 - num2
    fmt.Printf("%v-%v=%v\n", num1, num2, result)

    result = num1 * num2
    fmt.Printf("%v*%v=%v\n", num1, num2, result)

    var result2 = float32(num1) / float32(num2)
    fmt.Printf("%v/%v=%v\n", num1, num2, result2)

    result = num1 % num2
    fmt.Printf("%v%%%v=%v\n", num1, num2, result)
}
```

---

## 12. Find the Area of Triangle and Circle

```go
package main

import "fmt"

func main() {
    var radius, area float32

    fmt.Printf("Enter Radius : ")
    fmt.Scan(&radius)

    area = 3.1416 * radius * radius
    fmt.Printf("Area of circle: %v\n", area)

    // a program to calculate area of triangle
    // var base, height, area float32

    // fmt.Printf("Base = ")
    // fmt.Scan(&base)

    // fmt.Printf("Height = ")
    // fmt.Scan(&height)

    // area = 0.5 * base * height

    // fmt.Printf("Area of triangle = %v\n", area)
}
```

---

## 13. Assignment Operator

```go
package main

import "fmt"

func main() {
    // 2. Assignment Operators -> =, +=, -=, *=, /=, %=

    x := 4

    x = x + 1
    x += 1 // x = x + 1
    x -= 1 // x = x - 1
    x *= 1 // x = x * 1
    x /= 1 // x = x / 1
    x %= 1 // x = x % 1

    fmt.Printf("x = %v\n", x)
}
```

---

## 14. Unary Operator

- Unary Operators -> `++`, `--`

---

## 15. Relational Operator

- Relational Operators

---

## 16. Logical Operator

- Logical Operators -> `&&`, `||`, `!`

---

## 17. Bitwise Operator

```go
package main

import "fmt"

func main() {
    x := 18        // 18 = 10010
    y := 17        // 17 = 10001
    and := x & y   // 16 = 10000
    or := x | y    // 19 = 10011
    exor := x ^ y  // 3  = 00011

    fmt.Printf("x & y = %v\n", and)
    fmt.Printf("x | y = %v\n", or)
    fmt.Printf("x ^ y = %v\n", exor)
}
```

---

## 18. if, else if, else Control Statement

### With if only

```go
// with if only
package main

import "fmt"

func main() {
    // a program to determine positive / negative / zero
    number := 10
    if number > 0 {
        fmt.Printf("Positive\n")
    }
    if number < 0 {
        fmt.Printf("Negative\n")
    }
    if number == 0 {
        fmt.Printf("Zero\n")
    }
}
```

### With if, else if

```go
// with if, else if
package main

import "fmt"

func main() {
    // a program to determine positive / negative / zero
    number := 10
    if number > 0 {
        fmt.Printf("Positive\n")
    } else if number < 0 {
        fmt.Printf("Negative\n")
    } else if number == 0 {
        fmt.Printf("Zero\n")
    }
}
```

### With if, else if, else

```go
// with if, else if, else
package main

import "fmt"

func main() {
    // a program to determine positive / negative / zero
    number := 10
    if number > 0 {
        fmt.Printf("Positive\n")
    } else if number < 0 {
        fmt.Printf("Negative\n")
    } else {
        fmt.Printf("Zero\n")
    }
}
```

---

## 19. A Go Program to Determine Whether a Number is Even/Odd

```go
// a go program to determine even/odd
package main

import "fmt"

func main() {
    var number int
    fmt.Printf("Enter any integer: ")
    fmt.Scanf("%v", &number)

    if number%2 == 0 {
        fmt.Printf("Even\n")
    } else {
        fmt.Printf("Odd\n")
    }
}
```

---

## 20. A Go Program to Find Largest Number Among 3 Numbers Using Relational Operators

```go
package main

import "fmt"

func main() {
    var number1, number2, number3 int
    fmt.Printf("Enter 3 numbers: ")
    fmt.Scanf("%v %v %v", &number1, &number2, &number3)

    if number1 > number2 {
        if number1 > number3 {
            fmt.Printf("%v is the largest number\n", number1)
        } else {
            fmt.Printf("%v is the largest number\n", number3)
        }
    } else if number2 > number1 {
        if number2 > number3 {
            fmt.Printf("%v is the largest number\n", number2)
        } else {
            fmt.Printf("%v is the largest number\n", number3)
        }
    } else {
        fmt.Printf("Numbers are equal\n")
    }
}
```

---

## 21. A Go Program to Find Largest Number Among 3 Numbers Using Logical Operators

```go
package main

import "fmt"

func main() {
    var number1, number2, number3 int
    fmt.Printf("Enter 3 numbers: ")
    fmt.Scanf("%v %v %v", &number1, &number2, &number3)

    if (number1 > number2) && (number1 > number3) {
        fmt.Printf("%v is the largest number\n", number1)
    } else if (number2 > number1) && (number2 > number3) {
        fmt.Printf("%v is the largest number\n", number2)
    } else if (number3 > number1) && (number3 > number2) {
        fmt.Printf("%v is the largest number\n", number3)
    } else {
        fmt.Printf("Numbers are equal\n")
    }
}
```

---

## 22. A Go Program to Spell Digits

```go
// a go program to spell digits with if, else if, else
package main

import "fmt"

func main() {
    var digit int
    fmt.Printf("Enter a digit (0-9): ")
    fmt.Scanf("%v", &digit)

    if digit == 0 {
        fmt.Printf("Zero\n")
    } else if digit == 1 {
        fmt.Printf("One\n")
    } else if digit == 2 {
        fmt.Printf("Two\n")
    } else if digit == 3 {
        fmt.Printf("Three\n")
    } else if digit == 4 {
        fmt.Printf("Four\n")
    } else if digit == 5 {
        fmt.Printf("Five\n")
    } else if digit == 6 {
        fmt.Printf("Six\n")
    } else if digit == 7 {
        fmt.Printf("Seven\n")
    } else if digit == 8 {
        fmt.Printf("Eight\n")
    } else if digit == 9 {
        fmt.Printf("Nine\n")
    } else {
        fmt.Printf("Not a digit\n")
    }
}
```

---

## 23. switch Control Statement

- Same as C, C++, Java, JS but only difference you do not use break statement. Only run the matched statement.

```go
// a go program to spell digits with switch statement
package main

import "fmt"

func main() {
    var digit int
    fmt.Printf("Enter a digit (0-9): ")
    fmt.Scanf("%v", &digit)

    switch digit {
    case 0:
        fmt.Printf("Zero\n")
    case 1:
        fmt.Printf("One\n")
    case 2:
        fmt.Printf("Two\n")
    case 3:
        fmt.Printf("Three\n")
    case 4:
        fmt.Printf("Four\n")
    case 5:
        fmt.Printf("Five\n")
    case 6:
        fmt.Printf("Six\n")
    case 7:
        fmt.Printf("Seven\n")
    case 8:
        fmt.Printf("Eight\n")
    case 9:
        fmt.Printf("Nine\n")

    // multiple cases possible
    // case 0, 1, 2, 3, 4:
    //   fmt.Printf("less than 5")
    default:
        fmt.Printf("Not a digit\n")
    }
}
```

---

## 24. Loop Control Statement

- In Go we have only 1 loop which is `for` loop; can take form of `while` loop as well
- It helps us to execute some statements again and again

---

## 25. Series Related Programs

```go
package main

import "fmt"

func main() {
    // series -> 2 + 4 + 6 + 8 + ... N
    var startNumber, endNumber int

    fmt.Printf("Enter the starting number of series: ")
    fmt.Scan(&startNumber)

    fmt.Printf("Enter the last number of series: ")
    fmt.Scan(&endNumber)

    sum := 0
    for i := startNumber; i <= endNumber; i = i + 2 {
        sum = sum + i
    }
    fmt.Printf("sum = %v\n", sum)
}
```

---

## 26. break and continue Statement

```go
package main

import "fmt"

func main() {
    // break, continue statement

    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Printf("%v\n", i)
    }
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            break
        }
        fmt.Printf("%v\n", i)
    }
}
```

---

## 27. Functions

```go
package main

import "fmt"

func square(number int) int {
    return number * number
}

func displayMessage(country string) {
    fmt.Printf("I love my motherland: %v\n", country)
}

func main() {
    displayMessage("Bangladesh")
    displayMessage("India")
    fmt.Printf("Square of 6 = %v\n", square(6))
    fmt.Printf("Square of 7 = %v\n", square(7))
}
```

---

## 28. Create a Function for Calculator

```go
package main

import "fmt"

func add(x, y float32) float32 {
    return x + y
}

func sub(x, y float32) float32 {
    return x - y
}

func mul(x, y float32) float32 {
    return x * y
}

func div(x, y float32) float32 {
    return x / y
}

func main() {
    var num1, num2, result float32
    var option string

    i := true
    for i == true {
        fmt.Printf("num1 = ")
        fmt.Scan(&num1)

        fmt.Printf("num2 = ")
        fmt.Scan(&num2)

        fmt.Printf("choose an option ( + - * / ) : ")
        fmt.Scan(&option)

        switch option {
        case "+":
            result = add(num1, num2)
        case "-":
            result = sub(num1, num2)
        case "*":
            result = mul(num1, num2)
        case "/":
            result = div(num1, num2)
        default:
            fmt.Println("Invalid option")
            continue
        }
        fmt.Printf("Result = %v\n", result)
    }
}
```

---

## 29. Pointers

- **dereferencing operator**
- A pointer is a variable whose value is the address of another variable
- **declaration**: `var varName *varType`

```go
package main

import "fmt"

func change(val int) {
    val = 8
}

func changeWihPointer(val *int) {
    *val = 8
}

func main() {
    x := 10
    change(x)
    fmt.Println(x)

    // call by reference
    changeWihPointer(&x)
    fmt.Println(x)
}
```

---

## 30. Structure

- `class` is not supported here but we can use `structs`

### Example

```go
package main

import "fmt"

type Student struct {
    name string
    age  int
    id   int
}

func main() {
    <!-- type Student struct{
      name string
      age int
      id int
    } -->

    rahim := Student{"Anisul", 32, 1302020017}
    rahim.age = 30
    fmt.Println(rahim)
    fmt.Println(rahim.id)
    fmt.Println(rahim.name)
    fmt.Println(rahim.age)
}
```

### Passing Struct in a Function

```go
// passing struct in a function
package main

import "fmt"

type Student struct {
    name string
    age  int
    id   int
}

func displayInfo(s Student) {
    fmt.Println(s.name)
    fmt.Println(s.age)
}

func main() {
    s1 := Student{"Anisul", 32, 1302020017}
    displayInfo(s1)
}
```

### If We Need to Change Data

```go
// if we need to change data
package main

import "fmt"

type Student struct {
    name string
    age  int
    id   int
}

func (x *Student) increaseAge(val int) {
    x.age += val
}

func main() {
    s1 := Student{"Anisul Islam", 32}
    s1.increaseAge(1)
    fmt.Println(s1.age)
}
```

```go
// timer app
// Employee management app
```

---

## 31. Arrays & Slices Data Types

- Array is fixed in Go
- `var arrayName[size] dataType`
- Only the same data type in one array
- `var array_name = [length]datatype{values}` // here length is defined
- Array index starts with 0
- What if we dont know the size of array then we will use slice; allows us to use dynamic size
- `append(sliceName, sliceNewValue)`