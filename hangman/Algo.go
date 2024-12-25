package hangman

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Level(level int) {
	switch level {
	case 9:
		fmt.Println("=========")
		break
	case 8:
		fmt.Println(`      |  
      |  
      |  
      |  
      |  
=========`)
		break
	case 7:
		fmt.Println(`  +---+  
      |  
      |  
      |  
      |  
      |  
=========`)
		break
	case 6:
		fmt.Println(`  +---+  
  |   |  
      |  
      |  
      |  
      |  
=========`)
		break
	case 5:
		fmt.Println(`  +---+  
  |   |  
  O   |  
      |  
      |  
      |  
=========`)
		break
	case 4:
		fmt.Println(`  +---+  
  |   |  
  O   |  
  |   |  
      |  
      |  
=========`)
		break
	case 3:
		fmt.Println(`  +---+  
  |   |  
  O   |  
 /|   |  
      |  
      |  
=========`)
		break
	case 2:
		fmt.Println(`  +---+  
  |   |  
  O   |  
 /|\  |  
      |  
      |  
=========`)
		break
	case 1:
		fmt.Println(`  +---+  
  |   |  
  O   |  
 /|\  |  
 /    |  
      |  
=========`)
		break
	case 0:
		fmt.Println(`  +---+  
  |   |  
  O   |  
 /|\  |  
 / \  |  
      |  
=========`)

	}
}

func LoadWord(level string) Word {
	//Get all word in files
	content, err := os.ReadFile("./src/word" + level + ".txt")

	if err != nil {
		log.Fatal(err)
	}

	//split all word
	word := strings.Split(string(content), "\n")

	//init Struct Word with value
	var newWord Word
	newWord.word = word[rand.Intn((len(word)))]
	newWord.size = len(word)
	newWord.sizeDisplay = len(word) / 2
	newWord.letters = []rune(newWord.word)

	return newWord
}

func LoadPlayer() Player {
	var nameUser string
	fmt.Println("Entrer votre pseudo : ")
	fmt.Scanln(&nameUser)
	ClearT()
	var p Player
	p.name = nameUser
	p.life = 10
	p.difficult = 1

	return p
}

func CheckWord(w Word, value string, p Pos) Pos {
	r := []rune(value)
	v := w.letters
	ru := p.result
	//Check is correct
	for i := 0; i < len(v); i++ {
		if v[i] == r[0] {
			ru[i] = r[0]
			p.find += 1
		}
	}

	return p
}

func CheckWin(w Word, p Pos) bool {
	ru := p.result
	word := w.letters
	for i := 0; i < len(ru); i++ {
		if ru[i] == '_' {
			return false
		} else if ru[i] != word[i] {
			return false
		}
	}
	return true
}

func CheckWords(value string, w Word) bool {
	r := []rune(value)
	v := w.letters

	if len(r) == len(v)-1 {
		for i := 0; i < len(r); i++ {
			if r[i] != v[i] {
				return false
			}
		}
	} else {
		return false
	}

	return true
}

func ClearT() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Stop(p Pos, w Word, pl Player) {

	h := &HangmanGame{
		words:  w,
		player: pl,
		pos:    p,
	}

	b, _ := json.Marshal(h)
	fmt.Println(string(b))
	os.Exit(0)
}

func RandomLetters(w Word, p Pos) []rune {
	var usedLetters []rune
	for i := 0; i < len(w.letters)/2; i++ {
		r := rand.Intn(len(w.letters) / 2)
		if len(usedLetters) != 0 {
			for x := 0; x < len(usedLetters); x++ {
				if usedLetters[x] != w.letters[r] {
					usedLetters = append(usedLetters, w.letters[r])
				}
			}
		} else {
			usedLetters = append(usedLetters, w.letters[r])
		}
	}
	return usedLetters
}

type HangmanGame struct {
	words  Word
	player Player
	pos    Pos
}

type Word struct {
	word        string
	size        int
	letters     []rune
	sizeDisplay int
}

type Player struct {
	name      string
	life      int
	difficult int
}

type Pos struct {
	result []rune
	revel  string
	find   int
}
