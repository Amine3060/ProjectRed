package base

import (
	"fmt"
	"strings"

	personage "projectred/Personage"
)

func creerJoueur() personage.Joueur {
	var nom string

	fmt.Println("\n=== CREATION DU PERSONNAGE ===")
	fmt.Print("Choisis ton nom : ")
	fmt.Scanln(&nom)

	nom = strings.TrimSpace(nom)
	if nom == "" {
		nom = "Krag"
	}

	for {
		fmt.Println("\nChoisis une classe :")
		fmt.Println("1. Elfe    (90 PV, 80 mana)")
		fmt.Println("2. Gobelin (120 PV, 30 mana)")
		fmt.Println("3. Humain  (100 PV, 50 mana)")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			return personage.NouveauJoueur(nom, "Elfe")
		case 2:
			return personage.NouveauJoueur(nom, "Gobelin")
		case 3:
			return personage.NouveauJoueur(nom, "Humain")
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
