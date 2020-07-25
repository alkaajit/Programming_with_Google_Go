package main

import (
	"fmt"
	"sync"
	"time"
)

//ChopS struct
type ChopStick struct{ sync.Mutex }

//Philo struct
type Philosopher struct {
	leftCS, rightCS *ChopStick
	number          int
}

var on sync.Once

func setup() {
	fmt.Println("Init philosopher")
}

func host(ch chan int) {
	ch <- 1
	ch <- 2
	<-ch
}

func (p Philosopher) eat(ch chan int) {

	on.Do(setup)

	for i := 0; i < 3; i++ {

		<-ch

		fmt.Printf("starting to eat %d\n", p.number)

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("finishing eating %d\n", p.number)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		ch <- i
	}
	return
}

func main() {
	ch := make(chan int, 2)

	chop_stick := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chop_stick[i] = new(ChopStick)
	}

	phil := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		phil[i] = &Philosopher{chop_stick[i], chop_stick[(i+1)%5], i + 1}
	}

	go host(ch)

	for i := 0; i < 5; i++ {
		go phil[i].eat(ch)
	}

	time.Sleep(1000 * time.Millisecond)
}
