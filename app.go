// Erasmo Cardoso - Software Engineer | Electronics Technician
package main

import (
	"bufio"
	"context"
	"embed"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed bin/*
var embeddedBinaries embed.FS

type App struct {
	ctx context.Context
}

type DependencyStatus struct {
	DdInstalled      bool   `json:"ddInstalled"`
	Iso2GodInstalled bool   `json:"iso2godInstalled"`
	Platform         string `json:"platform"`
}

type DriveInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Size string `json:"size"`
}

type ProgressEvent struct {
	Type    string  `json:"type"`
	Message string  `json:"message"`
	Percent float64 `json:"percent"`
	Bytes   int64   `json:"bytes"`
	Total   int64   `json:"total"`
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetPlatform() string {
	return runtime.GOOS
}

func (a *App) getIso2GodPath() string {
	binName := "iso2god"
	if runtime.GOOS == "windows" {
		binName = "iso2god.exe"
	}

	data, err := embeddedBinaries.ReadFile("bin/" + binName)
	if err == nil {
		tmpDir := filepath.Join(os.TempDir(), "xboxforgod_bin")
		os.MkdirAll(tmpDir, 0755)
		
		outPath := filepath.Join(tmpDir, binName)
		if stat, err := os.Stat(outPath); err == nil && stat.Size() == int64(len(data)) {
			return outPath
		}
		
		err = os.WriteFile(outPath, data, 0755)
		if err == nil {
			return outPath
		}
	}

	return binName
}

func (a *App) VerificarDependencias() DependencyStatus {
	status := DependencyStatus{
		Platform: runtime.GOOS,
	}

	isoCmd := a.getIso2GodPath()

	switch runtime.GOOS {
	case "linux":
		_, err := exec.LookPath("dd")
		status.DdInstalled = err == nil

		_, err = exec.LookPath(isoCmd)
		status.Iso2GodInstalled = err == nil

	case "windows":
		status.DdInstalled = true
		_, err := exec.LookPath(isoCmd)
		status.Iso2GodInstalled = err == nil
	}

	return status
}

func (a *App) ListarDispositivosCDROM() ([]DriveInfo, error) {
	var drives []DriveInfo

	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("lsblk", "-rno", "NAME,SIZE,TYPE")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return nil, fmt.Errorf("erro ao listar dispositivos: %s", err)
		}
		lines := strings.Split(strings.TrimSpace(string(output)), "\n")
		for _, line := range lines {
			fields := strings.Fields(line)
			if len(fields) >= 3 && fields[2] == "rom" {
				drives = append(drives, DriveInfo{
					Name: fields[0],
					Path: "/dev/" + fields[0],
					Size: fields[1],
				})
			}
		}

	case "windows":
		cmd := exec.Command("wmic", "cdrom", "get", "Drive,Size,VolumeName", "/format:csv")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return nil, fmt.Errorf("erro ao listar dispositivos: %s", err)
		}
		lines := strings.Split(strings.TrimSpace(string(output)), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "Node") {
				continue
			}
			fields := strings.Split(line, ",")
			if len(fields) >= 3 {
				driveLetter := strings.TrimSpace(fields[1])
				if driveLetter != "" {
					drives = append(drives, DriveInfo{
						Name: driveLetter,
						Path: driveLetter,
						Size: strings.TrimSpace(fields[2]),
					})
				}
			}
		}
	}

	return drives, nil
}

func (a *App) SelecionarDiretorio() (string, error) {
	dir, err := wailsRuntime.OpenDirectoryDialog(a.ctx, wailsRuntime.OpenDialogOptions{
		Title: "Selecionar diretório de saída",
	})
	if err != nil {
		return "", err
	}
	return dir, nil
}

func (a *App) ListarISOs(dirPath string) ([]string, error) {
	if dirPath == "" {
		dirPath = "."
	}

	matches, err := filepath.Glob(filepath.Join(dirPath, "*.iso"))
	if err != nil {
		return nil, fmt.Errorf("erro ao listar ISOs: %s", err)
	}

	var isoFiles []string
	for _, m := range matches {
		isoFiles = append(isoFiles, filepath.Base(m))
	}
	return isoFiles, nil
}

func (a *App) SelecionarArquivoISO() (string, error) {
	file, err := wailsRuntime.OpenFileDialog(a.ctx, wailsRuntime.OpenDialogOptions{
		Title: "Selecionar arquivo ISO",
		Filters: []wailsRuntime.FileFilter{
			{DisplayName: "Arquivos ISO", Pattern: "*.iso"},
		},
	})
	if err != nil {
		return "", err
	}
	return file, nil
}

func (a *App) emitProgress(evType string, msg string, percent float64, bytes int64, total int64) {
	wailsRuntime.EventsEmit(a.ctx, "progress", ProgressEvent{
		Type:    evType,
		Message: msg,
		Percent: percent,
		Bytes:   bytes,
		Total:   total,
	})
}

func (a *App) getDeviceSize(devicePath string) int64 {
	if runtime.GOOS == "linux" {
		cmd := exec.Command("blockdev", "--getsize64", devicePath)
		out, err := cmd.Output()
		if err == nil {
			size, parseErr := strconv.ParseInt(strings.TrimSpace(string(out)), 10, 64)
			if parseErr == nil {
				return size
			}
		}
	}
	return 0
}

func (a *App) CopiarISO(devicePath string, outputDir string, outputName string) (string, error) {
	if devicePath == "" || outputName == "" {
		return "", fmt.Errorf("dispositivo e nome de saída são obrigatórios")
	}

	if !strings.HasSuffix(outputName, ".iso") {
		outputName += ".iso"
	}

	outputPath := filepath.Join(outputDir, outputName)
	a.emitProgress("start", "Iniciando cópia da ISO...", 0, 0, 0)

	switch runtime.GOOS {
	case "linux":
		if !strings.HasPrefix(devicePath, "/dev/") {
			devicePath = "/dev/" + devicePath
		}

		totalSize := a.getDeviceSize(devicePath)

		cmd := exec.Command("dd", "if="+devicePath, "of="+outputPath, "bs=4M", "status=progress")

		stderr, err := cmd.StderrPipe()
		if err != nil {
			return "", fmt.Errorf("erro ao configurar dd: %s", err)
		}

		if err := cmd.Start(); err != nil {
			return "", fmt.Errorf("erro ao iniciar dd: %s", err)
		}

		re := regexp.MustCompile(`(\d+)\s+bytes`)
		scanner := bufio.NewScanner(stderr)
		scanner.Split(splitDdOutput)

		for scanner.Scan() {
			line := scanner.Text()
			matches := re.FindStringSubmatch(line)
			if len(matches) >= 2 {
				copiedBytes, _ := strconv.ParseInt(matches[1], 10, 64)
				pct := float64(0)
				if totalSize > 0 {
					pct = (float64(copiedBytes) / float64(totalSize)) * 100
					if pct > 100 {
						pct = 100
					}
				}
				a.emitProgress("progress", fmt.Sprintf("Copiados: %.1f MB", float64(copiedBytes)/(1024*1024)), pct, copiedBytes, totalSize)
			}
		}

		if err := cmd.Wait(); err != nil {
			a.emitProgress("error", fmt.Sprintf("Erro: %s", err), 0, 0, 0)
			return "", fmt.Errorf("erro ao copiar ISO: %s", err)
		}

	case "windows":
		src, err := os.Open(`\\.\` + devicePath)
		if err != nil {
			a.emitProgress("error", fmt.Sprintf("Erro ao abrir dispositivo: %s", err), 0, 0, 0)
			return "", fmt.Errorf("erro ao abrir dispositivo: %s", err)
		}
		defer src.Close()

		dst, err := os.Create(outputPath)
		if err != nil {
			a.emitProgress("error", fmt.Sprintf("Erro ao criar arquivo: %s", err), 0, 0, 0)
			return "", fmt.Errorf("erro ao criar arquivo: %s", err)
		}
		defer dst.Close()

		buf := make([]byte, 4*1024*1024)
		totalBytes := int64(0)
		for {
			n, readErr := src.Read(buf)
			if n > 0 {
				_, writeErr := dst.Write(buf[:n])
				if writeErr != nil {
					return "", fmt.Errorf("erro ao gravar: %s", writeErr)
				}
				totalBytes += int64(n)
				a.emitProgress("progress", fmt.Sprintf("Copiados: %.1f MB", float64(totalBytes)/(1024*1024)), 0, totalBytes, 0)
			}
			if readErr != nil {
				if readErr != io.EOF {
					a.emitProgress("error", fmt.Sprintf("Erro de leitura: %s", readErr), 0, 0, 0)
				}
				break
			}
		}
	}

	a.emitProgress("complete", "ISO copiada com sucesso!", 100, 0, 0)
	return outputPath, nil
}

func (a *App) ConverterISOparaGOD(isoPath string, outputDir string) (string, error) {
	if isoPath == "" {
		return "", fmt.Errorf("caminho do arquivo ISO é obrigatório")
	}

	if outputDir == "" {
		outputDir = filepath.Dir(isoPath)
	}

	a.emitProgress("start", "Iniciando conversão ISO → GOD...", 0, 0, 0)

	iso2godCmd := a.getIso2GodPath()

	cmd := exec.Command(iso2godCmd, isoPath, outputDir)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("erro ao configurar iso2god: %s", err)
	}
	cmd.Stderr = cmd.Stdout

	if err := cmd.Start(); err != nil {
		a.emitProgress("error", fmt.Sprintf("Erro ao iniciar iso2god: %s", err), 0, 0, 0)
		return "", fmt.Errorf("erro ao iniciar iso2god: %s", err)
	}

	re := regexp.MustCompile(`(\d+)%`)
	var lastOutput strings.Builder
	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		line := scanner.Text()
		lastOutput.WriteString(line + "\n")

		matches := re.FindStringSubmatch(line)
		if len(matches) >= 2 {
			pct, _ := strconv.ParseFloat(matches[1], 64)
			a.emitProgress("progress", line, pct, 0, 0)
		} else {
			a.emitProgress("log", line, -1, 0, 0)
		}
	}

	if err := cmd.Wait(); err != nil {
		a.emitProgress("error", fmt.Sprintf("Erro na conversão: %s", err), 0, 0, 0)
		return "", fmt.Errorf("erro ao converter: %s\n%s", err, lastOutput.String())
	}

	a.emitProgress("complete", "Conversão concluída com sucesso!", 100, 0, 0)
	return lastOutput.String(), nil
}

func splitDdOutput(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	for i := 0; i < len(data); i++ {
		if data[i] == '\r' || data[i] == '\n' {
			return i + 1, data[:i], nil
		}
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}
