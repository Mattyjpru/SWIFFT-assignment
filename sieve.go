package main

import "fmt"

func main() {
	fmt.Println(SieveOfEratosthenes(101))
}

func SieveOfEratosthenes(n uint) []uint {

	var isPrime []bool
	
	for i := 0; uint(i) <= n; i++ {
		isPrime[i] = false
	}
	isPrime[0] = true
	isPrime[1] = true

	for j, v := range isPrime {
		if v {
			for k := uint(2); uint(j)*uint(k) <= n; k++ {
				isPrime[uint(j)*k] = true
			}
		}
	}

	var primes []uint
	for p, _ := range isPrime {
		primes = append(primes, uint(p+1))
	}

	return primes
}

//func TwinPrimes(n){

//}