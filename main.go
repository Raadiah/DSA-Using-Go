package main

import (
	"bufio"
	"fmt"
	"os"

	stack "algorithms.com/Stack"
	"github.com/fatih/color"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	color.Cyan("Enter 's/S' to run stack, 'x/X' to exit")

	for char, _, err := reader.ReadRune(); char != 'x' && char != 'X'; char, _, err = reader.ReadRune() {
		if err != nil {
			fmt.Println(err)
		}

		if char == 's' || char == 'S' {
			color.Green("Intitalizing Stack...")
			stack.RunStack()
			break
		} else if char == '\n' {
			continue
		} else {
			color.Red("Invalid Command")
			color.Yellow("Enter 's' to run stack, 'x' to exit")
		}
	}

	color.Green("Exiting program...")
}
