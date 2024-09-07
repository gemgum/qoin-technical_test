package repository

import (
	"dicegame/internal/player"
	"math/rand"
)

type Player struct {
	Player  int
	Point   int
	Dice    []int
	DiceOne int
}

type playerMod struct {
}

func NewPlayerMod() player.PlayerMod {
	return &playerMod{}
}

func (pm *playerMod) DiceNumbGenerator(min int, max int) (int, error) {
	RandomIntegerwithinRange := rand.Intn(max-min) + min
	return RandomIntegerwithinRange, nil

}
