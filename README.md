
# Gateway de Pagamento

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)

Este é um projeto de estudo para implementação de um gateway de pagamento simples, com integração à [Cielo](https://www.cielo.com.br/). 

O objetivo deste projeto é demonstrar a integração básica com a Cielo, bem como o uso das bibliotecas Gin Web Framework e GORM para roteamento HTTP e manipulação de banco de dados SQLite.

## Pré-requisitos
- Go (versão X.X.X)
- Credenciais de acesso à API da Cielo (Merchant ID e Merchant Key)
    - Para conseguir as chaves de acesso, basta acessar a documentação da API da Cielo. [Documentação API](https://developercielo.github.io/manual/cielo-ecommerce) 

## Configuração
Clone este repositório para o seu ambiente local:

```bash
git clone https://github.com/Luan-max/payment-gateway.git
```
Acesse o diretório do projeto:
```bash
cd payment-gateway
```
Abra o arquivo .env e insira suas credenciais da Cielo:
```env
MERCHANT_KEY=
MERCHANT_ID=
CIELO_URL=
SECRET=
GIN_MODE=
```
## Instalação das dependências

Este projeto usa o gerenciador de pacotes Go Modules para lidar com as dependências. Execute o seguinte comando para instalar as dependências necessárias:

```bash
go mod download
```
## Execução

Após a configuração e a instalação das dependências, você pode iniciar o servidor local executando o seguinte comando:

```bash
go run main.go
```
O servidor será iniciado e estará ouvindo as requisições HTTP na porta 8080.

Uso
Você pode interagir com o gateway de pagamento por meio de solicitações HTTP. Aqui estão alguns endpoints disponíveis:

```bash
POST api/v1/transaction - Cria uma nova transação de pagamento.
```