// Файл, хранящий информацию по Спасброскам/Испытаниям
package herocharsandskills

import (
	class "github.com/AsyaBiryukova/DND_RandomCharacter/MainHeroInfo"
)

func BaseSaves() map[string]int { //Функция, в которой хранятся базовые модификаторы спасбросков
	base_saves := map[string]int{
		"Сила":         Strenght_Modifiers(), // Все базовые Спасброски равны Модификатору одноимённой Характеристики
		"Ловкость":     Dexterity_Modifiers(),
		"Выносливость": Constitution_Modifiers(),
		"Интеллект":    Intelligense_Modifiers(),
		"Мудрость":     Wisdom_Modifiers(),
		"Харизма":      Charisma_Modifiers(),
	}
	return base_saves
}

func ClassUpdateSaves() map[string]int { // Функция повышает модификаторы в зависимости от Класса
	class_ubdate_saves := BaseSaves() // Вызов результата функции с базовыми Модификаторами Спасбросков
	hero_class := class.RandomClass() // Вызов результата функции с рандомным классом Героя
	// Перебор всех Классов и апдейт Спасбросков +2 от их умений
	// Бард: Ловкость, Харизма
	// Варвар: Сила, Выносливость
	// Воин: Сила, Выносливость
	// Волшебник: Интеллект, Мудрость
	// Друид: Интеллект, Мудрость
	// Жрец: Мудрость, Харизма
	// Колдун: Мудрость, Харизма
	// Монах: Сила, Ловкость
	// Паладин: Мудрость, Харизма
	// Плут: Ловкость, Интеллект
	// Следопыт: Сила, Ловкость
	// Чародей: Выносливость, Харизма
	switch {
	case hero_class == "Воин" || hero_class == "Варвар" || hero_class == "Монах" || hero_class == "Следопыт":
		class_ubdate_saves["Сила"] += 2
	case hero_class == "Бард" || hero_class == "Монах" || hero_class == "Плут" || hero_class == "Следопыт":
		class_ubdate_saves["Ловкость"] += 2
	case hero_class == "Варвар" || hero_class == "Воин" || hero_class == "Чародей":
		class_ubdate_saves["Выносливость"] += 2
	case hero_class == "Волшебник" || hero_class == "Друид" || hero_class == "Плут":
		class_ubdate_saves["Интеллект"] += 2
	case hero_class == "Волшебник" || hero_class == "Друид" || hero_class == "Жрец" || hero_class == "Колдун" || hero_class == "Паладин" || hero_class == "Чародей":
		class_ubdate_saves["Мудрость"] += 2
	case hero_class == "Бард" || hero_class == "Жрец" || hero_class == "Колдун" || hero_class == "Паладин":
		class_ubdate_saves["Харизма"] += 2
	}
	return class_ubdate_saves // Возврат результат выполнения функции
}
