
/*(only class attributes needed for the game logic will be included here) */
package gamelogic

import(
	"math/rand"
//	"time"
	"bufio"
  	"fmt"
  	"os"
  	"strings"
  	//"sort"
  	"strconv"
)

/* In casino play the first betting round begins with the player to the left of the big blind, 
and subsequent rounds begin with the player to the dealer's left. Home games typically use an ante; 
the first betting round begins with the player to the dealer's left, and the second round begins with 
the player who opened the first round.

Play begins with each player being dealt five cards, one at a time, all face down. The remaining deck
 is placed aside, often protected by placing a chip or other marker on it. Players pick up the cards and
  hold them in their hands, being careful to keep them concealed from the other players, then a round 
  of betting occurs.

If more than one player remains after the first round, the "draw" phase begins. Each player specifies
 how many of their cards they wish to replace and discards them. The deck is retrieved, and each player 
 is dealt in turn from the deck the same number of cards they discarded so that each player again has 
 five cards.

A second "after the draw" betting round occurs beginning with the player to the dealer's left or else 
beginning with the player who opened the first round (the latter is common when antes are used instead
 of blinds). This is followed by a showdown, if more than one player remains, in which the player with 
 the best hand wins the pot. */

 func Game(players []Player, ante int, minBet int, maxBet int, dealerToken int)*Player{
 	/* additional arguments might be added to function to denote display specifiations 
 	/* show() main game page
 	/* create and shuffle deck of cards */
 	//reader := bufio.NewReader(os.Stdin)
 	cheat := true
 	cardTypes, suites := Init_card_cat()
 	pot := 0
 	deck := createDeck(cardTypes, suites)
 	deck = shuffle(deck)
 	fmt.Printf("Deck test: \n")
	for _, d := range deck{
		fmt.Printf("%s of %s ", d.Face, d.Suit)
	}
 	/*maybe show() some shuffle gif animation
 	/* each player pays the ante (may later swich to 'blind') */
 	for i := 0; i < len(players); i++{
 		players[i].Money -= ante
 		pot += ante
 		fmt.Printf("%s pays %d for ante \n", players[i].Name, ante)
	}
	if cheat == true{
		reader := bufio.NewReader(os.Stdin)
		for i := 0; i < len(players); i++{
			fmt.Printf("Time for %s to chose cards...\n", players[i].Name)
			for j := 0; j < 5; j++{
				fmt.Printf("Please enter face of next card:\n")
				face, _ := reader.ReadString('\n')
				fmt.Printf("Please enter the suit of the next card:\n")
				suit, _ := reader.ReadString('\n')
				face = strings.Replace(face, "\r\n", "", -1)
				suit = strings.Replace(suit, "\r\n", "", -1)
				card := newCard(face, suit, cardTypes)
				players[i].Hand = append(players[i].Hand, *card)
			}
		}
	}else{
 
 		/*first round dealing */
	 	fmt.Printf("The dealer suffles the cards and begins dealing... \n")
	 	bufio.NewReader(os.Stdin).ReadBytes('\n')

	 	d := 0
	 	for d < 5{
	 		for i := 0; i < len(players); i++{
	 			card := draw(deck)
	 			deck = deck[1:]
	 			players[i].Hand = append(players[i].Hand, card)
	 			fmt.Printf(" %s is delt a %s of %s \n ", players[i].Name, card.Face, card.Suit)
	 		}
	 		d++
	 	}
	}
 	/*first round betting */
 	pot = betting_round(players, minBet, maxBet, pot)
 	remaining := check_num_players_remaining(players)
 	if remaining < 2 {
 		winner := find_winner(players)
 		return winner
 	}
 	/* first draw */ 
 	deck = redraw(players, deck)
 	/* second round of betting */
 	pot = betting_round(players, minBet, maxBet, pot)
 	remaining = check_num_players_remaining(players)
 	if remaining < 2 {
 		winner := find_winner(players)
 		return winner
 	}
 	/* second draw */ 
 	deck= redraw(players, deck)
 	/* Third and final round of betting */
 	pot = betting_round(players, minBet, maxBet, pot)
 	remaining = check_num_players_remaining(players)
 	if remaining < 2 {
 		winner := find_winner(players)
 		return winner
 	}
 	/* sort hands by rank to prepare for hand comparisons */
 	for i := 0; i < len(players); i++{
 		if players[i].Folded == false{
 			players[i].sort_hand_by_rank()
 			fmt.Printf("Sorted hand: \n")
 			players[i].show_hand()
 			players[i].card_histogram()
 			}
 		}
 	score_board :=rank_hands(players)	
 	winner := showdown(players, score_board)
 	fmt.Printf("%s win a pot worth %d \n", winner.Name, pot)
 	winner.Money += pot
 	return winner

 		
 }

 func createDeck(cardTypes []string, suites []string)[]Card{
	/* create deck, adding each card looping through type and suite */
	deck := make([]Card, 52)
	count := 0
	for _, t := range cardTypes{
		for _, s := range suites{
			crd := newCard(t, s, cardTypes)
			deck[count] = *crd
			count++

		}
	}
	//fmt.Printf("Deck test: \n")
	//for _, d := range deck{
	//	fmt.Printf("%s of %s ", d.Face, d.Suit)
	//}
	return deck
}


func shuffle(d []Card)[]Card{
	/*....   randomly re-order the array */
	for i := len(d) - 1; i > 0; i-- {
		selection := rand.Intn(i + 1)
		d[i], d[selection] = d[selection], d[i]
	}
	return d
}

func draw(d []Card)Card{
	crd := d[0]
	return crd
}


 func check_num_players_remaining(players []Player)int{
 	remaining := 0
 	for _, p := range players{
 		if p.Folded == false{
 			remaining ++
 		}
 	}
 	fmt.Printf("%d players remaining in this game\n", remaining)
 	return remaining
 }

func find_winner(players []Player)*Player{
	//function assumes only one players remains in the game
	for _, p := range players{
		if p.Folded == false{
			return &p
		}
	}
	p := players[0]		//just to make go happy
	return &p
}

func place_bet_test(p Player, current int, min_bet int, max_bet int) int{
	reader := bufio.NewReader(os.Stdin)
	p.show_hand()
	fmt.Printf("%s, please place your bet \n", p.Name)
	fmt.Printf("%s has %d dollars and current bet is %d \n", p.Name, p.Money, current)
	fmt.Printf("Input -1 to fold, 0 to call or the amount you wish to raise \n")
	var bet int
    //_, err := fmt.Scanf("%d", &bet)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\r\n", "", -1)
	fmt.Printf("Input = %s", input)
	bet, err := strconv.Atoi(input)
	fmt.Println(bet)
	fmt.Println(err)
	return bet

}

 	

func betting_round(players []Player, minBet int, maxBet int, pot int)int{
	//reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(players); i++{
		players[i].Bet = 0
	}
	bet := 0
	for i := 0; i < len(players); i++{
 		if len(players) == 1{
 			return pot
 		}
 		bet = bet + players[i].place_bet(bet, minBet, maxBet)
 		pot = players[i].pay_bet(bet, pot)
 		players[i].Bet = bet
 		for j := 0; j < len(players); j++{
 			fmt.Printf("If raised, other players will need to match the bet of %d \n", bet)
 			if bet > players[j].Bet{
 				difference := bet - players[j].Bet
 				stay := players[j].stay_in(difference)
 				if stay == false{
 					fmt.Printf("%s is folding \n", players[j].Name)
 					players[j].Folded = true
 				}
 				if stay == true{
 					fmt.Printf("%s is staying ing the game \n", players[j].Name)
 					pot = players[j].pay_bet(difference, pot)
 				}
 			}
 		}
 		//bet += amount
 		fmt.Printf("bet is currently %d \n", bet)
 	}
 	fmt.Printf("Returning from betting round \n")
 	return pot
}


func stringToIntSlice(initial string) []int{
 	strs := strings.Split(initial," ")
 	fmt.Printf("Len of int-string: %d \n", len(strs))
 	fmt.Printf("%s \n", strs)
    ary := make([]int, len(strs))
    for i := range ary {
        ary[i], _ = strconv.Atoi(strs[i])
    }
    return ary
}

func redraw(players []Player, deck []Card)[]Card{
	reader := bufio.NewReader(os.Stdin)
	/*Each player may discard cards */
	for i := 0; i < len(players); i++{
 		if players[i].Folded == false{
 			/* remove = p.show_func() ask player which cards to remove return array of cards to be removed 
 			the array may be empty */
 			players[i].show_hand()
 			fmt.Printf("Would you like to discard any cards (y/n)? \n")
 			//var choice string
 			choice, _ := reader.ReadString('\n')
 			choice = strings.Replace(choice, "\r\n", "", -1)
 			if choice == "n"{
 				continue
 			}
 			fmt.Printf("Which cards would you like to discard? \n")
 			var input string
 			//_, err := fmt.Scanf("%s\n", &input)
 			input, err := reader.ReadString('\n')
 			fmt.Printf("err: %s \n", err)
 			input = strings.Replace(input, "\r\n", "", -1)
 			fmt.Printf("input: %s \n", input)
 			discard_index := stringToIntSlice(input)
 			fmt.Printf("Discard index %v: \n", discard_index)
 			players[i].Hand = discarded_hand(players[i].Hand, discard_index)	
 			}
 		}
 	/* Deal new cards to players */
 	fmt.Printf("The dealer will now deal new cards to the players... (press 'enter')\n")
 	bufio.NewReader(os.Stdin).ReadBytes('\n')

 	for i := 0; i < len(players); i++{
 		if players[i].Folded == false{
 			//hand_size := 5 - len(discard_index)
 			fmt.Printf("%s's hand len after discard: %d", players[i].Name, len(players[i].Hand))
 			replace := 5 - len(players[i].Hand) 
 			for j := 0; j < replace; j++{
 				card := draw(deck)
 				deck = deck[1:]
 				fmt.Printf("%s draws a %s of %s \n", players[i].Name, card.Face, card.Suit)
 				players[i].Hand = append(players[i].Hand, card)
 			}
 		}
 	}
 	return deck
}

func discarded_hand(hand []Card, discard_index []int)[]Card{
	//temp_hand := make([]Card, 5)
	var temp_hand []Card
	for _, t := range discard_index {
		//temp_hand[i] = hand[t]
		temp_hand = append(temp_hand, hand[t])
	}
	fmt.Printf("Print temp_hand\n")
	for _, t := range temp_hand{
		fmt.Println(t)
	}
	var discarded_hand []Card
	//discarded_hand := make([]Card, 5)
	index := 0
	for _, c := range hand{
		fmt.Printf("%s of %s \n", c.Face, c.Suit)
		 		check := true
		for _, d := range temp_hand{
			fmt.Printf("%s of %s \n", d.Face, d.Suit)
			if c.Face == d.Face && c.Suit == d.Suit{
				fmt.Printf("Do not include %s of %s \n", c.Face, c.Suit )
				check = false
			}
		}
		if check == true{
			fmt.Printf("Include %s of %s \n", c.Face, c.Suit)
			discarded_hand = append(discarded_hand, c)
			index ++
			}
		}
	fmt.Printf("Discarded Hand:\n")
	for _, c := range discarded_hand{
		fmt.Println(c)
	}
	return discarded_hand
}

func (p *Player)card_histogram(){
	/*p.Card_Hist = map[string]int{
		"1":	0,
		"2":	0,
		"3":	0,
		"4":	0,
		"5":	0,
		"6":	0,
		"7":	0,
		"8":	0,
		"9":	0,
		"10":	0,
		"11":   0,
		"12":	0,
		"13":	0, 
	} */
	for _, crd := range p.Hand{
		p.Card_Hist[crd.Rank]++
	}
	fmt.Printf("Card Hist of %s: \n", p.Name)
	for _, i := range p.Card_Hist{
		fmt.Printf(" %d \n", i)
	}

}

func max_dict(dict map[string]int){
	max := 1
	for _, v := range dict{
		if v > max{
		max = v
		}
	}
}

/*func (p *Player)rank_hand(){
	first 	:= 1
	second  := 1
	for k, v := range p.Card_Hist{
		if v > 1{
			if first !=1{
				second = first

			}
		}
	}
	first_tier := 1
	second_tier := 1
	larg_group := 0
	small_group := 0
	for r := 13; r > 0; r--{
		//if r > first_tier
		if first_tier != 1{
			second_tier := first_tier
			two_rank := three_rank  
		}
		first_tier = x
		large_group = x 
	}
} */


	/* hand ranking :
		straight flush
		four of a kind
		full house
		straight
		flush
		three of a kind
		two pairs
		nothing
	when determining the winner: consider hand rank of each player. If no tie of rank select winner
	if two or more players are tied for top rank, call a second function that compares the hands based on 
	rank of the individual cards */

func rank_hands(players []Player)map[string][]int{
	fmt.Printf("About to begin ranking hands...\n")
	score_category_map := make(map[string][]int)
	var value []int
	for rank := 0; rank < 9; rank ++{
		score_category_map["straight_flush"] = value
		score_category_map["four_of_a_kind"] = value
		score_category_map["full_house"] = value
		score_category_map["flush"] = value
		score_category_map["straight"] = value
		score_category_map["three_of_a_kind"] = value
		score_category_map["two_pairs"] = value
		score_category_map["pair"] = value
		score_category_map["nothing"] = value
	}
	//	"straight_flush": []Player,
	//	"four_of_a_kind": []Player,
	//	"full_house": []Player,
	//	"flush": []Player,
	///	"straight": []Player,
	//"three_of_a_kind": []Player,
	//	"two_pairs": []Player,
	//	"pair": []Player,
	//	"nothing": []Player,
	//}
	for i := 0; i < len(players); i++{
		if players[i].Folded == false{
			hand := players[i].Hand
			if check_straight_flush(hand){
				//score_category_map["straight_flush"]++
				score_category_map["straight_flush"] = append(score_category_map["straight_flush"], i )
				fmt.Printf("%s has a straight flush\n", players[i].Name)
			}else if check_four_of_a_kind(hand){
				//score_category_map["four_of_a_kind"]++
				score_category_map["four_of_a_kind"] = append(score_category_map["four_of_a_kind"], i)
				fmt.Printf("%s has a four of a kind\n", players[i].Name)
			}else if check_full_house(hand){
				//score_category_map["full_house"]++
				score_category_map["full_house"] = append(score_category_map["full_house"], i)
				fmt.Printf("%s has a full house\n", players[i].Name)
			}else if check_flush(hand){
				fmt.Printf("%s has a flush\n", players[i].Name)
				score_category_map["flush"] = append(score_category_map["flush"], i)
				//player.Hand_Rank = 4
			}else if check_stright(hand){
				//score_category_map["straight"]++
				score_category_map["straight"] = append(score_category_map["straight"], i)
				fmt.Printf("%s has a straight\n ", players[i].Name)
			}else if check_three_of_a_kind(hand){
				//score_category_map["three_of_a_kind"]++
				score_category_map["three_of_a_kind"] = append(score_category_map["three_of_a_kind"], i)
				fmt.Printf("%s has a three of a kind \n", players[i].Name)
			}else if check_two_pairs(hand){
				//score_category_map["two_pairs"] ++
				score_category_map["two_pairs"] = append(score_category_map["two_pairs"], i)
				fmt.Printf("%s has a two pairs\n ", players[i].Name)
			}else if check_pair(hand){
				//score_category_map["pair"]++
				score_category_map["pair"] = append(score_category_map["pair"], i)
				fmt.Printf("%s has one pair \n", players[i].Name)
			}else{
				//score_category_map["nothing"]++
				score_category_map["nothing"] = append(score_category_map["nothing"], i)
				fmt.Printf("%s has nothing \n", players[i].Name)
			}
		/* not yet complete - will modify return values so that it is easier to determine the winner
			when the best two hands belong to the same category */	
		}
	}
	return score_category_map
}

func check_flush(hand []Card)bool{
	suit := hand[0].Suit
	for _, c := range hand{
		if c.Suit != suit{
			return false
		}
	}
	return true
}

func check_stright(hand []Card)bool{
	for i := 1; i < len(hand); i++{
		if hand[i].Rank != hand[i-1].Rank + 1{
			return false
		}
	}
	return true
}

func check_straight_flush(hand []Card)bool{
	flush := check_flush(hand)
	straight := check_stright(hand)
	if flush == true && straight == true{
		return true
	}else{
		return false
	}
}

func check_four_of_a_kind(hand []Card)bool{
	if hand[0].Rank == hand[3].Rank || hand[1].Rank == hand[4].Rank{
		return true
	}else{
  		return false
  	}
}


func check_full_house(hand []Card)bool{
	if hand[0].Rank == hand[1].Rank && hand[2].Rank == hand[4].Rank{
		return true
	}
	if hand[0].Rank == hand[2].Rank && hand[3].Rank == hand[4].Rank{
		return true
	}else{
		return false
	}
}

func check_three_of_a_kind(hand []Card)bool{
	if hand[0].Rank == hand[2].Rank || hand[1].Rank == hand[3].Rank || hand[2].Rank == hand[4].Rank{
		return true
	}else{
		return false
	}
}


/*A lower ranked unmatched card + 2 cards of the same rank + 2 cards of the same rank
2 cards of the same rank + a middle ranked unmatched card + 2 cards of the same rank
2 cards of the same rank + 2 cards of the same rank + a higher ranked unmatched card
*/


func check_two_pairs(hand []Card)bool{
	if hand[1].Rank == hand[2].Rank && hand[3].Rank == hand[4].Rank{
		return true
	}
	if hand[0].Rank == hand[1].Rank && hand[3].Rank == hand[4].Rank{
		return true
	}
	if hand[0].Rank == hand[1].Rank && hand[2].Rank == hand[3].Rank{
		return true
	}
	return false
}

/* this function must be used after the others, as it does not make sure that the hand has
nothing greater than two of a kind */
func check_pair(hand []Card)bool{
	try := 0
	for try < len(hand)-1{
		for i := try+1; i < len(hand); i++{
			if hand[try].Rank == hand[i].Rank{
				return true
			}
		}
		try ++
	}
	return false
}

/*func get_player_index_from_name(players []Player, name string)int{
	for i, p := range players{
		if p.Name == name{
		return i
		}	
	}
	return 0
} */

func showdown(players []Player, score_board map[string][]int)*Player{
	fmt.Printf("About to begin showdown \n")
	fmt.Printf("score_board:")
	for key, value := range score_board{
		fmt.Printf("%s: %v \n", key, value)
	}
	best := 0
	//hand_rank := 1
	var winner Player
	if len(score_board["straight_flush"]) == 1{
		win := score_board["straight_flush"][0]
		winner = players[win]
		//return winner
	}else if len(score_board["straight_flush"]) > 1{
		contenders := score_board["straight_flush"]
		for c := 0; c < len(contenders); c++{
			if players[c].Hand[4].Rank > best{
				best = players[c].Hand[4].Rank
				winner = players[contenders[c]]
			}
		}
	}else if len(score_board["four_of_a_kind"]) == 1{
		win := score_board["four_of_a_kind"][0]
		winner = players[win]
	}else if len(score_board["four_of_a_kind"]) > 1{
		contenders := score_board["four_of_a_kind"]
		for c := 0; c < len(contenders); c++{
			if players[contenders[c]].find_four_of_kind_rank() > best{
				best = players[contenders[c]].find_four_of_kind_rank()
				winner = players[contenders[c]]
				}
			}
	}else if len(score_board["full_house"]) == 1{
		win := score_board["full_house"][0]
		winner = players[win]
	}else if len(score_board["full_house"]) > 1{
		contenders := score_board["full_house"]
		for c := 0; c < len(contenders); c++{
			if players[contenders[c]].find_three_of_kind_rank() > best{
				best = players[contenders[c]].find_three_of_kind_rank()
				winner = players[contenders[c]]
				}
			}
	}else if len(score_board["flush"]) == 1{
		win := score_board["full_house"][0]
		winner = players[win]
	}else if len(score_board["flush"]) > 1{
		contenders := score_board["flush"]
		win := find_best_nothing(contenders, players)
		winner = players[win]
	
	}else if len(score_board["straight"]) == 1{
		win := score_board["straight"][0]
		winner = players[win]
	}else if len(score_board["straight"]) > 1{
		contenders := score_board["straight"]
		win := find_best_nothing(contenders, players)
		winner = players[win]
	}else if len(score_board["three_of_a_kind"]) == 1{
		win := score_board["three_of_a_kind"][0]
		winner = players[win]
	}else if len(score_board["three_of_a_kind"]) > 1{
		contenders := score_board["three_of_a_kind"]
		for c := 0; c < len(contenders); c++{
			if players[contenders[c]].find_three_of_kind_rank() > best{
				best = players[contenders[c]].find_three_of_kind_rank()
				winner = players[contenders[c]]
			} 
		}
	}else if len(score_board["two_pairs"]) == 1{
		win := score_board["two_pairs"][0]
		winner = players[win]
	}else if len(score_board["two_pairs"]) > 1{
		contenders := score_board["two_pairs"]
		best2 := 0
		for i := 0; i < len(contenders); i++{
			rank := players[contenders[i]].best_pair()
			fmt.Printf("%s best pair rank is %d \n", players[contenders[i]].Name, rank)
			if rank > best{
				best = rank
				winner = players[contenders[i]]
			}else if rank == best && best > 0{
				fmt.Printf("tie detected \n")
				best = rank
				for j := 0; j < len(contenders); j++{
					if players[contenders[j]].best_pair() < best{
						continue
					}
					rank2 := players[contenders[j]].second_best_pair()
					fmt.Printf("%s second best pair rank is %d \n", players[contenders[j]].Name, rank2)
					if rank2 > best2{
						best2 = rank2
						winner = players[contenders[j]]
					}else if rank2 == best2 && best2 > 0{
						fmt.Printf("Second tie!! \n")
						win := find_best_nothing(contenders, players)
						winner = players[win]							
						fmt.Printf("current winner: %s \n", winner.Name)
						return &winner
						}		
					}
				} 
			}
		}else if len(score_board["pair"]) == 1{
			win := score_board["pair"][0]
			winner = players[win]
		}else if len(score_board["pair"]) > 1{
			contenders := score_board["two_pairs"]
			fmt.Printf("Best is %d \n", best)
			for i := 0; i < len(contenders); i++{
				fmt.Printf("i is %d \n", i)
				rank := players[contenders[i]].best_pair()
				fmt.Printf("Rank is %d \n", rank)
				if rank > best{
					best = rank
					winner = players[contenders[i]]
				}else if rank == best && best > 0{
					index := find_best_nothing(contenders, players)
					winner = players[index]
					return &winner
					}
				}
		}else{
			contenders := score_board["nothing"]
			index := find_best_nothing(contenders, players)
			winner = players[index]
			}
return &winner 
}


func find_best_nothing(indexes []int, players []Player)int{

	for i := 4; i > 0; i--{
		win_indx := 0
		highest := 0
		tie := false
		for _, j := range indexes{
			fmt.Printf("%s has a rank %d card \n", players[j].Name, players[j].Hand[i].Rank)
			if players[j].Hand[j].Rank > highest{
				if players[j].Hand[j].Rank == highest && highest > 0{
					tie = true
				}
				win_indx = j
			}
		}
		if tie == false{
			return win_indx
		}
	}
	return -1
}


func Unit_Test()*Player{

	cardTypes, _ := Init_card_cat()

	Red := new(Player)
	Red.Name = "Red"
	Red.Money = 100
	Red.Folded = false
	Red.Bet = 0


	Green := new(Player)
	Green.Name = "Green"
	Green.Money = 100
	Green.Folded = false
	Green.Bet = 0

	players := make([]Player, 2)
	players[0] = *Red
	players[1] = *Green
	


	//deck := make([]Card, 52)

	
	a := newCard("King", "diamonds", cardTypes)
	b := newCard("Ace", "hearts", cardTypes)
	c := newCard("Jack", "spades", cardTypes)
	d := newCard("7", "clubs", cardTypes)
	e := newCard("7", "diamonds", cardTypes)

	arr := []Card{*a, *b, *c, *d, *e}
	for i := 0; i < len(arr); i++ {
		players[1].Hand = append(players[1].Hand, arr[i])
	}



	aa := newCard("10", "clubs", cardTypes)
	bb := newCard("Queen", "spades", cardTypes)
	cc := newCard("7", "spades", cardTypes)
	dd := newCard("Queen", "clubs", cardTypes)
	ee := newCard("Jack", "diamonds", cardTypes)

	arrb := []Card{*aa, *bb, *cc, *dd, *ee}
	for i := 0; i  < len(arrb); i++ {
		players[0].Hand = append(players[0].Hand, arrb[i])
	}


	//Red.show_hand()
	//Green.show_hand()

	for i := 0; i < len(players); i++{
		players[i].show_hand()
	}
	

	value := make([]int, 0)
	value = append(value, 0)
	value = append(value, 1)

	for i := 0; i < len(players); i++{
 	if players[i].Folded == false{
 		players[i].sort_hand_by_rank()
 		players[i].show_hand()
 		players[i].card_histogram()
 			}
 		}
 	

	best := 0
	winner := players[0]

	for i := 0; i < len(value); i++{
		fmt.Printf("i is %d", i)
		rank := players[value[i]].best_pair()
		fmt.Printf("Rank is %d", rank)
		if rank > best{
			if rank == best && best > 0{
				best = rank
				index := find_best_nothing(value, players)
				winner = players[index]
			}
		}
	}
	return &winner
}
