package main

type Player struct {
	hp, mana, armor int
	spells          []*Spell
}

type Boss struct {
	hp, damage int
}

type Spell struct {
	name    string
	mana    int
	damage  int
	heal    int
	effects int
}

type Effect struct {
	Duration, Remaining int
	Active              bool
}

func main() {

}
