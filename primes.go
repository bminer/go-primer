package primes

import (
	"math"
)

// Odds writes a sequence of odd numbers starting with `start` to the channel `ch`
func Odds(start uint, ch chan<- uint) {
	for i := start; ; i += 2 {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Filter copies the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan uint, out chan<- uint, prime uint) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// Nth computes the n-th prime using a daisy-chain filter process, which
// filters composite numbers concurrently
func Nth(n uint) uint {
	// Create a new channel and write 2, the only even prime number
	ch := make(chan uint, 10)
	ch <- 2
	// Launch Odds goroutine
	go Odds(3, ch)

	var prime uint
	for i := uint(0); i < n; i++ {
		// Read a prime number from `ch`
		prime = <-ch
		if prime > 2 && float64(i) < math.Sqrt(float64(n)) {
			// Filter to ensure that `ch` always returns prime numbers
			ch1 := make(chan uint, 100)
			go Filter(ch, ch1, prime)
			ch = ch1
		}
	}
	return prime
}
