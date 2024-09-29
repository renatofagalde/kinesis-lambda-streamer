const AWS = require('aws-sdk');
const s3 = new AWS.S3();

const BUCKET_NAME = 'poc44'; // Nome do bucket S3

// Função para adicionar uma espera de 2 segundos
const wait = (ms) => new Promise(resolve => setTimeout(resolve, ms));

exports.handler = async (event) => {
    const promises = event.Records.map(async (record) => {
        // Espera por 2 segundos
//        await wait(1000);

        // Decodifica o payload do Kinesis
        const payload = Buffer.from(record.kinesis.data, 'base64').toString('ascii');
        console.log('Decoded payload:', payload);

        // Configura o arquivo para ser salvo no S3
        const s3Params = {
            Bucket: BUCKET_NAME,
            Key: `${record.eventID}.txt`, // Nome do arquivo baseado no ID do evento
            Body: payload // Conteúdo do arquivo será o payload
        };

        // Salva o arquivo no S3
        try {
            await s3.putObject(s3Params).promise();
            console.log(`Arquivo salvo no S3: ${record.eventID}.txt`);
        } catch (error) {
            console.error(`Erro ao salvar no S3: ${error.message}`);
        }
    });

    // Espera todas as operações de gravação terminarem
    await Promise.all(promises);

    return `Successfully processed ${event.Records.length} records.`;
};

