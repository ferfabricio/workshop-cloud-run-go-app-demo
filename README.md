# Colocando meus containers em produção

Este repositório é um exemplo utilizado no live coding do workshop realizado no Grupo Integrado no dia 14/02/2025.

## Comandos golang

Para rodar a [aplicação na sua máquina você precisa do golang instalado](https://go.dev/doc/install).

- Rodar aplicação: `go run cmd/main.go`
- Passar o APP_NAME como variável: `APP_NAME=<qualquer valor> go run cmd/main.go`
- Criar binário: `go build -o app cmd/main.go`
- Baixar dependências: `go mod download`
- Atualizar dependências: `go mod tidy`

## Comandos docker

Para rodar os comandos do docker, [você precisa do docker instalado na sua máquina](https://docs.docker.com/desktop/).

- Fazer build da imagem: `docker build -t nome-do-container .`
- Executar uma instância da imagem: `docker run -p 8080:8080 nome-do-container`

## Slides

[Aqui estão os slides em PDF.](./slides.pdf)
