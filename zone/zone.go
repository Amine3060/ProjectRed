package zone

import (
	"fmt"
	"math/rand"

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
	joueur.Materiaux[fer] += 3
	joueur.Materiaux[charbon] += 10
	fmt.Println("Tu as extrait 3 Fer et 10 Charbon.")

	if rand.Intn(2) == 0 {
		joueur.Or++
		fmt.Println("Tu as trouvé 1 pièce d'or !")
	}
}

func afficherMateriaux(joueur *personage.Joueur) {
	fmt.Println("Fer :", joueur.Materiaux[fer])
	fmt.Println("Charbon :", joueur.Materiaux[charbon])
	fmt.Println("Or :", joueur.Or)
}
