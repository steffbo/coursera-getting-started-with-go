package main

import (
	"fmt"
	"sync"
	"time"
)

const TimesToEat = 3
const ChopstickAndPhilosopherCount = 5

var eatWg sync.WaitGroup
var host = make(chan bool, 2)

func main() {

	chopsticks := make([]*Chopstick, ChopstickAndPhilosopherCount)
	for i := range chopsticks {
		chopsticks[i] = new(Chopstick)
	}

	// give some personality
	icons := []string{"🧐", "🤓", "🤪", "🥳", "🥴"}

	philosophers := make([]*Philosopher, ChopstickAndPhilosopherCount)
	for i := range philosophers {
		philosophers[i] = &Philosopher{
			num:   i + 1,
			icon:  icons[i],
			left:  chopsticks[i],
			right: chopsticks[(i+1)%ChopstickAndPhilosopherCount],
		}
		eatWg.Add(1)
		go philosophers[i].Eat()
	}

	eatWg.Wait()
}

func (p Philosopher) Eat() {
	defer eatWg.Done()
	for i := 0; i < TimesToEat; i++ {
		host <- true

		p.left.Lock()
		p.right.Lock()
		p.count++
		fmt.Printf("%s starting to eat %d (%d/%d)\n", p.icon, p.num, p.count, TimesToEat)
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Printf("%s finishing eating %d (%d/%d)\n", p.icon, p.num, p.count, TimesToEat)
		p.right.Unlock()
		p.left.Unlock()

		<-host
	}
}

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	num, count  int
	icon        string
	left, right *Chopstick
}
