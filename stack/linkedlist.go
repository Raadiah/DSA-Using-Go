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
		return true
	}
	return false
}

func (stack *linkedListStack) Top() any {
	if stack.top == nil {
		return nil
	}
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

func GetLinkedListStack() stacker {
	arrayStack := createLinkedListStack()
	var stack stacker = arrayStack
	return stack
}
