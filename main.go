package main

import (
	"fmt"
	"sync"
	"time"
)

func funcionario(produto string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Tempo que cada funcionário leva para embalar um produto
	time.Sleep(3 * time.Second)
	ch <- produto
}

func main() {
	produtos := []string{
		"Mesa",
		"Armário",
		"Cafeteira",
		"Furadeira",
		"Panela Elétrica",
	}

	var wg sync.WaitGroup
	ch := make(chan string, len(produtos))

	start := time.Now()
	for _, produto := range produtos {
		wg.Add(1)
		go funcionario(produto, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	produtosProntos := []string{}
	for produto := range ch {
		produtosProntos = append(produtosProntos, produto)
	}

	fmt.Println("Caminhão enviado com os produtos: ", produtosProntos)
	fmt.Printf("Tempo gasto para embalar todos os produtos: %v\n", time.Since(start))
}
