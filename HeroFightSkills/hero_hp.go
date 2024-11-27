package herofightskills

import (
	"strconv"

	char "github.com/AsyaBiryukova/DND_RandomCharacter/HeroCharsAndSkills"
)

func Bard_HP() string { // Функция, определяющая здоровье героя
	bard_conctitution_modifier := char.Constitution_Modifiers() // Вызов итогового модификатора Выносливости
	bard_hill_dice := 8                                         // Кость здоровья барда 1d8
	max_bard_hp := bard_conctitution_modifier + bard_hill_dice  // Максимальное здоровья Барда на 1 уровне = максимальное значение кости здоровья + модификатор Выносливости

	bard_hp_info := "Ваше максимальное здоровье: " + strconv.Itoa(max_bard_hp) // Сообщение с информацией здоровья
	return bard_hp_info
}
