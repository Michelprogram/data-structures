package main

import (
	"fmt"
	"os"
)

func main() {

	stackInt := NewStack[int](3)

	fmt.Println(stackInt)
	value, err := stackInt.Push(1)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Println(stackInt)
	value, err = stackInt.Push(2)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Println(stackInt)
	value, err = stackInt.Push(3)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Println(stackInt)
	value, err = stackInt.Push(4)

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Println(stackInt.Peek())

	value, err = stackInt.Pop()
	value, err = stackInt.Pop()
	value, err = stackInt.Pop()

	value, err = stackInt.Pop()

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Removed : %v\n", value)

	value, err = stackInt.Push(5)

	fmt.Println(stackInt)

}
