package competences

import "testing"

func TestBuySkillSuccess(t *testing.T) {
	tree := NewSkillTree()
	points := 1

	if !tree.BuySkill("degats", &points) {
		t.Fatal("la compétence degats ne devrait pas être bloquée")
	}

	if points != 0 {
		t.Fatalf("il devait rester 0 point, obtenu %d", points)
	}

	if !tree.Skills["degats"].Purchased {
		t.Fatal("la compétence degats n'a pas été achetée")
	}
}

func TestBuySkillRequiresPrerequisite(t *testing.T) {
	tree := NewSkillTree()
	points := 1

	if tree.BuySkill("critique", &points) {
		t.Fatal("critique ne devrait pas être achetable sans le prérequis degats")
	}

	if points != 1 {
		t.Fatalf("les points ne doivent pas être consommés sans prérequis, obtenu %d", points)
	}
}

func TestBuySkillWithPrerequisiteWorks(t *testing.T) {
	tree := NewSkillTree()
	points := 2

	if !tree.BuySkill("degats", &points) {
		t.Fatal("degats devait être achetable")
	}

	if !tree.BuySkill("critique", &points) {
		t.Fatal("critique devait être achetable après degats")
	}

	if points != 0 {
		t.Fatalf("il devait rester 0 point après deux achats, obtenu %d", points)
	}
}
