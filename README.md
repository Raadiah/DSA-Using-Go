# Implementation of Different Data Structures and Algorithms using Go

In this project, we will implement different Data Structures and Algorithms using Go. We will be using the ***Fundamentals of Computer Algorithms***  by *E. Horowitz et.al* as a reference book.

This will be a command line tool programming. Also, you can import packages to your project.

#### Import Package

To add the module into your package enter into your project directory terminal the following code:

```go get github.com/Raadiah/DSA-Using-Go@v0.5.0```

Then using the following imports you can import the respective packages

### Data Structures

#### [Stack](/stack/stack.md)
```import "github.com/Raadiah/DSA-Using-Go@v0.5.0/stack"```

Then any stack operations can be done.

- stack.GetArrayStack() - To create an array stack
- stack.GetLinkedListStack() - To create a linked list stack

From a stacker object created from the above methods, the following methods can be called:

- Top() - To get top element. Nil is returned if no top is set
- Delete() - Delete the top element. Returns true for deleted, false for empty stack
- Insert(item) - Insert element to the top
