package stack

import (
	"bufio"
	"fmt"
	"os"
)

func RunStack() {
	fmt.Println("Initializing Stack..")
	fmt.Println("Do you want to build using Linked List or Array?")
	fmt.Println("Type 'L' for Linked List, 'A' for Array, 'X' to exit")

	reader := bufio.NewReader(os.Stdin)

	for char, _, err := reader.ReadRune(); char != 'X'; char, _, err = reader.ReadRune() {
		if err != nil {
			fmt.Println("An error occurred while getting input")
		}

		if char == 'A' {
			fmt.Println("Preparing your Array Stack")
			arrayStack()
			break
		} else if char == 'L' {
			fmt.Println("Preparing your Linked List Stack")
			linkedListStack()
			break
		} else if char == '\n' {
			continue
		} else {
			fmt.Println("Invalid Input")
		}
	}

	fmt.Println("Exiting Stack")
}
