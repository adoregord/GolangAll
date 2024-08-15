package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func total(price float64, quantity int, discount int) float64 {
	var result float64
	result = price * (float64(quantity))
	result -= (result * float64(discount) / 100)
	return result
}

func main() {
	// price
	// quantity
	// totalprice dan ada diskon
	Scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan Price dari Barang yang tertera: ")
	Scanner.Scan()
	priceStr := Scanner.Text()
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Masukkan jumlah barang yang ingin dibeli: ")
	Scanner.Scan()
	quantityStr := Scanner.Text()
	quantity, err := strconv.Atoi(quantityStr)

	fmt.Println("Masukkan nilai diskon pada barang tersbut: ")
	Scanner.Scan()
	discountStr := Scanner.Text()
	discount, err := strconv.Atoi(discountStr)

	fmt.Printf("Total nilai yang perlu dibayar: %.02f", total(price, quantity, discount))

}
