# MineServerTools

O **MineServerTools** é um conjunto de ferramentas essenciais para manter um servidor Minecraft sempre ativo em um servidor Linux. Ele oferece recursos que facilitam a administração e a otimização do seu servidor. Aqui estão alguns dos benefícios e funcionalidades principais:

1. **Monitoramento e Reinicialização Automática**: O **MineServerTools** verifica regularmente o status do servidor Minecraft e reinicia automaticamente o servidor caso ele pare de responder. Isso ajuda a manter o servidor online e minimiza o tempo de inatividade.

2. **Backup Automático**: O projeto inclui um sistema de backup automático que cria cópias de segurança regulares dos arquivos do servidor. Isso é crucial para proteger seus dados e evitar perda de progresso.

3. **Agendamento de Tarefas**: Com o **MineServerTools**, você pode agendar tarefas específicas, como reinicializações programadas, backups ou outras ações. Isso permite uma gestão mais eficiente do servidor.


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
  make package-deb
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
# Comandos via Interface

- Para acessar a interface da ferramenta, você pode utilizar o comando:

  ```bash
  mine-server-tools
  ```

# Comandos via terminal

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

## Informações do Sistema

- Para verificar as informações importantes do servidor que está hospedado o seu server minecraft como: Memoria RAM, Uso de CPU e etc, use o comando:

  ```bash
  info-bedrock
  ```

# Agradecimentos

Gostaria de expressar minha gratidão aos colaboradores que tornaram o **MineServerTools** possível. Seus esforços e contribuições são inestimáveis para a comunidade Minecraft e para todos os administradores de servidores.

## Colaboradores

- [Isac Gondim (GomdimApps)](https://github.com/GomdimApps/)
- [João Pedro Maciel (JoaoPedr0Maciel)](https://github.com/JoaoPedr0Maciel/)
- [Pedro Felipe (PedroFelipeCS)](https://github.com/PedroFelipeCS)

## Como Contribuir

Se você também deseja ajudar a melhorar o **MineServerTools**, considere as seguintes maneiras de contribuir:

[![Doe](https://img.shields.io/badge/Doe-Agora-brightgreen)](https://pag.ae/7-LKKsoXa)

1. **Envie Problemas (Issues)**: Se encontrar algum problema ou tiver uma ideia para aprimorar o projeto, abra um problema no repositório oficial.

2. **Envie Solicitações de Pull (Pull Requests)**: Se você tem correções ou melhorias específicas, envie uma solicitação de pull. Sua contribuição será revisada e considerada para inclusão no projeto.

3. **Compartilhe com Outros**: Espalhe a palavra sobre o **MineServerTools** para outros administradores de servidores Minecraft. Quanto mais pessoas usarem e contribuírem, melhor o projeto se tornará!
