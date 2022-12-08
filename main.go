package main

import (
	"fmt"
	"math/rand"
	"time"
)

// lempar dadu RNG yg menghasilkan 1 sampe 6
func rollDice() int {
	return rand.Intn(6) + 1 // biar jadi 1 - 6
}

func cekWinnerAndLastPlayer(players []Player) (int, int) {
	maxPoints := 0
	winner := 0
	lastPlayer := 0

	for i, p := range players {
		if p.Points > maxPoints {
			maxPoints = p.Points
			winner = i
		}
		if !p.IsFinished {
			lastPlayer = i
		}
	}

	return winner, lastPlayer
}

func PlayDice(jumlahPemain int, jumlahDadu int) {
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
				players[i].Dice[j] = rollDice()
			}
			fmt.Printf("Pemain#%v (%v): %v\n", i+1, players[i].Points, players[i].Dice)
		}

		// evaluasi
		for i := range players {
			for j := 0; j < len(players[i].Dice); j++ {
				if players[i].Dice[j] == 1 {
					var y = i
					for {
						if y == len(players)-1 {
							// kalo misalnya orang terakhir dapat 1 dan orang pertama juga belum selesai maka orang pertama dapat 1 (index 0)
							if !players[0].IsFinished {
								players[0].AcquiredDice = append(players[0].AcquiredDice, 1)
								break
							} else {
								// kalo misalnya orang terakhir dapat 1 dan orang pertama selesai maka index diset 0
								y = 0
							}
						} else if y != len(players)-1 && !players[y+1].IsFinished {
							players[y+1].AcquiredDice = append(players[y+1].AcquiredDice, 1)
							break
						} else {
							y++
						}
					}
					players[i].RemoveADice(j)
					j--

				} else if players[i].Dice[j] == 6 {
					players[i].AddPoint()
					players[i].RemoveADice(j)
					j--
				}
			}
		}

		// ini buat ngeprint hasil setelah evaluasinya
		sisaPlayer := jumlahPemain
		fmt.Println("\nSetelah evaluasi:")
		for i := range players {
			players[i].Dice = append(players[i].Dice, players[i].AcquiredDice...)
			players[i].AcquiredDice = nil
			fmt.Printf("Pemain#%v (%v): %v\n", i+1, players[i].Points, players[i].Dice)
			if len(players[i].Dice) == 0 {
				sisaPlayer--
				players[i].IsFinished = true
			}
		}

		fmt.Println("=================================")

		// break kalo udah sisa 1 player
		if sisaPlayer <= 1 {
			winner, lastPlayer := cekWinnerAndLastPlayer(players)

			fmt.Printf("Game berakhir karena hanya pemain #%v yang memiliki dadu.\n", lastPlayer+1)
			fmt.Printf("Game dimenangkan oleh pemain #%v karena memiliki poin lebih banyak dari pemain lainnya.\n", winner+1)

			break
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// for i := 0; i < 10; i++ {
	// PlayDice(rand.Intn(10)+1, rand.Intn(10)+1)
	// PlayDice(10, 9)
	// }

	var m, n int
	fmt.Println("masukkan jumlah pemain dan jumlah dadu: ")
	fmt.Scanln(&m, &n)
	PlayDice(m, n)

}
