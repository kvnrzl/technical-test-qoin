package main

// inisiasi struct player
type Player struct {
	Dice         []int
	Points       int
	IsFinished   bool
	AcquiredDice []int // ini buat nampung dadu yg didapat dari pemain lain
}

// tambah point kalo abis lempar dadu 6
func (p *Player) AddPoint() {
	p.Points++
}

func (p *Player) RemoveADice(i int) {
	p.Dice = append(p.Dice[:i], p.Dice[i+1:]...)
}
