package main

import "fmt"

func main() {
	fmt.Println(SieveOfEratosthenes(101))
}

func SieveOfEratosthenes(n uint) []uint {

	var isPrime []bool

	for i := 0; i <= n; i++ {
		isPrime[i] = false
	}
	isPrime[0] = true
	isPrime[1] = true

	for i, v := range isPrime {
		if v {
			for k := uint(2); uint(i)*k <= n; k++ {
				isPrime[uint(i)*k] = true
			}
		}
	}

	var primes []uint
	for p, _ := range isPrime {
		primes = append(primes, uint(p))
	}

	return primes
}

func TwinPrimes(n){

}