package base

import (
	"testing"

	personage "projectred/Personage"
)

func TestFabriquerArmeRareConsommeSesRessources(t *testing.T) {
	joueur := personage.NouveauJoueur("Krag", "Humain")
	joueur.Or = rareOr
	joueur.Materiaux[fer] = rareFer
	joueur.Materiaux[charbon] = rareCharbon

	fabriquer(&joueur, personage.ArmeEpeeRenforcee, map[string]int{fer: rareFer, charbon: rareCharbon}, rareOr)

	if !possedeObjetForge(joueur.Inventaire, personage.ArmeEpeeRenforcee) {
		t.Fatal("l'épée renforcée devrait être ajoutée à l'inventaire")
	}
	if joueur.Or != 0 || joueur.Materiaux[fer] != 0 || joueur.Materiaux[charbon] != 0 {
		t.Fatalf("ressources restantes = %d or, %d fer, %d charbon; attendu 0, 0, 0", joueur.Or, joueur.Materiaux[fer], joueur.Materiaux[charbon])
	}
}
