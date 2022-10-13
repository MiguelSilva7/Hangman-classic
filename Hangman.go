package main

import (
	"bufio"
	"fmt"
	"os"
)

type InfosUser struct {
	LastName  string
	FirstName string
	Age       int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hello What's your Lastname ?. \n" + "\n")
	LastName, _ := reader.ReadString('\n')
	fmt.Print(" What's your Firstname ?. \n" + "\n")
	FirstName, _ := reader.ReadString('\n')
	fmt.Print(" How old are you ?. \n" + "\n")
	Age, _ := reader.ReadString('\n')
	fmt.Print("\n")
	fmt.Print("Good Luck " + "\n" + LastName + FirstName + Age + "\n" + "you have 10 attempts.\n")
	//Hidden()
}

//func calcul_x() string {
//	Select := read()
//	x := Wrd_al(Select)
//	return Select[x]
//}

//func Hidden() { //le mot = _ _ _ _
//	s := string(convert(calcul_x()))
//	fmt.Println(s)
//	fmt.Print("\n")
//}
