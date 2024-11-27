// Файл, хранящий информацию по Навыкам Героя
package herocharsandskills

import (
	"math/rand"

	class "github.com/AsyaBiryukova/DND_RandomCharacter/MainHeroInfo"
)

func StrenhgtSkills() map[string]int { // Функция, определяющая чему будет равен Модификатор навыков, основанных на Cиле
	strenhgt_skills := map[string]int{ // Мапа, хранящая в себе названия навыков и значение Модификаторов
		"Атлетика": Strenght_Modifiers(), // "Навык": модификатор
	}
	return strenhgt_skills
}

func DexteritySkills() map[string]int { // Функция, определяющая чему будет равен модификатор навыков, основанных на Ловкости
	dexterity_skills := map[string]int{
		"Акробатика":   Dexterity_Modifiers(),
		"Ловкость рук": Dexterity_Modifiers(),
		"Скрытность":   Dexterity_Modifiers(),
	}
	return dexterity_skills
}

func IntelligenseSkills() map[string]int { // Функция, определяющая чему будет равен модификатор навыков, основанных на Интеллекте
	intelligense_skills := map[string]int{
		"История":       Intelligense_Modifiers(),
		"Магия":         Intelligense_Modifiers(),
		"Природа":       Intelligense_Modifiers(),
		"Расследование": Intelligense_Modifiers(),
		"Религия":       Intelligense_Modifiers(),
	}
	return intelligense_skills
}

func WisdomSkills() map[string]int { // Функция, определяющая чему будет равен модификатор навыков, основанных на Мудрости
	wisdom_skills := map[string]int{
		"Внимание":         Wisdom_Modifiers(),
		"Выживание":        Wisdom_Modifiers(),
		"Дрессировка":      Wisdom_Modifiers(),
		"Медицина":         Wisdom_Modifiers(),
		"Проницательность": Wisdom_Modifiers(),
	}
	return wisdom_skills
}

func CharismaSkills() map[string]int { // Функция, определяющая чему будет равен модификатор навыков, основанных на Харизме
	charisma_skills := map[string]int{
		"Запугивание": Charisma_Modifiers(),
		"Исполнение":  Charisma_Modifiers(),
		"Обман":       Charisma_Modifiers(),
		"Убеждение":   Charisma_Modifiers(),
	}
	return charisma_skills
}

func ClassUpdateSkills() map[string]int { // Функция, прокачивающая Модификаторы навыка по Классу
	hero_class := class.RandomClass() // Вызов рандомного Класса
	class_update_skills := map[string]int{
		"Акробатика":       Dexterity_Modifiers(),
		"Атлетика":         Strenght_Modifiers(),
		"Внимание":         Wisdom_Modifiers(),
		"Выживание":        Wisdom_Modifiers(),
		"Дрессировка":      Wisdom_Modifiers(),
		"Запугивание":      Charisma_Modifiers(),
		"Исполнение":       Charisma_Modifiers(),
		"История":          Intelligense_Modifiers(),
		"Ловкость рук":     Dexterity_Modifiers(),
		"Магия":            Intelligense_Modifiers(),
		"Медицина":         Wisdom_Modifiers(),
		"Обман":            Charisma_Modifiers(),
		"Природа":          Intelligense_Modifiers(),
		"Проницательность": Wisdom_Modifiers(),
		"Расследование":    Intelligense_Modifiers(),
		"Религия":          Intelligense_Modifiers(),
		"Скрытность":       Dexterity_Modifiers(),
		"Убеждение":        Charisma_Modifiers(),
	}
	if hero_class == "Бард" { // Поднятие модификаторов навыков от Барда (3 случайные)
		bard_skills := class_update_skills
		for i := 0; i < 3; i++ {
			skill_index := rand.Intn(len(bard_skills))
			bard_skills_slice := make([]string, 0)
			for bard_key := range bard_skills {
				bard_skills_slice = append(bard_skills_slice, bard_key)
			}
			bard_skills[bard_skills_slice[skill_index]] += 2
		}
		return bard_skills

	} else if hero_class == "Варвар" { // Поднятие модификаторов навыков от варвара
		barbarian_skills := map[string]int{
			"Атлетика":    Dexterity_Modifiers(),
			"Внимание":    Wisdom_Modifiers(),
			"Выживание":   Wisdom_Modifiers(),
			"Дрессировка": Wisdom_Modifiers(),
			"Запугивание": Charisma_Modifiers(),
			"Природа":     Intelligense_Modifiers(),
		}
		for i := 0; i < 2; i++ {
			barbarian_skill_index := rand.Intn(len(barbarian_skills))
			barbarian_skills_slice := make([]string, 0)
			for barbarian_key := range barbarian_skills {
				barbarian_skills_slice = append(barbarian_skills_slice, barbarian_key)
			}
			barbarian_skills[barbarian_skills_slice[barbarian_skill_index]] += 2
		}
		barbarian_skills["Акробатика"] = class_update_skills["Акробатика"]
		barbarian_skills["Исполнение"] = class_update_skills["Исполнение"]
		barbarian_skills["История"] = class_update_skills["История"]
		barbarian_skills["Ловкость рук"] = class_update_skills["Ловкость рук"]
		barbarian_skills["Магия"] = class_update_skills["Магия"]
		barbarian_skills["Медицина"] = class_update_skills["Медицина"]
		barbarian_skills["Обман"] = class_update_skills["Обман"]
		barbarian_skills["Проницательность"] = class_update_skills["Проницательность"]
		barbarian_skills["Расследование"] = class_update_skills["Расследование"]
		barbarian_skills["Религия"] = class_update_skills["Религия"]
		barbarian_skills["Скрытность"] = class_update_skills["Скрытность"]
		barbarian_skills["Убеждение"] = class_update_skills["Убеждение"]
		return barbarian_skills
	} else if hero_class == "Воин" { // От Воина
		warrior_skills := map[string]int{
			"Акробатика":       Dexterity_Modifiers(),
			"Атлетика":         Strenght_Modifiers(),
			"Внимание":         Wisdom_Modifiers(),
			"Выживание":        Wisdom_Modifiers(),
			"Дрессировка":      Wisdom_Modifiers(),
			"Запугивание":      Charisma_Modifiers(),
			"История":          Intelligense_Modifiers(),
			"Проницательность": Wisdom_Modifiers(),
		}

		for i := 0; i < 2; i++ {
			warrior_skill_index := rand.Intn(len(warrior_skills))
			warrior_skills_slice := make([]string, 0)
			for warrior_key := range warrior_skills {
				warrior_skills_slice = append(warrior_skills_slice, warrior_key)
			}
			warrior_skills[warrior_skills_slice[warrior_skill_index]] += 2
		}
		warrior_skills["Исполнение"] = class_update_skills["Исполнение"]
		warrior_skills["Ловкость рук"] = class_update_skills["Ловкость рук"]
		warrior_skills["Магия"] = class_update_skills["Магия"]
		warrior_skills["Медицина"] = class_update_skills["Медицина"]
		warrior_skills["Обман"] = class_update_skills["Обман"]
		warrior_skills["Природа"] = class_update_skills["Природа"]
		warrior_skills["Расследование"] = class_update_skills["Расследование"]
		warrior_skills["Религия"] = class_update_skills["Религия"]
		warrior_skills["Скрытность"] = class_update_skills["Скрытность"]
		warrior_skills["Убеждение"] = class_update_skills["Убеждение"]
		return warrior_skills
	} else if hero_class == "Волшебник" { // От Волшебника
		wizard_skills := map[string]int{
			"История":          Intelligense_Modifiers(),
			"Магия":            Intelligense_Modifiers(),
			"Медицина":         Wisdom_Modifiers(),
			"Проницательность": Wisdom_Modifiers(),
			"Расследование":    Intelligense_Modifiers(),
			"Религия":          Intelligense_Modifiers(),
		}
		for i := 0; i < 2; i++ {
			wizard_skill_index := rand.Intn(len(wizard_skills))
			wizard_skills_slice := make([]string, 0)
			for wizard_key := range wizard_skills {
				wizard_skills_slice = append(wizard_skills_slice, wizard_key)
			}
			wizard_skills[wizard_skills_slice[wizard_skill_index]] += 2
		}
		wizard_skills["Акробатика"] = class_update_skills["Акробатика"]
		wizard_skills["Атлетика"] = class_update_skills["Атлетика"]
		wizard_skills["Внимание"] = class_update_skills["Внимание"]
		wizard_skills["Выживание"] = class_update_skills["Выживание"]
		wizard_skills["Дрессировка"] = class_update_skills["Дрессировка"]
		wizard_skills["Запугивание"] = class_update_skills["Запугивание"]
		wizard_skills["Исполнение"] = class_update_skills["Исполнение"]
		wizard_skills["Ловкость рук"] = class_update_skills["Ловкость рук"]
		wizard_skills["Обман"] = class_update_skills["Обман"]
		wizard_skills["Природа"] = class_update_skills["Природа"]
		wizard_skills["Скрытность"] = class_update_skills["Скрытность"]
		wizard_skills["Убеждение"] = class_update_skills["Убеждение"]
		return wizard_skills
	} else if hero_class == "Друид" { // От Друида
		druid_skills := map[string]int{
			"Внимание":         Wisdom_Modifiers(),
			"Выживание":        Wisdom_Modifiers(),
			"Дрессировка":      Wisdom_Modifiers(),
			"Магия":            Intelligense_Modifiers(),
			"Медицина":         Wisdom_Modifiers(),
			"Природа":          Intelligense_Modifiers(),
			"Проницательность": Wisdom_Modifiers(),
			"Религия":          Intelligense_Modifiers(),
		}
		for i := 0; 1 < 2; i++ {
			druid_skill_index := rand.Intn(len(druid_skills))
			druid_skills_slice := make([]string, 0)
			for druid_key := range druid_skills {
				druid_skills_slice = append(druid_skills_slice, druid_key)
			}
			druid_skills[druid_skills_slice[druid_skill_index]] += 2
		}
		druid_skills["Акробатика"] = class_update_skills["Акробатика"]
		druid_skills["Атлетика"] = class_update_skills["Атлетика"]
		druid_skills["Запугивание"] = class_update_skills["Запугивание"]
		druid_skills["Исполнение"] = class_update_skills["Исполнение"]
		druid_skills["История"] = class_update_skills["История"]
		druid_skills["Ловкость рук"] = class_update_skills["Ловкость рук"]
		druid_skills["Обман"] = class_update_skills["Обман"]
		druid_skills["Расследование"] = class_update_skills["Расследование"]
		druid_skills["Скрытность"] = class_update_skills["Убеждение"]
		return druid_skills
	} else if hero_class == "Жрец" { // От Жреца
		cleric_skills := map[string]int{
			"Внимание":         Wisdom_Modifiers(),
			"Выживание":        Wisdom_Modifiers(),
			"Дрессировка":      Wisdom_Modifiers(),
			"Магия":            Intelligense_Modifiers(),
			"Медицина":         Wisdom_Modifiers(),
			"Природа":          Intelligense_Modifiers(),
			"Проницательность": Wisdom_Modifiers(),
			"Религия":          Intelligense_Modifiers(),
		}
		for i := 0; 1 < 2; i++ {
			cleric_skill_index := rand.Intn(len(cleric_skills))
			cleric_skills_slice := make([]string, 0)
			for cleric_key := range cleric_skills {
				cleric_skills_slice = append(cleric_skills_slice, cleric_key)
			}
			cleric_skills[cleric_skills_slice[cleric_skill_index]] += 2
		}
		cleric_skills["Акробатика"] = class_update_skills["Акробатика"]
		cleric_skills["Атлетика"] = class_update_skills["Атлетика"]
		cleric_skills["Запугивание"] = class_update_skills["Запугивание"]
		cleric_skills["Исполнение"] = class_update_skills["Исполнение"]
		cleric_skills["История"] = class_update_skills["История"]
		cleric_skills["Ловкость рук"] = class_update_skills["Ловкость рук"]
		cleric_skills["Обман"] = class_update_skills["Обман"]
		cleric_skills["Расследование"] = class_update_skills["Расследование"]
		cleric_skills["Скрытность"] = class_update_skills["Скрытность"]
		cleric_skills["Убеждение"] = class_update_skills["Убеждение"]
		return cleric_skills

	} else if hero_class == "Колдун" { // От Чародея
		warloсk_skills := map[string]int{
			"Запугивание":   Charisma_Modifiers(),
			"История":       Intelligense_Modifiers(),
			"Магия":         Intelligense_Modifiers(),
			"Обман":         Charisma_Modifiers(),
			"Природа":       Intelligense_Modifiers(),
			"Расследование": Intelligense_Modifiers(),
			"Религия":       Intelligense_Modifiers(),
		}
		for i := 0; 1 < 2; i++ {
			warloсk_skill_index := rand.Intn(len(warloсk_skills))
			warloсk_skills_slice := make([]string, 0)
			for warloсk_key := range warloсk_skills {
				warloсk_skills_slice = append(warloсk_skills_slice, warloсk_key)
			}
			warloсk_skills[warloсk_skills_slice[warloсk_skill_index]] += 2
		}
		warloсk_skills["Акробатика"] = class_update_skills["Акробатика"]
		warloсk_skills["Атлетика"] = class_update_skills["Атлетика"]
		warloсk_skills["Внимание"] = class_update_skills["Внимание"]
		warloсk_skills["Выживание"] = class_update_skills["Выживание"]
		warloсk_skills["Дрессировка"] = class_update_skills["Дрессировка"]
		warloсk_skills["Исполнение"] = class_update_skills["Исполнение"]
		warloсk_skills["Ловкость рук"] = class_update_skills["Ловкость рук"]
		warloсk_skills["Медицина"] = class_update_skills["Медицина"]
		warloсk_skills["Проницательность"] = class_update_skills["Проницательность"]
		warloсk_skills["Скрытность"] = class_update_skills["Скрытность"]
		warloсk_skills["Убеждение"] = class_update_skills["Убеждение"]
		return warloсk_skills

	} else if hero_class == "Монах" { // От Монаха
		monk_skills := map[string]int{
			"Акробатика":       Dexterity_Modifiers(),
			"Атлетика":         Strenght_Modifiers(),
			"История":          Intelligense_Modifiers(),
			"Проницательность": Wisdom_Modifiers(),
			"Религия":          Intelligense_Modifiers(),
			"Скрытность":       Dexterity_Modifiers(),
		}
		for i := 0; 1 < 2; i++ {
			monk_skill_index := rand.Intn(len(monk_skills))
			monk_skills_slice := make([]string, 0)
			for monk_key := range monk_skills {
				monk_skills_slice = append(monk_skills_slice, monk_key)
			}
			monk_skills[monk_skills_slice[monk_skill_index]] += 2
		}
		monk_skills["Внимание"] = class_update_skills["Внимание"]
		monk_skills["Выживание"] = class_update_skills["Выживание"]
		monk_skills["Дрессировка"] = class_update_skills["Дрессировка"]
		monk_skills["Запугивание"] = class_update_skills["Запугивание"]
		monk_skills["Исполнение"] = class_update_skills["Исполнение"]
		monk_skills["Ловкость рук"] = class_update_skills["Ловкость рук"]
		monk_skills["Магия"] = class_update_skills["Магия"]
		monk_skills["Медицина"] = class_update_skills["Медицина"]
		monk_skills["Обман"] = class_update_skills["Обман"]
		monk_skills["Природа"] = class_update_skills["Природа"]
		monk_skills["Расследование"] = class_update_skills["Расследование"]
		monk_skills["Убеждение"] = class_update_skills["Убеждение"]
		return monk_skills

	} else if hero_class == "Паладин" { // От Паладина
		paladin_skills := map[string]int{
			"Атлетика":         Strenght_Modifiers(),
			"Запугивание":      Charisma_Modifiers(),
			"Медицина":         Wisdom_Modifiers(),
			"Проницательность": Wisdom_Modifiers(),
			"Религия":          Intelligense_Modifiers(),
			"Убеждение":        Charisma_Modifiers(),
		}
		for i := 0; 1 < 2; i++ {
			paladin_skill_index := rand.Intn(len(paladin_skills))
			paladin_skills_slice := make([]string, 0)
			for paladin_key := range paladin_skills {
				paladin_skills_slice = append(paladin_skills_slice, paladin_key)
			}
			paladin_skills[paladin_skills_slice[paladin_skill_index]] += 2
		}
		paladin_skills["Акробатика"] = class_update_skills["Акробатика"]
		paladin_skills["Внимание"] = class_update_skills["Внимание"]
		paladin_skills["Выживание"] = class_update_skills["Выживание"]
		paladin_skills["Дрессировка"] = class_update_skills["Дрессировка"]
		paladin_skills["Исполнение"] = class_update_skills["Исполнение"]
		paladin_skills["История"] = class_update_skills["История"]
		paladin_skills["Ловкость рук"] = class_update_skills["Ловкость рук"]
		paladin_skills["Магия"] = class_update_skills["Магия"]
		paladin_skills["Обман"] = class_update_skills["Обман"]
		paladin_skills["Природа"] = class_update_skills["Природа"]
		paladin_skills["Расследование"] = class_update_skills["Расследование"]
		paladin_skills["Скрытность"] = class_update_skills["Скрытность"]
		return paladin_skills

	} else if hero_class == "Плут" { // От Плута
		crook_skills := map[string]int{
			"Акробатика":       Dexterity_Modifiers(),
			"Атлетика":         Strenght_Modifiers(),
			"Внимание":         Wisdom_Modifiers(),
			"Запугивание":      Charisma_Modifiers(),
			"Исполнение":       Charisma_Modifiers(),
			"Ловкость рук":     Dexterity_Modifiers(),
			"Обман":            Charisma_Modifiers(),
			"Проницательность": Wisdom_Modifiers(),
			"Расследование":    Intelligense_Modifiers(),
			"Скрытность":       Dexterity_Modifiers(),
			"Убеждение":        Charisma_Modifiers(),
		}
		for i := 0; 1 < 2; i++ {
			crook_skill_index := rand.Intn(len(crook_skills))
			crook_skills_slice := make([]string, 0)
			for crook_key := range crook_skills {
				crook_skills_slice = append(crook_skills_slice, crook_key)
			}
			crook_skills[crook_skills_slice[crook_skill_index]] += 2
		}
		crook_skills["Выживание"] = class_update_skills["Выживание"]
		crook_skills["Дрессировка"] = class_update_skills["Дрессировка"]
		crook_skills["История"] = class_update_skills["История"]
		crook_skills["Магия"] = class_update_skills["Магия"]
		crook_skills["Медицина"] = class_update_skills["Медицина"]
		crook_skills["Природа"] = class_update_skills["Природа"]
		crook_skills["Религия"] = class_update_skills["Религия"]
		return crook_skills
	} else if hero_class == "Следопыт" { // От Следопыта
		ranger_skills := map[string]int{
			"Атлетика":         Strenght_Modifiers(),
			"Внимание":         Wisdom_Modifiers(),
			"Выживание":        Wisdom_Modifiers(),
			"Дрессировка":      Wisdom_Modifiers(),
			"Природа":          Intelligense_Modifiers(),
			"Проницательность": Wisdom_Modifiers(),
			"Расследование":    Intelligense_Modifiers(),
			"Скрытность":       Dexterity_Modifiers(),
		}
		for i := 0; 1 < 2; i++ {
			ranger_skill_index := rand.Intn(len(ranger_skills))
			ranger_skills_slice := make([]string, 0)
			for ranger_key := range ranger_skills {
				ranger_skills_slice = append(ranger_skills_slice, ranger_key)
			}
			ranger_skills[ranger_skills_slice[ranger_skill_index]] += 2
		}
		ranger_skills["Акробатика"] = class_update_skills["Акробатика"]
		ranger_skills["Запугивание"] = class_update_skills["Запугивание"]
		ranger_skills["Исполнение"] = class_update_skills["Исполнение"]
		ranger_skills["История"] = class_update_skills["История"]
		ranger_skills["Ловкость рук"] = class_update_skills["Ловкость рук"]
		ranger_skills["Магия"] = class_update_skills["Магия"]
		ranger_skills["Медицина"] = class_update_skills["Медицина"]
		ranger_skills["Обман"] = class_update_skills["Обман"]
		ranger_skills["Религия"] = class_update_skills["Религия"]
		ranger_skills["Убеждение"] = class_update_skills["Убеждение"]
		return ranger_skills

	} else if hero_class == "Чародей" { // От Чародея
		sorcerer_skills := map[string]int{
			"Запугивание":      Charisma_Modifiers(),
			"Магия":            Wisdom_Modifiers(),
			"Обман":            Charisma_Modifiers(),
			"Проницательность": Wisdom_Modifiers(),
			"Религия":          Intelligense_Modifiers(),
			"Убеждение":        Charisma_Modifiers(),
		}
		for i := 0; 1 < 2; i++ {
			sorcerer_skill_index := rand.Intn(len(sorcerer_skills))
			sorcerer_skills_slice := make([]string, 0)
			for sorcerer_key := range sorcerer_skills {
				sorcerer_skills_slice = append(sorcerer_skills_slice, sorcerer_key)
			}
			sorcerer_skills[sorcerer_skills_slice[sorcerer_skill_index]] += 2
		}
		sorcerer_skills["Акробатика"] = class_update_skills["Акробатика"]
		sorcerer_skills["Атлетика"] = class_update_skills["Атлетика"]
		sorcerer_skills["Внимание"] = class_update_skills["Внимание"]
		sorcerer_skills["Выживание"] = class_update_skills["Выживание"]
		sorcerer_skills["Дрессировка"] = class_update_skills["Дрессировка"]
		sorcerer_skills["Исполнение"] = class_update_skills["Исполнение"]
		sorcerer_skills["История"] = class_update_skills["История"]
		sorcerer_skills["Ловкость рук"] = class_update_skills["Ловкость рук"]
		sorcerer_skills["Медицина"] = class_update_skills["Медицина"]
		sorcerer_skills["Природа"] = class_update_skills["Природа"]
		sorcerer_skills["Расследование"] = class_update_skills["Расследование"]
		sorcerer_skills["Скрытность"] = class_update_skills["Скрытность"]
		return sorcerer_skills
	}
	return class_update_skills
}
