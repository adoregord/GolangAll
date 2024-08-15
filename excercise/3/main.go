package main

import "fmt"

func main() {
	lorem := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

	var count int
	for i := 0; i < len(lorem)-1; i++ {
		if lorem[i] == 'a' || lorem[i] == 'i' || lorem[i] == 'u' || lorem[i] == 'e' || lorem[i] == 'o' || lorem[i] == 'A' || lorem[i] == 'I' || lorem[i] == 'U' || lorem[i] == 'E' || lorem[i] == 'O' {
			count++
		}
	}
	fmt.Printf("Jumlah huruf fokal yang ada pada kalimat tersebut adalah: %d", count)
}
