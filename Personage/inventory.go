package personage

import "fmt"

func (joueur *Joueur) OuvrirInventaire() {
	for {
		fmt.Println("\n=== INVENTAIRE ===")
		if len(joueur.Inventaire) == 0 {
			fmt.Println("L'inventaire est vide.")
		} else {
			for numero, objet := range joueur.Inventaire {
				fmt.Println(numero+1, ".", objet)
			}
		}

		fmt.Println("Arme équipée :", joueur.Arme)
		fmt.Println("Armure équipée :", joueur.ArmureEquipee)
		fmt.Println("1. Équiper l'épée")
		fmt.Println("2. Équiper l'épée renforcée")
		fmt.Println("3. Équiper l'armure")
		fmt.Println("4. Équiper l'armure renforcée")
		fmt.Println("5. Utiliser une potion de soins")
		fmt.Println("6. Voir l'utilisation de la potion de poison")
		fmt.Println("7. Retour")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			if !possedeObjet(joueur.Inventaire, ArmeEpee) {
				fmt.Println("Tu ne possèdes pas d'épée.")
				continue
			}
			joueur.Arme = ArmeEpee
			fmt.Println("Épée équipée.")
		case 2:
			if !possedeObjet(joueur.Inventaire, ArmeEpeeRenforcee) {
				fmt.Println("Tu ne possèdes pas d'épée renforcée.")
				continue
			}
			joueur.Arme = ArmeEpeeRenforcee
			fmt.Println("Épée renforcée équipée.")
		case 3:
			if !possedeObjet(joueur.Inventaire, Armure) {
				fmt.Println("Tu ne possèdes pas d'armure.")
				continue
			}
			joueur.ArmureEquipee = true
			joueur.ArmureRenforceeEquipee = false
			fmt.Println("Armure équipée.")
		case 4:
			if !possedeObjet(joueur.Inventaire, ArmureRenforcee) {
				fmt.Println("Tu ne possèdes pas d'armure renforcée.")
				continue
			}
			joueur.ArmureEquipee = false
			joueur.ArmureRenforceeEquipee = true
			fmt.Println("Armure renforcée équipée.")
		case 5:
			utiliserPotionSoins(joueur)
		case 6:
			if possedeObjet(joueur.Inventaire, PotionPoison) {
				fmt.Println("La potion de poison s'utilise pendant un combat.")
			} else {
				fmt.Println("Tu ne possèdes pas de potion de poison.")
			}
		case 7:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func utiliserPotionSoins(joueur *Joueur) {
	if !possedeObjet(joueur.Inventaire, PotionSoins) {
		fmt.Println("Tu ne possèdes pas de potion de soins.")
		return
	}

	if joueur.PointsVie == joueur.VieMax {
		fmt.Println("Tes PV sont déjà au maximum.")
		return
	}

	joueur.PointsVie += 50
	if joueur.PointsVie > joueur.VieMax {
		joueur.PointsVie = joueur.VieMax
	}
	joueur.Inventaire = retirerObjet(joueur.Inventaire, PotionSoins)
	fmt.Println("Tu récupères 50 PV.")
}

func possedeObjet(inventaire []string, objetRecherche string) bool {
	for _, objet := range inventaire {
		if objet == objetRecherche {
			return true
		}
	}
	return false
}

func retirerObjet(inventaire []string, objetRecherche string) []string {
	for i, objet := range inventaire {
		if objet == objetRecherche {
			return append(inventaire[:i], inventaire[i+1:]...)
		}
	}

	return inventaire
}
