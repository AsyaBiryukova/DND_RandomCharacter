package herootherskills

import (
	"math/rand"

	random "github.com/AsyaBiryukova/DND_RandomCharacter/MainHeroInfo"
)

func Languages() map[int]string {
	all_languages := map[int]string{
		1:  "Всеобщий",
		2:  "Великаний",
		3:  "Гномий",
		4:  "Дварфийский",
		5:  "Орочий",
		6:  "Полурослицкий",
		7:  "Глубинный",
		8:  "Демонический",
		9:  "Драконий",
		10: "Дьявольский",
		11: "Небесный",
		12: "Первозданный",
		13: "Подземный",
		14: "Сильван",
		15: "Эльфийский",
	}
	return all_languages
}

func DwarfLang() string {
	hero_race := random.RandomRace()
	langs := Languages()
	var total_languages string
	if hero_race == "Холмовой Дварф" || hero_race == "Горный Дварф" {
		total_languages = "Языки, которые вы знаете: " + langs[4] + langs[1]
	}
	return total_languages
}

func HalflingLang() string {
	hero_race := random.RandomRace()
	langs := Languages()
	var total_languages string
	if hero_race == "Крепкий Полурослик" || hero_race == "Легконогий Полурослик" {
		total_languages = "Языки, которые вы знаете: " + langs[6] + langs[1]
	}
	return total_languages
}

func HumanLang() string {
	hero_race := random.RandomRace()
	langs := Languages()
	var total_languages string

	if hero_race == "Человек" {
		human_langs := map[int]string{
			1:  langs[2],
			2:  langs[3],
			3:  langs[4],
			4:  langs[5],
			5:  langs[6],
			6:  langs[7],
			7:  langs[8],
			8:  langs[9],
			9:  langs[10],
			10: langs[11],
			11: langs[12],
			12: langs[13],
			13: langs[14],
			14: langs[15],
		}
		rl := rand.Intn(14) + 1
		human_bonusLang := human_langs[rl]

		total_languages = "Языки, которые вы знаете: " + langs[1] + human_bonusLang
	}
	return total_languages
}

func ElfLang() string {
	hero_race := random.RandomRace()
	langs := Languages()
	var total_languages string

	if hero_race == "Высший Эльф" || hero_race == "Полуэльф" {
		highElf_langs := map[int]string{
			1:  langs[2],
			2:  langs[3],
			3:  langs[4],
			4:  langs[5],
			5:  langs[6],
			6:  langs[7],
			7:  langs[8],
			8:  langs[9],
			9:  langs[10],
			10: langs[11],
			11: langs[12],
			12: langs[13],
			13: langs[14],
		}
		rl := rand.Intn(13) + 1
		highElf_bonusLang := highElf_langs[rl]

		total_languages = "Языки, которые вы знаете: " + langs[1] + langs[15] + highElf_bonusLang

	} else if hero_race == "Лесной Эльф" || hero_race == "Дроу" {
		total_languages = "Языки, которые вы знаете: " + langs[1] + langs[15]
	}
	return total_languages
}

func GnomeLang() string {
	hero_race := random.RandomRace()
	langs := Languages()
	var total_languages string

	if hero_race == "Лесной Гном" || hero_race == "Скальный Гном" {
		total_languages = "Языки, которые вы знаете: " + langs[1] + langs[3]
	}
	return total_languages
}

func DragonbornLang() string {
	hero_race := random.RandomRace()
	langs := Languages()
	var total_languages string

	if hero_race == "Драконорождённый" {
		total_languages = "Языки, которые вы знаете: " + langs[1] + langs[9]
	}
	return total_languages
}

func HalfOrcLang() string {
	hero_race := random.RandomRace()
	langs := Languages()
	var total_languages string

	if hero_race == "Полуорк" {
		total_languages = "Языки, которые вы знаете: " + langs[1] + langs[5]
	}
	return total_languages
}

func TieflingLang() string {
	hero_race := random.RandomRace()
	langs := Languages()
	var total_languages string

	if hero_race == "Тифлинг" {
		total_languages = "Языки, которые вы знаете: " + langs[1] + langs[9]
	}
	return total_languages
}
