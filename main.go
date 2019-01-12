package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	cards := CreateCards()
	deck := Deck{cards}
	p1 := Player{name: "p1"}
	dealer := Player{name: "dealer"}

	deck.Shuffle()

	p1.Draws(&deck)
	p1.Draws(&deck)
	dealer.Draws(&deck)
	dealer.Draws(&deck)

	for p1.Score() < 22 || dealer.Score() < 22 {
		p1.Draws(&deck)
		if p1.Score() > 21 {
			fmt.Println("BUST dealer wins with", dealer.Score())
			for _, card := range dealer.hand {
				fmt.Println(card.value.String(), card.suit.String())
			}
			os.Exit(0)
		}

		dealer.Draws(&deck)
		if dealer.Score() > 21 {
			fmt.Println("BUST p1 wins with", p1.Score())
			for _, card := range p1.hand {
				fmt.Println(card.value.String(), card.suit.String())
			}
			os.Exit(0)
		}
	}

}

type Suit int

type Arcania int

type Card struct {
	value Arcania
	suit  Suit
}

type Cards []Card

type Deck struct {
	cards Cards
}

type Player struct {
	score int
	hand  []Card
	name  string
}

const (
	HEARTS Suit = iota
	SPADES
	DIAMONDS
	CLUBS
)

const (
	ACE Arcania = iota + 1
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

func (s Suit) String() string {
	switch s {
	case 0:
		return "HEARTS"
	case 1:
		return "SPADES"
	case 2:
		return "DIAMONDS"
	case 3:
		return "SPADES"
	default:
		panic("Cant render suit")
	}
}

func (c Arcania) String() string {
	switch c {
	case 1:
		return "ACE"
	case 2:
		return "TWO"
	case 3:
		return "THREE"
	case 4:
		return "FOUR"
	case 5:
		return "FIVE"
	case 6:
		return "SIX"
	case 7:
		return "SEVEN"
	case 8:
		return "EIGHT"
	case 9:
		return "NINE"
	case 10:
		return "TEN"
	case 11:
		return "JACK"
	case 12:
		return "QUEEN"
	case 13:
		return "KING"
	default:
		panic("Card did not match")
	}
}

func ReadCard(card Card) int {
	switch card.value {
	case ACE:
		return 11
	case TWO:
		return 2
	case THREE:
		return 3
	case FOUR:
		return 4
	case FIVE:
		return 5
	case SIX:
		return 6
	case SEVEN:
		return 7
	case EIGHT:
		return 8
	case NINE:
		return 9
	case TEN:
		return 10
	case JACK:
		return 10
	case QUEEN:
		return 10
	case KING:
		return 10
	default:
		panic("Can't read card value")
	}
}

func (p *Player) Draws(d *Deck) {
	p.hand = append(p.hand, d.Deal())
}

func (p *Player) Score() int {
	var score int
	for _, card := range p.hand {
		score = score + ReadCard(card)
	}
	return score
}

func (d *Deck) Shuffle() {
	shuffled := append(d.cards[:0:0], d.cards...)
	for i, _ := range d.cards {
		rand.Seed(time.Now().UnixNano())
		span := len(shuffled) - i
		roll := rand.Intn(span)
		card := shuffled[roll]
		deck := append(shuffled[:roll], shuffled[1+roll:]...)
		shuffled = append(deck, card)
	}
	d.cards = shuffled
}

func (d *Deck) Deal() Card {
	if len(d.cards) == 0 {
		panic("All the cards have been dealt")
	}
	card, cards := d.cards[len(d.cards)-1], d.cards[:len(d.cards)-1]
	d.cards = cards
	return card
}

func CreateCards() Cards {
	cards := Cards{}

	for _, minor := range []Arcania{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN} {
		for _, suit := range []Suit{HEARTS, CLUBS, SPADES, DIAMONDS} {
			cards = append(cards, Card{minor, suit})
		}
	}

	for _, court := range []Arcania{JACK, QUEEN, KING} {
		for _, suit := range []Suit{HEARTS, CLUBS, SPADES, DIAMONDS} {
			cards = append(cards, Card{court, suit})
		}
	}

	return cards
}
