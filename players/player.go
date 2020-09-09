package players

type Player struct {
	name   string
	symbol byte
}

func New(name string, symbol byte) *Player {
	return &Player{name, symbol}
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Symbol() byte {
	return p.symbol
}
