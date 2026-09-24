# ProjectRed

Petit jeu en ligne de commande écrit en Go.

## Organisation

- `main.go` : démarre le jeu.
- `Base/base.go` : affiche le menu principal.
- `Base/adventure.go` : contient les scènes de l'aventure.
- `Personage/player.go` : décrit le joueur et ses statistiques.
- `Combat/combat.go` : contient les combats.

## Lancer le jeu

Depuis le dossier du projet :

```bash
go run .
```

Le jeu ne sauvegarde pas les parties. Une nouvelle partie recommence toujours au début.
