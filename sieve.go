package main

import "fmt"

func main() {
	fmt.Println(SieveOfEratosthenes(20))
}

func SieveOfEratosthenes(n uint) []uint {
	isPrime:=make([]bool, n+1)
	
	for i := 0; uint(i) <=n; i++ {
		isPrime[i] = true
	}
	for j, v := range isPrime {
		if v {
			j+=2
			for k := uint(2); uint(j)*uint(k) <=n; k+=uint(j) {
				isPrime[uint(j)*k] = false
			}
		}
	}

	var primes []uint
	for p, _ := range isPrime {
		if(isPrime[p]==true){
			primes = append(primes, uint(p))
		}
		
	}
//	TwinPrimes(primes)
	return primes
}

//func TwinPrimes([]uint) []uint{

//}

