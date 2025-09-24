package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var Hangman = []string{
	`
  +---+
  |   |
      |
      |
      |
      |
========= 
	`,
	`
  +---+
  |   |
  O   |
      |
      |
      |
========= 
	`,
	`
  +---+
  |   |
  O   |
  |   |
      |
      |
========= 
	`,
	`
  +---+
  |   |
  O   |
 /|   |
      |
      |
========= 
	`,
	`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
========= 

	`,
	`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
========= 
	`,
	`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
========= 
	`,
}

func main() {
	var Word = []rune(CreatWord())
	var WordTry []rune
	Try := 6
	StartGame(Word, WordTry, Try)
}

func StartGame(Word []rune, WordTry []rune, Try int) {
	Game(Word, &WordTry, &Try)
}

func CreatWord() string {
	Dictionnaire := []string{
		"arbre", "avion", "banjo", "beaux", "blanc", "bleus", "boute", "bruit", "carte",
		"chaud", "chien", "choix", "cinqs", "clair", "coeur", "corps", "cours", "crane",
		"dents", "droit", "ecran", "fleur", "force", "fruit", "geste", "glace", "grand",
		"jouer", "livre", "lundi", "mains", "merde", "motif", "neige", "noire", "nuage",
		"paris", "pelle", "perle", "piano", "plage", "pluie", "poche", "porte", "riree",
		"rouge", "sable", "table", "faims", "pause", "vague",
	}
	return Dictionnaire[rand.Intn(50)]
}

func PrintHangman(Try *int) {
	fmt.Println(Hangman[6-*Try])
}

func Game(Word []rune, WordTry *[]rune, Try *int) {
	GameWord := []rune{'_', '_', '_', '_', '_'}
	for !GameOver(Word, GameWord, Try) {
		ClearScreen()
		fmt.Println(`
		██╗  ██╗ █████╗ ███╗   ██╗ ██████╗ ███╗   ███╗ █████╗ ███╗   ██╗
		██║  ██║██╔══██╗████╗  ██║██╔════╝ ████╗ ████║██╔══██╗████╗  ██║
		███████║███████║██╔██╗ ██║██║  ███╗██╔████╔██║███████║██╔██╗ ██║
		██╔══██║██╔══██║██║╚██╗██║██║   ██║██║╚██╔╝██║██╔══██║██║╚██╗██║
		██║  ██║██║  ██║██║ ╚████║╚██████╔╝██║ ╚═╝ ██║██║  ██║██║ ╚████║
		╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝`)
		PrintHangman(Try)
		fmt.Println(string(GameWord))
		fmt.Println("Vous avez deja marquer : ", string(*WordTry))
		TapeLetter := ScanKeyboard(*WordTry)
		*WordTry = append(*WordTry, TapeLetter)
		if VerifLetter(TapeLetter, Word) {
			for i := 0; i <= 4; i++ {
				if TapeLetter == (Word)[i] {
					GameWord[i] = TapeLetter
				}
			}
		} else {
			(*Try)--
		}
	}
}

func ScanKeyboard(WordTry []rune) rune {
	scanner := bufio.NewScanner(os.Stdin)
	// Créer un nouveau scanner
	fmt.Println("Tapes une lettre : ")
	scanner.Scan()
	input := scanner.Text()
	// Gère le texte écrit
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	if len(input) == 1 {
		Runeletter := []rune(input)
		letter := Runeletter[0]
		if letter >= 'a' && letter <= 'z' {
			return letter
		} else {
			fmt.Println("Tu dois écrire une lettre")
			return 0
		}
	} else {
		fmt.Println("Tu dois écrire une lettre")
		return 0
	}
}

func VerifLetter(TapeLetter rune, Word []rune) bool {
	for i := 0; i <= 4; i++ {
		if TapeLetter == (Word)[i] {
			return true
		}
	}
	return false
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func GameOver(Word []rune, GameWord []rune, Try *int) bool {
	if string(Word) == string(GameWord) {
		ClearScreen()
		fmt.Println("Vous avez gagné car le mot était", string(Word))
		fmt.Println(`
		██╗   ██╗ ██████╗ ██╗   ██╗    ██╗    ██╗██╗███╗   ██╗
        ╚██╗ ██╔╝██╔═══██╗██║   ██║    ██║    ██║██║████╗  ██║
         ╚████╔╝ ██║   ██║██║   ██║    ██║ █╗ ██║██║██╔██╗ ██║
          ╚██╔╝  ██║   ██║██║   ██║    ██║███╗██║██║██║╚██╗██║
           ██║   ╚██████╔╝╚██████╔╝    ╚███╔███╔╝██║██║ ╚████║
           ╚═╝    ╚═════╝  ╚═════╝      ╚══╝╚══╝ ╚═╝╚═╝  ╚═══╝`)
		return true
	} else if *Try == 0 {
		ClearScreen()
		fmt.Println("Vous avez perdu c'est GameOver car le mot était", string(Word), "c'est ciao kobuchaw")
		fmt.Println(` 
		 ██████╗  █████╗ ███╗   ███╗███████╗     ██████╗ ██╗   ██╗███████╗██████╗ 
		██╔════╝ ██╔══██╗████╗ ████║██╔════╝    ██╔═══██╗██║   ██║██╔════╝██╔══██╗
		██║  ███╗███████║██╔████╔██║█████╗      ██║   ██║██║   ██║█████╗  ██████╔╝
		██║   ██║██╔══██║██║╚██╔╝██║██╔══╝      ██║   ██║╚██╗ ██╔╝██╔══╝  ██╔══██╗
		╚██████╔╝██║  ██║██║ ╚═╝ ██║███████╗    ╚██████╔╝ ╚████╔╝ ███████╗██║  ██║
		╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝     ╚═════╝   ╚═══╝  ╚══════╝╚═╝  ╚═╝`)
		return true
	} else {
		return false
	}
}
