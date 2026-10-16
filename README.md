

# Encurtador de Link
O projeto é um encurtador de link para praticar requisições HTTP, operações CRUD, manipulação de JSON, tratamento de erros e operações com banco de dados. O projeto, por fim, ainda não está finalizado.

## Tecnologias utilizadas
- Golang
- JSON

## Estrutura do projeto

```
Encurtador/
├── server/ # código principal
├── .gitignore
├── go.mod
├── go.sum

```
## Dentro de EncurtadorLink/
```
cmd/
├── main.go # entry point do projeto
database/
├── database.go # arquivo para conexão e funções do database
db/
├── local.db # arquivo gerado automaticamente por database.go
handler/
├── add_data.go # valida adiciona os dados "POST /links" ao banco de dados
├── read_data.go # lê o banco de dados com base em uma função de database.go
├── redirect_url # redireciona o usuário por "GET /{code}"
├── remove_data.go # removerá um código / url do banco de dados (a implementar)
├── update_data.go # atualiza uma url de um código (a implementar)
model/
├── model.go # o modelo utilizado pelos handlers e database
utils/
├── random_code.go # gera um código aleatório para a simplificaçã do link
```
## Pré requisitos
- [Golang](https://go.dev/doc/install)
- [Git](https://git-scm.com/install)
- [Postman](https://www.postman.com/downloads/)
## Como rodar

### 1. Clone o repositório
```bash
    git clone https://github.com/zikadaz/Encurtador.git
    cd Encurtador
```
### 2. Inicie o servidor
```bash
    cd cmd
    go run main.go
```
### 3. Abra o postman
- Faça uma requisição `POST` para `http://127.0.0.1:8080/links`
- Adicione o corpo no formato JSON seguindo essa regra:
    ```json
    {
        "url":"sua_url_aqui" //digite sua url
    }
    ```
- Abra outra requisição, dessa vez `GET` para a mesma URL, retornará algo parecido com
    ```json
    {
        "kgTY2": {
            "url":"sua_url", // a url da requisicao post
            "code":"kgTY2",
            "clicked": 0 // O tanto de vezes que alguem foi redirecionado por esse código
        }
    }
- Abra o seu navegador
    
    Na barra de pesquisa, coloque: 
        `http://127.0.0.1:8080/kgTY2`
        
    lembrando que esse código é ilustrativo.

    Você será automaticamente redirecionado para a URL que digitou anteriormente. 
    
    O contador será incrementado uma vez.

## Autor

Feito por [zikadaz](https://github.com/zikadaz) como um projeto de aprendizado! xD
