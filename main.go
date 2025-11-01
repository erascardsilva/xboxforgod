// Erasmo Cardoso - Dev
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	isoconvert "xboxforgod/features/iso_convert"
	isocopy "xboxforgod/features/iso_copy"
	"xboxforgod/utils"
)

func main() {
	if os.Getenv("USER") != "root" {
		fmt.Println("Este programa precisa ser executado com privilégios de root para acessar dispositivos de DVD e instalar dependências.")
		fmt.Println("Por favor, execute com 'sudo go run main.go' ou 'sudo ./seu_programa'.")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	utils.VerificarEInstalarDependencias()

	for {
		exibirMenuPrincipal()

		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, digite um número.")
			continue
		}

		switch choice {
		case 1:
			isocopy.CopiarISO(reader)
		case 2:
			isoconvert.TransformarISOemGOD(reader)
		case 3:
			fmt.Println("Saindo do programa. Até logo!")
			return
		default:
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
		}
	}
}

func exibirMenuPrincipal() {
	fmt.Println("\n--- Menu Principal XboxForGOD ---")
	fmt.Println("1. Copiar ISO de DVD")
	fmt.Println("2. Transformar ISO em GOD")
	fmt.Println("3. Sair")
	fmt.Print("Escolha uma opção: ")
}
