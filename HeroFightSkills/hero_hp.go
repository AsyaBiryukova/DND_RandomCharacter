package herofightskills

import (
	char "github.com/AsyaBiryukova/DND_RandomCharacter/HeroCharsAndSkills"
	random "github.com/AsyaBiryukova/DND_RandomCharacter/MainHeroInfo"
)

func HeroHP() int {
	hero_class := random.RandomClass()
	var hero_hp int

	switch {
	case hero_class == "Волшебник" || hero_class == "Чародей":
		hero_hp = 6 + char.Constitution_Modifiers()
	case hero_class == "Бард" || hero_class == "Друид" || hero_class == "Жрец" || hero_class == "Колдун" || hero_class == "Монах" || hero_class == "Плут":
		hero_hp = 8 + char.Constitution_Modifiers()
	case hero_class == "Воин" || hero_class == "Паладин" || hero_class == "Следопыт":
		hero_hp = 10 + char.Constitution_Modifiers()
	case hero_class == "Варвар":
		hero_hp = 12 + char.Constitution_Modifiers()
	}
	return hero_hp
}
