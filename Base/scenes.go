package base

import (
	"fmt"

	combat "projectred/Combat"
	personage "projectred/Personage"
)

func StartAdventure(player *personage.Player) {
	IntroductionScene(player)
	EscapeScene(player)
	combat.WarriorFight(player)
	DefeatTransitionScene()
	WakeAfterDefeatScene(player)
	ForestSearchScene(player)
	combat.StartFight(player)
}

func IntroductionScene(player *personage.Player) {
	clearScreen()
	printSceneHeader("INTRODUCTION", "Le village de Krag")
	fmt.Println("  À l'écart des grandes routes, un petit village gobelin vit paisiblement.")
	fmt.Println("  ", player.Name, "est un jeune gobelin qui partage chaque jour avec sa petite sœur Nima.")
	fmt.Println("  Leur vie est simple : des journées calmes, un village uni et la promesse de lendemains tranquilles.")
	fmt.Println("  Mais ce matin-là, Krag se réveille brusquement au son des cris et des armes.")
	fmt.Println("  Une mystérieuse armée vient d'attaquer le village.")
	waitForEnter()
}

func EscapeScene(player *personage.Player) {
	clearScreen()
	printSceneHeader("LA FUITE", "Quitter le village")
	fmt.Println("  Au milieu de l'attaque,", player.Name, "essaie de fuir le village avec Nima.")
	fmt.Println("  Ils cherchent un passage entre les maisons en flammes.")
	fmt.Println("  Mais un guerrier ennemi se dresse devant eux.")
	waitForEnter()
}

func WakeAfterDefeatScene(player *personage.Player) {
	player.HP = player.MaxHP
	clearScreen()
	printSceneHeader("LE RÉVEIL", "Après la défaite")
	fmt.Println("  ", player.Name, "se réveille au milieu des ruines du village.")
	fmt.Println("  La douleur du combat est encore présente, mais Nima n'est plus là.")
	fmt.Println("  Krag comprend qu'elle a été capturée et décide de partir à sa recherche.")
	waitForEnter()
}

func DefeatTransitionScene() {
	clearScreen()
	fmt.Println(dim + "                 .   .   ." + reset)
	fmt.Println(dim + "           .   .   .   .   .   ." + reset)
	fmt.Println(dim + "      .   .   .   .   .   .   .   ." + reset)
	fmt.Println(dim + "           .   .   .   .   .   ." + reset)
	fmt.Println(dim + "                 .   .   ." + reset)
	fmt.Scanln()
}

func ForestSearchScene(player *personage.Player) {
	clearScreen()
	printSceneHeader("DANS LA FORÊT", "Sur la piste de Nima")
	fmt.Println("  ", player.Name, "s'avance dans la forêt à la recherche de Nima.")
	fmt.Println("  Après plusieurs heures, il aperçoit un soldat isolé sur le chemin.")
	waitForEnter()
}

func waitForEnter() {
	fmt.Println()
	fmt.Println("  " + dim + "Appuyez sur Entrée pour continuer..." + reset)
	fmt.Scanln()
}
