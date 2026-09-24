package personage

import "fmt"

// Player contient les informations importantes du personnage.
type Player struct {
	Name    string
	HP      int
	MaxHP   int
	Mana    int
	MaxMana int
	Gold    int
}

func NewPlayer() Player {
	return Player{
		Name:    "Krag",
		HP:      100,
		MaxHP:   100,
		Mana:    50,
		MaxMana: 50,
		Gold:    0,
	}
}

func (player *Player) UseMana(amount int) bool {
	if player.Mana < amount {
		return false
	}

	player.Mana -= amount
	return true
}

func (player Player) DisplayInfo() {
	fmt.Println("\n=== PERSONNAGE ===")
	fmt.Println("Nom :", player.Name)
	fmt.Println("PV :", player.HP, "/", player.MaxHP)
	fmt.Println("Mana :", player.Mana, "/", player.MaxMana)
	fmt.Println("Or :", player.Gold)
}
