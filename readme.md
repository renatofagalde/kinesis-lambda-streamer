
# Node.js Lambda Kinesis Test Project

Este projeto foi desenvolvido para testar o comportamento do AWS Kinesis com um número superior de eventos ao permitido, visando observar o comportamento das Lambdas rodando em paralelo.

## Estrutura do Projeto

```bash
.
├── carga.sh
├── golang
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── index.js
├── infra
│   └── cf-infra.yaml
├── kinesis02.zip
├── package-lock.json
└── package.json
```

## Descrição dos Arquivos

### `carga.sh`
Este script Bash é responsável por enviar múltiplas requisições para o AWS Kinesis. Ele solicita ao usuário o número de requisições e confirma se deve prosseguir com o envio dos eventos. As requisições são enviadas em paralelo usando o comando `aws kinesis put-record`.

### `golang/main.go`
Este arquivo contém uma aplicação Go que realiza uma operação semelhante ao `carga.sh`, mas usando o SDK da AWS para Go (`aws-sdk-go-v2`). A aplicação solicita ao usuário a quantidade de requisições, confirma o envio e dispara as requisições em paralelo para o Kinesis. Cada requisição utiliza uma goroutine para garantir o envio em paralelo.

### `index.js`
Este arquivo define uma Lambda em Node.js que processa eventos recebidos do Kinesis, decodificando os dados e salvando-os no S3. A Lambda utiliza o SDK da AWS para interagir com o S3, salvando o payload dos eventos em arquivos cujo nome é baseado no `eventID` do evento Kinesis.

### `infra/cf-infra.yaml`
Este arquivo contém a definição do CloudFormation para a infraestrutura necessária, como o stream do Kinesis e a Lambda, facilitando o provisionamento do ambiente AWS.

### `kinesis02.zip`
Um arquivo compactado contendo a Lambda que deve ser implementada para processar os eventos recebidos no Kinesis.

### `package.json` e `package-lock.json`
Contêm as dependências e configurações do projeto Node.js para rodar a Lambda.

### `go.mod` e `go.sum`
Definem as dependências do projeto Go, incluindo a versão do SDK AWS para Go.

## Como Executar

### Pré-requisitos

- AWS CLI configurado com as permissões necessárias
- Go e Node.js instalados
- AWS SDKs para Go e Node.js configurados
- CloudFormation para provisionamento da infraestrutura

### Passo a Passo

1. **Provisionar a Infraestrutura**: Utilize o arquivo `infra/cf-infra.yaml` para provisionar o stream do Kinesis e a Lambda no AWS.

2. **Enviar Eventos ao Kinesis**:
   - Para enviar eventos ao Kinesis usando o script Bash:  
     ```bash
     ./carga.sh
     ```
   - Para enviar eventos ao Kinesis usando o programa Go:
     ```bash
     go run golang/main.go
     ```

3. **Monitorar a Lambda**: Após o envio dos eventos ao Kinesis, a Lambda em Node.js (`index.js`) processará os eventos e os armazenará no S3. Monitore o log da Lambda e os arquivos no S3.

4. **Verificar o S3**: Confirme que os arquivos foram salvos corretamente no bucket S3.

## Contribuição

Sinta-se à vontade para abrir issues ou submeter pull requests para melhorias ou correções.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

