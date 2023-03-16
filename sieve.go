package main

import "fmt"

func main() {
	fmt.Println(SieveOfEratosthenes(101))
	fmt.Println(TwinPrimes(101))
}

func SieveOfEratosthenes(n uint) []uint {
	isPrime:=make([]bool, n+1)
	
	for i := 0; uint(i) <=n; i++ {
		isPrime[i] = true
	}

	for j, v := range isPrime {
		if v {
			j+=2
			for k := uint(2); uint(j)*uint(k) <=n; k++ {
				isPrime[uint(j)*k] = false
			}
		}
	}

	var primes []uint
	for p, _ := range isPrime {
		if(isPrime[p]){
			primes = append(primes, uint(p))
		}
		
	}
	return primes
}

func TwinPrimes(n uint) []uint{
	primes:=SieveOfEratosthenes(n)
	var twins []uint
	

	for t, _ := range primes {
		if((primes[t]+2)==primes[t+1]){
			primes = append(twins, uint(t))
		}
	}
	return twins
}

