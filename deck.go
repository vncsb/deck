//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8
type Rank uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eigth
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Rank
	Suit
}

var suits = [...]Suit{Spade, Diamond, Club, Heart}

const (
	minRank = Ace
	maxRank = King
)

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}

	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func NewDeck(opts ...func([]Card) []Card) []Card {
	var deck []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}

	for _, opt := range opts {
		deck = opt(deck)
	}

	return deck
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func FilterOut(filter func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		filteredCards := make([]Card, 0)
		for _, card := range cards {
			if !filter(card) {
				filteredCards = append(filteredCards, card)
			}
		}

		return filteredCards
	}
}

func Shuffle(cards []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func AddJokers(nJokers int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < nJokers; i++ {
			cards = append(cards, Card{Rank: Rank(i), Suit: Joker})
		}
		return cards
	}
}

func Deck(nDecks int) func([]Card) []Card {
	return func(cards []Card) []Card {
		ret := make([]Card, 0)
		for i := 0; i < nDecks; i++ {
			ret = append(ret, cards...)
		}

		return ret
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}
