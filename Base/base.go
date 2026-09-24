package base

import (
	"fmt"

	personage "projectred/Personage"
)

func Menu(player *personage.Player) {
	for {
		fmt.Println("\n=== LES CENDRES DE L'AUBE ===")
		fmt.Println("1. Commencer une partie")
		fmt.Println("2. Voir le personnage")
		fmt.Println("3. Quitter")
		fmt.Print("Choix : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			*player = personage.NewPlayer()
			StartAdventure(player)
		case 2:
			player.DisplayInfo()
		case 3:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
