package main

import (
	"fmt"
	"reflect"
)

type deck struct {
	cards   []int
	history [][]int
	number  int
}

func (d *deck) copy(num int) deck {
	cp := make([]int, 0)
	rv := reverse(d.cards)
	for _, x := range rv {
		if len(cp) == num {
			break
		}

		cp = append(cp, x)
	}

	return deck{
		cards:   reverse(cp),
		history: make([][]int, 0),
		number:  d.number,
	}
}

func (d *deck) haveSeen() bool {
	for _, hist := range d.history {
		if len(hist) != len(d.cards) {
			continue
		}

		if reflect.DeepEqual(hist, d.cards) {
			return true
		}
	}

	return false
}

func (d *deck) addCards(add []int) {
	d.cards = append(add, d.cards...)
}

func (d *deck) addToHistory() {
	cp := make([]int, 0)
	for _, x := range d.cards {
		cp = append(cp, x)
	}

	d.history = append(d.history, cp)
}

func (d *deck) getLatest() int {
	last := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return last
}

func (d *deck) hasNext() bool {
	return len(d.cards) != 0
}

func day22(player1, player2 []int) {
	p1 := deck{
		history: make([][]int, 0),
		cards:   reverse(player1),
		number:  1,
	}
	p2 := deck{
		history: make([][]int, 0),
		cards:   reverse(player2),
		number:  2,
	}

	res := game(p1, p2)
	fmt.Println(res.cards)
	sum := int64(0)
	for ind, x := range res.cards {
		sum = sum + int64((ind+1)*x)
	}

	fmt.Println(sum)
}

func reverse(inp []int) []int {
	res := make([]int, 0)
	for i := len(inp) - 1; i >= 0; i-- {
		res = append(res, inp[i])
	}

	return res
}

func game(p1, p2 deck) deck {

	hasP1, hasP2 := true, true
	for hasP1 && hasP2 {
		if p1.haveSeen() && p2.haveSeen() {
			return p1
		}

		p1.addToHistory()
		p2.addToHistory()
		card, card2 := p1.getLatest(), p2.getLatest()
		if len(p1.cards) < card || len(p2.cards) < card2 {
			if card > card2 {
				p1.addCards([]int{card2, card})
			} else {
				p2.addCards([]int{card, card2})
			}
			hasP1, hasP2 = p1.hasNext(), p2.hasNext()
			continue
		} else {
			res := game(p1.copy(card), p2.copy(card2))
			if res.number == 1 {
				p1.addCards([]int{card2, card})
			} else {
				p2.addCards([]int{card, card2})
			}
		}

		hasP1, hasP2 = p1.hasNext(), p2.hasNext()
	}

	if len(p1.cards) != 0 && len(p2.cards) != 0 {
		panic("fail")
	}

	if len(p1.cards) > 0 {
		return p1
	}

	if len(p2.cards) > 0 {
		return p2
	}

	return p2
}
