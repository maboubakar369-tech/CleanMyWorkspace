package main

import (
	"fmt"

	"github.com/maboubakar369-tech/CleanMyWorkspace/Clean"
	cmw "github.com/maboubakar369-tech/CleanMyWorkspace/Mentor/Paris/CleanMyWorkspace"
)

func main() {
	// Génération du souk
	workspace := cmw.GenererateWorkSpace()

	fmt.Println("===== SOUK AVANT NETTOYAGE =====")
	printWorkspace(workspace)

	// Nettoyage du souk
	Clean.CleanWorkSpace(workspace)

	fmt.Println("\n===== SOUK APRÈS NETTOYAGE =====")
	printWorkspace(workspace)
}

func printWorkspace(workspace *[][]*string) {
	for _, row := range *workspace {
		fmt.Print("|")
		for _, item := range row {
			if item == nil {
				fmt.Print("nil|")
			} else {
				fmt.Printf("%s|", *item)
			}
		}
		fmt.Println()
	}
}
