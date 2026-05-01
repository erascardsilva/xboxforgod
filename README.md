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

O XboxForGOD é uma ferramenta desktop multiplataforma desenvolvida para simplificar o gerenciamento de arquivos de jogos do Xbox 360. A aplicação permite realizar a cópia de DVDs de jogos para imagens ISO e a conversão dessas imagens para o formato GOD (Games on Demand), prontas para uso em consoles com modificação RGH/JTAG.

## Funcionalidades

- **Extração de DVD para ISO:** Criação direta de imagem ISO a partir do disco original.
- **Conversão ISO para GOD:** Processamento de arquivos ISO existentes para o formato Games on Demand.
- **Interface Multi-idioma:** Suporte completo para Português (PT-BR) e Inglês.
- **Dependências Embutidas:** Os binários do `iso2god` estão integrados ao executável, dispensando instalações manuais.

## Arquitetura

O projeto utiliza o framework **Wails v2**, unindo o desempenho do Go no backend com a flexibilidade de tecnologias web no frontend.

```mermaid
graph TD
    subgraph Frontend [JS / HTML / CSS]
        UI[Interface Gráfica]
        I18N[Módulo I18N]
        Events[Wails Events]
    end

    subgraph Backend [Go]
        App[Estrutura App]
        Embed[Binários Embarcados]
        Syscall[Chamadas de Sistema: dd / OS Read]
    end
    
    subgraph Host OS [Sistema Operacional]
        FS[Sistema de Arquivos]
        DVD[Unidade Óptica]
    end

    UI <-->|Wails IPC| App
    Events <--|Progresso| App
    App --> Embed
    App --> Syscall
    Syscall --> FS
    Syscall --> DVD
```

## Funcionamento

1. **Detecção:** O backend em Go identifica as unidades ópticas disponíveis através de comandos nativos do sistema.
2. **Cópia:** A extração é realizada via `dd` (Linux) ou leitura direta do bloco de dispositivo (Windows).
3. **Conversão:** O utilitário `iso2god` é extraído para um local temporário e executado, com o progresso sendo enviado em tempo real para a interface.

---

### Apoio ao Projeto

Se esta ferramenta for útil para você e desejar apoiar o desenvolvimento contínuo, considere realizar uma doação via PayPal:

[**Doar via PayPal**](https://www.paypal.com/ncp/payment/8V6WQCGN6HDCQ)

---

### Desenvolvido por
**Erasmo Cardoso**
*Software Engineer | Electronics Technician*

---

### Sistemas Compatíveis

<p>
  <img src="https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black" alt="Linux" />
  <img src="https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white" alt="Windows" />
</p>

- **Linux (amd64):** Requer `dd` (disponível na maioria das distribuições).
- **Windows (amd64):** Totalmente independente, sem requisitos externos.

### Instalação e Downloads

#### Linux (via Snap Store)
A forma recomendada de instalação no Linux é através da Snap Store. O pacote é isolado e gerencia todas as atualizações automaticamente:

[![Disponível na Snap Store](https://snapcraft.io/pt/dark/install.svg)](https://snapcraft.io/xboxforgod)

*(Ou via terminal: `sudo snap install xboxforgod`)*

#### Windows (Instalador e Executável)
Os arquivos para Windows são gerados na pasta de build após a compilação:

```text
build/bin/
```

Neste diretório estão disponíveis o **Instalador** (`xboxforgod-amd64-installer.exe`) e o executável standalone.
