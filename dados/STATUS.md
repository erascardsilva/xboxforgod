# Status Atual do Projeto e Passos Seguintes

## O que já está pronto e implementado
1. **Interface (Wails):** Telas ajustadas, suporte multi-idioma (Português/Inglês) configurado e botão de doação integrado com a API nativa de abertura de links.
2. **Dependências (Go):** Binários (`iso2god`) embutidos no código fonte para não depender de instalação externa, suporte nativo ao Linux (`dd`) e Windows.
3. **Empacotamento (Snap):** `snapcraft.yaml` reestruturado para _strict confinement_, contendo plugs corretos (`optical-drive`, `removable-media`), arquivo `.desktop` gerado e ícone visual customizado alocado na pasta correta (`snap/gui`).
4. **CI/CD:** Pipeline do GitHub Actions (`.github/workflows/snap-publish.yml`) criado para compilação remota Ubuntu.

## Suas tarefas para continuar (Checklist de Deploy)
1. **Registrar o Nome:** Rodar `snapcraft register xboxforgod` no seu terminal local.
2. **Gerar Token:** Rodar o comando para exportar as credenciais localmente:
   ```bash
   snapcraft export-login --snaps=xboxforgod --acls=package_access,package_manage,package_push snapcraft.login
   ```
3. **Configurar GitHub:** Copiar o texto do arquivo `snapcraft.login` gerado e inseri-lo no repositório GitHub como um Secret em _Settings > Secrets and variables > Actions > New repository secret_ com o nome `SNAPCRAFT_STORE_CREDENTIALS`.
4. **Acionar Publicação:** Fazer o _commit_ de todas as alterações e realizar o _push_ para a branch `main`. Isso fará o GitHub Actions inicializar o pacote e publicá-lo na Snap Store.
