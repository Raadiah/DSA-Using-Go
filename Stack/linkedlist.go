package stack

import (
	"github.com/Raadiah/DSA-Using-Go/theme"
)

type linkedListStack struct {
	top *stackNode
}

type stackNode struct {
	data     any
	prevNode *stackNode
}

func (stack *linkedListStack) Insert(item any) {
	newNode := stackNode{item, stack.top}
	stack.top = &newNode
}

func (stack *linkedListStack) Delete() bool {
	if stack.top != nil {
		stack.top = stack.top.prevNode
	}
	return stack.top != nil
}

func (stack *linkedListStack) Top() any {
	return stack.top.data
}

func (stack *linkedListStack) checkTop() bool {
	if stack.top == nil {
		theme.ErrorMessage("The stack is empty.")
	}
	return stack.top != nil
}

func createLinkedListStack() *linkedListStack {
	stack := linkedListStack{nil}
	return &stack
}
