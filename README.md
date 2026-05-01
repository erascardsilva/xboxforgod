# XboxForGOD

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/Wails-ED2945?style=for-the-badge&logo=wails&logoColor=white" alt="Wails" />
  <img src="https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black" alt="JavaScript" />
  <img src="https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white" alt="HTML5" />
  <img src="https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white" alt="CSS3" />
  <img src="https://img.shields.io/badge/Xbox_360-107C10?style=for-the-badge&logo=xbox&logoColor=white" alt="Xbox 360" />
</p>

<p align="center">
  <a href="https://snapcraft.io/xboxforgod">
    <img alt="Disponível na Snap Store" src="https://snapcraft.io/pt/dark/install.svg" />
  </a>
</p>

XboxForGOD is a modern, cross-platform desktop application built with Wails that simplifies managing Xbox 360 game files. It allows users to effortlessly copy game DVDs into ISO images and convert those ISOs into the GOD (Games on Demand) format, ready to be played on RGH/JTAG modified consoles.

## ✨ Features

- **DVD to ISO Extraction:** Directly create an ISO image from your Xbox 360 game DVD.
- **ISO to GOD Conversion:** Convert existing ISO files into the GOD format for seamless execution from the console's hard drive.
- **Bilingual Interface:** Fully supports English and Portuguese (PT-BR).
- **Embedded Dependencies:** The `iso2god` binaries are bundled directly within the application, eliminating the need for manual installations.

## 🏗 Architecture

XboxForGOD follows a modern desktop application architecture leveraging the **Wails v2** framework, combining the performance of Go with the flexibility of web technologies.

```mermaid
graph TD
    subgraph Frontend [Vanilla JS / HTML / CSS]
        UI[Graphical User Interface]
        I18N[I18N Module]
        Events[Wails Event Listeners]
    end

    subgraph Backend [Go Application]
        App[App Struct]
        Embed[Embedded Binaries FS]
        Syscall[System Calls: dd / OS Read]
    end
    
    subgraph Host OS [Operating System]
        FS[File System]
        DVD[DVD Optical Drive]
    end

    UI <-->|Wails IPC Bindings| App
    Events <--|Progress Events| App
    App --> Embed
    App --> Syscall
    Syscall --> FS
    Syscall --> DVD
```

### Components
1. **Frontend:** A lightweight Vanilla JavaScript interface (`index.html`, `main.js`, `style.css`) providing a responsive and dynamic user experience without the overhead of heavy web frameworks.
2. **Backend:** A Go service that interacts natively with the OS. It lists optical drives, monitors extraction progress, and manages external processes.
3. **Dependency Manager:** The `iso2god` executables (for both Linux and Windows) are securely embedded in the Go binary using `//go:embed`. During runtime, they are extracted to a temporary folder and executed automatically. 

## 🚀 How it Works

1. **Insert DVD:** The user inserts the Xbox 360 game disc. The Go backend detects available optical drives using `lsblk` (Linux) or `wmic` (Windows).
2. **Copy ISO:** The application extracts the disc content. On Linux, it wraps the native `dd` command. On Windows, it reads directly from the device block `\\.\<DriveLetter>:`.
3. **Convert to GOD:** The application extracts the embedded `iso2god` utility to a temporary location and executes it against the selected ISO, piping the progress output back to the Wails frontend in real-time.

---

### ☕ Support the Project

If you find this tool useful and would like to support its development, you can make a donation via PayPal:

[**Donate via PayPal**](https://www.paypal.com/ncp/payment/8V6WQCGN6HDCQ)

---

### Developed by
**Erasmo Cardoso**
*Software Engineer | Electronics Technician*

---

### 💻 Compatible Systems

<p>
  <img src="https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black" alt="Linux" />
  <img src="https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white" alt="Windows" />
</p>

- **Linux (amd64):** Fully compatible. Requires `dd` (coreutils) which is natively present in almost all distributions.
- **Windows (amd64):** Fully compatible. Does not require external installations.

### 📂 Instalação e Downloads

#### Ubuntu / Linux (via Snap Store)
A forma mais fácil e recomendada de instalar no Linux é diretamente pela Snap Store. O aplicativo é isolado, auto-atualizável e já contém todas as dependências de sistema:

[![Disponível na Snap Store](https://snapcraft.io/pt/dark/install.svg)](https://snapcraft.io/xboxforgod)

*(Ou via terminal executando: `sudo snap install xboxforgod`)*

#### Windows (Instalador Nativo e Standalone)
Para usuários de Windows, os arquivos para instalação são gerados automaticamente durante o processo de build do Wails (Cross-Compiling via NSIS) e ficam disponíveis localmente no seguinte diretório:

```text
build/bin/
```

Dentro desta pasta você encontrará o **Instalador para Windows** (ex: `xboxforgod-amd64-installer.exe`), que instalará a aplicação e criará atalhos no Menu Iniciar/Desktop, além do binário standalone pronto para uso. O pacote já contém a interface, a lógica em Go e as dependências embarcadas num único arquivo.
