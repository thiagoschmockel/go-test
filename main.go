package main

import "fmt"

func hogwartsWizards() []string {
	return []string{"Harry Potter", "Hermione Granger", "Ronald Weasley"}
}

func main() {
	var wizards[]string = hogwartsWizards()
	
	for i:=0; len(wizards) > i; i++ {
	  fmt.Println(wizards[i])	
	}
}