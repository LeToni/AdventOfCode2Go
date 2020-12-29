package main

type Wizard struct {
	hp, mana, armor int
}

type Boss struct {
	hp, damage int
}

type Spell struct {
	name   string
	mana   int
	damage int
	heal   int
	effect string
}

type Effect struct {
	duration, remaining int
	active              bool
}

func (w *Wizard) isDead() bool {
	// wizard can be also considered dead if he has not enough mana to cast spell
	if w.hp <= 0 || w.mana < 53 {
		return true
	} else {
		return false
	}
}

func (b *Boss) isDead() bool {
	if b.hp <= 0 {
		return true
	} else {
		return false
	}
}

func (boss *Boss) attacks(wizard *Wizard) {
	if boss.damage-wizard.armor < 1 {
		wizard.hp = wizard.hp - 1
	} else {
		wizard.hp = wizard.hp - boss.damage
	}
}

var (
	spells = []*Spell{
		{name: "Magic Missole", mana: 53, damage: 4, heal: 0, effect: "None"},
		{name: "Drain", mana: 73, damage: 2, heal: 2, effect: "None"},
		{name: "Shield", mana: 113, damage: 0, heal: 0, effect: "Shield"},
		{name: "Poison", mana: 173, damage: 3, heal: 0, effect: "Poison"},
		{name: "Recharge", mana: 229, damage: 0, heal: 0, effect: "Recharge"},
	}

	effects = map[string]*Effect{
		"Shield":   {duration: 6, remaining: 6, active: false},
		"Poison":   {duration: 6, remaining: 6, active: false},
		"Recharge": {duration: 5, remaining: 5, active: false},
	}

	wizard Wizard
	boss   Boss
)

func main() {

}
