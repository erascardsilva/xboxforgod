// Erasmo Cardoso - Dev
package isoconvert

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

// TransformarISOemGOD é responsável por listar ISOs, solicitar arquivo de entrada e nome de saída, e converter para GOD.
func TransformarISOemGOD(reader *bufio.Reader) {
	fmt.Println("\n--- Transformar ISO em GOD ---")

	fmt.Println("Listando arquivos ISO no diretório atual (ls *.iso):")
	cmdLs := exec.Command("bash", "-c", "ls -1 *.iso")
	outputLs, errLs := cmdLs.CombinedOutput()

	if errLs != nil {
		if exitErr, ok := errLs.(*exec.ExitError); ok && exitErr.ExitCode() == 2 {
			fmt.Println("Nenhum arquivo .iso encontrado no diretório atual.")
		} else {
			fmt.Printf("Erro ao executar 'ls *.iso': %s\n", errLs)
			fmt.Printf("Saída de erro do ls:\n%s\n", string(outputLs))
			fmt.Println("Verifique se o comando 'ls' está disponível e se você tem permissão para listar arquivos.")
		}
		return
	}

	if len(strings.TrimSpace(string(outputLs))) > 0 {
		fmt.Printf("%s\n", string(outputLs))
	} else {
		fmt.Println("Nenhum arquivo .iso encontrado no diretório atual.")
		return
	}

	fmt.Print("Digite o nome do arquivo ISO de entrada (ex: meu_jogo.iso): ")
	inputIso, _ := reader.ReadString('\n')
	inputIso = strings.TrimSpace(inputIso)

	if inputIso == "" {
		fmt.Println("Nome do arquivo ISO de entrada não pode ser vazio.")
		return
	}

	fmt.Print("Digite o nome da pasta GOD de saída (ex: meu_jogo): ")
	outputGod, _ := reader.ReadString('\n')
	outputGod = strings.TrimSpace(outputGod)

	if outputGod == "" {
		fmt.Println("Nome da pasta GOD de saída não pode ser vazio.")
		return
	}

	fmt.Printf("Convertendo '%s' para a pasta '%s'...\n", inputIso, outputGod)
	cmdIso2God := exec.Command("iso2god", inputIso, outputGod)

	outputConvert, errConvert := cmdIso2God.CombinedOutput()
	if errConvert != nil {
		fmt.Printf("Erro ao converter ISO para GOD: %s\n", errConvert)
		fmt.Printf("Saída do iso2god:\n%s\n", string(outputConvert))
		return
	}

	fmt.Println("ISO convertida para GOD com sucesso!")
	fmt.Printf("Saída do iso2god:\n%s\n", string(outputConvert))
}
