// Erasmo Cardoso - Dev
package isocopy

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// CopiarISO é responsável por listar unidades, solicitar caminho do DVD e nome da ISO, e criar a ISO.
func CopiarISO(reader *bufio.Reader) {
	fmt.Println("\n--- Copiar ISO de DVD ---")

	fmt.Println("Listando unidades disponíveis (lsblk):")
	cmdLsblk := exec.Command("lsblk")
	outputLsblk, errLsblk := cmdLsblk.CombinedOutput()
	if errLsblk != nil {
		fmt.Printf("Erro ao executar lsblk: %s\n", errLsblk)
		return
	}
	fmt.Printf("%s\n", string(outputLsblk))

	fmt.Print("Digite o caminho do dispositivo DVD (ex: sr0 ou /dev/sr0): ")
	dvdPath, _ := reader.ReadString('\n')
	dvdPath = strings.TrimSpace(dvdPath)

	if !strings.HasPrefix(dvdPath, "/dev/") {
		dvdPath = "/dev/" + dvdPath
	}

	if dvdPath == "" {
		fmt.Println("Caminho do DVD não pode ser vazio.")
		return
	}

	fmt.Print("Digite o nome do arquivo ISO a ser criado (ex: minha_iso): ")
	isoName, _ := reader.ReadString('\n')
	isoName = strings.TrimSpace(isoName)

	if isoName == "" {
		fmt.Println("Nome da ISO não pode ser vazio.")
		return
	}

	if !strings.HasSuffix(isoName, ".iso") {
		isoName += ".iso"
	}

	fmt.Printf("Criando ISO '%s' a partir de '%s'...\n", isoName, dvdPath)
	cmdDd := exec.Command("dd", "if="+dvdPath, "of="+isoName, "bs=4M", "status=progress")

	cmdDd.Stdout = os.Stdout
	cmdDd.Stderr = os.Stderr

	errDd := cmdDd.Run()
	if errDd != nil {
		fmt.Printf("Erro ao criar ISO: %s\n", errDd)
		return
	}

	fmt.Println("ISO criada com sucesso!")
}
