package stack

import (
	"bufio"
	"os"

	"algorithms.com/theme"
)

type stackMethods interface {
	Top() any
	Delete() bool
	Insert(item any)
	checkTop() bool
}

func writeActionValuesOnTerminal() {
	theme.WriteInstruction("\nEnter command: Exit[x/X], Insert [i/I], Delete [d/D], Peek [p/P]")
}

func RunStack() {
	theme.WriteInstruction("\nType 'l/L' for Linked List, 'a/A' for Array, 'x/X' to exit")

	reader := bufio.NewReader(os.Stdin)

forLoop:
	for char, _, err := reader.ReadRune(); char != 'X' && char != 'x'; char, _, err = reader.ReadRune() {
		if err != nil {
			theme.ErrorMessage("An error occurred while getting input")
		}

		switch char {
		case 'A', 'a':
			runArrayStack()
			break forLoop
		case 'L', 'l':
			linkedListStack()
			break forLoop
		default:
			theme.ErrorMessage("Invalid Command")
		}

		reader.Discard(reader.Buffered())
	}
}
