package main

import "fmt"

func main() {
	var makanan map[string]int
	makanan = map[string]int{}

	makanan["Ayam Goreng"] = 10
	makanan["Bebek Goreng"] = 1
	fmt.Println("Ayam Goreng", makanan["Ayam Goreng"])
	fmt.Println("Bebek Goreng", makanan["Bebek Goreng"])
	fmt.Println()


	var minuman = map[string]int{
		"minuman 1": 10,
		"minuman 2": 2,
		"minuman 3": 100,
		"minuman 4": 1,
	}

	for key, val := range minuman{
		fmt.Println(key, "\t:",val)
	}

	var tempat = map[int]string{
		1: "Jakarta",
		2: "Bandung",
		3: "Semarang",
		4: "Surabaya",
		5: "Asem",
	}
	for key, val := range tempat{
		fmt.Println(key, "\t:",val)
	}
	fmt.Println()

	//delete item in map
	delete(tempat, 5)
	for key, val := range tempat{
		fmt.Println(key, "\t:",val)
	}

	//make map
	var permainan = make(map[string]int)
	var permainan2 = *new(map[string]int)
	var permainan3 = map[string]int{"A":1}
	fmt.Println(permainan, permainan2, permainan3)

}
