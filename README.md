# Gerador de Senha Aleatória

Este projeto implementa um gerador de senhas aleatórias utilizando a linguagem Go. A senha gerada inclui caracteres minúsculos, maiúsculos, números e símbolos especiais, com a possibilidade de personalizar o comprimento da senha entre 8 e 32 caracteres.

## Tecnologias Utilizadas
- Go

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

4. **A aplicação pedirá para inserir o número de caracteres desejado para a senha. O número deve ser entre 8 e 32 caracteres.**

5. **A senha gerada será exibida no terminal.**

## Como Funciona

O programa solicita ao usuário a quantidade de caracteres para a senha, validando se o valor está entre 8 e 32. Em seguida, gera uma senha com os seguintes critérios:
- Pelo menos 1 caractere minúsculo.
- Pelo menos 1 caractere maiúsculo.
- Pelo menos 1 caractere numérico.
- Pelo menos 1 caractere especial (como `!@#$%^&*`).
- O restante dos caracteres são aleatórios, escolhidos de uma lista de caracteres válidos.

A senha gerada é embaralhada para garantir que os caracteres obrigatórios não fiquem em posições fixas.

## Exemplo de Uso

1. O programa solicitará: **"Quantos caracteres você deseja para a senha?"**
2. O usuário insere um número, por exemplo: **12**.
3. A senha gerada será exibida, por exemplo:
   ```
   Sua senha gerada é:
   X1!vG7*bQz
   ```

## Melhorias Futuras
- Adicionar a opção de personalizar os tipos de caracteres incluídos na senha.
- Implementar uma interface gráfica para facilitar o uso.
- Adicionar a opção de salvar senhas geradas.

## Licença
Este projeto está sob a licença MIT.

---

Esse `README.md` agora reflete o funcionamento do gerador de senhas em Go, conforme o código fornecido. Se precisar de mais alguma alteração ou detalhes, é só avisar!
