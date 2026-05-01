Coloque os binários das dependências nesta pasta:
- Linux: `iso2god` (binário nativo ou script configurado)
- Windows: `iso2god.exe`

Ao compilar a aplicação, estes binários serão "embutidos" (embedded) diretamente no executável final do Wails, garantindo que o programa tenha "dentro de si" tudo que é necessário para rodar sem pedir ao usuário para instalar nada.

Nota: `dd` já vem instalado por padrão no Linux, e o Windows usa APIs nativas de leitura do Go, então não precisam ser colocados aqui.
