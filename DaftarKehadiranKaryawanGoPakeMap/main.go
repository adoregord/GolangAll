package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"main.go/model"
)

// function to add karyawan list
func add(maps *map[int]model.Karyawan, nama string) {
	var ID int
	length := len(*maps)

	if length == 0 {
		ID = 1
	} else {
		// Ambil ID terbesar dan tambahkan 1
		for key := range *maps {
			if key >= ID {
				ID = key + 1
			}
		}
	}

	(*maps)[ID] = model.Karyawan{
		ID:         ID,
		Nama:       nama,
		Kehardiran: false,
	}

}

// function to change status of karyawan to true
func changeStatus(maps *map[int]model.Karyawan, ID int) {
	for index, isi := range *maps {
		if isi.ID == ID {
			isi.Kehardiran = true
			(*maps)[index] = isi
			return
		}
	}
	fmt.Println("There's no employee with such ID")
}

// function to chnage the status back to false
func revertStatus(maps *map[int]model.Karyawan, ID int) {
	for index, isi := range *maps {
		if isi.ID == ID {
			isi.Kehardiran = false
			(*maps)[index] = isi
			return
		}
	}
	fmt.Println("There's no employee with such ID")
}

// function to delete karyawan from list
func deleteKaryawan(maps *map[int]model.Karyawan, ID int) {
	// for index, value := range *maps {
	// 	if value.ID == ID {
	// 		*maps = append((*maps)[:index], (*maps)[index+1:]...)
	// 		for i := index; i < len(*maps); i++ {
	// 			(*maps)[i].ID--
	// 		}
	// 		return
	// 	}
	// }
	// fmt.Println("There's no employee with such ID")
	delete(*maps, ID)
	// for index, value := range *maps{
		
	// 
}

// function to show the list of karyawan
func showKaryawan(map1 *map[int]model.Karyawan) {
	fmt.Println("--------------DAFTAR KARYAWAN--------------")
	fmt.Printf("ID \t Nama \t\t\t Kehadiran \n")
	if len(*map1) == 0 {
		fmt.Printf("~~~~~~No employee recorded~~~~~~\n\n")
		return
	}
	for _, value := range *map1 {
		fmt.Printf("%d \t %s \t\t\t %t \n", value.ID, value.Nama, value.Kehardiran)
	}
	fmt.Println()
}

func main() {
	// var list []model.Karyawan
	// map1 := make(map[int]model.Karyawan)
	map1 := map[int]model.Karyawan{}

	Scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please input number from 1-5 to change employee list")
	fmt.Println("Input [1] to add new employee to the list")
	fmt.Println("Input [2] to change the employee status to true")
	fmt.Println("Input [3] to change the employee status to false")
	fmt.Println("Input [4] to delete the emploee from the list")
	fmt.Println("Input [5] to find by id")
	fmt.Println("Input [6] to show all employee")
	fmt.Println("Input [7] to close the app")

	for {
		// input for selecting the action
		fmt.Print("Your input action: ")
		Scanner.Scan()
		optionStr := Scanner.Text()
		option, err := strconv.Atoi(optionStr)
		if err != nil {
			fmt.Printf("ERROR\n\n")
			continue
		}

		// the option will be go through switch case
		switch option {
		case 1: // add karyawan to the list
			fmt.Print("Employee you want to add :")
			Scanner.Scan()
			valueStr := Scanner.Text()
			add(&map1, valueStr)
		
		case 2: // change karyawan status to true
			fmt.Print("Employee ID you want to change the status to true:")
			Scanner.Scan()
			idStr := Scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Printf("ERROR COK\n\n")
				continue
			}
			changeStatus(&map1, id)
		case 3:
			fmt.Print("Employee ID you want to change the status to false:")
			Scanner.Scan()
			idStr := Scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Printf("ERROR COK\n\n")
				continue
			}
			revertStatus(&map1, id)
		case 4: // delete karyawan from the list
			fmt.Print("Employee ID you want to change the delete :")
			Scanner.Scan()
			idStr := Scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Printf("ERROR COK\n\n")
				continue
			}
			deleteKaryawan(&map1, id)
		// case 5: // find by id
		// 	fmt.Print("Employee ID you want to find by id :")
		// 	Scanner.Scan()
		// 	idStr := Scanner.Text()
		// 	var karyawan model.Karyawan
		// 	id, err := strconv.Atoi(idStr)
		// 	if err != nil {
		// 		fmt.Printf("ERROR COK\n\n")
		// 		continue
		// }
		// 	findById(list, id, &karyawan)
		// 	fmt.Println(karyawan)
		case 6: // show all karyawan list
			showKaryawan(&map1)
		}
		// case 7: // break from the for loop
		// 	return
		// default:
		// 	fmt.Printf("GABISA GITU COK\n\n")
		// 	continue
		// }
		// fmt.Println()
		
	}

}
