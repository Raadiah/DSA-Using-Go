package stack

import (
	"bufio"
	"fmt"
	"os"

	"algorithms.com/theme"
)

type arrayStack struct {
	top   int
	array []any
}

func (stack *arrayStack) Top() any {
	return stack.array[stack.top]
}

func (stack *arrayStack) Delete() bool {
	if stack.top < 0 {
		return false
	}
	stack.top--
	stack.array = stack.array[:len(stack.array)-1]
	return true
}

func (stack *arrayStack) Insert(item any) {
	stack.top++
	stack.array = append(stack.array, item)
}

func (stack *arrayStack) checkTop() bool {
	if stack.top < 0 {
		theme.ErrorMessage("The stack is empty.")
	}
	return stack.top >= 0
}

func runArrayStack() {
	theme.SuccessMessage("Created Linked List Stack. Enter commands to operate on the stack.")
	stack := arrayStack{-1, nil}
	var arrayStackMethods stackMethods = &stack
	reader := bufio.NewReader(os.Stdin)

	writeActionValuesOnTerminal()

	for char, _, err := reader.ReadRune(); char != 'X' && char != 'x'; char, _, err = reader.ReadRune() {
		if err != nil {
			fmt.Println(err)
		}

		switch char {
		case 'I', 'i':
			var num int
			theme.WriteInstruction("Enter numeric value to append to stack: ")
			_, e := fmt.Scanf("%d", &num)
			if e != nil {
				print(e)
			}
			arrayStackMethods.Insert(num)
			theme.SuccessMessage("Appended successfully")
		case 'P', 'p':
			if arrayStackMethods.checkTop() {
				theme.SuccessMessage("Current top: ")
				fmt.Println(arrayStackMethods.Top())
			}
		case 'D', 'd':
			if arrayStackMethods.checkTop() {
				arrayStackMethods.Delete()
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
	theme.SuccessMessage("Exiting array stack")
}
