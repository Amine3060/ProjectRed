package base

import (
	"fmt"

	personage "projectred/Personage"
)

func Menu(joueur *personage.Joueur) {
	for {
		fmt.Println("\n=== LES CENDRES DE L'AUBE ===")
		fmt.Println("1. Commencer une partie")
		fmt.Println("2. Voir le personnage")
		fmt.Println("3. Quitter")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			*joueur = creerJoueur()
			StartAdventure(joueur)
		case 2:
			joueur.AfficherInfos()
		case 3:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
