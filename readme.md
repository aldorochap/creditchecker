# Credit Checker - Microserviço de Crédito Assíncrono

Este projeto simula um sistema de análise de crédito de alta performance, inspirado na arquitetura de sistemas distribuídos do **Aldo Créditos**. O foco é demonstrar o uso de mensageria para garantir escalabilidade e resiliência.

## Tecnologias Utilizadas
*   **Go (Golang):** Escolhido pela baixa latência e alta performance em concorrência.
*   **Apache Kafka:** Utilizado como motor de mensageria para desacoplamento de serviços.
*   **Docker & Docker Compose:** Para orquestração da infraestrutura (Kafka/Zookeeper).

## Arquitetura do Projeto
O fluxo segue o padrão de **Sistemas Orientados a Eventos (EDA)**:
1. O Microserviço recebe uma solicitação de crédito.
2. Aplica uma regra de negócio (Score Analysis).
3. Produz um evento JSON no tópico `credit-results` do Kafka para que outros serviços (como notificações ou antifraude) possam consumir o dado de forma assíncrona.

## Como Executar o Projeto

### Pré-requisitos
*   Docker Desktop instalado.
*   Go (Golang) instalado.

### 1. Subir a Infraestrutura
Na raiz do projeto, execute o comando para subir o Kafka e o Zookeeper:
```powershell
docker-compose up -d

### 2. Configurar o Projeto Go
Instale as dependências necessárias:

PowerShell
go mod tidy

### 3. Executar o Microserviço
PowerShell
go run main.go

📊 O que este projeto demonstra?
Resiliência: O sistema continua operando mesmo com picos de carga graças ao buffer do Kafka.

Separação de Responsabilidades: A lógica de negócio está isolada da infraestrutura de comunicação.

Prontas para Produção: Uso de healthchecks e versionamento de imagens Docker.