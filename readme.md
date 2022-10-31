# Trying out GO ... !!! 

![Golang image](./go.png "GO")

#### Hello GO ...
```go 
package main 
include "fmt"

func main(){
    fmt.Println("Say Hi to GO ")
}

```
#### diffrent printing methods 
<p> fmt.Println("hello go") --> print the sring and moves cursor to the new line 
</p>

<p> print("hello go")
<br>
fmt.Printf("hello go") //formatted printing
</p>

```go
var num1=5
fmt.Println("hello go")
print("Hello GO")
fmt.Printf("hello GO %v",num1)

```



#### understanding variables and namings
``` go 

   variables can be declared using var keyword followed by variable name and the type of varaible
   var num1 int = 10  //int
   var name string = "kushal"  //string
   var num2 float64=9.55  //float
   var state bool = false  //boolean
```

#### The implicit declaration of variable type

``` go
	name := "kushal"  // string	
    number :=5        // integer	
    number2 :=8.8     // float	
    this is useful, since we can define variable without explicitly  defining their datatypes


```

```go 
the printf and Scan

   val := 5
   name:="kushal"
   fmt.Printf("Name is %v ",name)
   fmt.Printf("the value is %v ",val)
  // --> the %v refers to the value of the variable
   // taking user input in go
   fmt.Scan()
   var name string
   fmt.Scan(&name)  //-> taking user input and storing in the memory location
   fmt.Printf("name is %v", name)
   var lastname string 
   fmt.Printf("Enter your last name")
   fmt.Scan(&lastname)
   fmt.Printf("your last name is %v",lastname)

```
```go 

name := "kushal"
surname := "subedi"
fmt.Println("Full Name is", name+"", surname)
val1 := 2
val2 := 5
fmt.Printf("%v \n", val1+val2)
```

#### <u> conditional statement </u>

```go 

var age int = 20
if age >= 20 {
	fmt.Printf(" you are an adult \n")
} else { // else should be started right where if ends
	fmt.Printf("you are a kid")
}
```
#### Scan() & Scanf()
```go 

var fullname string
var fullname2 string 
fmt.Printf("Enter your full name  ")
fmt.Scan(&fullname)
fmt.Printf("fullname = %v",fullname)

fmt.Printf("Again Enter your full name ")
fmt.Scanf("%s",&fullname2)
fmt.Printf("full name =%v",fullname2)
// when you use Scan() & Scanf() it reads only words 
// if your string have multiple words like "Kushal Subedi"
// the Scan() & Scanf() reads only first word and store that in the variabble and reads another word in another variable 
```

``` go
// we can use bufio package to read entire line with white spaces 

package main
import (
    "bufio"
 
    "fmt"
    "os"
)
func main(){
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
    input := scanner.Text()
    fmt.Println(input)
}


```

