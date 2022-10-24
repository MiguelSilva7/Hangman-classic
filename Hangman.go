package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type InfosUser struct {
	LastName  string
	FirstName string
	Age       int
}

type Hangman struct { //structure du jeu
	Word             []rune   //mot = "_ _ _ _ _"
	ToFind           []rune   //trouve le mot
	Attempts         int      // essaie pour le mot
	HangmanPositions []string // positions du pendu
}

func main() {
	data, err := ioutil.ReadFile("hangman.txt") // lire le fichier text.txt
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data)) // conversion de byte en string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hello what's your Lastname ? \n" + "\n")
	LastName, _ := reader.ReadString('\n')
	fmt.Print("What's your Firstname ? \n" + "\n")
	FirstName, _ := reader.ReadString('\n')
	fmt.Print("How old are you ? \n" + "\n")
	Age, _ := reader.ReadString('\n')
	fmt.Print("\n")
	fmt.Print("Good Luck " + "\n" + LastName + FirstName + Age + "\n" + "You have 10 attempts.\n")
	Hidden()
}

func readpostition(Game *Hangman) {
	f, err := os.Open("jose.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewScanner(f)
	count := 0
	temp := ""
	var tab []string
	for reader.Scan() {
		if count != 0 && count%8 == 0 {
			tab = append(tab, temp)
			temp = ""
		}
		temp = temp + reader.Text() + "\n"
		count++
	}
	tab = append(tab, temp)
	Game.HangmanPositions = tab
	fmt.Print(Game.HangmanPositions[-(Game.Attempts - 9)])
}

func Read() []string { //fonction qui lit le fichier
	content, err := ioutil.ReadFile("words.txt") //lecture du fichier words.txt

	if err != nil { //message d'erreur
		log.Fatal(err)
	}
	str := string(content)          //convertis bytes strings
	Mot := strings.Split(str, "\n") // sépare les mots
	return Mot                      //renvoie un tableau de string
}

func convert(s string) []rune { //convertisseur string tableau de rune
	var Game Hangman
	Game.Attempts = 10
	Game.ToFind = []rune(s) //mot transformé en tableau de rune & lien à la structure
	var Hiddentab []rune    //définition d'un nouveau tableau de rune
	for i := 0; i < len(s); i++ {
		Hiddentab = append(Hiddentab, 95) //tableau de rune correspondant au mot caché
	}
	Game.Word = Hiddentab // attribution du tableau de rune au mot caché dans la structure
	compare(&Game)
	return Game.Word
}

func Wrd_al(x []string) int { // Fonction nombre aléatoire
	rand.Seed(time.Now().UnixNano()) // random
	min := 0                         //numéro minimum
	max := len(x)                    // numéro max
	return rand.Intn(max-min) + min  // numéro aléatoire
}

func calcul_x() string {
	Select := Read()
	x := Wrd_al(Select)
	return Select[x]
}

func Hidden() { //le mot = _ _ _ _
	s := string(convert(calcul_x()))
	fmt.Println(s)
	fmt.Print("\n")
}

func compare(Game *Hangman) { //Fonction qui permet de faire fonctionner le jeu
	var st rune
	count := 0
	Word := Game.ToFind
	Hideword := Game.Word
	fmt.Print(string(Hideword))
	fmt.Print("\n")
	for {
		if Game.Attempts <= 0 {
			data, err := ioutil.ReadFile("lose.txt")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(data))
			restart()
		} else {
			reader := bufio.NewReader(os.Stdin)
			pickaletter, _ := reader.ReadString('\n')
			st = rune(pickaletter[0])
			for i := 0; i < len(Word); i++ {
				if st == Word[i] {
					Hideword[i] = st
				}
			}
			fmt.Println(string(Hideword))
			for k := 0; k < len(Word); k++ {
				if st != Word[k] {
					count++
				}
			}

			//println("count : ", count)
			if string(Word) == string(Hideword) {
				data, err := ioutil.ReadFile("win.txt") // lire le fichier text.txt
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(string(data)) // conversion de byte en string
				restart()
			} else if count == len(Word) {
				Game.Attempts--
				readpostition(Game)
				fmt.Print("\n")
				fmt.Println("Try again, lives= ", Game.Attempts)
			} else if count != len(Word) {
				fmt.Print("\n")
				fmt.Println("Continue, lives= ", Game.Attempts)
			}
			count = 0
		}

	}
}

func restart() {
	fmt.Println("Restart a party? (y/n)")
	for {
		reader := bufio.NewReader(os.Stdin)
		Renew, _ := reader.ReadString('\n')
		if rune(Renew[0]) != 110 && rune(Renew[0]) != 121 {
			fmt.Println("Wrong answer !")
		} else if rune(Renew[0]) == 121 {
			fmt.Print("\n")
			Hidden()
		} else if rune(Renew[0]) == 110 {
			os.Exit(1)
			break
		}
	}
}
