package charactercard

func CharacterInfo() map[string]string {
	char_info := map[string]string{
		"Класс":         "",
		"Раса":          "",
		"Происхождение": "",
		"Мировоззрение": "",
	}
	return char_info
}

func CharacterPersonality() map[string]string {
	char_pers := map[string]string{
		"Черты Характера": "",
		"Идеалы":          "",
		"Привязанности":   "",
		"Слабости":        "",
	}
	return char_pers
}

func FightAbilitys() map[string]int {
	fight_ab := map[string]int{
		"Класс Брони": 0,
		"Инициатива":  0,
		"Скорость":    0,
	}
	return fight_ab
}

func OtherAbilitys() map[string]string {
	oth_ab := map[string]string{
		"Языки":       "",
		"Броня":       "",
		"Оружие":      "",
		"Инструменты": "",
	}
	return oth_ab
}
