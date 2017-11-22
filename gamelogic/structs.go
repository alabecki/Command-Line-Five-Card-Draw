package gamelogic

import(
	"math/rand"
	"time"
	"fmt"
	"bufio"
	"os"
	"sort"
	"strings"
)

func Init_card_cat() ([]string, []string){
	cardTypes := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	suites := []string{"hearts", "spades", "clubs", "diamonds"}
	return cardTypes, suites
}


func getIndex(array []string, item string)int{
    for i := 0; i < len(array); i++ {
        if  array[i] == item {
            return i
        }
    }
    return -1
}


func Rand_init() {
    rand.Seed(time.Now().UTC().UnixNano())
}

/*classes:   */

type Globals struct {
	Bets	map[string]int 	/*A map containing the bets of all players */
}

type Player struct {			/* A more complete player struct will likely be someplace else in repo */
	Name	string
	Money	int
	Hand 	[]Card
	Folded  bool    /*default value false */
	Bet     int
	Hand_Rank int
	Card_Hist [14]int
}

 func (p *Player) pay_bet(amount int, pot int)int{
 	p.Money -= amount 
 	pot += amount
 	fmt.Printf("%s places %d dollars into the pot \n", p.Name, amount)
 	fmt.Printf("The pot now contains %s dollars \n", pot)
 	return pot
 }

 func (p *Player) stay_in(difference int) bool {
 	reader := bufio.NewReader(os.Stdin)
 	if difference > p.Money{
 		p.Folded = true
 		return false
 		fmt.Print("%s is unable to meet the raised bet and is out of the game \n", p.Name)
 	}
	fmt.Printf("The bet has been rased by %d \n", difference)
	fmt.Printf("Will %s stay in the game? (Y/N) \n", p.Name)
	var input string
	input, err := reader.ReadString('\n')
	fmt.Println(err)
	input = strings.Replace(input, "\r\n", "", -1)
	stay := false
	if input == "N" || input == "n"{
		stay = false
	}else if input == "Y" || input == "y"{
		stay = true
		}

 	/* show_fun() letting player know that he or she is out of the game */
 	
 	/* stay = show_func(difference) player p gets a pop up asking if he or she wishes to keep up with
 	the latest bet in order to remain in the game */
 	if stay == false{
 		p.Folded = true
 		return false
 	}else{
 		return true
 	}
 }

 func (p *Player)show_hand(){
 	fmt.Printf("%s's Hand: \n", p.Name)
	for i, crd := range p.Hand{
		fmt.Printf("%d %s of %s \n", i, crd.Face, crd.Suit)
	}
}


func (p *Player)sort_hand_by_rank(){
	fmt.Printf("About to sort %s hand by rank \n", p.Name)
	hand := p.Hand
	sort.Slice(hand[:], func(i, j int) bool {
    return hand[i].Rank < hand[j].Rank
	})
	p.Hand = hand
}


 func (p *Player) place_bet(current int, max_bet int, min_bet int) int {
 	//options := []string {"call", "fold", "raise"}
 	//if current == 0{
 	//	options := append(options, "check")
 	//}
 	if p.Money < current{
 		p.Folded = true
 		return current
 		fmt.Printf("You need to have %d dollars to stay in the game and only have %d \n", current, p.Money)
 		fmt.Printf("You have no choice but to fold \n")
 	}
 	fmt.Printf("Place bet for player %s \n", p.Name)
 	value := place_bet_test(*p, current, min_bet, max_bet)
 	fmt.Printf("Value is %d \n", value)

 	/* 
 	value = function(options, max_bet, min_bet)
 	function should show (or call something to show) the appropriate player a pop-up or something with 
 	the options listed and ok button if call, return 0, if raise, return the amount added to bet,
 	if fold, return -1. Do not let player bet more than his current money or the maximum bet*/
 	if value == -1{
 		p.Folded = true
 		return current
 	}
    return value
}

func (p *Player)find_four_of_kind_rank()int{
	for k, v := range p.Card_Hist{
		if v == 4{
			fmt.Printf("For of kind rank: %d \n", k)
			return k
		}
	}
	return 0 
}

func (p *Player)find_three_of_kind_rank()int{
	for k, v := range p.Card_Hist{
		if v == 3{
			return k
		}
	}
	return 0
}

func (p *Player)best_pair()int{
	for k := len(p.Card_Hist) - 1; k >= 0; k--{
		if p.Card_Hist[k] == 2{
			return k
		}
	}
	return 0
}

func (p *Player)second_best_pair()int{
	for k, _ := range p.Card_Hist{
		if p.Card_Hist[k] == 2{
			return k 
		}
	}
	return 0
}	






 /*func (p *Player)discarded_hand(discard_index []int){

	discard := make([]Card, 5)
	for _, d := range discard_index {
		discard = append(discard, p.Hand[d])
	}

	discarded_hand := make([]Card, 5)
	check := true
	for _, c := range p.hand{
		for _, d := range p.hand{
			if c.Face == d.Face && c.Suit == d.Suit{
				check = false
			}
		if check == true{
			discarded_hand = append(discarded_hand, c)
			}
		}
	}
} */

//func (p *Player) remove_card(Card){
//	index := getIndex(p.Hand, card)
//	p.Hand = append(p.Hand[:index], p.Hand[index+1:]...)
 //}



type Card struct {
	Face	string
	Suit	string
	Rank	int
	}


func newCard(face string, suit string, cardTypes []string)*Card{
	crd := new(Card)
	crd.Face = face
	crd.Suit = suit
	rank := getIndex(cardTypes, face)
	crd.Rank = rank
	return crd
	}





