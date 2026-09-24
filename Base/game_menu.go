package base

import (
	"bufio"
	"fmt"
	"os"

	personage "projectred/Personage"
)

func attendreEntree(joueur *personage.Joueur) {
	for {
		fmt.Println("\nAppuyez sur Entrée pour continuer ou tapez 0 pour ouvrir le menu.")
		entree, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if entree != "0\n" && entree != "0\r\n" {
			return
		}
		ouvrirMenuJeu(joueur)
	}
}

func ouvrirMenuJeu(joueur *personage.Joueur) {
	for {
		fmt.Println("\n=== MENU DU JEU ===")
		fmt.Println("1. Voir le personnage")
		fmt.Println("2. Voir l'inventaire")
		fmt.Println("3. Aller chez le forgeron")
		fmt.Println("4. Reprendre la partie")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			joueur.AfficherInfos()
		case 2:
			joueur.OuvrirInventaire()
		case 3:
			ouvrirForgeron(joueur)
		case 4:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
