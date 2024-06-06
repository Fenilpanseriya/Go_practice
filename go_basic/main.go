package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")
	var name = "fenil"
	fmt.Printf("my name is %s\n", name)
	fmt.Println("which language is you learning")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("hello", input)
	hello()
}
