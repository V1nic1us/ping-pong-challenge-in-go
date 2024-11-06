package main

import (
	"fmt"
	"time"
)

func ping(pings chan<- string, pongs <-chan string) {
	for {
		// Recebe uma mensagem do canal "pongs" (espera pela resposta de "pong")
		msg := <-pongs
		fmt.Println("Ping received:", msg)

		// Envia a mensagem "ping" para o canal "pings"
		time.Sleep(1 * time.Second)
		pings <- "ping"
	}
}

func pong(pings <-chan string, pongs chan<- string) {
	for {
		// Recebe uma mensagem do canal "pings" (espera pela mensagem de "ping")
		msg := <-pings
		fmt.Println("Pong received:", msg)

		// Envia a resposta "pong" para o canal "pongs"
		time.Sleep(1 * time.Second)
		pongs <- "pong"
	}
}

func main() {
	// Cria dois canais para comunicação entre "ping" e "pong"
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// Inicializa os canais com uma mensagem para começar o ciclo
	pongs <- "start"

	// Inicia as funções "ping" e "pong" como goroutines
	go ping(pings, pongs)
	go pong(pings, pongs)

	// Mantém o programa em execução
	select {}
}
