package main

import "fmt"

type Stack struct{
	data []int
}

func (stack *Stack) push(k int){
	stack.data = append(stack.data,k)
}

func (stack *Stack) pop() int{
	last:= stack.peek();
	stack.data = stack.data[:len(stack.data)-1]
	return last
}

func (stack *Stack) peek() int{
	index := len(stack.data)-1
	last := stack.data[index]
	return last
}

func (stack *Stack) printSlice(){
	fmt.Println("Size : ",len(stack.data))
	for i:=0;i<len(stack.data);i++{
		fmt.Println(stack.data[i]);
	}
	fmt.Println();
}

func main(){
	stack := Stack{};
	stack.printSlice();
	stack.push(1);
	stack.printSlice();
	defer stack.printSlice();
	
	for i:=0;i<5;i++{
		defer stack.push(i);
	}

	fmt.Println("Pop ",stack.pop());
}