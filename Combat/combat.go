package combat

import (
	"fmt"

	personage "projectred/Personage"
)

type Ennemi struct {
	Nom     string
	Vie     int
	Attaque int
}

func CombatGuerrier(joueur *personage.Joueur, ouvrirMenu func(*personage.Joueur)) {
	fmt.Println("\n=== PREMIER COMBAT ===")
	fmt.Println("Le guerrier est trop fort. Krag est vaincu.")
	joueur.PointsVie = 0
	fmt.Println("PV de", joueur.Nom, ":", joueur.PointsVie)
	attendreEntree(joueur, ouvrirMenu)
}

func CommencerCombat(joueur *personage.Joueur, ouvrirMenu func(*personage.Joueur)) {
	ennemi := Ennemi{
		Nom:     "Soldat ennemi",
		Vie:     40,
		Attaque: 8,
	}
	toursPoison := 0

	fmt.Println("\n=== SECOND COMBAT ===")
	fmt.Println("Un", ennemi.Nom, "attaque !")

	for joueur.PointsVie > 0 && ennemi.Vie > 0 {
		fmt.Println("\nTes PV :", joueur.PointsVie, "| PV ennemi :", ennemi.Vie)
		fmt.Println("1. Attaquer")
		fmt.Println("2. Lancer Fireball (10 mana)")
		fmt.Println("3. Utiliser une potion")
		fmt.Println("4. Fuir")
		fmt.Println("0. Ouvrir le menu")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 0:
			ouvrirMenu(joueur)
			continue
		case 1:
			degats := 10
			if joueur.Arme == personage.ArmeEpee {
				degats = 20 + (joueur.NiveauEpee-1)*5
			}
			ennemi.Vie -= degats
			fmt.Println("Tu infliges", degats, "dégâts.")
		case 2:
			if joueur.UtiliserMana(10) {
				ennemi.Vie -= 25
				fmt.Println("Fireball inflige 25 dégâts.")
			} else {
				fmt.Println("Tu n'as pas assez de mana.")
				continue
			}
		case 3:
			if !utiliserPotion(joueur, &ennemi, &toursPoison) {
				continue
			}
		case 4:
			fmt.Println("Tu prends la fuite.")
			return
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		if ennemi.Vie <= 0 {
			break
		}

		if toursPoison > 0 {
			ennemi.Vie -= 15
			toursPoison--
			fmt.Println("Le poison inflige 15 dégâts.")
		}

		if ennemi.Vie <= 0 {
			break
		}

		degatsRecus := joueur.DegatsRecus(ennemi.Attaque)
		joueur.PointsVie -= degatsRecus
		fmt.Println("L'ennemi inflige", degatsRecus, "dégâts.")
	}

	if joueur.PointsVie <= 0 {
		fmt.Println("Tu as perdu le combat.")
	} else {
		joueur.Or += 100
		joueur.Inventaire = append(joueur.Inventaire, personage.ArmeEpee)
		fmt.Println("Tu as gagné ! Tu trouves 100 pièces d'or.")
		fmt.Println("Le soldat laisse tomber une épée.")
	}

	attendreEntree(joueur, ouvrirMenu)
}

func attendreEntree(joueur *personage.Joueur, ouvrirMenu func(*personage.Joueur)) {
	fmt.Println("Appuie sur Entrée pour continuer ou tape 0 pour ouvrir le menu.")
	var choix string
	fmt.Scanln(&choix)
	if choix == "0" {
		ouvrirMenu(joueur)
	}
}
