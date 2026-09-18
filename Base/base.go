package base

import (
	"fmt"

	personage "projectred/Personage"
)

const (
	reset  = "\033[0m"
	bold   = "\033[1m"
	dim    = "\033[2m"
	cyan   = "\033[36m"
	green  = "\033[32m"
	red    = "\033[31m"
	yellow = "\033[33m"
)

func Menu(player *personage.Player) {
	for {
		clearScreen()
		printTitle()
		fmt.Println("  ", bold+"MENU PRINCIPAL"+reset)
		fmt.Println()
		fmt.Println("  "+cyan+"[1]"+reset, "Nouvelle partie")
		fmt.Println("  "+cyan+"[2]"+reset, "Personnage")
		fmt.Println("  "+cyan+"[3]"+reset, "Charger une partie")
		fmt.Println("  "+red+"[4]"+reset, "Quitter")
		fmt.Println()
		fmt.Println("  " + dim + "----------------------------------------" + reset)
		fmt.Print("  " + yellow + "> " + reset)

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println(red + "  Entrez un nombre entre 1 et 4." + reset)
			waitForEnter()
			continue
		}

		switch choice {
		case 1:
			NouvellePartie(player)
		case 2:
			player.DisplayInfo()
		case 3:
			ChargerPartie()
		case 4:
			fmt.Println(green + "  À bientôt, aventurier." + reset)
			return
		default:
			fmt.Println(red + "  Choix incorrect. Sélectionnez une option de 1 à 4." + reset)
			waitForEnter()
		}
	}
}

func printTitle() {
	fmt.Println()
	fmt.Println("  " + cyan + "╔══════════════════════════════════════╗" + reset)
	fmt.Println("  " + cyan + "║" + reset + bold + "       LES CENDRES DE L'AUBE       " + reset + cyan + "║" + reset)
	fmt.Println("  " + cyan + "╚══════════════════════════════════════╝" + reset)
}

func NouvellePartie(player *personage.Player) {
	fmt.Println()
	clearScreen()
	printSceneHeader("NOUVELLE PARTIE", "Le début de l'aventure")
	fmt.Println("  Bienvenue,", player.Name+".")
	*player = personage.NewPlayer()
	StartAdventure(player)
}

func ChargerPartie() {
	clearScreen()
	printSceneHeader("CHARGER UNE PARTIE", "Sauvegarde")
	fmt.Println("  Aucune sauvegarde disponible pour le moment.")
	waitForEnter()
}

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func printSceneHeader(title string, subtitle string) {
	fmt.Println()
	fmt.Println("  " + cyan + "┌──────────────────────────────────────┐" + reset)
	fmt.Printf("  %s│%s %-36s %s│%s\n", cyan, reset, bold+title+reset, cyan, reset)
	fmt.Printf("  %s│%s %-36s %s│%s\n", cyan, reset, dim+subtitle+reset, cyan, reset)
	fmt.Println("  " + cyan + "└──────────────────────────────────────┘" + reset)
	fmt.Println()
}
