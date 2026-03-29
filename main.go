package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// 1. REGRA DE NEGÓCIO (O que estava no analise.go)
// Separamos a lógica para que ela possa ser testada facilmente.
func avaliarStatusCredito(score int) string {
	if score > 700 {
		return "APROVADO_PREMIUM"
	} else if score >= 500 {
		return "APROVADO_PADRAO"
	}
	return "REPROVADO"
}

func main() {
	// 2. CONFIGURAÇÃO (O que aprendemos com o Docker)
	topic := "credit-results"
	partition := 0
	
	// Conectando ao container que você confirmou o nome: projetogo-kafka-1
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("Erro de conexão: ", err)
	}
	defer conn.Close()

	// 3. EXECUÇÃO DO PROCESSO
	usuarioID := "Aldo-Rocha-PL"
	scoreSimulado := 850 // Imagine que este dado veio de um banco de dados
	
	// Chamamos nossa regra de negócio
	resultado := avaliarStatusCredito(scoreSimulado)
	
	// Montamos o "envelope" (JSON) para o Kafka
	mensagem := fmt.Sprintf(`{"user_id": "%s", "score": %d, "status": "%s"}`, usuarioID, scoreSimulado, resultado)

	// 4. ENVIO PARA O KAFKA
	fmt.Printf("Analisando crédito para %s... Resultado: %s\n", usuarioID, resultado)
	
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(mensagem)},
	)

	if err != nil {
		log.Fatal("Falha ao enviar para o Kafka: ", err)
	}

	fmt.Println("✅ Sucesso! O evento de decisão foi postado no Kafka.")
}