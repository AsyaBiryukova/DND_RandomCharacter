package Dices

import (
	"math/rand"
)

func RollDiceD4() int {
	return rand.Intn(4) + 1
}

func RollDiceD6() int {
	return rand.Intn(6) + 1
}

func RollDiceD8() int {
	return rand.Intn(8) + 1
}

func RollDiceD10() int {
	return rand.Intn(10) + 1
}

func RollDiceD12() int {
	return rand.Intn(12) + 1
}

func RollDiceD20() int {
	return rand.Intn(20) + 1
}

func RollDiceFullD100() int {
	return rand.Intn(100) + 1
}
