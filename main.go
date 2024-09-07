package main

import (
	"dicegame/internal/player/repository"
	"dicegame/internal/player/service"
	"fmt"
)

func main() {
	playMood := repository.NewPlayerMod()
	playServ := service.NewPlayerServ(playMood)

	var dice int
	var player int
	fmt.Print("Masukkan Jumlah Pemain: ")
	fmt.Scan(&player)
	fmt.Print("Masukkan Jumlah Dadu yang dimainkan: ")
	fmt.Scan(&dice)

	dataPlayer := playServ.PlayerPos(player, dice)
	playServ.AutoRunGame(dataPlayer)

}
