package combat

import (
	"fmt"
	"strings"

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

type Monster struct {
	Name   string
	MaxHP  int
	HP     int
	Attack int
}

type Spell struct {
	Name     string
	ManaCost int
	Damage   int
}

var fireball = Spell{
	Name:     "Fireball",
	ManaCost: 10,
	Damage:   25,
}

func WarriorFight(player *personage.Player) {
	warrior := Monster{
		Name:   "Guerrier ennemi",
		MaxHP:  1000,
		HP:     1000,
		Attack: 35,
	}

	printBattleHeader("COMBAT IMPOSSIBLE", "Le guerrier ennemi est trop puissant")
	fmt.Println("  "+cyan+"Krag :"+reset, "Nous ne voulons pas nous battre. Laissez-nous partir !")
	fmt.Println("  "+red+warrior.Name+" :"+reset, "Vous ne quitterez pas ce village vivants.")
	fmt.Println("  Le combat commence, mais ce guerrier est bien trop puissant pour Krag.")

	turn := 1
	for player.HP > 0 {
		printBattleStatus(player, warrior, turn)
		printBattleOptions()

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println(red + "  Veuillez choisir entre 1 et 6." + reset)
			continue
		}

		defending := false
		switch choice {
		case 1:
			damage := 10
			warrior.HP -= damage
			fmt.Println(green+"  "+player.Name, "inflige", damage, "dégâts."+reset)
		case 5:
			if !castFireball(player, &warrior) {
				continue
			}
		case 2:
			defending = true
			fmt.Println(yellow+"  "+player.Name, "se met en position défensive."+reset)
		case 3:
			fmt.Println(dim + "  Le guerrier refuse de discuter." + reset)
			continue
		case 4:
			fmt.Println(dim + "  L'inventaire sera disponible plus tard." + reset)
			continue
		case 6:
			fmt.Println(dim + "  La fuite est impossible face à ce guerrier." + reset)
			continue
		default:
			fmt.Println(red + "  Choix incorrect. Veuillez choisir entre 1 et 6." + reset)
			continue
		}

		damage := warrior.Attack
		if defending {
			damage /= 2
		}
		fmt.Println(yellow + "  TOUR DE L'ENNEMI" + reset)
		fmt.Println("  "+red+warrior.Name+reset, "attaque !")
		player.HP -= damage
		if player.HP < 0 {
			player.HP = 0
		}
		fmt.Println(red+"  Vous avez reçu", damage, "dégâts."+reset)
		restoreMana(player)
		turn++
	}

	fmt.Println(red + "  " + player.Name + " est vaincu." + reset)
}

func StartFight(player *personage.Player) {
	enemy := Monster{
		Name:   "Soldat mystérieux",
		MaxHP:  40,
		HP:     40,
		Attack: 5,
	}

	printBattleHeader("PREMIER COMBAT", "Un adversaire sur la piste de Nima")
	fmt.Println("  Un soldat ennemi bloque le chemin de", player.Name+".")
	fmt.Println("  "+cyan+player.Name+" :"+reset, "Nous n'avons pas besoin de nous battre. Pourquoi attaquez-vous notre village ?")
	fmt.Println("  "+red+enemy.Name+" :"+reset, "Je ne parle pas avec les gens de ton espèce.")
	fmt.Println("  Le soldat lève son arme. Le combat se déroulera au tour par tour.")

	turn := 1
	for player.HP > 0 && enemy.HP > 0 {
		printBattleStatus(player, enemy, turn)
		printBattleOptions()

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println(red + "  Veuillez choisir entre 1 et 6." + reset)
			continue
		}

		defending := false
		switch choice {
		case 1:
			damage := 10
			enemy.HP -= damage
			if enemy.HP < 0 {
				enemy.HP = 0
			}
			fmt.Println(green+"  "+player.Name, "inflige", damage, "dégâts."+reset)
		case 5:
			if !castFireball(player, &enemy) {
				continue
			}
		case 2:
			defending = true
			fmt.Println(yellow+"  "+player.Name, "se met en position défensive."+reset)
		case 3:
			fmt.Println(dim + "  La discussion sera disponible plus tard." + reset)
			continue
		case 4:
			fmt.Println(dim + "  L'inventaire sera disponible plus tard." + reset)
			continue
		case 6:
			escapeFirstFight(player)
			continue
		default:
			fmt.Println(red + "  Choix incorrect. Veuillez choisir entre 1 et 6." + reset)
			continue
		}

		if enemy.HP == 0 {
			break
		}

		damage := enemy.Attack
		if defending {
			damage /= 2
		}
		fmt.Println(yellow + "  TOUR DE L'ENNEMI" + reset)
		fmt.Println("  "+red+enemy.Name+reset, "attaque !")
		player.HP -= damage
		if player.HP < 0 {
			player.HP = 0
		}
		fmt.Println(red+"  Vous avez reçu", damage, "dégâts."+reset)
		restoreMana(player)
		turn++
	}

	fmt.Println()
	if player.HP == 0 {
		fmt.Println(red + "  " + player.Name + " est vaincu." + reset)
		return
	}
	fmt.Println(green + "  Le soldat ennemi est vaincu." + reset)
	fmt.Println(green + "  " + player.Name + " a remporté son premier combat." + reset)
}

func printBattleHeader(title string, subtitle string) {
	fmt.Println()
	fmt.Println("  " + cyan + "┌──────────────────────────────────────┐" + reset)
	fmt.Printf("  %s│%s %-36s %s│%s\n", cyan, reset, bold+title+reset, cyan, reset)
	fmt.Printf("  %s│%s %-36s %s│%s\n", cyan, reset, dim+subtitle+reset, cyan, reset)
	fmt.Println("  " + cyan + "└──────────────────────────────────────┘" + reset)
}

func printBattleStatus(player *personage.Player, enemy Monster, turn int) {
	fmt.Println()
	fmt.Println("  " + yellow + "TOUR " + fmt.Sprint(turn) + " · TOUR DE KRAG" + reset)
	fmt.Printf("  %-20s %s\n", player.Name, healthBar(player.HP, player.MaxHP))
	fmt.Printf("  %-20s %s\n", "Mana", manaBar(player.Mana, player.MaxMana))
	fmt.Printf("  %-20s %s\n", enemy.Name, healthBar(enemy.HP, enemy.MaxHP))
}

func printBattleOptions() {
	fmt.Println()
	fmt.Println("  " + cyan + "[1]" + reset + " Attaquer       " + cyan + "[2]" + reset + " Défendre")
	fmt.Println("  " + cyan + "[3]" + reset + " Discuter        " + cyan + "[4]" + reset + " Inventaire")
	fmt.Println("  " + cyan + "[5]" + reset + " Sort : Fireball  " + dim + "(10 mana)" + reset)
	fmt.Println("  " + cyan + "[6]" + reset + " Fuir")
	fmt.Print("  " + yellow + "> " + reset)
}

func escapeFirstFight(player *personage.Player) {
	clearScreen()
	printBattleHeader("FUITE IMPOSSIBLE", "Le soldat vous a rattrapé")
	fmt.Println("  Krag tente de s'enfuir dans la forêt.")
	fmt.Println("  Mais le soldat le rattrape avant qu'il puisse disparaître.")

	damage := player.HP / 2
	player.HP -= damage
	fmt.Println(red+"  Vous avez perdu", damage, "PV en tentant de fuir."+reset)
	fmt.Println()
	fmt.Println(dim + "  Retour au combat..." + reset)
	fmt.Println("  Appuyez sur Entrée pour continuer...")
	fmt.Scanln()
}

func castFireball(player *personage.Player, enemy *Monster) bool {
	if !player.SpendMana(fireball.ManaCost) {
		fmt.Println(red + "  Mana insuffisant pour lancer Fireball." + reset)
		return false
	}

	enemy.HP -= fireball.Damage
	if enemy.HP < 0 {
		enemy.HP = 0
	}
	fmt.Println(cyan+"  "+player.Name, "lance", fireball.Name, "et inflige", fireball.Damage, "dégâts."+reset)
	return true
}

func restoreMana(player *personage.Player) {
	const manaRecovery = 5
	previousMana := player.Mana
	player.RestoreMana(manaRecovery)
	recoveredMana := player.Mana - previousMana
	if recoveredMana > 0 {
		fmt.Println(cyan+"  Mana récupéré : +", recoveredMana, "PM."+reset)
	}
}

func healthBar(current int, maximum int) string {
	return resourceBar(current, maximum, "PV")
}

func manaBar(current int, maximum int) string {
	return resourceBar(current, maximum, "PM")
}

func resourceBar(current int, maximum int, label string) string {
	const width = 20
	filled := current * width / maximum
	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}
	return green + "[" + strings.Repeat("█", filled) + dim + strings.Repeat("░", width-filled) + green + "]" + reset + fmt.Sprintf(" %d/%d %s", current, maximum, label)
}

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}
