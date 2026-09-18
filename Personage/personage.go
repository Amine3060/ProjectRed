package personage

import "fmt"

type Player struct {
	Name           string
	HP             int
	MaxHP          int
	Level          int
	XP             int
	MaxXP          int
	SkillPoints    int
	EquippedWeapon string
}

func NewPlayer() Player {
	return Player{
		Name:           "Krag",
		HP:             100,
		MaxHP:          100,
		Level:          1,
		XP:             0,
		MaxXP:          100,
		SkillPoints:    0,
		EquippedWeapon: "Aucune",
	}
}

func (p Player) DisplayInfo() {
	fmt.Println()
	fmt.Println("================================")
	fmt.Println("           PERSONNAGE")
	fmt.Println("================================")
	fmt.Println("Nom :", p.Name)
	fmt.Printf("PV : %d / %d\n", p.HP, p.MaxHP)
	fmt.Println("Niveau :", p.Level)
	fmt.Printf("XP : %d / %d\n", p.XP, p.MaxXP)
	fmt.Println("Points de compétence :", p.SkillPoints)
	fmt.Println("Arme équipée :", p.EquippedWeapon)
	fmt.Println("================================")
}
