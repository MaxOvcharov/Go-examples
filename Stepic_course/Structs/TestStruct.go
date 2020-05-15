package main

import "fmt"

type Gun struct {
	On bool
	Ammo int
	Power int
}

func (g *Gun) Shoot() bool {
	if g.On != false {
		if g.Ammo > 0 {
			g.Ammo--
		} else {
			return !g.On
		}
	}

	return g.On

}

func (g *Gun) RideBike() bool {
	if g.On != false {
		if g.Power > 0 {
			g.Power--
		} else {
			return !g.On
		}
	}

	return g.On
}

func main() {
	ts := Gun{true, 12, 1}
	fmt.Println(ts.Shoot())
	fmt.Println(ts.RideBike())
	fmt.Println(ts)
}
