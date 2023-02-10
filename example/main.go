package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	primes "github.com/bminer/go-primer"
)

func usage(err error) {
	fmt.Println("error:", err)
	fmt.Println("usage: " + os.Args[0] + " [nth_prime]\n" +
		"\tComputes the n-th prime.")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		usage(errors.New("invalid argument count"))
	}
	n64, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		usage(err)
	}
	n := uint(n64)
	fmt.Println(n, "-th prime:", primes.Nth(n))
}
