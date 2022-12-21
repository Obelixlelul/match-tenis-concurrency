package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	waiting  bool
	throwing bool
}

type Game struct {
	game  int
	set   int
	match int
}

// rand de forma alearória para determinar se jogador perdeu a bola ou não
/*
Por questões de simplicidade, admita que o match possui um único set composto de um único game,
ganhando o jogo o jogador que alcançar um número de pontos fixo P desse game.
*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

/*
	GAME: pelo menos 4 pontos no total e 2 a mais que o adversáio 15 30 40 win the game
	SET: Conjunto de games, 6 games e dois a mais que o oponente
	MATCH: Uma partida como um todo;
	-> É o vencedor quem ganhar 3 de 5 sets.
*/
