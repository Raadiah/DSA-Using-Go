package stack

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func checkTop(top int) bool {
	if top < 0 {
		color.Red("The stack is empty.")
	}
	return top >= 0
}

func arrayStack() {
	color.Green("Created Linked List Stack. Enter commands to operate on the stack.")
	reader := bufio.NewReader(os.Stdin)
	top := -1
	var stack []int

	writeActionValuesOnTerminal()

forLoop:
	for char, _, err := reader.ReadRune(); char != 'X' && char != 'x'; char, _, err = reader.ReadRune() {
		if err != nil {
			fmt.Println(err)
		}

		switch char {
		case '\n':
			continue forLoop
		case 'I', 'i':
			top++
			var num int
			color.Cyan("Enter numeric value to append to stack: ")
			_, e := fmt.Scanf("%d", &num)
			if e != nil {
				print(e)
			}
			stack = append(stack, num)
			color.Green("Appended successfully")
		case 'P', 'p':
			if checkTop(top) {
				color.Green("Top value: %d", stack[top])
			}
		case 'D', 'd':
			if checkTop(top) {
				top--
				stack = stack[:len(stack)-1]
				color.Green("Deleted from stack")
			}
		case 'H', 'h':
			writeActionValuesOnTerminal()
		default:
			color.Red("Invalid Command.")
		}

		writeActionValuesOnTerminal()
	}
	color.Green("Exiting array stack")
}
