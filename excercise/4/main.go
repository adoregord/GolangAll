package main

import (
	"fmt"
)

// func Factorial(n int64) *big.Int {
// 	result := big.NewInt(1)
// 	for i := int64(n); i > 0; i-- {
// 		result.Mul(result, big.NewInt(i))
// 	}
// 	return result
// }

//	func main() {
//		n := int64(30)
//		factorial := Factorial(n)
//		fmt.Printf("%d! = %s\n", n, factorial)
//	}
func main() {
	value := float64(1)
	for i := float64(1); i <= 30; i++ {
		value *= i
	}
	fmt.Printf("%.00f", value)
}
