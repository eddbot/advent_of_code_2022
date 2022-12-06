package main

type instruction struct {
	move int
	from int
	to   int
}

func (i *instruction) populate() (move, to, from int) {
	return i.move, i.to, i.from
}
