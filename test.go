package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	var random_number int = rand.Intn(100)

	fmt.Println("eh viens on joue, devine un nombre en 0 et 100")
	var user_choice int

	for user_choice != random_number {

		fmt.Scanf("%d", &user_choice)

		if user_choice < random_number {
			fmt.Println("c'est plus ma gueule!")

		} else if user_choice > random_number {
			fmt.Println("c'est moins rhey")
		}
	}
	fmt.Println("tu as trouvé bravo")
}
