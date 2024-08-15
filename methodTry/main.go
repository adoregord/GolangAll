package main

import "fmt"

type student struct {
	Name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("halo", s.Name, "dengan grade", s.grade)
}

func main(){
	student1 := student{
		Name:  "aa",
		grade: 100,
	}

	student1.sayHello()
}
