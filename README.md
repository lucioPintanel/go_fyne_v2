# CRUD Fyne V2

[![SQLite](https://img.shields.io/badge/SQLite-%2307405e.svg?logo=sqlite&logoColor=white)](#)
[![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?&logo=go&logoColor=white)](#)
[![Git](https://img.shields.io/badge/Git-F05032?logo=git&logoColor=fff)](#) 
![GitHub Org's stars](https://img.shields.io/github/stars/camilafernanda?style=social)

![Badge em Desenvolvimento](http://img.shields.io/static/v1?label=STATUS&message=EM%20DESENVOLVIMENTO&color=GREEN&style=for-the-badge)

## Visão Geral

CRUD Fyne V2 é um projeto que demonstra a implementação de operações CRUD (Criar, Ler, Atualizar e Excluir) utilizando o framework [Fyne](https://fyne.io/) para [Go](https://go.dev/). Este projeto é ideal para quem busca criar interfaces gráficas de usuário (GUIs) elegantes e eficientes em Go.

## Funcionalidades

- **CRUD Completo**: Implementação completa de operações CRUD.
- **UI Moderna**: Utilização do framework Fyne para uma interface moderna e responsiva.
- **Arquitetura Limpa**: Estrutura modular e organizada para fácil manutenção e escalabilidade.

## Instalação

Para clonar e executar este projeto, siga estas etapas:

```bash
git clone https://dev.azure.com/luciopintanel/_git/crud_fyne_v2
cd crud_fyne_v2
go mod tidy
go run internal/cmd/main.go
```

## Como Contribuir
Para contribuir, por favor, siga os passos abaixo:
1. Faça um fork do repositório
2. Crie uma branch para suas modificações (`git checkout -b minha-branch`)
3. Faça o commit das suas mudanças (`git commit -am 'Adicionei uma nova funcionalidade'`)
4. Envie suas mudanças para a branch original (`git push origin minha-branch`)
5. Crie uma Pull Request

## Issues
Aqui estão algumas issues abertas no projeto:

### [#3 Teste Unitário para Funções do 'Repository'](https://github.com/lucioPintanel/go_fyne_v2/issues/3#issue-2616197499)
- **Descrição**: Criação de testes unitários para as funções do repositório.
- **Objetivos**:
1. Testar função de criação de registros.
2. Testar função de leitura de registros.
3. Testar função de atualização de registros.
4. Testar função de exclusão de registros.

### [#4 Teste Unitário para Funções do 'Services'](https://github.com/lucioPintanel/go_fyne_v2/issues/4#issue-2616198373)
- **Descrição**: Criação de testes unitários para as funções do serviço.
- **Objetivos**:
1. Testar lógica de negócios.
2. Validar comportamento de serviços.
