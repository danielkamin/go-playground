package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func mutexes() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, inc int) {
		for range inc {
			c.inc(name)
		}
	}

	wg.Go(func() {
		doIncrement("a", 1002)
	})
	wg.Go(func() {
		doIncrement("a", 987)
	})
	wg.Go(func() {
		doIncrement("b", 2001)
	})

	wg.Wait()
	fmt.Println(c.counters)
}
