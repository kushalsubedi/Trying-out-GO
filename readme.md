# Trying out GO ... !!! 

![Golang image](./golang.png "GO").

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

	   var num1 int = 10
	   var name string = "kushal"
	   var num2 float64=9.55
	   var state bool = false
```

#### The implicit declaration of variable type

``` go

		name := "kushal"  // string
		number :=5        // integer
		number2 :=8.8     // float

		this is useful, since we can define 
                variable without explicitly  defining their datatypes


```