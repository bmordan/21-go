package main

import (
	"reflect"
	"testing"
)

func TestCreateCards(t *testing.T) {
	card := Card{ACE, HEARTS}

	if card.suit != HEARTS {
		t.Error("Is not the Ace of hearts!")
	}

	if card.value != ACE {
		t.Error("No. the value is not 1")
	}
}

func TestReadCard(t *testing.T) {
	card := Card{ACE, HEARTS}
	points := ReadCard(card)

	if points != 11 {
		t.Error("No ACE card should be equal to 11 points")
	}
}

func TestDeck(t *testing.T) {
	cards := CreateCards()
	deck := Deck{cards}

	if len(deck.cards) != 52 {
		t.Error("No. Cards must equal 52")
	}

	preShuffle := deck.cards
	deck.Shuffle()
	postShuffle := deck.cards

	if reflect.DeepEqual(preShuffle, postShuffle) {
		t.Error("Shuffle didn't work")
	}
}

func TestDeckDeal(t *testing.T) {
	cards := CreateCards()
	deck := Deck{cards}
	card := deck.Deal()
	kingOfDiamonds := Card{KING, DIAMONDS}

	if len(deck.cards) != 51 {
		t.Error("Dealing must remove cards from the deck")
	}

	if kingOfDiamonds != card {
		t.Error("Should be the * of *")
	}

	points := ReadCard(card)

	if points != 10 {
		t.Error("King should have points equal to 10")
	}
}

func TestPlayer(t *testing.T) {
	p1 := Player{name: "p1"}

	if p1.name != "p1" {
		t.Error("Player one should have a name equal to p1")
	}

	if len(p1.hand) != 0 {
		t.Error("Player one should have an empty hand")
	}

	if p1.score != 0 {
		t.Error("Player one should have a score of zero")
	}
}

func TestPlayerDrawing(t *testing.T) {
	deck := Deck{cards: CreateCards()}
	p1 := Player{name: "p1"}
	p1.Draws(&deck)

	if len(deck.cards) != 51 {
		t.Error("Player must remove cards from the deck")
	}

	if len(p1.hand) != 1 {
		t.Error("Player should hold a card")
	}
}

func TestPlayerScore(t *testing.T) {
	deck := Deck{cards: CreateCards()}
	p1 := Player{name: "p1"}
	p1.Draws(&deck)
	p1.Draws(&deck)

	if p1.Score() != 20 {
		t.Error("Players score is not correct")
	}
}
