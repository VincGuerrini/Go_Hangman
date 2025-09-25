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
	fmt.Println("\033[38;2;231;222;121m", `
		██╗  ██╗ █████╗ ███╗   ██╗ ██████╗ ███╗   ███╗ █████╗ ███╗   ██╗
		██║  ██║██╔══██╗████╗  ██║██╔════╝ ████╗ ████║██╔══██╗████╗  ██║
		███████║███████║██╔██╗ ██║██║  ███╗██╔████╔██║███████║██╔██╗ ██║
		██╔══██║██╔══██║██║╚██╗██║██║   ██║██║╚██╔╝██║██╔══██║██║╚██╗██║
		██║  ██║██║  ██║██║ ╚████║╚██████╔╝██║ ╚═╝ ██║██║  ██║██║ ╚████║
		╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝`, "\033[0m")
	Menu(&Word, WordTry, Try)
}

func Menu(Word *[]rune, WordTry []rune, Try int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1 - Jouer")
		fmt.Println("2 - Quitter")
		fmt.Print("Choisis une option : ")

		Input, _ := reader.ReadString('\n')
		Choice := strings.TrimSpace(Input)

		switch Choice {
		case "1":
			fmt.Println("Lets go Play !")
			*Word = []rune(CreatWord())
			StartGame(*Word, WordTry, Try)
		case "2":
			fmt.Println("Bye bye !!")
			os.Exit(0)
		default:
			ClearScreen()
			fmt.Println("Ton choix n'est pas correct retry !")
		}
	}
}

func StartGame(Word []rune, WordTry []rune, Try int) {
	Game(Word, &WordTry, &Try)
}

func Game(Word []rune, WordTry *[]rune, Try *int) {
	GameWord := []rune{'_', '_', '_', '_', '_'}
	for !GameOver(Word, GameWord, Try) {
		ClearScreen()
		fmt.Println("\033[38;2;231;222;121m", `
		██╗  ██╗ █████╗ ███╗   ██╗ ██████╗ ███╗   ███╗ █████╗ ███╗   ██╗
		██║  ██║██╔══██╗████╗  ██║██╔════╝ ████╗ ████║██╔══██╗████╗  ██║
		███████║███████║██╔██╗ ██║██║  ███╗██╔████╔██║███████║██╔██╗ ██║
		██╔══██║██╔══██║██║╚██╗██║██║   ██║██║╚██╔╝██║██╔══██║██║╚██╗██║
		██║  ██║██║  ██║██║ ╚████║╚██████╔╝██║ ╚═╝ ██║██║  ██║██║ ╚████║
		╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝`, "\033[0m")
		PrintHangman(Try)
		fmt.Println(strings.ToTitle(string(GameWord)))
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

func ScanKeyboard(WordTry []rune) rune {
	scanner := bufio.NewScanner(os.Stdin)
	// Créer un nouveau scanner
	fmt.Print("Tapes une lettre : ")
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
		fmt.Println("Vous avez gagné car le mot était", "\033[35m"+strings.ToTitle(string(Word))+"\033[0m")
		fmt.Println("\033[32m", `
	██╗   ██╗ ██████╗ ██╗   ██╗    ██╗    ██╗██╗███╗   ██╗
        ╚██╗ ██╔╝██╔═══██╗██║   ██║    ██║    ██║██║████╗  ██║
         ╚████╔╝ ██║   ██║██║   ██║    ██║ █╗ ██║██║██╔██╗ ██║
          ╚██╔╝  ██║   ██║██║   ██║    ██║███╗██║██║██║╚██╗██║
           ██║   ╚██████╔╝╚██████╔╝    ╚███╔███╔╝██║██║ ╚████║
           ╚═╝    ╚═════╝  ╚═════╝      ╚══╝╚══╝ ╚═╝╚═╝  ╚═══╝`, "\033[0m")
		return true
	} else if *Try == 0 {
		ClearScreen()
		fmt.Println("Vous avez perdu c'est GameOver car le mot était", "\033[35m"+strings.ToTitle(string(Word))+"\033[0m", "c'est ciao kobuchaw")
		fmt.Println("\033[31m", `
		 ██████╗  █████╗ ███╗   ███╗███████╗     ██████╗ ██╗   ██╗███████╗██████╗ 
		██╔════╝ ██╔══██╗████╗ ████║██╔════╝    ██╔═══██╗██║   ██║██╔════╝██╔══██╗
		██║  ███╗███████║██╔████╔██║█████╗      ██║   ██║██║   ██║█████╗  ██████╔╝
		██║   ██║██╔══██║██║╚██╔╝██║██╔══╝      ██║   ██║╚██╗ ██╔╝██╔══╝  ██╔══██╗
		╚██████╔╝██║  ██║██║ ╚═╝ ██║███████╗    ╚██████╔╝ ╚████╔╝ ███████╗██║  ██║
		╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝     ╚═════╝   ╚═══╝  ╚══════╝╚═╝  ╚═╝`, "\033[0m")
		return true
	} else {
		return false
	}
}
