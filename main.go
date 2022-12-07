package main

import (
	"fmt"
	"math/rand"
	"time"
)

// inisiasi struct player
type Player struct {
	Dice         []int
	Points       int
	IsFinished   bool
	AcquiredDice []int
}

// tambah point kalo abis lempar dadu 6
func (p *Player) AddPoint() {
	p.Points++
}

// remove dadu kalo dapat 6
func (p *Player) RemoveOneDice() {
	p.Dice = p.Dice[1:]
	// p.Dice = p.Dice[:len(p.Dice)-1]
}

// lempar dadu RNG yg menghasilkan 1 sampe 6
func RollDice() int {
	return rand.Intn(6) + 1 // biar jadi 1 - 6
}

func PlayDice2(jumlahPemain int, jumlahDadu int) {
	// inisiasi players
	players := make([]Player, jumlahPemain)

	// inisiasi dadu masing2 player
	for i := range players {
		players[i].Dice = make([]int, jumlahDadu)
	}
	fmt.Printf("Pemain = %v, Dadu = %v\n", jumlahPemain, jumlahDadu)

	fmt.Println("==================")
	for i := 0; ; i++ {
		// lempar dadu
		fmt.Printf("Giliran %v lempar dadu:\n", i+1)
		for i := range players {
			for j := range players[i].Dice {
				players[i].Dice[j] = RollDice()
			}
			fmt.Printf("Pemain#%v (%v): %v\n", i+1, players[i].Points, players[i].Dice)
		}

		// evaluasi
		for i := range players {
			// fmt.Println("Player ke", i+1, "dice", players[i].Dice)
			// fmt.Printf("Pemain#%v (%v): %v", i+1, players[i].Points, players[i].Dice)
			for j := 0; j < len(players[i].Dice); j++ {
				// for j := range players[i].Dice {
				// fmt.Println("len dice (", len(player.Dice), "), j ke-", j)
				if players[i].Dice[j] == 1 {
					// fmt.Println("INI I NYA : ", i)
					var y = i
					for {
						if y == len(players)-1 && !players[0].IsFinished {
							// fmt.Println("ERROR A INI I: ", y)

							players[0].AcquiredDice = append(players[0].AcquiredDice, 1)
							players[0].Dice = append(players[0].Dice, players[0].AcquiredDice...)
							players[0].AcquiredDice = nil
							break
						} else if y != len(players)-1 && !players[y+1].IsFinished {
							// fmt.Println("ERROR B INI I: ", y)

							players[y+1].AcquiredDice = append(players[y+1].AcquiredDice, 1)
							break
						} else {
							// fmt.Println("ERROR C INI I: ", y)
							if y == len(players)-1 {
								y = 0
							} else {
								y++
							}
							// fmt.Println("INI I++: ", y)
						}
					}
					// if i == len(players)-1 && !players[0].IsFinished {
					// 	players[0].AcquiredDice = append(players[0].AcquiredDice, 1)
					// 	players[0].Dice = append(players[0].Dice, players[0].AcquiredDice...)
					// 	players[0].AcquiredDice = nil
					// } else {
					// 	players[i+1].AcquiredDice = append(players[i+1].AcquiredDice, 1)
					// }
					// fmt.Println("ERROR D INI I: ", i)

					players[i].Dice = append(players[i].Dice[:j], players[i].Dice[j+1:]...)
					j--
					// fmt.Println("dapat 1")
					// fmt.Println(players[i])
				} else if players[i].Dice[j] == 6 {
					players[i].Points++
					// players[i].RemoveOneDice()
					players[i].Dice = append(players[i].Dice[:j], players[i].Dice[j+1:]...)
					j--
					// fmt.Println("dapat 6")
					// fmt.Println(players[i])
					// player.Dice = append(player.Dice[:j], player.Dice[j+1:]...)
				}
			}
			players[i].Dice = append(players[i].Dice, players[i].AcquiredDice...)
			players[i].AcquiredDice = nil
			// fmt.Println("New Dice : ", players[i].Dice)

		}

		// ini buat ngeprint hasil setelah evalnya
		sisaPlayer := jumlahPemain
		fmt.Println("\nSetelah evaluasi:")
		for i := range players {
			fmt.Printf("Pemain#%v (%v): %v\n", i+1, players[i].Points, players[i].Dice)
			if len(players[i].Dice) == 0 {
				sisaPlayer--
				players[i].IsFinished = true
				// fmt.Printf("Player ke-%v telah selesai bermain\n", i+1)
			}
		}

		fmt.Println("=================================")
		fmt.Println("sisa player: ", sisaPlayer)
		fmt.Println("=================================")

		// break kalo udah sisa 1 player
		if sisaPlayer <= 1 {
			// cek pemenang
			maxPoints := 0
			winner := 0

			for i, p := range players {
				if p.Points > maxPoints {
					maxPoints = p.Points
					winner = i
				}
			}

			fmt.Println("Pemenangnya adalah pemain ke-", winner+1, "dengan skor", maxPoints, "poin")
			break
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// for i := 0; i < 10; i++ {
	PlayDice2(rand.Intn(10)+1, rand.Intn(10)+1)
	// }

	// var m, n int
	// fmt.Println("masukkan jumlah pemain dan jumlah dadu: ")
	// fmt.Scanln(&m, &n)
	// PlayDice2(m, n)

}
