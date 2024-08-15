package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"main.go/model"
)

// function to add karyawan list
func add(list *[]model.Karyawan, nama string) {
	var ID int
	length := len(*list)
	if length == 0 {
		ID = length + 1
	} else {
		ID = (*list)[len(*list)-1].ID + 1
	}

	*list = append(
		*list, model.Karyawan{
			ID:         ID,
			Nama:       nama,
			Kehardiran: false,
		})
}

// function to change status of karyawan to true
func changeStatus(list *[]model.Karyawan, ID int) {
	for index, isi := range *list {
		if isi.ID == ID {
			(*list)[index].Kehardiran = true
			return
		}
	}
	fmt.Println("There's no employee with such ID")
}

// function to chnage the status back to false
func revertStatus(list *[]model.Karyawan, ID int) {
	for index, isi := range *list {
		if isi.ID == ID {
			(*list)[index].Kehardiran = false
			return
		}
	}
	fmt.Println("There's no employee with such ID")
}

// function to delete karyawan from list
func deleteKaryawan(list *[]model.Karyawan, ID int) {
	for index, value := range *list {
		if value.ID == ID {
			*list = append((*list)[:index], (*list)[index+1:]...)
			// for i := index; i < len(*list); i++ {
			// 	(*list)[i].ID--
			// }
			return
		}
	}
	fmt.Println("There's no employee with such ID")

}

// function to show the list of karyawan
func showKaryawan(list *[]model.Karyawan) {
	fmt.Println("--------------DAFTAR KARYAWAN--------------")
	fmt.Printf("ID \t Nama \t\t\t Kehadiran \n")
	if len(*list) == 0 {
		fmt.Printf("~~~~~~No employee recorded~~~~~~\n\n")
		return
	}
	for _, value := range *list {
		fmt.Printf("%d \t %s \t\t\t %t \n", value.ID, value.Nama, value.Kehardiran)
	}
	fmt.Println()
}

// //coba coba
// func findById(list []model.Karyawan, id int, karyawan *model.Karyawan) {
// 	for _, value := range list {
// 		if(value.ID ==  id){
// 			*karyawan = value
// 		}
// 	}
// }

func main() {
	var list []model.Karyawan

	Scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please input number from 1-5 to change employee list")
	fmt.Println("Input [1] to add new employee to the list")
	fmt.Println("Input [2] to change the employee status to true")
	fmt.Println("Input [3] to change the employee status to false")
	fmt.Println("Input [4] to delete the emploee from the list")
	// fmt.Println("Input [5] to find by id")
	fmt.Println("Input [5] to show all employee")
	fmt.Println("Input [6] to close the app")

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
			add(&list, valueStr)
		case 2: // change karyawan status to true
			fmt.Print("Employee ID you want to change the status to true:")
			Scanner.Scan()
			idStr := Scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Printf("ERROR\n\n")
				continue
			}
			changeStatus(&list, id)
		case 3:
			fmt.Print("Employee ID you want to change the status to false:")
			Scanner.Scan()
			idStr := Scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Printf("ERROR\n\n")
				continue
			}
			revertStatus(&list, id)
		case 4: // delete karyawan from the list
			fmt.Print("Employee ID you want to change the delete :")
			Scanner.Scan()
			idStr := Scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Printf("ERROR\n\n")
				continue
			}
			deleteKaryawan(&list, id)
		// case 5: // find by id
		// 	fmt.Print("Employee ID you want to find by id :")
		// 	Scanner.Scan()
		// 	idStr := Scanner.Text()
		// 	var karyawan model.Karyawan
		// 	id, err := strconv.Atoi(idStr)
		// 	if err != nil {
		// 		fmt.Printf("ERROR\n\n")
		// 		continue
		// 	}
		// 	findById(list, id, &karyawan)
		// 	fmt.Println(karyawan)
		case 5: // show all karyawan list
			showKaryawan(&list)
		case 6: // break from the for loop
			return
		default:
			fmt.Printf("Please pick a number from 1 to 6\n\n")
			continue
		}
		fmt.Println()
	}

}
