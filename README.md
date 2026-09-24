# ProjectRed

Petit jeu en ligne de commande écrit en Go.

## Organisation

- `main.go` : démarre le jeu.
- `Base/base.go` : affiche le menu principal.
- `Base/creation.go` : crée le personnage et demande son nom et sa classe.
- `Base/adventure.go` : contient les scènes de l'aventure.
- `Base/game_menu.go` : menu accessible pendant l'aventure.
- `Base/merchant.go` : achats et vente d'objets.
- `Base/forgeron.go` : fabrication d'armes et d'armures, amélioration de l'épée et gestion des matériaux.
- `Personage/player.go` : décrit le joueur et ses statistiques.
- `Personage/inventory.go` : affiche et gère l'inventaire.
- `Personage/items.go` : contient les noms des objets.
- `Combat/combat.go` : contient le déroulement des combats.
- `Combat/potions.go` : contient l'utilisation des potions.

## Lancer le jeu

Depuis le dossier du projet :

```bash
go run .
```

Le jeu ne sauvegarde pas les parties. Une nouvelle partie recommence toujours au début.
