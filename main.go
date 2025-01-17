package main

import (
	"bufio"
	"fmt"
	"os"

	stack "algorithms.com/Stack"
	"algorithms.com/theme"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	theme.WriteInstruction("Enter 's/S' to run stack, 'x/X' to exit")

	for char, _, err := reader.ReadRune(); char != 'x' && char != 'X'; char, _, err = reader.ReadRune() {
		if err != nil {
			fmt.Println(err)
		}

		if char == 's' || char == 'S' {
			theme.SuccessMessage("Intitalizing Stack...")
			stack.ChooseStack()
			break
		} else {
			theme.ErrorMessage("Invalid Command")
			theme.WriteInstruction("Enter 's' to run stack, 'x' to exit")
		}
		reader.Discard(reader.Buffered())
	}

	theme.SuccessMessage("Exiting program...")
}
