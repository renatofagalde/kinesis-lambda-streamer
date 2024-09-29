package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

const streamName = "poc"
const partitionKey = "1"
const data = "HelloKinesis"

// Função para enviar o registro ao Kinesis usando o SDK
func sendKinesisRecord(client *kinesis.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	// Cria o registro
	input := &kinesis.PutRecordInput{
		StreamName:   aws.String(streamName),
		PartitionKey: aws.String(partitionKey),
		Data:         []byte(data),
	}

	// Envia o registro para o Kinesis
	_, err := client.PutRecord(context.TODO(), input)
	if err != nil {
		fmt.Printf("Erro ao enviar o registro: %v\n", err)
	} else {
		fmt.Println("Registro enviado com sucesso")
	}
}

func main() {
	// Carrega a configuração padrão da AWS
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("sandbox"))
	if err != nil {
		fmt.Printf("Erro ao carregar configuração da AWS: %v\n", err)
		return
	}

	// Cria o cliente do Kinesis
	client := kinesis.NewFromConfig(cfg)

	reader := bufio.NewReader(os.Stdin)

	// Solicita ao usuário a quantidade de requisições
	fmt.Print("Quantas requisições você deseja enviar ao Kinesis? ")
	numRequestsStr, _ := reader.ReadString('\n')
	numRequestsStr = strings.TrimSpace(numRequestsStr)

	var numRequests int
	fmt.Sscanf(numRequestsStr, "%d", &numRequests)

	// Confirmação do usuário
	fmt.Printf("Você quer realmente enviar %d requisições? (s/n): ", numRequests)
	confirmation, _ := reader.ReadString('\n')
	confirmation = strings.TrimSpace(confirmation)

	// Se o usuário confirmar com 's' (sim), o script prossegue
	if confirmation == "s" || confirmation == "S" {
		fmt.Printf("Enviando %d requisições ao Kinesis...\n", numRequests)

		var wg sync.WaitGroup

		// Loop para enviar os registros em paralelo
		for i := 0; i < numRequests; i++ {
			wg.Add(1)
			go sendKinesisRecord(client, &wg)
		}

		// Espera todas as requisições terminarem
		wg.Wait()

		fmt.Println("As requisições foram enviadas com sucesso.")
	} else {
		fmt.Println("Operação cancelada.")
	}
}

