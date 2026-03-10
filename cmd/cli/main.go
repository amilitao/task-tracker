package main

import (
	"fmt"
	"os"

	"github.com/amilitao/task-tracker/internal/domain"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Digite o comando a ser executado...")
		os.Exit(1)
	}

	arg1 := os.Args[1]
	arg2 := os.Args[2]

	switch arg1 {
	case "add":
		domain.Add(arg2)
	case "update":

	case "delete":

	case "mark-in-progress":

	case "mark-done":

	case "list":

	default:
		fmt.Println("Opção não encontrada")

	}

}
