package base

import (
	"fmt"

	personage "projectred/Personage"
)

func marchand(joueur *personage.Joueur) {
	for {
		fmt.Println("\n=== LE MARCHAND ===")
		fmt.Println("Or disponible :", joueur.Or)
		fmt.Println("1. Vendre l'épée : 50 pièces")
		fmt.Println("2. Acheter une potion de soins : 50 pièces")
		fmt.Println("3. Acheter une potion de poison : 60 pièces")
		fmt.Println("4. Partir")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			vendreEpee(joueur)
		case 2:
			acheterObjet(joueur, personage.PotionSoins, 50)
		case 3:
			acheterObjet(joueur, personage.PotionPoison, 60)
		case 4:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func vendreEpee(joueur *personage.Joueur) {
	if !possedeEpee(joueur.Inventaire) {
		fmt.Println("Tu n'as pas d'épée à vendre.")
		return
	}

	joueur.Or += 50
	joueur.Inventaire = retirerEpee(joueur.Inventaire)
	if joueur.Arme == personage.ArmeEpee {
		joueur.Arme = "Aucune"
	}
	fmt.Println("Tu as vendu l'épée pour 50 pièces d'or.")
}

func acheterObjet(joueur *personage.Joueur, objet string, prix int) {
	if joueur.Or < prix {
		fmt.Println("Tu n'as pas assez d'or.")
		return
	}

	joueur.Or -= prix
	joueur.Inventaire = append(joueur.Inventaire, objet)
	fmt.Println("Tu as acheté :", objet)
}

func retirerEpee(inventaire []string) []string {
	for i, objet := range inventaire {
		if objet == personage.ArmeEpee {
			return append(inventaire[:i], inventaire[i+1:]...)
		}
	}

	return inventaire
}

func possedeEpee(inventaire []string) bool {
	for _, objet := range inventaire {
		if objet == personage.ArmeEpee {
			return true
		}
	}

	return false
}
