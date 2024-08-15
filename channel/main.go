package main

import (
	"fmt"
	"runtime"
)

// membuat fungsi  untuk mencari nilai tertinggi :)
func getMax(nums []int, ch chan int) {
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	ch <- max
}

// / membuat fungsi untuk mencari rata-rata
func getAverage(nums []int, ch chan float64) {
	var sum = 0
	for _, num := range nums {
		sum += num
	}
	ch <- float64(sum) / float64(len(nums))
}

func main() {
	runtime.GOMAXPROCS(2)

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("number", numbers)

	ch1 := make(chan float64)
	go getAverage(numbers, ch1)

	ch2 := make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f\n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %d\n", max)
		}
	}
}
