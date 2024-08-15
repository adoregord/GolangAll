package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Barang struct {
	Name  string
	Stock int
}

func add(list *[]Barang, a ...Barang) {
	for _, val := range a {
		*list = append(*list, val)
	}

}
func add2(list *[]Barang, nama string, value int) {
	*list = append(*list, Barang{Name: nama, Stock: value})
}

func reduce(list *[]Barang, index int) {
	if index < 0 || index >= len(*list) {
		return
	}
	*list = append((*list)[:index], (*list)[index+1:]...)
}

func edit(list *[]Barang, index int, nama string, value int) {
	(*list)[index].Name = nama
	(*list)[index].Stock = value
}

func main() {
	list := []Barang{}
	// var name string
	// var value int

	barang1 := Barang{
		Name:  "Hape",
		Stock: 1,
	}
	barang2 := Barang{
		Name:  "Laptop",
		Stock: 10,
	}
	barang3 := Barang{
		Name:  "Keyboard",
		Stock: 30,
	}
	barang4 := Barang{
		Name:  "Mouse",
		Stock: 100,
	}

	add(&list, barang1, barang2, barang3, barang4)
	fmt.Println(list)

	// fmt.Print("Masukkan nama barang yang ingin ditambahkan: ")
	// fmt.Scanf("%[^\n]s", &name)
	// fmt.Print("Masukkan stock barang yang ingin ditambahkan: ")
	// fmt.Scanln(&value)

	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan nama barang yang ingin ditambahkan: ")
	Scanner.Scan()
	name := Scanner.Text()

	fmt.Print("Masukkan stock barang yang ingin ditambahkan: ")
	Scanner.Scan()
	valueStr := Scanner.Text()
	value, _ := strconv.Atoi(valueStr)
	

	add2(&list, name, value)
	fmt.Println(list)

	edit(&list, 0, "Hape", 0)
	fmt.Println(list)

	reduce(&list, 2)
	fmt.Println(list)
}
