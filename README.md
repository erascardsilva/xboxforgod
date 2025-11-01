# XboxForGOD

Este é um utilitário Go para gerenciar arquivos de jogos do Xbox 360, permitindo copiar DVDs de jogos para imagens ISO e converter essas ISOs para o formato GOD (Games on Demand).

## Funcionalidades

-   **Copiar ISO de DVD:** Crie arquivos ISO a partir de seus DVDs de jogos do Xbox 360.
-   **Transformar ISO em GOD:** Converta arquivos ISO existentes para o formato GOD, ideal para uso em consoles Xbox 360 modificados (RGH/JTAG).

## Requisitos

-   **Sistema Operacional:** Linux (compatível com a maioria das distribuições Linux).
-   **Privilégios de Root:** O programa deve ser executado com `sudo` para acessar o dispositivo de DVD e instalar dependências.
-   **Comandos:**
    -   `dd`: Para copiar DVDs. Será instalado automaticamente se não estiver presente (requer `coreutils`).
    -   `iso2god`: Para converter ISOs para GOD. Este comando **não é instalado automaticamente** e deve ser instalado manualmente. Pesquise por "iso2god install linux" para obter instruções específicas para sua distribuição.

## Como Usar

1.  **Clone o Repositório:**
    ```bash
    git clone https://github.com/erasmo/xboxforgod.git
    cd xboxforgod
    ```

2.  **Compile o Programa (Opcional, mas recomendado para uso contínuo):**
    ```bash
    go build -o xboxforgod
    ```
    Após a compilação, você terá um executável chamado `xboxforgod` no diretório atual. Você pode copiar este executável para qualquer pasta em seu sistema (por exemplo, `/usr/local/bin` para torná-lo acessível globalmente).

3.  **Execute o Programa:**
    **Importante:** As ISOs e pastas GOD serão criadas no diretório de onde você executar o programa. Certifique-se de estar no diretório desejado antes de executar.

    ```bash
    sudo go run main.go
    ```
    Ou, se você compilou o programa:
    ```bash
    sudo ./xboxforgod
    ```
    Se você copiou o executável para uma pasta no seu PATH (como `/usr/local/bin`), pode executá-lo de qualquer lugar, mas lembre-se que os arquivos de saída serão criados no diretório atual de onde o comando foi invocado:
    ```bash
    sudo xboxforgod
    ```

4.  **Navegue pelo Menu:**
    Siga as opções do menu para copiar ISOs ou converter para GOD.

    ```
    --- Menu Principal XboxForGOD ---
    1. Copiar ISO de DVD
    2. Transformar ISO em GOD
    3. Sair
    Escolha uma opção:
    ```

    **Importante para a conversão de ISO para GOD:** O arquivo ISO que você deseja converter deve estar no mesmo diretório de onde você está executando o programa `xboxforgod` ou você deve fornecer o caminho completo para o arquivo ISO.

## Por que o formato GOD no Xbox 360?

O formato GOD (Games on Demand) é um formato de arquivo específico para jogos do Xbox 360, projetado para ser executado diretamente do disco rígido do console. Ele é amplamente utilizado em consoles Xbox 360 modificados (com RGH/JTAG) por várias razões:

-   **Conveniência:** Permite que os jogos sejam armazenados e executados diretamente do HD interno ou externo, eliminando a necessidade de ter o disco físico inserido.
-   **Velocidade de Carregamento:** Jogos carregam mais rapidamente do HD do que de um DVD.
-   **Preservação do Leitor de DVD:** Reduz o desgaste do leitor de DVD do console.
-   **Organização:** Facilita a organização de uma grande biblioteca de jogos digitais.

## Estrutura do Projeto

-   `main.go`: Ponto de entrada principal do programa, gerencia o menu e a lógica de execução.
-   `utils/dependencies.go`: Contém a lógica para verificar e instalar dependências (`dd`, `iso2god`).
-   `features/iso_copy/iso_copy.go`: Implementa a funcionalidade de copiar um DVD para um arquivo ISO.
-   `features/iso_convert/iso_convert.go`: Implementa a funcionalidade de converter um arquivo ISO para o formato GOD.

## Contribuição

Sinta-se à vontade para contribuir com melhorias, correções de bugs ou novas funcionalidades.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
