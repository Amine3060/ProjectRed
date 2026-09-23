package personage

import (
	"fmt"
	"strings"

	competences "projectred/competences"
)

const (
	reset = "\033[0m"
	bold  = "\033[1m"
	dim   = "\033[2m"
	cyan  = "\033[36m"
	green = "\033[32m"
)

type Player struct {
	Name           string
	HP             int
	MaxHP          int
	Mana           int
	MaxMana        int
	Level          int
	XP             int
	MaxXP          int
	SkillPoints    int
	EquippedWeapon string
	Spells         []string
	Skills         *competences.SkillTree
}

func NewPlayer() Player {
	return Player{
		Name:           "Krag",
		HP:             100,
		MaxHP:          100,
		Mana:           50,
		MaxMana:        50,
		Level:          1,
		XP:             0,
		MaxXP:          100,
		SkillPoints:    0,
		EquippedWeapon: "Aucune",
		Spells:         []string{"Fireball"},
		Skills:         competences.NewSkillTree(),
	}
}

func (p *Player) SpendMana(cost int) bool {
	if p.Mana < cost {
		return false
	}
	p.Mana -= cost
	return true
}

func (p *Player) RestoreMana(amount int) {
	p.Mana += amount
	if p.Mana > p.MaxMana {
		p.Mana = p.MaxMana
	}
}

func (p *Player) BuySkill(skillID string) bool {
	if p.Skills == nil {
		p.Skills = competences.NewSkillTree()
	}
	return p.Skills.BuySkill(skillID, &p.SkillPoints)
}

func (p Player) DisplayInfo() {
	fmt.Println()
	fmt.Println("  " + cyan + "┌──────────────────────────────────────┐" + reset)
	fmt.Println("  " + cyan + "║" + reset + bold + "             PERSONNAGE              " + reset + cyan + "║" + reset)
	fmt.Println("  " + cyan + "└──────────────────────────────────────┘" + reset)
	fmt.Println()
	fmt.Println("  " + dim + "Identité" + reset)
	fmt.Println("  Nom                 ", bold+p.Name+reset)
	fmt.Println("  Niveau              ", p.Level)
	fmt.Println()
	fmt.Println("  " + dim + "État" + reset)
	fmt.Printf("  PV                  %s\n", healthBar(p.HP, p.MaxHP))
	fmt.Printf("  Mana                %s %d/%d\n", healthBar(p.Mana, p.MaxMana), p.Mana, p.MaxMana)
	fmt.Printf("  XP                  %s %d/%d\n", healthBar(p.XP, p.MaxXP), p.XP, p.MaxXP)
	fmt.Println("  Points de compétence", p.SkillPoints)
	fmt.Println()
	fmt.Println("  " + dim + "Équipement" + reset)
	fmt.Println("  Arme équipée        ", p.EquippedWeapon)
	fmt.Println("  Sorts               ", strings.Join(p.Spells, ", "))
	fmt.Println()
	if p.Skills != nil {
		p.Skills.DisplaySkills()
	}
	fmt.Println()
}

func healthBar(current int, maximum int) string {
	const width = 20
	filled := current * width / maximum
	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}
	return green + "[" + strings.Repeat("█", filled) + dim + strings.Repeat("░", width-filled) + green + "]" + reset
}
