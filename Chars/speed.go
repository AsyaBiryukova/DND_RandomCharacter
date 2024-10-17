package chars

import (
	r "github.com/AsyaBiryukova/DND_RandomCharacter/Randoms"
)

func HeroSpeed() int {
	race := r.RandomRace()
	var base_speed int

	if race == "Холмовой Дварф" || race == "Горный Дварф" || race == "Крепкий Полурослик" || race == "Легконогий Полурослик" || race == "Лесной Гном" || race == "Скальный Гном" {
		base_speed = 25
	}
	if race == "Человек" || race == "Высший Эльф" || race == "Дроу" || race == "Драконорождённый" || race == "Полуорк" || race == "Полуэльф" || race == "Тифлинг" {
		base_speed = 30
	}
	if race == "Лесной Эльф" {
		base_speed = 35
	}
	return base_speed
}
