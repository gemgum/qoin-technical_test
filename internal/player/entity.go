package player

type Player struct {
	Player  int
	Point   int
	Dice    []int
	DiceOne int
}

type PlayerMod interface {
	DiceNumbGenerator(min int, max int) (int, error)
}

type PlayerServ interface {
	PlayerPos(total int, diceRules int) []Player
	PlayerTurn(data []Player) ([]Player, error)
	GameRules(dataPlayer []Player) ([]Player, error)
	AutoRunGame(dataPlayer []Player) error
	TheWinner(dataPlayer []Player) error
}
