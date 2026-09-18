package base

import (
	"fmt"

	personage "projectred/Personage"
)

func Menu(player *personage.Player) {
	for {
		printTitle()
		fmt.Println("1. Nouvelle partie")
		fmt.Println("2. Personnage")
		fmt.Println("3. Charger une partie")
		fmt.Println("4. Quitter")
		fmt.Println("================================")
		fmt.Print("Votre choix : ")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Veuillez entrer un nombre.")
			continue
		}

		switch choice {
		case 1:
			NouvellePartie(player)
		case 2:
			player.DisplayInfo()
		case 3:
			ChargerPartie()
		case 4:
			fmt.Println("Fermeture du jeu...")
			return
		default:
			fmt.Println("Choix incorrect. Veuillez choisir 1, 2, 3 ou 4.")
		}
	}
}

func printTitle() {
	fmt.Println()
	fmt.Println("================================")
	fmt.Println("       LES CENDRES DE L'AUBE")
	fmt.Println("================================")
}

func NouvellePartie(player *personage.Player) {
	fmt.Println()
	fmt.Println("================================")
	fmt.Println("         NOUVELLE PARTIE")
	fmt.Println("================================")
	fmt.Println("Début de votre aventure...")
	fmt.Println("Bienvenue,", player.Name+".")
	player.DisplayInfo()
}

func ChargerPartie() {
	fmt.Println()
	fmt.Println("================================")
	fmt.Println("        CHARGER UNE PARTIE")
	fmt.Println("================================")
	fmt.Println("Aucune sauvegarde disponible pour le moment.")
}
