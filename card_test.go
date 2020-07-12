package deck

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	deck := NewDeck()
	if len(deck) != 13*4 {
		t.Error("Wrong number of cards in the deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := NewDeck(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received: ", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := NewDeck(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received: ", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	shuffledCards := NewDeck(Shuffle)
	defaultCards := NewDeck()

	if reflect.DeepEqual(shuffledCards, defaultCards) {
		t.Error("Shuffled cards and default cards are the same.")
	}
}

func TestJokers(t *testing.T) {
	nJokers := 5
	cards := NewDeck(AddJokers(nJokers))
	jokers := cards[len(cards)-nJokers:]

	for _, card := range jokers {
		if card.Suit != Joker {
			t.Error("Incorrect number of jokers added to the deck.")
		}
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := NewDeck(FilterOut(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out.")
		}
	}
}

func TestDeck(t *testing.T) {
	nDecks := 5
	cards := NewDeck(Deck(nDecks))
	if len(cards) != 13*4*nDecks {
		t.Errorf("Expected %d cards, found: %d", 13*4*nDecks, len(cards))
	}
}
