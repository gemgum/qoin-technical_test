package service

import (
	"dicegame/internal/player"
	"fmt"
)

type playerServ struct {
	mod player.PlayerMod
}

func NewPlayerServ(m player.PlayerMod) player.PlayerServ {
	return &playerServ{
		mod: m,
	}
}

func removeElement(slice []int) []int {
	// Hasil slice baru setelah elemen dihapus
	result := []int{}

	// Iterasi setiap elemen dalam slice asli
	for _, v := range slice {
		if v != 6 {
			result = append(result, v)
		}
	}

	return result
}

func (ps *playerServ) PlayerPos(total int, diceRules int) []player.Player {
	rv := []player.Player{}
	for i := 0; i < total; i++ {
		result := player.Player{
			Player: i,
			Dice:   make([]int, diceRules),
		}
		rv = append(rv, result)
	}
	return rv
}

func (ps *playerServ) PlayerTurn(dataPlayer []player.Player) ([]player.Player, error) {

	rv := dataPlayer

	for i, v := range dataPlayer {
		result := v
		result.Dice = make([]int, len(v.Dice))
		for j := 0; j < len(result.Dice); j++ {
			numbDice, _ := ps.mod.DiceNumbGenerator(1, 7)
			result.Dice[j] = numbDice
			result.Player = v.Player
			result.Point = v.Point

		}
		rv[i] = result
	}
	fmt.Println("rv", rv)

	return rv, nil
}

func (ps *playerServ) GameRules(data []player.Player) ([]player.Player, error) {
	dataPlayer, _ := ps.PlayerTurn(data)
	rv := dataPlayer
	for i, v := range rv {
		result := v
		for _, r := range result.Dice {
			if r == 1 {
				result.DiceOne++
			}
			if r == 6 {
				result.Point = result.Point + 1
			}

		}

		result.Dice = removeElement(result.Dice)

		if result.DiceOne != 0 {
			var pos int
			pos = i
			for k := 0; k < len(rv)-1; k++ {
				pos++
				if pos > len(rv)-1 {
					pos = 0
				}
				if pos != i && len(rv[pos].Dice) > 0 {
					for l := 0; l < result.DiceOne; l++ {
						rv[pos].Dice = append(rv[pos].Dice, 0)
					}
					break
				}

			}
			result.DiceOne = 0
		}
		rv[i] = result
	}

	for i, v := range data {
		fmt.Println("player ", i, v)
	}
	fmt.Println()

	return rv, nil
}

func (ps *playerServ) AutoRunGame(dataPlayer []player.Player) error {
	countTurn := len(dataPlayer) - 1
	for countTurn > 0 {
		tempDataPlayer, _ := ps.GameRules(dataPlayer)
		for _, v := range tempDataPlayer {
			if len(v.Dice) == 0 {
				countTurn--
			}
		}
		if countTurn > 0 {
			countTurn = len(dataPlayer) - 1
		}
	}
	ps.TheWinner(dataPlayer)
	return nil
}

func (ps *playerServ) TheWinner(dataPlayer []player.Player) error {

	for i := 0; i < len(dataPlayer); i++ {
		var index int
		max := dataPlayer[0].Point
		posPlay := dataPlayer[0].Player
		for j := 1; j < len(dataPlayer); j++ {
			// fmt.Println("sw", max, dataPlayer[j].Point)
			if dataPlayer[j].Point > max {
				dataPlayer[index].Point = dataPlayer[j].Point
				dataPlayer[index].Player = dataPlayer[j].Player
				dataPlayer[j].Player = posPlay
				dataPlayer[j].Point = max

			}
			index++

		}
	}

	// fmt.Println(dataPlayer)
	for i, v := range dataPlayer {
		fmt.Println("Juara", i+1, "Player ke", v.Player, "dengan Poin ", v.Point)
	}
	return nil
}
