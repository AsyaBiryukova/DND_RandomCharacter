// Файл с Характеристиками Героя
package herocharsandskills

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/AsyaBiryukova/DND_RandomCharacter/Dices"

	r "github.com/AsyaBiryukova/DND_RandomCharacter/MainHeroInfo"
)

func BaseCharacteristics() map[string]int { // Функция, определяющая базовые Характеристики Героя
	// Установка базовых Характеристики Силы
	strenght := make([]int, 4) // Заводим слайс из 4-х элементов типа int
	for i := range strenght {  // Начинаем цикл, выполняющийся для каждого элемента слайса strenght
		strenght[i] = Dices.RollDiceD6() // В каждой итерации присваиваем значение из функции RollDiceD6
	}
	sort.Ints(strenght)                        // Сортируем значения strenght по возрастанию
	lastStrenght := strenght[len(strenght)-3:] // Выбираем 3 последние элемента слайса и заносим их в новую переменную
	totalStrenght := 0                         // Заводим "пустую" переменную, в которой будем храниться итоговое значение базовой Силы
	for _, s_num := range lastStrenght {       // Перебираем элементы в lastStrenght
		totalStrenght += s_num // Суммируем текущее значение элемента lastStrenght и отправляем храниться в totalStrenght
	}
	// И так с каждой Характеристикой
	// Установка базовых Характеристики Ловкости
	dexterity := make([]int, 4)
	for i := range dexterity {
		dexterity[i] = Dices.RollDiceD6()
	}
	sort.Ints(dexterity)
	lastDexterity := dexterity[len(dexterity)-3:]
	totalDexterity := 0
	for _, d_num := range lastDexterity {
		totalDexterity += d_num
	}
	//Устанавливаем базовые характеристики Выносливости
	constitution := make([]int, 4)
	for i := range constitution {
		constitution[i] = Dices.RollDiceD6()
	}
	sort.Ints(constitution)
	lastConstitution := constitution[len(constitution)-3:]
	totalConstitution := 0
	for _, c_num := range lastConstitution {
		totalConstitution += c_num
	}
	//Устанавливаем базовые характеристики Интеллекта
	intelligense := make([]int, 4)
	for i := range intelligense {
		intelligense[i] = Dices.RollDiceD6()
	}
	sort.Ints(intelligense)
	lastIntelligense := intelligense[len(intelligense)-3:]
	totalIntelligense := 0
	for _, i_num := range lastIntelligense {
		totalIntelligense += i_num
	}
	//Устанавливаем базовые характеристики Мудрости
	wisdom := make([]int, 4)
	for i := range wisdom {
		wisdom[i] = Dices.RollDiceD6()
	}
	sort.Ints(wisdom)
	lastwisdom := wisdom[len(wisdom)-3:]
	totalWisdom := 0
	for _, i_num := range lastwisdom {
		totalWisdom += i_num
	}
	//Устанавливаем базовые характеристики Харизмы
	charisma := make([]int, 4)
	for i := range charisma {
		charisma[i] = Dices.RollDiceD6()
	}
	sort.Ints(charisma)
	lastcharisma := charisma[len(charisma)-3:]
	totalCharisma := 0
	for _, i_num := range lastcharisma {
		totalCharisma += i_num
	}

	basecharacteristics := map[string]int{ // Создаём мапу и заносим туда значения базовых Характеристик
		"Сила":         totalStrenght,
		"Ловкость":     totalDexterity,
		"Выносливость": totalConstitution,
		"Интеллект":    totalIntelligense,
		"Мудрость":     totalWisdom,
		"Харизма":      totalCharisma,
	}

	return basecharacteristics // Возврат конечной мапы
}

func RaceUpdateCharacteristic() map[string]int { // Функция, которая апдейтит базовые Характеристики по Расе
	race := r.RandomRace()         // Вызов расы из функции RandomRace
	fmt.Println("Ваша раса", race) // Вывод расы (убрать из этой функции?)

	ch := BaseCharacteristics() // Вызов базовых Характеристик из функции BaseCharacteristics
	fmt.Println(ch)             // ДЛЯ ТЕСТОВ ПОТОМ УБРАТЬ!!!
	// Сортируем Характеристики и Модификаторы по Расам и добавляем их к базовым значеням
	// Холмовой Дварф: Выносливость +2, Мудрость +1
	// Горный Дварф: Выносливость +2, Сила +2
	// Крепкий Полурослик: Ловкость +2, Выносливость +1
	// Легконогий Полурослик: Ловкость +2, Харизма +1
	// Человек: Все характеристики +1
	// Высший Эльф: Ловкость +2, Интеллект +1
	// Лесной Эльф: Ловкость +2, Мудрость +1
	// Дроу: Ловкость +2, Харизма +1
	// Лесной Гном: Интеллект +2, Ловкость +1
	// Скальный Гном: Интеллект +2, Выносливость +1
	// Драконорождённый: Сила +2, Харизма +1
	// Полуорк: Сила +2, Выносливость +1
	// Полуэльф: Харизма +2, две рандомные +1
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
	if race == "Полуэльф" { // Апдейт Характеристик от Полуэльфа
		halfElv := map[string]int{ // Заводим мапу, в которой будем перебирать те значения, которые можно прокачать рандомно
			"Сила":         ch["Сила"], // Приравниваем их к базовым Характеристикам
			"Ловкость":     ch["Ловкость"],
			"Выносливость": ch["Выносливость"],
			"Интеллект":    ch["Интеллект"],
			"Мудрость":     ch["Мудрость"],
		}

		for i := 0; i < 2; i++ { // Начинаем цикл, который выберет 2 Характеристики на прокачку
			char_index := rand.Intn(len(halfElv)) // Генерируем случайное число от 0 до максимальной длины halfElv
			slice := make([]string, 0)            // Заводим пустой слайс типа string
			for key := range halfElv {            // Начинаем ещё один цикл, который проходит по всем элементам halfElv
				slice = append(slice, key) // В каждой итерации добавляем элемент halfElv в слайс
			}
			halfElv[slice[char_index]] += 1 // Увеличиваем значение элемента на 1
		}
		halfElv["Харизма"] = ch["Харизма"] + 2 // Добавляем в мапу и увеличиваем Харизму на 2
		return halfElv                         // Возвращаем итоговые Характеристики Полуэльфа
	}
	fmt.Println(ch) // Выводим итоговые Характеристики (убрать?)
	return ch

}
