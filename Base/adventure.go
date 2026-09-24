package base

import (
	"fmt"

	combat "projectred/Combat"
	personage "projectred/Personage"
)

func StartAdventure(joueur *personage.Joueur) {
	introduction(joueur)
	escape(joueur)
	combat.CombatGuerrier(joueur, ouvrirMenuJeu)
	reveil(joueur)
	foret(joueur)
	combat.CommencerCombat(joueur, ouvrirMenuJeu)
	marchand(joueur)
	finAventure(joueur)
}

func introduction(joueur *personage.Joueur) {
	fmt.Println("\n=== INTRODUCTION ===")
	fmt.Println(joueur.Nom, "habite dans un petit village avec sa sœur Nima.")
	fmt.Println("Un jour, une armée attaque le village.")
	attendreEntree(joueur)
}

func escape(joueur *personage.Joueur) {
	fmt.Println("\n=== LA FUITE ===")
	fmt.Println(joueur.Nom, "essaie de fuir avec Nima.")
	fmt.Println("Un guerrier ennemi bloque le chemin.")
	attendreEntree(joueur)
}

func reveil(joueur *personage.Joueur) {
	joueur.PointsVie = joueur.VieMax / 2
	fmt.Println("\n=== LE REVEIL ===")
	fmt.Println(joueur.Nom, "se réveille dans les ruines du village.")
	fmt.Println("Nima a été capturée. Il part à sa recherche.")
	attendreEntree(joueur)
}

func foret(joueur *personage.Joueur) {
	fmt.Println("\n=== LA FORET ===")
	fmt.Println(joueur.Nom, "avance dans la forêt et rencontre un soldat.")
	attendreEntree(joueur)
}

func finAventure(joueur *personage.Joueur) {
	fmt.Println("\nL'aventure continue...")
	fmt.Println("Or gagné :", joueur.Or)
	attendreEntree(joueur)
}
