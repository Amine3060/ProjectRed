package main

import (
	base "projectred/Base"
	personage "projectred/Personage"
)

func main() {
	player := personage.NewPlayer()
	base.Menu(&player)
}
