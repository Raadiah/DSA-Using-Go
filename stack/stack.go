package stack

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Raadiah/DSA-Using-Go/theme"
)

type stacker interface {
	Top() any
	Delete() bool
	Insert(item any)
	checkTop() bool
}

func writeActionValuesOnTerminal() {
	theme.WriteInstruction("\nEnter command: Exit[x/X], Insert [i/I], Delete [d/D], Peek [p/P]")
}

func runStack(stack stacker) {
	reader := bufio.NewReader(os.Stdin)

	writeActionValuesOnTerminal()

	for char, _, err := reader.ReadRune(); char != 'X' && char != 'x'; char, _, err = reader.ReadRune() {
		if err != nil {
			fmt.Println(err)
		}

		switch char {
		case 'I', 'i':
			var item string
			theme.WriteInstruction("Enter numeric value to append to stack: ")
			_, e := fmt.Scanf("%s", &item)
			if e != nil {
				fmt.Println("Error ", e)
			} else {
				stack.Insert(item)
				theme.SuccessMessage("Appended successfully")
			}
		case 'P', 'p':
			if stack.checkTop() {
				top := stack.Top()
				theme.SuccessMessage("Current top: %v", top)
			}
		case 'D', 'd':
			if stack.checkTop() {
				stack.Delete()
				theme.SuccessMessage("Deleted successfully")
			}

		case 'H', 'h':
			writeActionValuesOnTerminal()
		default:
			theme.ErrorMessage("Invalid Command.")
		}

		reader.Discard(reader.Buffered())

		writeActionValuesOnTerminal()
	}
}

func ChooseStack() {
	theme.WriteInstruction("\nType 'l/L' for Linked List Stack, 'a/A' for Array Stack, 'x/X' to exit stack functions")

	reader := bufio.NewReader(os.Stdin)

	for char, _, err := reader.ReadRune(); char != 'X' && char != 'x'; char, _, err = reader.ReadRune() {
		if err != nil {
			theme.ErrorMessage("An error occurred while getting input")
		}

		switch char {
		case 'A', 'a':
			var arrayStack = createArrayStack()
			theme.SuccessMessage("Created array stack. Enter commands to operate on the stack.")
			runStack(arrayStack)
			theme.SuccessMessage("Exiting array stack")
		case 'L', 'l':
			var linkedListStack = createLinkedListStack()
			theme.SuccessMessage("Created linked list stack. Enter commands to operate on the stack.")
			runStack(linkedListStack)
			theme.SuccessMessage("Exiting linked list stack")
		default:
			theme.ErrorMessage("Invalid Command")
		}

		reader.Discard(reader.Buffered())
		theme.WriteInstruction("\nType 'l/L' for Linked List Stack, 'a/A' for Array Stack, 'x/X' to exit stack functions")
	}
}
