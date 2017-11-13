
package main

import(
//	"math/rand"
//	"time"
//	"fmt"
	"./gamelogic"
	



)

func main(){

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
	players = append(players, *Red)
	players = append(players, *Green)
	ante := 1
	minBet := 0
	maxBet := 10
	dealerToken := 0


	gamelogic.Game(players, ante, minBet, maxBet, dealerToken)




}