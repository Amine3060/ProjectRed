package combat

import (
	"fmt"

	personage "projectred/Personage"
)

func utiliserPotion(joueur *personage.Joueur, ennemi *Ennemi, toursPoison *int) bool {
	fmt.Println("1. Potion de soins (+50 PV)")
	fmt.Println("2. Potion de poison (15 dégâts pendant 3 tours)")
	fmt.Println("3. Retour")
	fmt.Print("Choix : ")

	var choix int
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		if !retirerObjet(joueur, personage.PotionSoins) {
			fmt.Println("Tu n'as pas cette potion.")
			return false
		}
		joueur.PointsVie += 50
		if joueur.PointsVie > joueur.VieMax {
			joueur.PointsVie = joueur.VieMax
		}
		fmt.Println("Tu récupères 50 PV.")
		return true
	case 2:
		if !retirerObjet(joueur, personage.PotionPoison) {
			fmt.Println("Tu n'as pas cette potion.")
			return false
		}
		*toursPoison = 3
		fmt.Println("L'ennemi est empoisonné.")
		return true
	case 3:
		return false
	default:
		fmt.Println("Choix invalide.")
		return false
	}
}

func retirerObjet(joueur *personage.Joueur, objetRecherche string) bool {
	for i, objet := range joueur.Inventaire {
		if objet == objetRecherche {
			joueur.Inventaire = append(joueur.Inventaire[:i], joueur.Inventaire[i+1:]...)
			return true
		}
	}

	return false
}
