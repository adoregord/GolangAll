package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(score float64) string {
	var result string
	if score >= 90 && score <= 100 {
		result = "Selamat! Anda mendapatkan nilai A."
	} else if score >= 80 && score <= 89 {
		result = "Anda mendapatkan nilai B."
	} else if score >= 70 && score <= 79 {
		result = "Anda mendapatkan nilai C."
	} else if score >= 60 && score <= 69 {
		result = "Anda mendapatkan nilai D."
	} else if score >= 0 && score <= 59 {
		result = "Anda mendapatkan nilai E"
	} else {
		result = "Mohon memasukkan nilai dari 0 sampai 100 saja"
	}
	return result
}

func main() {

	Scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan nilai skor: ")
	Scanner.Scan()
	nilaistr := Scanner.Text()
	nilai, err := strconv.ParseFloat(nilaistr, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(check(nilai))

}
