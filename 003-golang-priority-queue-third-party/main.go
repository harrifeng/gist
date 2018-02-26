package main

import (
	"fmt"
	"os"

	lane "gopkg.in/oleiade/lane.v1"
)

func main() {

	// Let's create a new max ordered priority queue
	var priorityQueue *lane.PQueue = lane.NewPQueue(lane.MINPQ)
	// And push some prioritized content into it
	priorityQueue.Push("easy as", 3)
	priorityQueue.Push("123", 2)
	priorityQueue.Push("do re mi", 4)
	priorityQueue.Push("abc", 1)
	// Now let's take a look at the min element in
	// the priority queue
	for i := 0; i < 5; i++ {
		headValue, headPriority := priorityQueue.Pop()
		fmt.Println(headValue, headPriority) // 1
	}

	fmt.Println()
	os.Exit(0)
}

// <===================OUTPUT===================>
// abc 1
// 123 2
// easy as 3
// do re mi 4
// <nil> 0
