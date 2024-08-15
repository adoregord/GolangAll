package main

import (
	"fmt"
	"math/rand"
	"time"

	"pokemon/model" // Import the model package containing the Pokemon struct
)

// catchProbability checks if catching is successful given a percentage rate
func catchProbability(rate float32) string {
	if rate < 0 || rate > 100 {
		return "Invalid rate. Please provide a rate between 0 and 100."
	}

	// rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	// chance := rand.Intn(100) + 1     // Generate a random number between 1 and 100
	chance := rand.Float32() * (100 - 0) // Generate a random number between 1 and 100

	if chance <= rate {
		return "SUCCESS, you got it"
	}
	return "FAIL, it got away"
}

func main() {
	now := time.Now()

	// Create sample Pokemon
	// pokemons := []model.Pokemon{
	// 	{ID: 1, Name: "Pikachu", Type: "Electric", CatchRate: 70, IsRare: false, RegisteredDate: now},
	// 	{ID: 2, Name: "Charmander", Type: "Fire", CatchRate: 40, IsRare: true, RegisteredDate: now.AddDate(0, -1, 0)},
	// 	{ID: 3, Name: "Bulbasaur", Type: "Grass/Poison", CatchRate: 10, IsRare: true, RegisteredDate: now.AddDate(0, -6, 0)},
	// 	{ID: 4, Name: "Dragonite", Type: "Dragon/Flying", CatchRate: 30, IsRare: true, RegisteredDate: now.AddDate(0, -8, 0)},
	// 	{ID: 5, Name: "Mew", Type: "Psychic", CatchRate: 1, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0)},
	// 	{ID: 6, Name: "Mew Two", Type: "Psychic", CatchRate: 1, IsRare: true, RegisteredDate: now.AddDate(0, -11, 0)},
	// 	{ID: 7, Name: "Lugia", Type: "Psychic", CatchRate: 0.1, IsRare: true, RegisteredDate: now.AddDate(0, -11, 0)},
	// }

	characters := []model.StarRail{
		{ID: 1, Name: "March 7th", Type: "Ice", CatchRate: 70, IsRare: false, RegisteredDate: now},
		{ID: 2, Name: "Trailblazer", Type: "Physical", CatchRate: 99.99, IsRare: false, RegisteredDate: now.AddDate(0, -1, 0)},
		{ID: 3, Name: "Himeko", Type: "Fire", CatchRate: 10, IsRare: true, RegisteredDate: now.AddDate(0, -6, 0)},
		{ID: 4, Name: "Welt", Type: "Imaginary", CatchRate: 30, IsRare: true, RegisteredDate: now.AddDate(0, -8, 0)},
		{ID: 5, Name: "Danheng", Type: "Wind", CatchRate: 10, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0)},
		{ID: 6, Name: "Topaz", Type: "Fire", CatchRate: 0.01, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0)},
		{ID: 7, Name: "Firefly", Type: "Fire", CatchRate: 0.01, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0)},
		{ID: 8, Name: "Acheron", Type: "Lightning", CatchRate: 0.001, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0)},
		{ID: 9, Name: "Argenti", Type: "Physical", CatchRate: 0.001, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0)},
	}

	// fmt.Println("Available Pokémon:")
	// for _, pokemon := range pokemons {
	// 	fmt.Printf("ID: %d, Name: %s, Type: %s, Catch Rate: %f%%, Is Rare: %t, Registered Date: %s\n",
	// 		pokemon.ID, pokemon.Name, pokemon.Type, pokemon.CatchRate, pokemon.IsRare, pokemon.RegisteredDate.Format(time.RFC1123))
	// }

	fmt.Println("Available Star-Rail Characters:")
	for _, character := range characters {
		fmt.Printf("ID: %d, Name: %s, Type: %s, Catch Rate: %f%%, Is Rare: %t, Registered Date: %s\n",
			character.ID, character.Name, character.Type, character.CatchRate, character.IsRare, character.RegisteredDate.Format(time.RFC1123))
	}

	var id [10]int
	// var selectedPokemon [10]*model.Pokemon

	// for i := 0; i < 10; i++ {
	// 	flag := false
	// 	fmt.Print("Enter Pokémon ID to attempt to catch: ")
	// 	for {
	// 		var _, err = fmt.Scanf("%d\n", &id[i])
	// 		if err != nil {
	// 			fmt.Println("Err " + err.Error())
	// 			fmt.Print("Enter Pokémon ID to attempt to catch: ")
	// 			var s string
	// 			fmt.Scanln(&s)
	// 			fmt.Print(s)
	// 		} else {

	// 			for _, pokemon := range pokemons {
	// 				if pokemon.ID == id[i] {
	// 					selectedPokemon[i] = &pokemon
	// 					flag = true

	// 					break
	// 				}
	// 			}
	// 			if flag {
	// 				break
	// 			} else {
	// 				fmt.Printf("ID pokemon is not available for id %d please enter pokemon ID again: ", id[i])
	// 			}
	// 		}
	// 	}

	// }
	var selectedCharacter [10]*model.StarRail

	for i := 0; i < 10; i++ {
		flag := false
		fmt.Print("Enter character ID to attempt to pull: ")
		for {
			_, err := fmt.Scanf("%d\n", &id[i])
			if err != nil {
				fmt.Println("Error because: " + err.Error())
				fmt.Print("Enter character ID to attempt to pull: ")
				var s string
				fmt.Scanln(&s)
				// fmt.Print(s)
			} else {

				for _, character := range characters {
					if character.ID == id[i] {
						selectedCharacter[i] = &character
						flag = true

						break
					}
				}
				if flag {
					break
				} else {
					fmt.Printf("StarRail character ID is not available for id %d please enter character ID again: ", id[i])
				}
			}
		}

	}

	// for i := 0; i < 10; i++ {
	// 	// if selectedPokemon[i] != nil {
	// 	result := catchProbability(selectedPokemon[i].CatchRate)
	// 	fmt.Printf("You attempted to catch %s (%s type) with a catch rate of %f%%: %s\n",
	// 		selectedPokemon[i].Name, selectedPokemon[i].Type, selectedPokemon[i].CatchRate, result)
	// 	// } //else {
	// 	// 	fmt.Println("Invalid Pokémon ID. Please try again.")
	// 	// }
	// }
	for i := 0; i < 10; i++ {
		// if selectedPokemon[i] != nil {
		result := catchProbability(selectedCharacter[i].CatchRate)
		fmt.Printf("You attempted to pull %s (%s type) with rate of %f%%: %s\n",
			selectedCharacter[i].Name, selectedCharacter[i].Type, selectedCharacter[i].CatchRate, result)
		// } //else {
		// 	fmt.Println("Invalid Pokémon ID. Please try again.")
		// }
	}
}

// for i := 0; i < len(id); i++ {
// 	for _, pokemon := range pokemons {
// 		if pokemon.ID == id[i] {
// 			selectedPokemon[i] = &pokemon
// 			break
// 		}
// 	}
// }

// for {
// 	if id[i] >= len(pokemons)+1 {
// 		fmt.Print("ID pokemon is not available please enter pokemon ID again: ")
// 		fmt.Scan(&id[i])
// 	} else {
// 		break
// 	}
// }
