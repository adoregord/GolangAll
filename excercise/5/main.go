package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			if i == 2 {
				fmt.Print("dua")
			} else if i == 4 {
				fmt.Print("empat")
			} else if i == 6 {
				fmt.Print("enam")
			} else if i == 8 {
				fmt.Print("delapan")
			} else if i == 10 {
				fmt.Print("sepuluh")
			}
			fmt.Printf(" Genap\n")
		} else if i%2 == 1 {
			if i == 1 {
				fmt.Print("satu")
			} else if i == 3 {
				fmt.Print("dua")
			} else if i == 5 {
				fmt.Print("lima")
			} else if i == 7 {
				fmt.Print("tujuh")
			} else if i == 9 {
				fmt.Print("sembilan")
			}
			fmt.Printf(" Ganjil\n")
		}
	}
}
