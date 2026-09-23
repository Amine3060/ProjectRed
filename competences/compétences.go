package competences

import (
	"fmt"
	"sort"
	"strings"
)

type Category string

const (
	CombatCategory    Category = "COMBAT"
	SurvieCategory    Category = "SURVIE"
	TechnologieCategory Category = "TECHNOLOGIE"
)

type Skill struct {
	ID             string
	Name           string
	Category       Category
	Description    string
	Cost           int
	Prerequisites  []string
	Purchased      bool
	Bonus          string
}

type SkillTree struct {
	Skills map[string]*Skill
}

func NewSkillTree() *SkillTree {
	tree := &SkillTree{Skills: map[string]*Skill{}}

	skills := []Skill{
		{ID: "degats", Name: "Dégâts", Category: CombatCategory, Description: "Augmente les dégâts infligés en combat.", Cost: 1, Bonus: "Dégâts +10%"},
		{ID: "critique", Name: "Critique", Category: CombatCategory, Description: "Augmente les chances de coup critique.", Cost: 1, Prerequisites: []string{"degats"}, Bonus: "Critique +10%"},
		{ID: "pv", Name: "PV", Category: SurvieCategory, Description: "Augmente les points de vie maximum.", Cost: 1, Bonus: "PV +15"},
		{ID: "soins", Name: "Soins", Category: SurvieCategory, Description: "Améliore les soins reçus.", Cost: 1, Prerequisites: []string{"pv"}, Bonus: "Soins +20%"},
		{ID: "rechargement", Name: "Rechargement", Category: SurvieCategory, Description: "Récupère plus vite le mana.", Cost: 1, Bonus: "Mana +5"},
		{ID: "armes_techno", Name: "Armes techno", Category: TechnologieCategory, Description: "Débloque des armes plus avancées.", Cost: 1, Bonus: "Armes sophistiquées"},
		{ID: "drones", Name: "Drones", Category: TechnologieCategory, Description: "Active le soutien par drones.", Cost: 1, Prerequisites: []string{"armes_techno"}, Bonus: "Drone de soutien"},
	}

	for i := range skills {
		skill := skills[i]
		tree.Skills[skill.ID] = &skill
	}

	return tree
}

func (tree *SkillTree) HasPrerequisites(skillID string) bool {
	skill, exists := tree.Skills[skillID]
	if !exists {
		return false
	}

	for _, prereqID := range skill.Prerequisites {
		prereq, found := tree.Skills[prereqID]
		if !found || !prereq.Purchased {
			return false
		}
	}

	return true
}

func (tree *SkillTree) BuySkill(skillID string, points *int) bool {
	if points == nil || *points <= 0 {
		return false
	}

	skill, exists := tree.Skills[skillID]
	if !exists || skill.Purchased {
		return false
	}

	if !tree.HasPrerequisites(skillID) {
		return false
	}

	if *points < skill.Cost {
		return false
	}

	skill.Purchased = true
	*points -= skill.Cost
	return true
}

func (tree *SkillTree) GetSkillsByCategory(category Category) []Skill {
	var result []Skill
	for _, skill := range tree.Skills {
		if skill.Category == category {
			result = append(result, *skill)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result
}

func (tree *SkillTree) DisplaySkills() {
	fmt.Println()
	fmt.Println("         KRAG")
	fmt.Println("          |")
	fmt.Println(" ┌────────┼────────┐")
	fmt.Println(" ↓        ↓        ↓")
	fmt.Println()

	categories := []Category{CombatCategory, SurvieCategory, TechnologieCategory}
	maxLen := 0
	for _, category := range categories {
		if count := len(tree.GetSkillsByCategory(category)); count > maxLen {
			maxLen = count
		}
	}

	for _, category := range categories {
		fmt.Printf("%-15s", string(category))
	}
	fmt.Println()

	for i := 0; i < maxLen; i++ {
		for _, category := range categories {
			skills := tree.GetSkillsByCategory(category)
			label := ""
			if i < len(skills) {
				status := " "
				if skills[i].Purchased {
					status = "✓"
				}
				label = fmt.Sprintf("%s %s", status, skills[i].Name)
			}
			fmt.Printf("%-15s", label)
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("  Compétences disponibles :")
	for _, category := range categories {
		fmt.Println("  " + string(category))
		for _, skill := range tree.GetSkillsByCategory(category) {
			status := "[ ]"
			if skill.Purchased {
				status = "[✓]"
			}
			prereqs := ""
			if len(skill.Prerequisites) > 0 {
				prereqs = "  Prérequis : " + strings.Join(skill.Prerequisites, ", ")
			}
			fmt.Printf("    %s %s - %s%s\n", status, skill.Name, skill.Bonus, prereqs)
		}
	}
}
