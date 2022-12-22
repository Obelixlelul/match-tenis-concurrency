package main

/*
Por questões de simplicidade, escolhemos admitir que o match possui um único set composto de um único game,
ganhando o jogo o jogador que alcançar um número de pontos fixo 6 desse game.

É o vencedor quem ganhar 3 de 5 sets.
game  - 4 pontos no total, 2 a mais que o adversário
set   - 6 games, 2 a mais que o adversário
match - É uma partidade como um todo que possui 3 a 5  sets
*/

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg é usado para esperar que todas as goroutines lançadas aqui terminem.
var wg sync.WaitGroup

// função invocada antes da execução do main
func init() {
	rand.Seed(time.Now().UnixNano()) // Definindo a seed do rand
}

// definindo os scores iniciais antes da partida
var p1Score int = 0
var p2Score int = 0

func main() {
	// Criando o channel que representará a partida
	partida := make(chan int)

	// Adicionando um as duas goroutines representando os jogadores
	wg.Add(2)

	go player("Rafael", partida, &p1Score)
	go player("João", partida, &p2Score)

	// Começando a partida
	partida <- 1

	// espera o jogo terminar
	wg.Wait()
}

func player(name string, partida chan int, pscore *int) {
	defer wg.Done()

	// Loop infinito simulando a partida até que algum jogador vença
	for {
		rodada, ok := <-partida

		if !ok {
			return
		}

		// Gera um numero aleatório entre 0 e 10
		n := rand.Intn(10)

		// Se maior que 6, acertou, senão, errou (40% de chance de acerto)
		if n >= 6 {
			*pscore++ // Incrementa o score do jogador que está na vez
			fmt.Printf("Jogador %s Acertou! - Rodada: %d\n", name, rodada)
			rodada++
		} else {
			fmt.Printf("Jogador %s Errou! - Rodada: %d\n", name, rodada)
			rodada++
		}

		// Imprime o placar
		fmt.Printf("Placar: [%d | %d]\n", p2Score, p1Score)

		/*
			Verifica se o jogador fez 6 pontos com a diferença de 2 pontos do adversário
			se sim, encerra a partica mudando o status do ok para false
		*/
		if (p1Score >= 6 && p1Score-p2Score >= 2) || (p2Score >= 6 && p2Score-p1Score >= 2) {
			fmt.Printf("Jogador %s ganhou!!!!\n", name)
			close(partida)
			return
		}

		partida <- rodada
	}

}
