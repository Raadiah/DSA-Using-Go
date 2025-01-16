package main

import (
	"bufio"
	"fmt"
	"os"

	"algorithms.com/stack"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter 's' to start Program")

	for char, err := reader.ReadByte(); char != 's'; char, err = reader.ReadByte() {
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Do you want to run stack? Type 'y' or 'n'")

	for char, _, err := reader.ReadRune(); char != 'n'; char, _, err = reader.ReadRune() {
		if err != nil {
			fmt.Println(err)
		}

		if char == 'y' {
			stack.LinkedListStack()
			stack.ArrayStack()
			fmt.Println("Do you want to run stack again? Type 'y' or 'n'")
		} else if char == '\n' {
			continue
		} else {
			fmt.Println("Invalid Input")
		}
	}

	fmt.Println("Exiting program...")
}
