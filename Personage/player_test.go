package personage

import "testing"

func TestDegatsRecusAvecArmure(t *testing.T) {
	joueur := NouveauJoueur("Krag", "Humain")
	joueur.ArmureEquipee = true

	if degats := joueur.DegatsRecus(20); degats != 15 {
		t.Fatalf("dégâts reçus avec armure = %d, attendu 15", degats)
	}
}

func TestDegatsRecusSansArmure(t *testing.T) {
	joueur := NouveauJoueur("Krag", "Humain")

	if degats := joueur.DegatsRecus(20); degats != 20 {
		t.Fatalf("dégâts reçus sans armure = %d, attendu 20", degats)
	}
}
