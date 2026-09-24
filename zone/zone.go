package zone

import (
	"fmt"

	personage "projectred/Personage"
)

const (
	fer     = "Fer"
	charbon = "Charbon"
)

func OuvrirZoneMiniere(joueur *personage.Joueur) {
	for {
		fmt.Println("\n=== ZONE MINIÈRE ===")
		fmt.Println("1. Miner des minéraux")
		fmt.Println("2. Voir les matériaux")
		fmt.Println("3. Quitter la mine")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			miner(joueur)
		case 2:
			afficherMateriaux(joueur)
		case 3:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func miner(joueur *personage.Joueur) {
	joueur.Materiaux[fer] += 2
	joueur.Materiaux[charbon]++
	fmt.Println("Tu as extrait 2 Fer et 1 Charbon.")
}

func afficherMateriaux(joueur *personage.Joueur) {
	fmt.Println("Fer :", joueur.Materiaux[fer])
	fmt.Println("Charbon :", joueur.Materiaux[charbon])
}
