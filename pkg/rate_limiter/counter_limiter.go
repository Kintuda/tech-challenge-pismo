package ratelimiter

import (
	"sync"

	"golang.org/x/time/rate"
)

type CounterLimiter struct {
	mutext  *sync.RWMutex
	counter map[string]*rate.Limiter
}

func NewCounterLimiter() *CounterLimiter {
	return &CounterLimiter{
		counter: map[string]*rate.Limiter{},
		mutext:  &sync.RWMutex{},
	}
}

func (c *CounterLimiter) GetInstance(ip string) *rate.Limiter {
	c.mutext.Lock()

	counter, ok := c.counter[ip]

	if !ok {
		limiter := rate.NewLimiter(1, 5)
		c.counter[ip] = limiter
		c.mutext.Unlock()

		return limiter
	}

	defer c.mutext.Unlock()

	return counter
}
