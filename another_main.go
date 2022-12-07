package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // Simulates rolling a dice and returns a random number between 1 and 6.
// func rollDice() int {
// 	rand.Seed(time.Now().UnixNano())
// 	return rand.Intn(6) + 1
// }

// // Represents a player in the game.
// type player struct {
// 	dice   []int // The dice owned by the player
// 	points int   // The points accumulated by the player
// }

// // Removes a dice from the given player's dice collection.
// func (p *player) removeDice() {
// 	if len(p.dice) > 0 {
// 		p.dice = p.dice[1:]
// 	}
// }

// // Adds a point to the given player's point total.
// func (p *player) addPoint() {
// 	p.points++
// }

// // Determines the winner of the game based on the players' point totals.
// func findWinner(players []player) int {
// 	maxPoints := 0
// 	winner := 0

// 	for i, p := range players {
// 		if p.points > maxPoints {
// 			maxPoints = p.points
// 			winner = i
// 		}
// 	}

// 	return winner
// }

// // Plays the dice game with the given number of players and rounds.
// func playGame(numPlayers, numRounds int) {
// 	// Initialize the players and the game counter
// 	players := make([]player, numPlayers)
// 	gameRound := 1

// 	// Give each player the specified number of dice
// 	for i := range players {
// 		players[i].dice = make([]int, numRounds)
// 	}

// 	// Play the game for the given number of rounds
// 	for gameRound <= numRounds {
// 		// In each round, each player rolls their dice
// 		for i := range players {
// 			for j := range players[i].dice {
// 				roll := rollDice()
// 				fmt.Printf("Player %d rolls a %d\n", i+1, roll)

// 				// Update the player's dice collection based on the roll
// 				if roll == 6 {
// 					players[i].removeDice()
// 					players[i].addPoint()
// 				} else if roll == 1 {
// 					// If the player is the last one, give the dice to the first player
// 					if i == len(players)-1 {
// 						players[0].dice = append(players[0].dice, roll)
// 					} else {
// 						players[i+1].dice = append(players[i+1].dice, roll)
// 					}
// 				}
// 			}
// 		}

// 		// Remove players who have no more dice
// 		remainingPlayers := make([]player, 0)
// 		for _, p := range players {
// 			if len(p.dice) > 0 {
// 				remainingPlayers = append(remainingPlayers, p)
// 			}
// 		}
// 		players = remainingPlayers

// 		// Move on to the next round
// 		gameRound++
// 	}

// 	// Determine the winner of the game
// 	winner := findWinner(players)
// }
