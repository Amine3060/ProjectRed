package base

import (
	"fmt"

	combat "projectred/Combat"
	personage "projectred/Personage"
)

type adventureStep struct {
	name string
	run  func(*personage.Player)
}

var adventureSteps = []adventureStep{
	{"introduction", IntroductionScene},
	{"escape", EscapeScene},
	{"warriorFight", combat.WarriorFight},
	{"defeatTransition", func(*personage.Player) { DefeatTransitionScene() }},
	{"wake", WakeAfterDefeatScene},
	{"forestSearch", ForestSearchScene},
	{"firstFight", combat.StartFight},
	{"merchant", MerchantScene},
}

func StartAdventure(player *personage.Player) {
	ResumeAdventure(player)
}

// ResumeAdventure runs the remaining adventure steps starting from player.Progress,
// saving the game after each completed step.
func ResumeAdventure(player *personage.Player) {
	for i := player.Progress; i < len(adventureSteps); i++ {
		adventureSteps[i].run(player)
		player.Progress = i + 1
		saveGame(player)
	}
}

func IntroductionScene(player *personage.Player) {
	clearScreen()
	printSceneHeader("INTRODUCTION", "Le village de Krag")
	fmt.Println(dim + "                                                  _" + reset)
	fmt.Println(dim + "  __                   ___                       ( )" + reset)
	fmt.Println(dim + " |\"\"|  ___    _   __  |\"\"\"|  __                   `" + reset)
	fmt.Println(dim + " |\"\"| |\"\"\"|  |\"| |\"\"| |\"\"\"| |\"\"|        _._ _" + reset)
	fmt.Println(dim + " |\"\"| |\"\"\"|  |\"| |\"\"| |\"\"\"| |\"\"|       (__((_(" + reset)
	fmt.Println(dim + " |\"\"| |\"\"\"|  |\"| |\"\"| |\"\"\"| |\"\"|      \\'-:--:-." + reset)
	fmt.Println(dim + " \"'''\"''\"'\"\"'\"\"\"''\"''''\"\"\"'\"\"'\"\"'~~~~~~'-----'~~~~  ldb" + reset)
	fmt.Println()
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
	player.HP = player.MaxHP / 2
	clearScreen()
	printSceneHeader("LE RÉVEIL", "Après la défaite")
	fmt.Println("  ", player.Name, "se réveille au milieu des ruines du village.")
	fmt.Println("  La douleur du combat est encore présente, mais Nima n'est plus là.")
	fmt.Println("  Krag comprend qu'elle a été capturée et décide de partir à sa recherche.")
	waitForEnter()
}

func DefeatTransitionScene() {
	clearScreen()
	fmt.Println(dim + "  Krag perd connaissance..." + reset)
	fmt.Println()
	fmt.Println(dim + "  ` : | | | |:  ||  :     `  :  |  |+|: | : : :|   .        `              ." + reset)
	fmt.Println(dim + "      ` : | :|  ||  |:  :    `  |  | :| : | : |:   |  .                    :" + reset)
	fmt.Println(dim + "         .' ':  ||  |:  |  '       ` || | : | |: : |   .  `           .   :." + reset)
	fmt.Println(dim + "                `'  ||  |  ' |   *    ` : | | :| |*|  :   :               :|" + reset)
	fmt.Println(dim + "        *    *       `  |  : :  |  .      ` ' :| | :| . : :         *   :.||" + reset)
	fmt.Println(dim + "             .`            | |  |  : .:|       ` | || | : |: |          | ||" + reset)
	fmt.Println(dim + "      '          .         + `  |  :  .: .         '| | : :| :    .   |:| ||" + reset)
	fmt.Println(dim + "         .                 .    ` *|  || :       `    | | :| | :      |:| |" + reset)
	fmt.Println(dim + " .                .          .        || |.: *          | || : :     :|||" + reset)
	fmt.Println(dim + "        .            .   . *    .   .  ` |||.  +        + '| |||  .  ||`" + reset)
	fmt.Println(dim + "     .             *              .     +:`|!             . ||||  :.||`" + reset)
	fmt.Println(dim + " +                      .                ..!|*          . | :`||+ |||`" + reset)
	fmt.Println(dim + "     .                         +      : |||`        .| :| | | |.| ||`     ." + reset)
	fmt.Println(dim + "       *     +   '               +  :|| |`     :.+. || || | |:`|| `" + reset)
	fmt.Println(dim + "                            .      .||` .    ..|| | |: '` `| | |`  +" + reset)
	fmt.Println(dim + "  .       +++                      ||        !|!: `       :| |" + reset)
	fmt.Println(dim + "              +         .      .    | .      `|||.:      .||    .      .    `" + reset)
	fmt.Println(dim + "          '                           `|.   .  `:|||   + ||'     `" + reset)
	fmt.Println(dim + "  __    +      *                         `'       `'|.    `:" + reset)
	fmt.Println(dim + "\"'  `---\"\"\"----....____,..^---`^``----.,.___          `.    `.  .    ____,.,-" + reset)
	fmt.Println(dim + "    ___,--'\"\"`---\"'   ^  ^ ^        ^       \"\"\"'---,..___ __,..---\"\"'" + reset)
	fmt.Println(dim + "--\"'                           ^                         ``--..,__ D. Rice" + reset)
	waitForEnter()
}

func ForestSearchScene(player *personage.Player) {
	clearScreen()
	printSceneHeader("DANS LA FORÊT", "Sur la piste de Nima")
	fmt.Println("  ", player.Name, "s'avance dans la forêt à la recherche de Nima.")
	fmt.Println("  Après plusieurs heures, il aperçoit un soldat isolé sur le chemin.")
	waitForEnter()
}

type potion struct {
	name string
	cost int
}

var merchantPotions = []potion{
	{name: "Potion de vie", cost: 55},
	{name: "Potion de poison", cost: 60},
	{name: "Potion explosive", cost: 100},
}

func MerchantScene(player *personage.Player) {
	clearScreen()
	printSceneHeader("LE MARCHAND", "Une rencontre sur le chemin")
	fmt.Println("  En s'aventurant plus loin,", player.Name, "croise un marchand ambulant.")
	fmt.Println("  "+cyan+"Le marchand :"+reset, "J'ai quelques potions qui pourraient vous être utiles.")

	for {
		fmt.Println()
		fmt.Println("  " + dim + "Or disponible : " + reset + fmt.Sprint(player.Gold))
		fmt.Println()
		for i, p := range merchantPotions {
			fmt.Printf("  %s[%d]%s %-20s %s(%d pièces)%s\n", cyan, i+1, reset, p.name, dim, p.cost, reset)
		}
		fmt.Println("  " + cyan + "[4]" + reset + " Partir")
		fmt.Print("  " + yellow + "> " + reset)

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println(red + "  Veuillez choisir entre 1 et 4." + reset)
			continue
		}

		if choice == 4 {
			fmt.Println("  "+cyan+"Le marchand :"+reset, "Bonne route, aventurier.")
			break
		}

		if choice < 1 || choice > len(merchantPotions) {
			fmt.Println(red + "  Choix incorrect. Veuillez choisir entre 1 et 4." + reset)
			continue
		}

		chosen := merchantPotions[choice-1]
		if player.Gold < chosen.cost {
			fmt.Println(red + "  Vous n'avez pas assez d'or pour cet achat." + reset)
			continue
		}

		player.Gold -= chosen.cost
		player.Inventory = append(player.Inventory, chosen.name)
		fmt.Println(green + "  Vous achetez : " + chosen.name + reset)
	}

	waitForEnter()
}

func waitForEnter() {
	fmt.Println()
	fmt.Println("  " + dim + "Appuyez sur Entrée pour continuer..." + reset)
	fmt.Scanln()
}
