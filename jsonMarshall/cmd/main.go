package main

import (
	"encoding/json"
	"fmt"

	"main.go/internal/domain"
)

var user domain.User

// change struct to json
func marshallJSON(user domain.User) string{
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return (string(jsonData))
}

// change json to struct
func unmarshallJSON(jsonData string) {
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Name: %s\nEmail: %s\nAge: %d\n", user.Name, user.Email, user.Age)
}


func main() {

	user1 := domain.User{
		Name:  "Didi",
		Email: "Didiganteng@gmail.com",
		Age:   22,
	}

	hasilMarshall := marshallJSON(user1)
	fmt.Println(hasilMarshall)

	unmarshallJSON(hasilMarshall)
}
