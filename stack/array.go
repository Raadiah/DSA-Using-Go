package stack

import (
	"github.com/Raadiah/DSA-Using-Go/theme"
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

func createArrayStack() *arrayStack {
	stack := arrayStack{-1, nil}
	return &stack
}

func GetArrayStack () stacker {
	arrayStack := createArrayStack()
	var stack stacker = arrayStack
	return stack
}
