package combat

import (
	"fmt"
	"math/rand"
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
	Name         string
	MaxHP        int
	HP           int
	Attack       int
	Accuracy     int
	LegsCrippled bool
	ArmsCrippled bool
	Poisoned     bool
	PoisonTurns  int
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

type bodyPart struct {
	name             string
	damageMultiplier float64
	baseHitChance    int
	cripplesLegs     bool
	cripplesArms     bool
}

var bodyParts = []bodyPart{
	{name: "Tête", damageMultiplier: 3.0, baseHitChance: 45},
	{name: "Torse", damageMultiplier: 1.6, baseHitChance: 65},
	{name: "Bras droit", damageMultiplier: 0.6, baseHitChance: 75, cripplesArms: true},
	{name: "Bras gauche", damageMultiplier: 0.6, baseHitChance: 75, cripplesArms: true},
	{name: "Jambes", damageMultiplier: 0.7, baseHitChance: 80, cripplesLegs: true},
}

// targetHitChance lowers accuracy against stronger enemies, but immobilized legs make any hit easier to land.
func targetHitChance(part bodyPart, enemy *Monster) int {
	chance := part.baseHitChance - enemy.Attack/5
	if enemy.LegsCrippled {
		chance += 30
	}
	if chance < 5 {
		chance = 5
	}
	if chance > 95 {
		chance = 95
	}
	return chance
}

// chooseBodyPart asks the player where to aim their attack.
func chooseBodyPart() (bodyPart, bool) {
	fmt.Println()
	fmt.Println("  " + dim + "Où voulez-vous frapper ?" + reset)
	for i, part := range bodyParts {
		fmt.Printf("  %s[%d]%s %s\n", cyan, i+1, reset, part.name)
	}
	fmt.Print("  " + yellow + "> " + reset)

	choice, ok := readChoice(len(bodyParts))
	if !ok {
		return bodyPart{}, false
	}
	if choice < 1 || choice > len(bodyParts) {
		fmt.Println(red + "  Choix incorrect." + reset)
		return bodyPart{}, false
	}
	return bodyParts[choice-1], true
}

// inventorySummary lists unique item names in first-seen order with their counts.
func inventorySummary(inventory []string) ([]string, map[string]int) {
	var names []string
	counts := map[string]int{}
	for _, item := range inventory {
		if counts[item] == 0 {
			names = append(names, item)
		}
		counts[item]++
	}
	return names, counts
}

// chooseItem lets the player pick an item from their inventory, or back out without using a turn.
func chooseItem(player *personage.Player) (string, bool) {
	names, counts := inventorySummary(player.Inventory)
	if len(names) == 0 {
		fmt.Println(dim + "  Votre inventaire est vide." + reset)
		return "", false
	}

	fmt.Println()
	fmt.Println("  " + dim + "Inventaire :" + reset)
	for i, name := range names {
		fmt.Printf("  %s[%d]%s %-20s x%d\n", cyan, i+1, reset, name, counts[name])
	}
	fmt.Println("  " + cyan + "[0]" + reset + " Retour")
	fmt.Print("  " + yellow + "> " + reset)

	choice, ok := readChoice(len(names))
	if !ok || choice == 0 {
		return "", false
	}
	if choice < 1 || choice > len(names) {
		fmt.Println(red + "  Choix incorrect." + reset)
		return "", false
	}
	return names[choice-1], true
}

// removeItem removes one occurrence of the named item from the inventory.
func removeItem(player *personage.Player, name string) {
	for i, item := range player.Inventory {
		if item == name {
			player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			return
		}
	}
}

// useItem applies the item's effect on the player or the enemy.
func useItem(player *personage.Player, enemy *Monster, name string) {
	switch name {
	case "Potion de vie":
		heal := 40
		player.HP += heal
		if player.HP > player.MaxHP {
			player.HP = player.MaxHP
		}
		fmt.Println(green+"  Vous buvez une Potion de vie et récupérez", heal, "PV."+reset)
	case "Potion de poison":
		enemy.Poisoned = true
		enemy.PoisonTurns = 3
		fmt.Println(green + "  Vous aspergez " + enemy.Name + " de poison." + reset)
	case "Potion explosive":
		damage := 35
		enemy.HP -= damage
		if enemy.HP < 0 {
			enemy.HP = 0
		}
		fmt.Println(green+"  L'explosion inflige", damage, "dégâts à "+enemy.Name+"."+reset)
	default:
		fmt.Println(dim + "  Cet objet n'a pas d'effet pour le moment." + reset)
	}
}

// tickPoison applies poison damage to a poisoned enemy at the end of the player's turn.
func tickPoison(enemy *Monster) {
	if !enemy.Poisoned {
		return
	}
	damage := 8
	enemy.HP -= damage
	if enemy.HP < 0 {
		enemy.HP = 0
	}
	fmt.Println(red+"  Le poison inflige", damage, "dégâts à "+enemy.Name+"."+reset)
	enemy.PoisonTurns--
	if enemy.PoisonTurns <= 0 {
		enemy.Poisoned = false
		fmt.Println(dim + "  Le poison se dissipe." + reset)
	}
}

func WarriorFight(player *personage.Player) {
	warrior := &Monster{
		Name:     "Guerrier ennemi",
		MaxHP:    1000,
		HP:       1000,
		Attack:   35,
		Accuracy: 90,
	}

	printBattleHeader("COMBAT", "Le guerrier ennemi de Krag")
	fmt.Println("  "+cyan+"Krag :"+reset, "Nous ne voulons pas nous battre. Laissez-nous partir !")
	fmt.Println("  "+red+warrior.Name+" :"+reset, "Vous ne quitterez pas ce village vivants.")
	fmt.Println("  Le combat commence.")
	pauseTurn()

	runBattle(player, warrior, battleConfig{
		title:       "COMBAT",
		subtitle:    "Le guerrier ennemi de Krag",
		winnable:    false,
		discussLine: "Le guerrier refuse de discuter.",
		onFlee: func(*personage.Player) {
			fmt.Println(dim + "  La fuite est impossible face à ce guerrier." + reset)
		},
	})

	fmt.Println(red + "  " + player.Name + " est vaincu." + reset)
}

func StartFight(player *personage.Player) {
	enemy := &Monster{
		Name:     "Soldat mystérieux",
		MaxHP:    40,
		HP:       40,
		Attack:   5,
		Accuracy: 80,
	}

	printBattleHeader("PREMIER COMBAT", "Un adversaire sur la piste de Nima")
	fmt.Println("  Un soldat ennemi bloque le chemin de", player.Name+".")
	fmt.Println("  "+cyan+player.Name+" :"+reset, "Nous n'avons pas besoin de nous battre. Pourquoi attaquez-vous notre village ?")
	fmt.Println("  "+red+enemy.Name+" :"+reset, "Je ne parle pas avec les gens de ton espèce.")
	fmt.Println("  Le soldat lève son arme.")
	offerPreFightEscape(player)
	pauseTurn()

	outcome := runBattle(player, enemy, battleConfig{
		title:       "PREMIER COMBAT",
		subtitle:    "Un adversaire sur la piste de Nima",
		winnable:    true,
		discussLine: "La discussion sera disponible plus tard.",
		onFlee:      escapeFirstFight,
	})

	fmt.Println()
	if outcome == outcomePlayerDefeated {
		fmt.Println(red + "  " + player.Name + " est vaincu." + reset)
		return
	}
	fmt.Println(green + "  Le soldat ennemi est vaincu." + reset)
	fmt.Println(green + "  " + player.Name + " a remporté son premier combat." + reset)
	player.Gold += 100
	fmt.Println(green + "  Vous trouvez 100 pièces d'or sur le soldat." + reset)
}

// offerPreFightEscape lets the player try to flee before the fight; the guard spots them and they lose 20 HP.
func offerPreFightEscape(player *personage.Player) {
	fmt.Println()
	fmt.Println("  " + cyan + "[1]" + reset + " Combattre       " + cyan + "[2]" + reset + " Tenter de fuir")
	fmt.Print("  " + yellow + "> " + reset)

	choice, ok := readChoice(2)
	if !ok || choice != 2 {
		return
	}

	fmt.Println(dim + "  Le soldat vous repère avant que vous ayez pu fuir." + reset)
	player.HP -= 20
	if player.HP < 0 {
		player.HP = 0
	}
	fmt.Println(red + "  Vous perdez 20 PV." + reset)
}

type battleOutcome int

const (
	outcomePlayerDefeated battleOutcome = iota
	outcomeEnemyDefeated
)

type battleConfig struct {
	title       string
	subtitle    string
	winnable    bool
	discussLine string
	onFlee      func(*personage.Player)
}

// runBattle drives one full turn-based fight, re-displaying the battle screen every turn.
func runBattle(player *personage.Player, enemy *Monster, cfg battleConfig) battleOutcome {
	turn := 1
	for player.HP > 0 && (!cfg.winnable || enemy.HP > 0) {
		clearScreen()
		printBattleHeader(cfg.title, cfg.subtitle)
		printBattleStatus(player, *enemy, turn)
		printBattleOptions()

		choice, ok := readChoice(6)
		if !ok {
			pauseTurn()
			continue
		}

		defending, actionTaken := resolvePlayerAction(player, enemy, choice, cfg)
		if !actionTaken {
			pauseTurn()
			continue
		}

		tickPoison(enemy)

		if cfg.winnable && enemy.HP == 0 {
			pauseTurn()
			break
		}

		enemyTurn(player, enemy, defending)
		restoreMana(player)
		pauseTurn()
		turn++
	}

	if player.HP == 0 {
		return outcomePlayerDefeated
	}
	return outcomeEnemyDefeated
}

// resolvePlayerAction applies the chosen action and reports whether it consumed the turn.
func resolvePlayerAction(player *personage.Player, enemy *Monster, choice int, cfg battleConfig) (defending bool, actionTaken bool) {
	switch choice {
	case 1:
		part, ok := chooseBodyPart()
		if !ok {
			return false, false
		}
		if rand.Intn(100) >= targetHitChance(part, enemy) {
			fmt.Println(dim + "  L'attaque visant " + strings.ToLower(part.name) + " rate sa cible." + reset)
			return false, true
		}
		damage := int(10 * part.damageMultiplier)
		enemy.HP -= damage
		if enemy.HP < 0 {
			enemy.HP = 0
		}
		fmt.Println(green+"  "+player.Name, "frappe", strings.ToLower(part.name), "et inflige", damage, "dégâts."+reset)
		if part.cripplesLegs && !enemy.LegsCrippled {
			enemy.LegsCrippled = true
			fmt.Println(yellow + "  " + enemy.Name + " est immobilisé, il sera plus facile à toucher." + reset)
		}
		if part.cripplesArms && !enemy.ArmsCrippled {
			enemy.ArmsCrippled = true
			fmt.Println(yellow + "  " + enemy.Name + " a du mal à manier son arme, sa précision baisse." + reset)
		}
		return false, true
	case 2:
		fmt.Println(yellow+"  "+player.Name, "se met en position défensive."+reset)
		return true, true
	case 3:
		fmt.Println(dim + "  " + cfg.discussLine + reset)
		return false, false
	case 4:
		name, ok := chooseItem(player)
		if !ok {
			return false, false
		}
		useItem(player, enemy, name)
		removeItem(player, name)
		return false, true
	case 5:
		if !castFireball(player, enemy) {
			return false, false
		}
		return false, true
	case 6:
		cfg.onFlee(player)
		return false, false
	default:
		fmt.Println(red + "  Choix incorrect. Veuillez choisir entre 1 et 6." + reset)
		return false, false
	}
}

func enemyTurn(player *personage.Player, enemy *Monster, defending bool) {
	fmt.Println()
	fmt.Println(yellow + "  TOUR DE L'ENNEMI" + reset)

	accuracy := enemy.Accuracy
	if enemy.ArmsCrippled {
		accuracy -= 30
	}
	if accuracy < 5 {
		accuracy = 5
	}

	if rand.Intn(100) >= accuracy {
		fmt.Println("  "+red+enemy.Name+reset, "attaque mais rate sa cible.")
		return
	}

	damage := enemy.Attack
	if defending {
		damage /= 2
	}
	fmt.Println("  "+red+enemy.Name+reset, "attaque !")
	player.HP -= damage
	if player.HP < 0 {
		player.HP = 0
	}
	fmt.Println(red+"  Vous avez reçu", damage, "dégâts."+reset)
}

func readChoice(max int) (int, bool) {
	var choice int
	if _, err := fmt.Scanln(&choice); err != nil {
		fmt.Println(red + fmt.Sprintf("  Veuillez choisir entre 1 et %d.", max) + reset)
		return 0, false
	}
	return choice, true
}

func pauseTurn() {
	fmt.Println()
	fmt.Println(dim + "  Appuyez sur Entrée pour continuer..." + reset)
	fmt.Scanln()
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
