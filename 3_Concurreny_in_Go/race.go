package main

import (
	"fmt"
	"time"
)

var count int

func race() {
	fmt.Println("count = ",count)
	count++
}

func main() {

	count = 1
	go race()
	go race()
	time.Sleep(1)
}


/* Explanation
------------------------------------------------------------------------------------------------------
Race Condition

A race condition occurs when two goroutines access the same variable concurrently, and at least one of the access is a write instruction.

To detect a race condition:
Use the command:
	go run -race routineOne.go to generate the race report

If there is no race condition, it should print count = 2
but here it shows count = 1, both times.
*/
