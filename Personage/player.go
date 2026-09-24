package personage

import "fmt"

type Joueur struct {
	Nom        string
	Classe     string
	PointsVie  int
	VieMax     int
	Mana       int
	ManaMax    int
	Or         int
	Inventaire []string
	Arme       string
	Materiaux  map[string]int
	NiveauEpee int
}

func NouveauJoueur(nom string, classe string) Joueur {
	joueur := Joueur{
		Nom:        nom,
		Classe:     classe,
		Or:         0,
		Inventaire: []string{},
		Arme:       "Aucune",
		Materiaux: map[string]int{
			"Fer":     5,
			"Cuir":    3,
			"Charbon": 2,
		},
		NiveauEpee: 0,
	}

	switch classe {
	case "Elfe":
		joueur.VieMax = 90
		joueur.ManaMax = 80
	case "Gobelin":
		joueur.VieMax = 120
		joueur.ManaMax = 30
	default:
		joueur.Classe = "Humain"
		joueur.VieMax = 100
		joueur.ManaMax = 50
	}

	joueur.PointsVie = joueur.VieMax
	joueur.Mana = joueur.ManaMax
	return joueur
}

func (joueur *Joueur) UtiliserMana(quantite int) bool {
	if joueur.Mana < quantite {
		return false
	}

	joueur.Mana -= quantite
	return true
}

func (joueur Joueur) AfficherInfos() {
	fmt.Println("\n=== PERSONNAGE ===")
	fmt.Println("Nom :", joueur.Nom)
	fmt.Println("Classe :", joueur.Classe)
	fmt.Println("PV :", joueur.PointsVie, "/", joueur.VieMax)
	fmt.Println("Mana :", joueur.Mana, "/", joueur.ManaMax)
	fmt.Println("Or :", joueur.Or)
	fmt.Println("Inventaire :", joueur.Inventaire)
	fmt.Println("Arme équipée :", joueur.Arme)
}
