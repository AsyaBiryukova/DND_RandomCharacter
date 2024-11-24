// Файл для симуляции бросков разных дайсов
package Dices

import (
	"math/rand"
)

func RollDiceD4() int { // Бросок 1d4
	return rand.Intn(4) + 1
}

func RollDiceD6() int { // Бросок 1d6
	return rand.Intn(6) + 1
}

func RollDiceD8() int { // Бросок 1d8
	return rand.Intn(8) + 1
}

func RollDiceD10() int { // Бросок 1d10
	return rand.Intn(10) + 1
}

func RollDiceD12() int { // Бросок 1d12
	return rand.Intn(12) + 1
}

func RollDiceD20() int { // Бросок 1d20
	return rand.Intn(20) + 1
}

func RollDiceFullD100() int { // Бросок 1d100
	return rand.Intn(100) + 1
}
