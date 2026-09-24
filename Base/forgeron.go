package base

import (
	"fmt"

	personage "projectred/Personage"
)

const (
	fer           = "Fer"
	cuir          = "Cuir"
	charbon       = "Charbon"
	niveauEpeeMax = 5
	rareOr        = 20
	rareFer       = 30
	rareCharbon   = 70
)

func ouvrirForgeron(joueur *personage.Joueur) {
	for {
		fmt.Println("\n=== LE FORGERON ===")
		fmt.Println("1. Fabriquer une arme")
		fmt.Println("2. Fabriquer une armure")
		fmt.Println("3. Améliorer l'épée")
		fmt.Println("4. Fabriquer une arme rare (20 or, 30 fer, 70 charbon)")
		fmt.Println("5. Fabriquer une armure rare (20 or, 30 fer, 70 charbon)")
		fmt.Println("6. Voir les matériaux")
		fmt.Println("7. Partir")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			fabriquer(joueur, personage.ArmeEpee, map[string]int{fer: 3, charbon: 1}, 0)
		case 2:
			fabriquer(joueur, personage.Armure, map[string]int{fer: 4, cuir: 3}, 0)
		case 3:
			ameliorerEpee(joueur)
		case 4:
			fabriquer(joueur, personage.ArmeEpeeRenforcee, map[string]int{fer: rareFer, charbon: rareCharbon}, rareOr)
		case 5:
			fabriquer(joueur, personage.ArmureRenforcee, map[string]int{fer: rareFer, charbon: rareCharbon}, rareOr)
		case 6:
			afficherMateriaux(joueur)
		case 7:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func fabriquer(joueur *personage.Joueur, objet string, cout map[string]int, prixOr int) {
	if possedeObjetForge(joueur.Inventaire, objet) {
		fmt.Println("Tu possèdes déjà", objet+".")
		return
	}
	if !consommerRessources(joueur, cout, prixOr) {
		return
	}

	joueur.Inventaire = append(joueur.Inventaire, objet)
	if objet == personage.ArmeEpee || objet == personage.ArmeEpeeRenforcee {
		joueur.NiveauEpee = 1
	}
	fmt.Println("Tu as fabriqué", objet+".")
}

func ameliorerEpee(joueur *personage.Joueur) {
	if !possedeObjetForge(joueur.Inventaire, personage.ArmeEpee) {
		fmt.Println("Tu dois posséder une épée pour l'améliorer.")
		return
	}
	if joueur.NiveauEpee >= niveauEpeeMax {
		fmt.Println("Ton épée est déjà au niveau maximum.")
		return
	}

	if !consommerMateriaux(joueur, map[string]int{fer: 2, charbon: 1}) {
		return
	}

	joueur.NiveauEpee++
	fmt.Println("Ton épée est maintenant niveau", joueur.NiveauEpee, ".")
}

func afficherMateriaux(joueur *personage.Joueur) {
	fmt.Println("\n=== MATÉRIAUX ===")
	fmt.Println("Fer     :", joueur.Materiaux[fer])
	fmt.Println("Cuir    :", joueur.Materiaux[cuir])
	fmt.Println("Charbon :", joueur.Materiaux[charbon])
}

func consommerMateriaux(joueur *personage.Joueur, cout map[string]int) bool {
	return consommerRessources(joueur, cout, 0)
}

func consommerRessources(joueur *personage.Joueur, cout map[string]int, prixOr int) bool {
	if joueur.Or < prixOr {
		fmt.Println("Tu n'as pas assez d'or.")
		return false
	}
	for materiau, quantite := range cout {
		if joueur.Materiaux[materiau] < quantite {
			fmt.Println("Tu n'as pas assez de", materiau, ".")
			return false
		}
	}

	for materiau, quantite := range cout {
		joueur.Materiaux[materiau] -= quantite
	}
	joueur.Or -= prixOr
	return true
}

func possedeObjetForge(inventaire []string, objetRecherche string) bool {
	for _, objet := range inventaire {
		if objet == objetRecherche {
			return true
		}
	}
	return false
}
