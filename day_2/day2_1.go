package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type choice string
type score int
type shape string
type result string

var leftColumnLookup = map[choice]shape{"A": "rock", "B": "paper", "C": "scissors"}
var rightColumnLookup = map[choice]shape{"X": "rock", "Y": "paper", "Z": "scissors"}
var resultScores = map[result]score{"lost": 0, "draw": 3, "win": 6}
var shapeScores = map[shape]score{"rock": 1, "paper": 2, "scissors": 3}

func lookupScores(plays [2]shape) score {
	roundScore := lookUpResultScore(plays) + shapeScores[plays[1]]

	return roundScore

}

func normalise(game_line []string) [2]shape {
	result := [...]shape{leftColumnLookup[choice(game_line[0])], rightColumnLookup[choice(game_line[1])]}
	return result
}

func lookUpResultScore(hand [2]shape) score {
	leftHand := hand[0]
	rightHand := hand[1]

	if leftHand == rightHand {
		return resultScores["draw"]
	}

	if leftHand == "rock" {
		if rightHand == "scissors" {
			return resultScores["lose"]
		}
	}

	if leftHand == "paper" {
		if rightHand == "rock" {
			return resultScores["lose"]
		}
	}

	if leftHand == "scissors" {
		if rightHand == "paper" {
			return resultScores["lose"]
		}
	}

	return resultScores["win"]
}

func main() {
	gamesRecord, err := os.Open("day_2/input/games_played.txt")
	score := score(0)

	if err != nil {
		log.Fatal(err)
	}

	defer gamesRecord.Close()

	scanner := bufio.NewScanner(gamesRecord)

	for scanner.Scan() {
		line := scanner.Text()

		gameAnswers := normalise(strings.Split(line, " "))

		score = score + lookupScores(gameAnswers)
		fmt.Println(score)
	}
}
