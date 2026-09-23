package competences

import (
	"fmt"
	"sort"
	"strings"
)

type Category string

const (
	CombatCategory        Category = "COMBAT"
	SurvieCategory        Category = "SURVIE"
	TechnologieCategory   Category = "TECHNOLOGIE"
)

type Skill struct {
	ID            string
	Name          string
	Category      Category
	Description   string
	Cost          int
	Prerequisites []string
	Purchased     bool
	Bonus         string
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

func (tree *SkillTree) InteractiveMenu(playerName string, points *int) {
	if points == nil {
		fmt.Println("  Aucun point de compétence disponible.")
		return
	}

	for {
		fmt.Println()
		fmt.Println("  ┌──────────────────────────────────────┐")
		fmt.Printf("  │ %-36s │\n", "COMPETENCES")
		fmt.Printf("  │ %-36s │\n", "Joueur : "+playerName)
		fmt.Printf("  │ %-36s │\n", fmt.Sprintf("Points : %d", *points))
		fmt.Println("  └──────────────────────────────────────┘")
		fmt.Println()
		fmt.Println("  [1] COMBAT")
		fmt.Println("  [2] SURVIE")
		fmt.Println("  [3] TECHNOLOGIE")
		fmt.Println("  [0] Retour")
		fmt.Print("  > ")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("  Choix invalide.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if choice == 0 {
			return
		}

		var category Category
		switch choice {
		case 1:
			category = CombatCategory
		case 2:
			category = SurvieCategory
		case 3:
			category = TechnologieCategory
		default:
			fmt.Println("  Choix incorrect.")
			continue
		}

		if !tree.displayCategorySkills(category, points) {
			fmt.Println("  Retour au menu des compétences.")
		}
	}
}

func (tree *SkillTree) displayCategorySkills(category Category, points *int) bool {
	skills := tree.GetSkillsByCategory(category)
	if len(skills) == 0 {
		fmt.Println("  Aucune compétence dans cette catégorie.")
		return false
	}

	for i, skill := range skills {
		status := "[ ]"
		if skill.Purchased {
			status = "[✓]"
		}
		prereqText := ""
		if len(skill.Prerequisites) > 0 {
			prereqText = " | Prérequis : " + strings.Join(skill.Prerequisites, ", ")
		}
		fmt.Printf("  [%d] %s %s - %s%s\n", i+1, status, skill.Name, skill.Bonus, prereqText)
	}
	fmt.Println("  [0] Retour")
	fmt.Print("  Sélection : ")

	var selection int
	if _, err := fmt.Scanln(&selection); err != nil {
		fmt.Println("  Entrez un numéro valide.")
		var discard string
		fmt.Scanln(&discard)
		return false
	}

	if selection == 0 {
		return false
	}

	if selection < 1 || selection > len(skills) {
		fmt.Println("  Numéro invalide.")
		return false
	}

	selectedSkill := skills[selection-1]
	if selectedSkill.Purchased {
		fmt.Println("  Vous possédez déjà cette compétence.")
		return true
	}

	if !tree.HasPrerequisites(selectedSkill.ID) {
		fmt.Println("  Prérequis manquants pour :", selectedSkill.Name)
		if len(selectedSkill.Prerequisites) > 0 {
			fmt.Println("  Il faut d'abord acheter :", strings.Join(selectedSkill.Prerequisites, ", "))
		}
		return true
	}

	if *points < selectedSkill.Cost {
		fmt.Println("  Pas assez de points de compétence.")
		return true
	}

	if tree.BuySkill(selectedSkill.ID, points) {
		fmt.Println("  Compétence acquise :", selectedSkill.Name)
		fmt.Println("  Bonus :", selectedSkill.Bonus)
	} else {
		fmt.Println("  Achat impossible.")
	}
	return true
}
