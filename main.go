package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Buatlah sebuah script permainan dadu yang menerima input N jumlah pemain dan M jumlah
	dadu, dengan peraturan sebagai berikut:
	1. Pada awal permainan, setiap pemain mendapatkan dadu sejumlah M unit.
	2. Semua pemain akan melemparkan dadu mereka masing-masing secara bersamaan
	3. Setiap pemain akan mengecek hasil dadu lemparan mereka dan melakukan evaluasi
	seperti berikut:
	a. Dadu angka 6 akan dikeluarkan dari permainan dan ditambahkan sebagai poin
	bagi pemain tersebut.
	b. Dadu angka 1 akan diberikan kepada pemain yang duduk disampingnya.
	Contoh, pemain pertama akan memberikan dadu angka 1 nya ke pemain kedua.
	c. Dadu angka 2,3,4 dan 5 akan tetap dimainkan oleh pemain.
	4. Setelah evaluasi, pemain yang masih memiliki dadu akan mengulangi step yang ke-2
	sampai tinggal 1 pemain yang tersisa.
	a. Untuk pemain yang tidak memiliki dadu lagi dianggap telah selesai bermain.
	5. Pemain yang memiliki poin terbanyak lah yang menang.
	Buatlah script ini menggunakan bahasa yang kamu kuasai.
*/

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
	for i := 0; i < 99; i++ {
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
					if i == len(players)-1 && !players[0].IsFinished {
						players[0].AcquiredDice = append(players[0].AcquiredDice, 1)
						players[0].Dice = append(players[0].Dice, players[0].AcquiredDice...)
						players[0].AcquiredDice = nil
					} else {
						players[i+1].AcquiredDice = append(players[i+1].AcquiredDice, 1)
					}
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
		remainingPlayer := jumlahPemain
		fmt.Println("Setelah evaluasi:")
		for i := range players {
			fmt.Printf("Pemain#%v (%v): %v\n", i+1, players[i].Points, players[i].Dice)
			if len(players[i].Dice) == 0 {
				remainingPlayer--
				players[i].IsFinished = true
				fmt.Printf("Player ke-%v telah selesai bermain\n", i+1)
			}
		}

		fmt.Println("=================================")
		fmt.Println("remainingPlayer: ", remainingPlayer)
		fmt.Println("=================================")

		// break kalo udah sisa 1 player
		if remainingPlayer <= 1 {
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

			// return winner
			// var winner Player
			// // fmt.Println(winner.Points)
			// for i := range players {
			// 	if players[i].Points > winner.Points {
			// 		winner = players[i]
			// 	}
			// }
			// fmt.Println("winner: ", winner)
			break
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// PlayDice1(3, 4)

	var m, n int
	fmt.Println("masukkan jumlah pemain dan jumlah dadu: ")
	fmt.Scanln(&m, &n)
	PlayDice2(m, n)

}
