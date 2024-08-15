package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type calc struct {
	value float32
}

func (c *calc) add(value float32) {
	c.value += value
}

func (c *calc) substract(value float32) {
	c.value -= value
}

func (c *calc) multiply(value float32) {
	c.value *= value
}

func (c *calc) divide(value float32) {
	c.value /= value
}

func (c *calc) power(value int) {
	temp := c.value
	if value == 0 {
		c.value = 1
	} else if value > 0 {
		for i := 1; i < value; i++ {
			c.value *= temp
		}
	} else if value < 0 {
		for i := value; i < 1; i++ {
			c.value /= temp
		}
	}
}

func (c *calc) clear() {
	c.value = 0
}

func main() {

	Scanner := bufio.NewScanner(os.Stdin)

	calc1 := calc{
		value: 0,
	}

	for {
		fmt.Println("Masukkan nilai angka untuk melakukan operasi matematika")
		fmt.Println("Masukkan [1] untuk menjumlah kan nilai")
		fmt.Println("Masukkan [2] untuk mengurangkan kan nilai")
		fmt.Println("Masukkan [3] untuk mengalikan nilai")
		fmt.Println("Masukkan [4] untuk membagikan nilai")
		fmt.Println("Masukkan [5] untuk menguadratkan nilai")
		fmt.Println("Masukkan [6] untuk clear nilai")
		fmt.Println("Masukkan [7] untuk keluar")
		fmt.Println("Nilai saat ini: ", calc1)
		fmt.Print("Masukkan pilihan: ")
		Scanner.Scan()
		optionstr := Scanner.Text()
		option, err := strconv.Atoi(optionstr)
		if err != nil {
			fmt.Printf("Error because: {%s}\n\n", err.Error())
			continue
		}

		if option >= 1 && option <= 5 {
			fmt.Print("Masukkan nilai: ")
			Scanner.Scan()
			valueStr := Scanner.Text()
			value, err := strconv.ParseFloat(valueStr, 32)
			if err != nil {
				fmt.Printf("Error because: {%s}\n\n", err.Error())
				continue
			}

			switch option {
			case 1:
				calc1.add(float32(value))
			case 2:
				calc1.substract(float32(value))
			case 3:
				calc1.multiply(float32(value))
			case 4:
				if value == 0 {
					fmt.Printf("Cannot divide by 0\n\n")
					continue
				}
				calc1.divide(float32(value))
			case 5:
				if value == 1 {
					continue
				}
				calc1.power(int(value))
			}
		} else if option == 6 {
			calc1.clear()
		} else if option == 7 {
			return
		} else {
			fmt.Println("Tolong masukkan pilihan yang benar!!")
		}
		fmt.Println()
	}

}
