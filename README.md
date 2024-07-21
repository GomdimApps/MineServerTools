Aqui está o texto corrigido:

---

# Compilação do projeto

## Instalação do make

- Para instalar o `make` no seu Debian, execute o seguinte comando caso não esteja instalado:

  ```bash
  sudo apt install build-essential -y
  ```
## Gerar versão
- Usar o git tag gera uma versão

  ```bash
  git tag v1.0.0
  ```
## Compilar o pacote DEBIAN
- Para compilar o projeto, use o comando:

  ```bash
  make installer
  ```

# Dependências necessárias

- A aplicação necessita de algumas ferramentas para funcionar corretamente, use este comando:

  ```bash
  sudo apt install tar wget rsync tmux pv zstd -y
  ```

# Instalação do pacote no Debian/Ubuntu

- Para instalar este pacote Deb, você pode usar o comando abaixo como no exemplo:

  ```bash
  sudo dpkg -i MineServerTools_1.1.1_all.deb
  ```

# Desinstalação do pacote

- Para desinstalar o pacote, use o comando:

  ```bash
  sudo apt purge mineservertools -y
  ```

# Configuração da aplicação

## Configuração de Backups

- Para que o comando de backup funcione corretamente, você precisa configurar os arquivos de configuração do servidor. Acesse o arquivo de configuração de acordo com o seu servidor:

  ```bash
  nano /etc/mineservertools/bedrock-server.conf
  ```

- Você deve colocar o diretório onde está rodando seu servidor, como no exemplo:

  ```bash
  server-dir="/opt/server/"
  ```

# Comandos

## Backups

### Fazer Backup

- Para realizar um backup do seu servidor, você pode utilizar o comando:

  ```bash
  backup-bedrock --backup
  ```

- Você pode visualizar os backups existentes usando o comando:

  ```bash
  backup-bedrock --view
  ```