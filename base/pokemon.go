package base

type Stats struct{
	hp, attack, defense, specialAttack, specialDefence, speed int
}

type Move struct{}

type Pokemon struct{
	name string
	learnset []Move
	stats Stats
}
