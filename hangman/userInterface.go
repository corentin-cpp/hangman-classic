package hangman

import (
	"fmt"
	"os"
	"strings"

	"github.com/01-edu/z01"
)

func InitInterface() {
	ClearT()
	fmt.Println(` /$$$$$$$  /$$$$$$ /$$$$$$$$ /$$   /$$ /$$    /$$ /$$$$$$$$ /$$   /$$ /$$   /$$ /$$$$$$$$
| $$__  $$|_  $$_/| $$_____/| $$$ | $$| $$   | $$| $$_____/| $$$ | $$| $$  | $$| $$_____/
| $$  \ $$  | $$  | $$      | $$$$| $$| $$   | $$| $$      | $$$$| $$| $$  | $$| $$      
| $$$$$$$   | $$  | $$$$$   | $$ $$ $$|  $$ / $$/| $$$$$   | $$ $$ $$| $$  | $$| $$$$$   
| $$__  $$  | $$  | $$__/   | $$  $$$$ \  $$ $$/ | $$__/   | $$  $$$$| $$  | $$| $$__/   
| $$  \ $$  | $$  | $$      | $$\  $$$  \  $$$/  | $$      | $$\  $$$| $$  | $$| $$      
| $$$$$$$/ /$$$$$$| $$$$$$$$| $$ \  $$   \  $/   | $$$$$$$$| $$ \  $$|  $$$$$$/| $$$$$$$$
|_______/ |______/|________/|__/  \__/    \_/    |________/|__/  \__/ \______/ |________/
                                                                                         
                                                                                         
                                                                                         `)

}

func Game() {
	fmt.Print("\n")
	fmt.Print("\n")

	//Init Level game
	var level string
	fmt.Println("Entrer la difficulté du jeu : 1 - Facile, 2 - Moyen, 3 - Difficile")
	fmt.Println("Entrer la commande stop pour quitter le jeu, la partie sauvegarder ne fonctionne pas encore pour des raisons techniques")
	fmt.Scanln(&level)

	//Init struct with default value
	w := LoadWord(level)
	pl := LoadPlayer()

	//Init Pos struct
	var p Pos
	for i := 0; i < len(w.letters)-1; i++ {
		p.result = append(p.result, '_')
	}

	//Random letters
	fmt.Println("Lettres du mot : ")
	z := RandomLetters(w, p)

	for i := 0; i < len(z); i++ {
		fmt.Print(string(z[i]))
		fmt.Print(" ")
	}

	for {
		fmt.Print("\n")
		fmt.Println("Entrer une lettre ou le mot : ")

		for i := 0; i < len(p.result); i++ {
			fmt.Print(string(p.result[i]))
		}
		fmt.Println(" ")
		var entry string
		fmt.Scanln(&entry)

		if entry == "exit" {
			Stop(p, w, pl)
		}

		chars := strings.ToLower(entry)
		p.revel += " " + chars

		fmt.Println("Caractère saisie : ", p.revel)
		fmt.Println("Nombre de vie restant : ", pl.life)

		if CheckScan(chars) == true {
			beforeFind := p.find
			if len(chars) > 1 {
				if CheckWords(chars, w) == false {
					pl.life = pl.life - 1
					Level(pl.life)
					if pl.life == 0 {
						ClearT()
						Level(pl.life)
						LoseGame(w.word, pl.name)
						os.Exit(0)
					}
				} else {
					Win(w.word, pl.name)
				}
			} else {
				result := CheckWord(w, chars, p)
				if result.find == beforeFind {
					pl.life = pl.life - 1
					Level(pl.life)
					if pl.life == 0 {
						ClearT()
						Level(pl.life)
						LoseGame(w.word, pl.name)
						os.Exit(0)
					}
				}
				if CheckWin(w, p) == true {
					Win(w.word, pl.name)
				}

				z01.PrintRune('\n')
			}
		}

	}
}

func CheckScan(value string) bool {
	if len(value) == 0 {
		return false
	}
	valueR := []rune(value)

	if valueR[0] > 'A' || valueR[0] > 'Z' || valueR[0] > 'z' || valueR[0] > 'a' {
		return true
	}

	return false
}

func LoseGame(word string, playerName string) {
	fmt.Println(` /$$$$$$$  /$$$$$$$$ /$$$$$$$  /$$$$$$$  /$$   /$$
	| $$__  $$| $$_____/| $$__  $$| $$__  $$| $$  | $$
	| $$  \ $$| $$      | $$  \ $$| $$  \ $$| $$  | $$
	| $$$$$$$/| $$$$$   | $$$$$$$/| $$  | $$| $$  | $$
	| $$____/ | $$__/   | $$__  $$| $$  | $$| $$  | $$
	| $$      | $$      | $$  \ $$| $$  | $$| $$  | $$
	| $$      | $$$$$$$$| $$  | $$| $$$$$$$/|  $$$$$$/
	|__/      |________/|__/  |__/|_______/  \______/ 
													  
													  
													  `)
	fmt.Println("Désolé ", playerName, " vous avez perdu le mot était : ", word)
	fmt.Println(" ")
}

func Win(word string, playerName string) {
	ClearT()
	fmt.Println("Bravo ", playerName, " vous avez trouvé le mot : ", word)
	fmt.Println(" ")
	fmt.Println(` /$$$$$$$$ /$$$$$$ /$$   /$$ /$$$$$$       /$$
| $$_____/|_  $$_/| $$$ | $$|_  $$_/      | $$
| $$        | $$  | $$$$| $$  | $$        | $$
| $$$$$     | $$  | $$ $$ $$  | $$        | $$
| $$__/     | $$  | $$  $$$$  | $$        |__/
| $$        | $$  | $$\  $$$  | $$            
| $$       /$$$$$$| $$ \  $$ /$$$$$$       /$$
|__/      |______/|__/  \__/|______/      |__/
                                              
                                              
                                              `)
	os.Exit(0)
}
