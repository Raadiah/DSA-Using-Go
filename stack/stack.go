package stack

import (
	"bufio"
	"os"

	"github.com/fatih/color"
)

func writeActionValuesOnTerminal() {
	color.Cyan("\nEnter command: Exit[x/X], Insert [i/I], Delete [d/D], Peek [p/P]")
}

func RunStack() {
	color.Cyan("\nType 'l/L' for Linked List, 'a/A' for Array, 'x/X' to exit")

	reader := bufio.NewReader(os.Stdin)

forLoop:
	for char, _, err := reader.ReadRune(); char != 'X' && char != 'x'; char, _, err = reader.ReadRune() {
		if err != nil {
			color.Red("An error occurred while getting input")
		}

		switch char {
		case 'A', 'a':
			arrayStack()
			break forLoop
		case 'L', 'l':
			linkedListStack()
			break forLoop
		case '\n':
			continue forLoop
		default:
			color.Red("Invalid Command")
		}
	}
}
