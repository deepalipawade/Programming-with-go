package fibonacci

import "fmt"

func main() {
	fmt.Println("Fibonacci, recursive n5 = ", fibo(5))
	fmt.Println("Fibonacci memorization n5 = ", memo(5))
	fmt.Println("Fibonacci memorization Optmized n5 = ", memoOptimized(5))
	fmt.Println("Fibonacci Using formulae n5 = ", fiboMath(5))
}

// recursive approach ... Time complexity -> O(n)
func fibo(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

// memorization ...  Time complexity -> O(n) space -> O(n)
var fiboMap = map[int]int{}

func memo(n int) int {
	fiboMap[0] = 0
	fiboMap[1] = 1
	for i := 2; i <= n; i++ {
		fiboMap[i] = fiboMap[i-1] + fiboMap[i-2]
	}
	return fiboMap[n]
}

// Optimizing memorization techinique by storing only previous 2 values ... Time complexity -> O(n) space -> O(1)
func memoOptimized(n int) int {
	a, b, c := 0, 1, 0
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

// Using mathematical formulae ... for power

var fibo1 = map[int]int{}

func fiboMath(n int) int {

	// base case
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		fibo1[n] = 1
		return 1
	}
	var k int
	if n%2 == 0 {
		// n is even
		k = n / 2
		fibo1[n] = (2*fiboMath(k-1) + fiboMath(k)) * fiboMath(k)
	} else {
		// n is odd
		k = (n + 1) / 2
		fibo1[n] = (fiboMath(k)*fiboMath(k) + fiboMath(k-1)*fiboMath(k-1))
	}

	return fibo1[n]
}
