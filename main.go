package main

import (
	base "projectred/Base"
	personage "projectred/Personage"
)

func main() {
	joueur := personage.NouveauJoueur("Krag", "Humain")
	base.Menu(&joueur)
}
