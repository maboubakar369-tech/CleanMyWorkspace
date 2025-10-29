package main 

import (
	"fmt"
	"github.com/maboubakar369/CleanMyWorkspace/Clean"
)

func main() {
	souk := GénérerWorkSpace()

	fmt.Println("=== Espace de travail AVANT le nettoyage ===")
	GenererWorkSpace(souk)

	soukNettoye := Clean.CleanWorkSpace(souk)

	fmt.Println("\n=== Espace de travail APRÈS le nettoyage ===")
	GenererWorkSpace(soukNettoye)
}

func GenererWorkSpace(workspace *[][]*string) {
	for _, ligne := range *workspace {
		fmt.Print("|")
		for _, case_ := range ligne {
			if case_ == nil {
				fmt.Print("nil|")
			} else {
				fmt.Printf("%s|", *case_)
			}
		}
		fmt.Println()
	}
}




