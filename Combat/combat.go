package combat

import (
	"fmt"
	"math/rand"

	personage "projectred/Personage"
)

type Monster struct {
	Name   string
	HP     int
	Attack int
}

func WarriorFight(player *personage.Player) {
	fmt.Println("\n=== PREMIER COMBAT ===")
	fmt.Println("Le guerrier est trop fort. Krag est vaincu.")
	player.HP = 0
	fmt.Println("PV de", player.Name, ":", player.HP)
	waitForEnter()
}

func StartFight(player *personage.Player) {
	enemy := Monster{
		Name:   "Soldat ennemi",
		HP:     40,
		Attack: 8,
	}

	fmt.Println("\n=== SECOND COMBAT ===")
	fmt.Println("Un", enemy.Name, "attaque !")

	for player.HP > 0 && enemy.HP > 0 {
		fmt.Println("\nTes PV :", player.HP, "| PV ennemi :", enemy.HP)
		fmt.Println("1. Attaquer")
		fmt.Println("2. Lancer Fireball (10 mana)")
		fmt.Println("3. Fuir")
		fmt.Print("Choix : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			damage := rand.Intn(11) + 10
			enemy.HP -= damage
			fmt.Println("Tu infliges", damage, "dégâts.")
		case 2:
			if player.UseMana(10) {
				enemy.HP -= 25
				fmt.Println("Fireball inflige 25 dégâts.")
			} else {
				fmt.Println("Tu n'as pas assez de mana.")
				continue
			}
		case 3:
			fmt.Println("Tu prends la fuite.")
			return
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		if enemy.HP <= 0 {
			break
		}

		player.HP -= enemy.Attack
		fmt.Println("L'ennemi inflige", enemy.Attack, "dégâts.")
	}

	if player.HP <= 0 {
		fmt.Println("Tu as perdu le combat.")
	} else {
		player.Gold += 100
		fmt.Println("Tu as gagné ! Tu trouves 100 pièces d'or.")
	}

	waitForEnter()
}

func waitForEnter() {
	fmt.Println("Appuie sur Entrée pour continuer.")
	fmt.Scanln()
}
