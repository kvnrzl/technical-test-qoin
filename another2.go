package main

import "fmt"

func PlayDice1(jumlahPemain int, jumlahDadu int) {
	// inisiasi players
	players := make([]Player, jumlahPemain)

	// inisiasi dadu masing2 player
	for i := range players {
		players[i].Dice = make([]int, jumlahDadu)
	}
	fmt.Println(players)
	fmt.Println("aman 1")

	// lempar dadu
	for i, player := range players {
		fmt.Println("player ke", i+1)
		for j := range player.Dice {
			fmt.Println("ini roll ke-", j+1)
			roll := RollDice()
			fmt.Println(roll)
			if roll == 1 {
				player.RemoveOneDice()
				// player.Dice = append(player.Dice[:j], player.Dice[j+1:]...)
				if i == len(players)-1 {
					players[0].Dice = append(players[0].Dice, 1)
				} else {
					players[i+1].Dice = append(players[i+1].Dice, 1)
				}
				fmt.Println("aman pindah ke player sebelah")

			} else if roll == 6 {
				player.RemoveOneDice()
				player.AddPoint()
				fmt.Println("aman point +1 dan dadu dihapus")

			} else {
				player.Dice[j] = roll
				// player.Dice = append(player.Dice, roll)
				fmt.Println("aman dadu diisi")

			}
			fmt.Println(player)
		}
	}

	// cek pemenang
	var winner Player
	// fmt.Println(winner.Points)
	for i := range players {
		if players[i].Points > winner.Points {
			winner = players[i]
		}
	}

	fmt.Println("Pemenangnya adalah player ke-", winner)
}
