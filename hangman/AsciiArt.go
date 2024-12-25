package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func Ascii(s string) {
	file, err := os.Open("./src/ascii-art.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterBlocks := make(map[rune][]string)
	var lines []string
	letterIndex := 0
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
		if len(lines) == 5 {
			if letterIndex < len(alphabet) {
				letter := rune(alphabet[letterIndex])
				letterBlocks[letter] = append([]string{}, lines...)
				letterIndex++
			}
			lines = []string{}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
	}

	for lineIdx := 0; lineIdx < 5; lineIdx++ {
		for _, letter := range s {
			if block, exists := letterBlocks[letter]; exists {
				fmt.Print(block[lineIdx] + " ")
			} else {
				fmt.Print("                        ")
			}
		}
		fmt.Println()
	}
}
