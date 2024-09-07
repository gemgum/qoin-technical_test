package main

import (
	"dicegame/internal/player/repository"
	"dicegame/internal/player/service"
)

func main() {
	playMood := repository.NewPlayerMod()
	playServ := service.NewPlayerServ(playMood)
	dataPlayer := playServ.PlayerPos(3, 4)
	// data, _ := playServ.GameRules(dataPlayer)
	playServ.AutoRunGame(dataPlayer)

}
