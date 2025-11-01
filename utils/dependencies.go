// Erasmo Cardoso - Dev
package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// VerificarEInstalarDependencias verifica se 'dd' e 'iso2god' estão instalados e tenta instalá-los.
func VerificarEInstalarDependencias() {
	fmt.Println("Verificando dependências...")

	// dd
	_, errDd := exec.LookPath("dd")
	if errDd != nil {
		fmt.Println("Comando 'dd' não encontrado. Tentando instalar...")
		fmt.Println("Será necessária a sua aprovação para a instalação (sudo).")

		cmdInstallDd := exec.Command("sudo", "bash", "-c", "apt-get update && apt-get install -y coreutils")
		cmdInstallDd.Stdin = os.Stdin
		output, installErr := cmdInstallDd.CombinedOutput()
		if installErr != nil {
			fmt.Printf("Erro ao instalar 'dd' (coreutils): %s\n%s\n", installErr, string(output))
			fmt.Println("Por favor, verifique sua conexão com a internet ou instale 'dd' manualmente e execute o programa novamente.")
			os.Exit(1)
		}
		fmt.Println("'dd' (coreutils) instalado com sucesso.")
	} else {
		fmt.Println("'dd' já está instalado.")
	}

	// iso2god
	_, errIso2God := exec.LookPath("iso2god")
	if errIso2God != nil {
		fmt.Println("Comando 'iso2god' não encontrado.")
		fmt.Println("A instalação de 'iso2god' geralmente requer passos manuais ou um PPA específico, pois não está nos repositórios padrão do Ubuntu/Debian.")
		fmt.Println("Por favor, instale 'iso2god' manualmente (pesquise por 'iso2god install linux' para instruções) e execute o programa novamente.")
	} else {
		fmt.Println("'iso2god' já está instalado.")
	}
}
