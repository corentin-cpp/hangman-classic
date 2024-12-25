
# Présentation
Le programme "Hangman" est un jeu du pendu en ligne de commande écrit en Go. Le but est de deviner le mot en proposant des lettres avant que la partie soit perdue en raison d'échecs répétés. Le programme permet de gérer les niveaux de difficulté et enregistre les informations du joueur et les résultats intermédiaires.

# Fonctions 
## `Level`

`func Level(level int)`
Affiche le dessin de la potence selon le nombre d'erreurs restant. Les différents cas de l'instruction switch représentent des étapes spécifiques dans la progression de la potence.

`level`: entier représentant le niveau de la potence (de 9 à 0).
Affiche le dessin de la potence correspondant au niveau actuel.

## `LoadWord`

`func LoadWord(level string) Word`

Charge un mot aléatoire dans un fichier texte en fonction du niveau spécifié.

`level`: chaîne de caractères représentant la difficulté (nom du fichier, par ex., "1" pour word1.txt).
Retourne une structure Word initialisée avec le mot choisi aléatoirement.

## `LoadPlayer`

`func LoadPlayer() Player`

Demande à l'utilisateur d'entrer son nom pour initialiser la structure du joueur avec des valeurs par défaut.

Demande un nom d'utilisateur en ligne de commande et le stocke.
Initialise la structure Player avec 10 vies et un niveau de difficulté par défaut de 1.


## `CheckWord`

`func CheckWord(w Word, value string, p Pos) Pos`

Vérifie si une lettre proposée existe dans le mot.

`w`: instance de la structure Word.
value: lettre proposée par le joueur.
`p`: état actuel de la structure Pos, qui contient les résultats intermédiaires.
Retourne la structure Pos mise à jour avec les lettres trouvées.

## `CheckWin`

`func CheckWin(w Word, p Pos) bool`
Vérifie si toutes les lettres du mot ont été trouvées.

`w`: instance de Word.
`p`: instance de Pos qui contient les résultats du jeu.
Retourne true si le joueur a gagné, sinon false.

## `CheckWords`

`func CheckWords(value string, w Word) bool`
Vérifie si le mot complet proposé par le joueur est correct.

`value`: mot proposé par le joueur.
`w`: structure Word contenant le mot correct.
Retourne true si le mot proposé est correct, sinon false.

## `ClearT`

`func ClearT()`
Efface l'écran de la console.

Utilise cmd sous Windows et clear pour les autres systèmes.
Nettoie l'affichage pour garder une interface de jeu propre.


## `Stop`

`func Stop(p Pos, w Word, pl Player)`
Arrête le jeu en affichant l'état final sous forme JSON.

`p`: état actuel de la structure Pos.
`w`: mot actuel sous forme de structure Word.
`pl`: joueur actuel sous forme de structure Player.
Affiche les informations en JSON et termine le programme.

## `RandomLetters`

func RandomLetters(w Word, p Pos) []rune
Retourne un sous-ensemble de lettres du mot pour aider le joueur.

`w`: mot actuel sous forme de structure Word.
`p`: position actuelle sous forme de structure Pos.
Retourne un tableau de rune contenant des lettres aléatoires du mot.

# Structures
## `HangmanGame`
Structure principale pour gérer les informations du jeu en cours.

```
type HangmanGame struct {
    words  Word
    player Player
    pos    Pos
}
```

`words`: mot actuel sous forme de Word.

`player`: joueur actuel sous forme de Player.

`pos`: position actuelle sous forme de Pos.

## `Word`
Représente le mot choisi pour le jeu.

```
type Word struct {
    word        string
    size        int
    letters     []rune
    sizeDisplay int
}
```

`word`: le mot à deviner.

`size`: la longueur du mot.

`letters`: tableau de lettres du mot.

`sizeDisplay`: nombre de lettres à afficher initialement.


## `Player`
Représente le joueur actuel.
```
type Player struct {
    name      string
    life      int
    difficult int
}
```

`name`: nom du joueur.

`life`: nombre de vies restantes.

`difficult`: niveau de difficulté.

## `Pos`
Gère les résultats intermédiaires du jeu.

``` golang
type Pos struct {
    result []rune
    revel  string
    find   int
}
```

`result`: état actuel des lettres trouvées.

`revel`: lettres révélées pour aider le joueur.

`find`: nombre de lettres trouvées.




[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/1YLV-els)

