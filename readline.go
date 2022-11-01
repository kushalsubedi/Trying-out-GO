package main

import (
	"bufio"

	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	input := scanner.Text()
	fmt.Println(input)
}
