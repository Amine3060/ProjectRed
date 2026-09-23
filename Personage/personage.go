package personage

import (
	"encoding/json"
	"fmt"
	"os"
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
	Gold           int
	EquippedWeapon string
	Spells         []string
	Skills         *competences.SkillTree
	Inventory      []string
	Progress       int
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
		SkillPoints:    3,
		Gold:           0,
		EquippedWeapon: "Aucune",
		Spells:         []string{"Fireball"},
		Skills:         competences.NewSkillTree(),
		Inventory:      []string{},
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

func (p *Player) OpenSkillsMenu() {
	if p.Skills == nil {
		p.Skills = competences.NewSkillTree()
	}
	p.Skills.InteractiveMenu(p.Name, &p.SkillPoints)
}

// Save writes the player's state to a JSON file.
func (p *Player) Save(path string) error {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// LoadPlayer reads a player's state from a JSON save file.
func LoadPlayer(path string) (Player, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Player{}, err
	}
	var p Player
	if err := json.Unmarshal(data, &p); err != nil {
		return Player{}, err
	}
	return p, nil
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
	fmt.Println("  Or                  ", p.Gold)
	fmt.Println()
	fmt.Println("  " + dim + "Équipement" + reset)
	fmt.Println("  Arme équipée        ", p.EquippedWeapon)
	fmt.Println("  Sorts               ", strings.Join(p.Spells, ", "))
	fmt.Println("  Inventaire          ", inventoryOrNone(p.Inventory))
	fmt.Println()
	if p.Skills != nil {
		p.Skills.DisplaySkills()
	}
	fmt.Println()
}

func inventoryOrNone(inventory []string) string {
	if len(inventory) == 0 {
		return "Vide"
	}
	return strings.Join(inventory, ", ")
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
