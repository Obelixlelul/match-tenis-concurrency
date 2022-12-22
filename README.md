# Um Jogo de Tênis com Programação concorrente

Esse é um programa escrito na linguagem Go que simula um jogo de tênis com dois jogadores, cada um deles podendo ter dois
possíveis estados, a saber (i) esperando para receber a bola ou (ii) mandando a bola de volta para o adversário.
Para determinar se o jogador da vez perdeu a bola ou não, foram utilizadas facilidades providas pelo pacote math/rand.

Cada jogador possui 40% de chance de rebater a bola com sucesso. O primeiro a obter, pelo menos, 6 pontos e, simultâneamente, dois pontos de diferença do adversário, vence a partida.


# Como executar

Para executar o programa, basta possuir o Go instalado na sua máquina. Em seguida executar, na raiz do projeto, o seguinte comando:

```
go run main.go
```