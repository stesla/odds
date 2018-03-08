package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	numDice   = flag.Int("dice", 1, "number of dice to roll")
	diff      = flag.Int("diff", 6, "difficulty to roll against")
	numTrials = flag.Int("trials", 100000, "number of trials to make")
	maxTrials = flag.Int("maxTrials", 10000, "maximum number of trials to use per thread")
)

type result struct {
	botch bool
	succs int
}

func main() {
	flag.Parse()

	ch := make(chan result, *numTrials)
	var wg sync.WaitGroup
	for n := *numTrials; n > 0; n -= *maxTrials {
		i := *maxTrials
		if n < i {
			i = n
		}
		wg.Add(1)
		go doTrials(i, &wg, ch)
	}
	wg.Wait()
	close(ch)

	var successes, botches, failures, totalSuccesses int
	for r := range ch {
		if r.botch {
			botches++
		} else if r.succs == 0 {
			failures++
		} else {
			successes++
			totalSuccesses += r.succs
		}
	}

	fmt.Printf("successes = %0.2f%%, failures = %0.2f%%, botches = %0.2f%%, average successes = %0.2f",
		100.0*float64(successes)/float64(*numTrials),
		100.0*float64(failures)/float64(*numTrials),
		100.0*float64(botches)/float64(*numTrials),
		float64(totalSuccesses)/float64(successes))
}

func doTrials(count int, wg *sync.WaitGroup, ch chan result) {
	defer wg.Done()
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	for i := 0; i < count; i++ {
		var roll result
		dice := Roll(r, *numDice)
		roll.succs, roll.botch = Interpret(*diff, dice)
		ch <- roll
	}
}
