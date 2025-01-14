# Task API

Task API é uma aplicação simples para gerenciar tarefas, construída com Go. Esta API permite criar, ler, atualizar e deletar tarefas.

## Funcionalidades

- **Listar Tarefas**: Obtenha uma lista de todas as tarefas.
- **Obter Tarefa por ID**: Obtenha os detalhes de uma tarefa específica pelo seu ID.
- **Criar Tarefa**: Adicione uma nova tarefa.
- **Atualizar Tarefa**: Atualize uma tarefa existente.
- **Remover Tarefa**: Delete uma tarefa existente.

## Endpoints

- `GET /tasks`: Lista todas as tarefas.
- `GET /tasks/{id}`: Obtém uma tarefa específica pelo ID.
- `POST /tasks`: Cria uma nova tarefa.
- `PUT /tasks/{id}`: Atualiza uma tarefa existente.
- `DELETE /tasks/{id}`: Remove uma tarefa existente.

## Instalação

1. Clone o repositório:
    ```sh
    git clone https://github.com/seu-usuario/task-api.git
    ```
2. Navegue até o diretório do projeto:
    ```sh
    cd task-api
    ```
3. Instale as dependências:
    ```sh
    go mod tidy
    ```

## Uso

Para iniciar o servidor, execute:
```sh
go run cmd/api/main.go
```

O servidor estará disponível em http://localhost:8080.

## Estrutura do Projeto

- **main.go**: Ponto de entrada da aplicação.
- **controller**: Contém os handlers para os endpoints.
- **task**: Contém a lógica de negócio para gerenciar tarefas.

## Contribuição
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.

## Licença
Este projeto está licenciado sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.