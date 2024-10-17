package chars

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/AsyaBiryukova/DND_RandomCharacter/Dices"

	r "github.com/AsyaBiryukova/DND_RandomCharacter/Randoms"
)

func BaseCharacteristics() map[string]int {

	strenght := make([]int, 4) //Устанавливаем базовые характеристики силы
	for i := range strenght {
		strenght[i] = Dices.RollDiceD6()
	}
	sort.Ints(strenght)
	lastStrenght := strenght[len(strenght)-3:]
	totalStrenght := 0
	for _, s_num := range lastStrenght {
		totalStrenght += s_num
	}

	dexterity := make([]int, 4) //Устанавливаем базовые характеристики ловкости
	for i := range dexterity {
		dexterity[i] = Dices.RollDiceD6()
	}
	sort.Ints(dexterity)
	lastDexterity := dexterity[len(dexterity)-3:]
	totalDexterity := 0
	for _, d_num := range lastDexterity {
		totalDexterity += d_num
	}

	constitution := make([]int, 4) //Устанавливаем базовые характеристики выносливости
	for i := range constitution {
		constitution[i] = Dices.RollDiceD6()
	}
	sort.Ints(constitution)
	lastConstitution := constitution[len(constitution)-3:]
	totalConstitution := 0
	for _, c_num := range lastConstitution {
		totalConstitution += c_num
	}

	intelligense := make([]int, 4) //Устанавливаем базовые характеристики интелекта
	for i := range intelligense {
		intelligense[i] = Dices.RollDiceD6()
	}
	sort.Ints(intelligense)
	lastIntelligense := intelligense[len(intelligense)-3:]
	totalIntelligense := 0
	for _, i_num := range lastIntelligense {
		totalIntelligense += i_num
	}

	wisdom := make([]int, 4) //Устанавливаем базовые характеристики мудрость
	for i := range wisdom {
		wisdom[i] = Dices.RollDiceD6()
	}
	sort.Ints(wisdom)
	lastwisdom := wisdom[len(wisdom)-3:]
	totalIwisdom := 0
	for _, i_num := range lastwisdom {
		totalIwisdom += i_num
	}

	charisma := make([]int, 4) //Устанавливаем базовые характеристики харизмы
	for i := range charisma {
		charisma[i] = Dices.RollDiceD6()
	}
	sort.Ints(charisma)
	lastcharisma := charisma[len(charisma)-3:]
	totalIcharisma := 0
	for _, i_num := range lastcharisma {
		totalIcharisma += i_num
	}

	// заносим значения полученных базовых характеристик в мапу
	basecharacteristics := map[string]int{
		"Сила":         totalStrenght,
		"Ловкость":     totalDexterity,
		"Выносливость": totalConstitution,
		"Интеллект":    totalIntelligense,
		"Мудрость":     totalIwisdom,
		"Харизма":      totalIcharisma,
	}

	return basecharacteristics // возвращаем конечную мапу
}

func RaceUpdateCharacteristic() map[string]int {
	race := r.RandomRace()
	fmt.Println("Ваша раса", race)

	ch := BaseCharacteristics()
	fmt.Println("Ваши базовые значения характеристик:", ch) // для удобства выводим базовые

	// в зависимости от рандомной рассы - накидываем сверху какой-то модификатор
	if race == "Горный дварф" || race == "Полуорк" || race == "Драконорождённый" {
		ch["Сила"] += 2
	}
	if race == "Человек" {
		ch["Сила"] += 1
	}
	if race == "Крепкий полурослик" || race == "Легконогий полурослик" || race == "Высший Эльф" || race == "Лесной эльф" || race == "Дроу" {
		ch["Ловкость"] += 2
	}
	if race == "Человек" || race == "Лесной гном" {
		ch["Ловкость"] += 1
	}
	if race == "Холмовой дварф" || race == "Горный дварф" {
		ch["Выносливость"] += 2
	}
	if race == "Крепкий полурослик" || race == "Человек" || race == "Скальный гном" || race == "Полуорк" {
		ch["Выносливость"] += 1
	}
	if race == "Лесной гном" || race == "Скальный гном" {
		ch["Интеллект"] += 2
	}
	if race == "Человек" || race == "Высший эльф" || race == "Тифлинг" {
		ch["Интеллект"] += 1
	}
	if race == "Холмовой дварф" || race == "Человек" || race == "Лесной эльф" {
		ch["Мудрость"] += 1
	}
	if race == "Полуэльф" || race == "Тифлинг" {
		ch["Харизма"] += 2
	}
	if race == "Легконогий полурослик" || race == "Человек" || race == "Дроу" || race == "Драконорождённый" {
		ch["Харизма"] += 1
	}
	if race == "Полуэльф" {
		halfElf := map[string]int{
			"Сила":         ch["Сила"],
			"Ловкость":     ch["Ловкость"],
			"Выносливость": ch["Выносливость"],
			"Интеллект":    ch["Интеллект"],
			"Мудрость":     ch["Мудрость"],
		}

		for i := 0; i < 2; i++ {
			char_index := rand.Intn(len(halfElf)) //4
			slice := make([]string, 0)
			for key := range halfElf {
				slice = append(slice, key)
			}
			halfElf[slice[char_index]] += 1
		}
		halfElf["Харизма"] = ch["Харизма"]
		return halfElf
	}
	fmt.Println(ch)
	return ch

}
