package gamelogic

import(
	"math/rand"
	"time"
	//"fmt"
)

func init_card_cat() ([]string, []string){
cardTypes := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
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


func init() {
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
}


type Card struct {
	Face	string
	Suit	string
	Rank	int
	}


func newCard(face string, suit string, cardTypes []string)*Card{
	crd := new(Card)
	crd.Face = face
	crd.Suit = suit
	rank := getIndex(cardTypes, face) + 1
	crd.Rank = rank
	return crd
	}




func createDeck(cardTypes []string, suites []string)[]Card{
	/* create deck, adding each card looping through type and suite */
	deck := make([]Card, 52)
	for _, t := range cardTypes{
		for _, s := range suites{
			crd := newCard(t, s, cardTypes)
			deck = append(deck, *crd)

		}
	}
	return deck
}

func shuffle(d []Card){
	/*....   randomly re-order the array */
	N := len(d)
	for i := 0; i < N; i++ {
		selection := rand.Intn(N - i)
		d[selection], d[i] = d[i], d[selection]
	}
}

func draw(d []Card) Card{
	crd := d[0]
	d = d[1:]
	return crd
}
