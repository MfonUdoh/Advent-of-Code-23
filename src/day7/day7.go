package day7

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day7() {
	file, err := os.Open(".input/day-7-input.txt")

	if err != nil {
		panic(err.Error())
	}

	scanner := bufio.NewScanner(file)

	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	var hands []Hand

	for _, line := range text {
		hands = append(hands, createHand(line))
	}

	game := Game{hands: hands}

	sort.Sort(game)

	score := game.getScore()

	fmt.Println(score)
}

func createHand(line string) Hand {
	split := strings.Split(line, " ")
	cards := split[0]
	bidString := split[1]
	bid, err := strconv.Atoi(bidString)

	if err != nil {
		panic(err.Error())
	}

	return Hand{
		cards: cards,
		bid:   bid,
	}
}

func rankHands(game Game) Game {
	// slices.Sort[Game](game)
	return game
}

type Game struct {
	hands []Hand
}

func (this Game) Len() int {
	return len(this.hands)
}

func (this Game) Less(i, j int) bool {
	return this.hands[i].beats(this.hands[j])
}

func (this Game) Swap(i, j int) {

}

func (this Game) getScore() int {
	var score int
	for i := 0; i < len(this.hands); i++ {
		score += this.hands[i].bid * (len(this.hands) - i)
	}
	return score
}

func calculateScores(rankedHands []Hand) int

func (this Hand) beats(otherHand Hand) bool

func (this Hand) beatsBasedOnType() bool
func (this Hand) beatsBasedOnHigh() bool

func (this *Hand) getType()

func (this Hand) find5Kind() (bool, int)
func (this Hand) find4Kind() (bool, int)
func (this Hand) findHouse() (bool, int)
func (this Hand) find3Kind() (bool, int)
func (this Hand) find2Pair() (bool, int)
func (this Hand) findPair() (bool, int)
func (this Hand) findHigh() (bool, int)

type Hand struct {
	cards    string
	bid      int
	handType int
}
