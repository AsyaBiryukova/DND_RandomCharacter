// Файл с Модификаторами харктеристик, которые дальше пойдут в Навыки и Испытания
package herocharsandskills

func Strenght_Modifiers() int { // Функция, определяющая модификатор Cилы
	hero_characteristics := RaceUpdateCharacteristic() // Вызов результата функции, определяющей значение характеристики
	var strenght_modifier int                          // "Пустая" переменная, в которую передаётся значение модификатора Силы

	switch { // Перебор всех возможных значений Силы и подбор к ним модификаторов из Таблицы (Книга Игрока стр. 13)
	case hero_characteristics["Сила"] == 3:
		strenght_modifier = -4
	case hero_characteristics["Сила"] == 4 || hero_characteristics["Сила"] == 5:
		strenght_modifier = -3
	case hero_characteristics["Сила"] == 6 || hero_characteristics["Сила"] == 7:
		strenght_modifier = -2
	case hero_characteristics["Сила"] == 8 || hero_characteristics["Сила"] == 9:
		strenght_modifier = -1
	case hero_characteristics["Сила"] == 10 || hero_characteristics["Сила"] == 11:
		strenght_modifier = +0
	case hero_characteristics["Сила"] == 12 || hero_characteristics["Сила"] == 13:
		strenght_modifier = +1
	case hero_characteristics["Сила"] == 14 || hero_characteristics["Сила"] == 15:
		strenght_modifier = +2
	case hero_characteristics["Сила"] == 16 || hero_characteristics["Сила"] == 17:
		strenght_modifier = +3
	case hero_characteristics["Сила"] == 18:
		strenght_modifier = +4
	}
	return strenght_modifier // Возврат итогого модификатора Силы
}

func Dexterity_Modifiers() int { // Функция, определяющая модификатор Ловоксти
	hero_characteristics := RaceUpdateCharacteristic()
	var dexterity_modifier int

	switch {
	case hero_characteristics["Ловкость"] == 3:
		dexterity_modifier = -4
	case hero_characteristics["Ловкость"] == 4 || hero_characteristics["Ловкость"] == 5:
		dexterity_modifier = -3
	case hero_characteristics["Ловкость"] == 6 || hero_characteristics["Ловкость"] == 7:
		dexterity_modifier = -2
	case hero_characteristics["Ловкость"] == 8 || hero_characteristics["Ловкость"] == 9:
		dexterity_modifier = -1
	case hero_characteristics["Ловкость"] == 10 || hero_characteristics["Ловкость"] == 11:
		dexterity_modifier = +0
	case hero_characteristics["Ловкость"] == 12 || hero_characteristics["Ловкость"] == 13:
		dexterity_modifier = +1
	case hero_characteristics["Ловкость"] == 14 || hero_characteristics["Ловкость"] == 15:
		dexterity_modifier = +2
	case hero_characteristics["Ловкость"] == 16 || hero_characteristics["Ловкость"] == 17:
		dexterity_modifier = +3
	case hero_characteristics["Ловкость"] == 18:
		dexterity_modifier = +4
	}

	return dexterity_modifier
}

func Constitution_Modifiers() int { // Функция, определяющая модификатор Выносливости
	hero_characteristics := RaceUpdateCharacteristic()
	var constitution_modifier int

	switch {
	case hero_characteristics["Выносливость"] == 3:
		constitution_modifier = -4
	case hero_characteristics["Выносливость"] == 4 || hero_characteristics["Выносливость"] == 5:
		constitution_modifier = -3
	case hero_characteristics["Выносливость"] == 6 || hero_characteristics["Выносливость"] == 7:
		constitution_modifier = -2
	case hero_characteristics["Выносливость"] == 8 || hero_characteristics["Выносливость"] == 9:
		constitution_modifier = -1
	case hero_characteristics["Выносливость"] == 10 || hero_characteristics["Выносливость"] == 11:
		constitution_modifier = +0
	case hero_characteristics["Выносливость"] == 12 || hero_characteristics["Выносливость"] == 13:
		constitution_modifier = +1
	case hero_characteristics["Выносливость"] == 14 || hero_characteristics["Выносливость"] == 15:
		constitution_modifier = +2
	case hero_characteristics["Выносливость"] == 16 || hero_characteristics["Выносливость"] == 17:
		constitution_modifier = +3
	case hero_characteristics["Выносливость"] == 18:
		constitution_modifier = +4
	}
	return constitution_modifier
}

func Intelligense_Modifiers() int { // Функция, определяющая модификатор Интеллекта
	hero_characteristics := RaceUpdateCharacteristic()
	var intelligense_modifier int

	switch {
	case hero_characteristics["Интеллект"] == 3:
		intelligense_modifier = -4
	case hero_characteristics["Интеллект"] == 4 || hero_characteristics["Интеллект"] == 5:
		intelligense_modifier = -3
	case hero_characteristics["Интеллект"] == 6 || hero_characteristics["Интеллект"] == 7:
		intelligense_modifier = -2
	case hero_characteristics["Интеллект"] == 8 || hero_characteristics["Интеллект"] == 9:
		intelligense_modifier = -1
	case hero_characteristics["Интеллект"] == 10 || hero_characteristics["Интеллект"] == 11:
		intelligense_modifier = +0
	case hero_characteristics["Интеллект"] == 12 || hero_characteristics["Интеллект"] == 13:
		intelligense_modifier = +1
	case hero_characteristics["Интеллект"] == 14 || hero_characteristics["Интеллект"] == 15:
		intelligense_modifier = +2
	case hero_characteristics["Интеллект"] == 16 || hero_characteristics["Интеллект"] == 17:
		intelligense_modifier = +3
	case hero_characteristics["Интеллект"] == 18:
		intelligense_modifier = +4
	}
	return intelligense_modifier
}

func Wisdom_Modifiers() int { // Функция, определяющая модификатор Мудрости
	hero_characteristics := RaceUpdateCharacteristic()
	var wisdom_modifier int

	switch {
	case hero_characteristics["Мудрость"] == 3:
		wisdom_modifier = -4
	case hero_characteristics["Мудрость"] == 4 || hero_characteristics["Мудрость"] == 5:
		wisdom_modifier = -3
	case hero_characteristics["Мудрость"] == 6 || hero_characteristics["Мудрость"] == 7:
		wisdom_modifier = -2
	case hero_characteristics["Мудрость"] == 8 || hero_characteristics["Мудрость"] == 9:
		wisdom_modifier = -1
	case hero_characteristics["Мудрость"] == 10 || hero_characteristics["Мудрость"] == 11:
		wisdom_modifier = +0
	case hero_characteristics["Мудрость"] == 12 || hero_characteristics["Мудрость"] == 13:
		wisdom_modifier = +1
	case hero_characteristics["Мудрость"] == 14 || hero_characteristics["Мудрость"] == 15:
		wisdom_modifier = +2
	case hero_characteristics["Мудрость"] == 16 || hero_characteristics["Мудрость"] == 17:
		wisdom_modifier = +3
	case hero_characteristics["Мудрость"] == 18:
		wisdom_modifier = +4
	}
	return wisdom_modifier
}

func Charisma_Modifiers() int { // Функция, определяющая модификатор Харизмы
	hero_characteristics := RaceUpdateCharacteristic()
	var charisma_modifier int

	switch {
	case hero_characteristics["Харизма"] == 3:
		charisma_modifier = -4
	case hero_characteristics["Харизма"] == 4 || hero_characteristics["Харизма"] == 5:
		charisma_modifier = -3
	case hero_characteristics["Харизма"] == 6 || hero_characteristics["Харизма"] == 7:
		charisma_modifier = -2
	case hero_characteristics["Харизма"] == 8 || hero_characteristics["Харизма"] == 9:
		charisma_modifier = -1
	case hero_characteristics["Харизма"] == 10 || hero_characteristics["Харизма"] == 11:
		charisma_modifier = +0
	case hero_characteristics["Харизма"] == 12 || hero_characteristics["Харизма"] == 13:
		charisma_modifier = +1
	case hero_characteristics["Харизма"] == 14 || hero_characteristics["Харизма"] == 15:
		charisma_modifier = +2
	case hero_characteristics["Харизма"] == 16 || hero_characteristics["Харизма"] == 17:
		charisma_modifier = +3
	case hero_characteristics["Харизма"] == 18:
		charisma_modifier = +4
	}
	return charisma_modifier
}
