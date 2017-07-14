package main

import "fmt"

type Queue struct{
	data []int
}

func (q *Queue) queue(k int){
	q.data = append(q.data,k);
}

func (q *Queue) dequeue() int{
	last := q.peek();
	q.data = q.data[1:];
	return last;
}

func (q *Queue) peek() int{
	return q.data[0];
}

func main(){
	q := Queue{};

	for i:=0;i<5;i++{
		q.queue(i)
	}

	for len(q.data)>0{
		fmt.Println(q.dequeue(),"Len",len(q.data))
	}
}