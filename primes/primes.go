package primes

import (
	"math"
)

// Seq generates the sequence 2, 3, 4, 5, ... and writes to the channel `ch`
func Seq(ch chan<- uint) {
	for i := uint(2); ; i++ {
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
	// Create a new channel
	ch := make(chan uint, 10)
	// Launch Seq goroutine
	go Seq(ch)

	var prime uint
	for i := uint(0); i < n; i++ {
		// Read a prime number from `ch`
		prime = <-ch
		if float64(i) < math.Sqrt(float64(n)) {
			// Filter to ensure that `ch` always returns prime numbers
			ch1 := make(chan uint, 100)
			go Filter(ch, ch1, prime)
			ch = ch1
		}
	}
	return prime
}
