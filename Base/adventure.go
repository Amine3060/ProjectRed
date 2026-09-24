package base

import (
	"bufio"
	"fmt"
	"os"

	combat "projectred/Combat"
	personage "projectred/Personage"
)

func StartAdventure(player *personage.Player) {
	introduction(player)
	escape(player)
	combat.WarriorFight(player)
	wakeUp(player)
	forest(player)
	combat.StartFight(player)
	finishAdventure(player)
}

func waitForEnter() {
	fmt.Println("\nAppuyez sur Entrée pour continuer.")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func introduction(player *personage.Player) {
	fmt.Println("\n=== INTRODUCTION ===")
	fmt.Println(player.Name, "habite dans un petit village avec sa sœur Nima.")
	fmt.Println("Un jour, une armée attaque le village.")
	waitForEnter()
}

func escape(player *personage.Player) {
	fmt.Println("\n=== LA FUITE ===")
	fmt.Println(player.Name, "essaie de fuir avec Nima.")
	fmt.Println("Un guerrier ennemi bloque le chemin.")
	waitForEnter()
}

func wakeUp(player *personage.Player) {
	player.HP = player.MaxHP / 2
	fmt.Println("\n=== LE REVEIL ===")
	fmt.Println(player.Name, "se réveille dans les ruines du village.")
	fmt.Println("Nima a été capturée. Il part à sa recherche.")
	waitForEnter()
}

func forest(player *personage.Player) {
	fmt.Println("\n=== LA FORET ===")
	fmt.Println(player.Name, "avance dans la forêt et rencontre un soldat.")
	waitForEnter()
}

func finishAdventure(player *personage.Player) {
	fmt.Println("\nL'aventure continue...")
	fmt.Println("Or gagné :", player.Gold)
	waitForEnter()
}
