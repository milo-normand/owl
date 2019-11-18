package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cast"
)

func main() {
	prompt := promptui.Prompt{
		Label: "Entre un chiffre",
	}

	input, err := prompt.Run()
	chiffre := cast.ToInt(input)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("Le double du chiffre que tu as choisi est %d\n", chiffre*2)
}
