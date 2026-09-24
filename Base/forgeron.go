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
)

func ouvrirForgeron(joueur *personage.Joueur) {
	for {
		fmt.Println("\n=== LE FORGERON ===")
		fmt.Println("1. Fabriquer une arme")
		fmt.Println("2. Fabriquer une armure")
		fmt.Println("3. Améliorer l'épée")
		fmt.Println("4. Voir les matériaux")
		fmt.Println("5. Partir")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			fabriquerArme(joueur)
		case 2:
			fabriquerArmure(joueur)
		case 3:
			ameliorerEpee(joueur)
		case 4:
			afficherMateriaux(joueur)
		case 5:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func fabriquerArme(joueur *personage.Joueur) {
	if possedeObjetForge(joueur.Inventaire, personage.ArmeEpee) {
		fmt.Println("Tu possèdes déjà une épée.")
		return
	}

	if !consommerMateriaux(joueur, map[string]int{fer: 3, charbon: 1}) {
		return
	}

	joueur.Inventaire = append(joueur.Inventaire, personage.ArmeEpee)
	joueur.NiveauEpee = 1
	fmt.Println("Tu as fabriqué une épée.")
}

func fabriquerArmure(joueur *personage.Joueur) {
	if possedeObjetForge(joueur.Inventaire, personage.Armure) {
		fmt.Println("Tu possèdes déjà une armure.")
		return
	}

	if !consommerMateriaux(joueur, map[string]int{fer: 4, cuir: 3}) {
		return
	}

	joueur.Inventaire = append(joueur.Inventaire, personage.Armure)
	fmt.Println("Tu as fabriqué une armure.")
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
	for materiau, quantite := range cout {
		if joueur.Materiaux[materiau] < quantite {
			fmt.Println("Tu n'as pas assez de", materiau, ".")
			return false
		}
	}

	for materiau, quantite := range cout {
		joueur.Materiaux[materiau] -= quantite
	}
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
