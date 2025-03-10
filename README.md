# Gerenciamento de Contatos 

Este projeto implementa uma API simples de gerenciamento de contatos utilizando a linguagem Go e o pacote `net/http`.

## Tecnologias Utilizadas
- Go
- `net/http`
- `encoding/json`

## Como Executar o Projeto

1. **Instale o Go**: Certifique-se de ter o Go instalado em sua máquina. Caso não tenha, baixe e instale pelo site oficial: [https://go.dev/](https://go.dev/).

2. **Clone o repositório**:
   ```sh
   git clone <URL_DO_REPOSITORIO>
   cd <NOME_DO_DIRETORIO>
   ```

3. **Execute a aplicação**:
   ```sh
   go run main.go
   ```

4. **A API será iniciada na porta 8080**.

## Endpoints Disponíveis

### Criar Contato
- **Rota:** `POST /contacts`
- **Descrição:** Cria um novo contato.
- **Requisição:**
  ```json
  {
    "name": "teste",
    "email": "teste@mail.com",
    "phone": "123456789"
  }
  ```
- **Resposta:**
  ```json
  {
    "id": 1,
    "name": "joao",
    "email": "joao@mail.com",
    "phone": "123456789"
  }
  ```

### Listar Contatos
- **Rota:** `GET /contacts`
- **Descrição:** Retorna todos os contatos.
- **Resposta:**
  ```json
  [
    {
      "id": 1,
      "name": "joao",
      "email": "joao@mail.com",
      "phone": "123456789"
    }
  ]
  ```

### Buscar Contato por ID
- **Rota:** `GET /contacts?id={id}`
- **Descrição:** Retorna um contato específico.
- **Exemplo de Resposta:**
  ```json
  {
    "id": 1,
    "name": "joao",
    "email": "joao@mail.com",
    "phone": "123456789"
  }
  ```
### Atualizar Contato
- **Rota:** `PUT /contacts?id={id}`
- **Descrição:** Atualiza um contato existente.
- **Requisição:**
  ```json
  {
    "name": "john",
    "email": "John@mail.com",
    "phone": "987654321"
  }
  ```
- **Resposta:** 200 OK (sem corpo de resposta)

### Excluir Contato
- **Rota:** `DELETE /contacts?id={id}`
- **Descrição:** Remove um contato do sistema.
- **Resposta:** 200 OK 

## Melhorias Futuras
- Persistência em banco de dados (exemplo: SQLite, PostgreSQL, MongoDB)
- Autenticação e autorização
- Implementação de testes unitários

## Licença
Este projeto está sob a licença MIT.
