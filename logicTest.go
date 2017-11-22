
package main

import(
//	"math/rand"
//	"time"
	"fmt"
	"./gamelogic"
	"bufio"
  	"os"

)

func main(){
	//reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Welcome to Five Card Draw! (press 'enter' between messages to continue)")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	Red := new(gamelogic.Player)
	Red.Name = "Red"
	Red.Money = 100
	Red.Folded = false
	Red.Bet = 0


	Green := new(gamelogic.Player)
	Green.Name = "Green"
	Green.Money = 100
	Green.Folded = false
	Green.Bet = 0

	players := make([]gamelogic.Player, 2)
	players[0] = *Red
	players[1] = *Green
	ante := 1
	minBet := 0
	maxBet := 10
	dealerToken := 0

	gamelogic.Rand_init()

	fmt.Printf("This prototype game will have two players, Red and Green.")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	gamelogic.Game(players, ante, minBet, maxBet, dealerToken)


}


