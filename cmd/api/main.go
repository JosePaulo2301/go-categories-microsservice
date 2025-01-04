package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Tarefa concluída")
	case <-ctx.Done():
		fmt.Println("Operação cancelada:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// Inicia uma goroutine
	go doSomething(ctx)

	// Cancela o contexto após 1 segundo
	time.Sleep(1 * time.Second)
	cancel()

	// Aguarda a goroutine terminar
	time.Sleep(2 * time.Second)
}
