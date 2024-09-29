#!/bin/bash

# Solicita ao usuário a quantidade de requisições
read -p "Quantas requisições você deseja enviar ao Kinesis? " num_requests

# Confirmação do usuário
read -p "Você quer realmente enviar $num_requests requisições? (s/n): " confirmation

# Se o usuário confirmar com 's' (sim), o script prossegue
if [[ "$confirmation" == "s" || "$confirmation" == "S" ]]; then
    echo "Enviando $num_requests requisições ao Kinesis..."

    # Função para enviar o registro ao Kinesis
    send_kinesis_record() {
        # Envia a requisição para o Kinesis e redireciona a saída para /dev/null
        aws kinesis put-record --stream-name poc --partition-key 1 --data "HelloKinesis" --profile sandbox > /dev/null 2>&1 &
    }

    # Loop para enviar os registros em paralelo
    for i in $(seq 1 $num_requests); do
        send_kinesis_record
    done

    echo "As requisições foram iniciadas em segundo plano."
else
    echo "Operação cancelada."
fi

