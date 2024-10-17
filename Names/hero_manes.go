package names

import (
	"math/rand"
)

// функция для определения имени вашего персонажа если он Холмовой или Горный дварф
func DwarfNames() string {
	dwarf_male_names := []string{"Адрик", "Альбрих", "Барендд", "Броттор", "Бренор", "Бэрн", // слайс с мужскими именами
		"Вейт", "Вондал", "Гардаин", "Даин", "Даррак", "Дельг", "Кильдрак", "Моргран", "Орсик",
		"Оскар", "Рангрим", "Рюрик", "Таклинн", "Торадин", "Тордек", "Торин", "Травок", "Траубон",
		"Ульфгар", "Фаргрим", "Флинт", "Харбек", "Эберг", "Эйнкиль"}
	dwarf_male_index := rand.Intn(len(dwarf_male_names)) // рандомайзер мужских имён

	dwarf_female_names := []string{"Амбер", "Артин", "Аудхильд", "Бардрин", "Вистра", // слайс с женскими именами
		"Гуннлода", "Гурдис", "Дагнал", "Диеза", "Ильде", "Катра", "Кристрит", "Лифтраза",
		"Мардред", "Рисвинн", "Саннл", "Торбера", "Торгга", "Фалькрунн", "Финеллен", "Хелья",
		"Хлин", "Эльдед"}
	dwarf_female_index := rand.Intn(len(dwarf_female_names)) //рандомайзер женских имён

	dwarf_tribes := []string{"Айронфист", "Бальдерк", "Батлхаммер", "Броненвил", "Горунн", "Данкилл", // слайс с названиями кланов
		"Лодерр", "Лютгер", "Румнахейм", "Стракельн", "Торунн", "Унгарт", "Фаерфордж", "Фростберд", "Хольдерхек"}
	dwarf_tribe_index := rand.Intn(len(dwarf_tribes)) //рандомайзер кланов

	var dwarf_names_var string //вывод результатов рандомайзера
	dwarf_names_var = "Вот имена, подходящие вашему персонажу, если вы Холмовой Дварф или Горный Дварф: \nMужское: " + dwarf_male_names[dwarf_male_index] +
		"\nЖенское: " + dwarf_female_names[dwarf_female_index] + "\nВаш клан: " + dwarf_tribes[dwarf_tribe_index]
	return dwarf_names_var
}

// функция для определения имени вашего персонажа если он полурослик
func HalflingNames() string {
	halfling_male_names := []string{"Альтон", "Андер", "Веллби", "Гаррет", "Кейд", "Коррин",
		"Лайль", "Линдал", "Майло", "Меррик", "Осборн", "Перрин", "Рид", "Финнан", "Эльдон",
		"Эррик"}
	halfling_male_index := rand.Intn(len(halfling_male_names))

	halfling_female_names := []string{"Бри", "Верна", "Вэни", "Джиллиан", "Келли",
		"Китри", "Кора", "Ливиния", "Лидда", "Мерла", "Недда", "Паэла",
		"Порция", "Сераифна", "Трим", "Шейна", "Эндри", "Юфимия"}
	halfling_female_index := rand.Intn(len(halfling_female_names))

	halfling_surnames := []string{"Валигора", "Висельчак", "Вышегор", "Добробочка", "Зеленобутыль",
		"Камнемёт", "Кистебор", "Подветка", "Чаелист", "Шипомер"}
	halfling_surnames_index := rand.Intn(len(halfling_surnames))

	var halfling_names_var string
	halfling_names_var = "Вот имена, подходящие вашему персонажу, если вы Легконогий Полурослик или Крепкий Полурослик: \nМужское: " + halfling_male_names[halfling_male_index] +
		"\nЖенское: " + halfling_female_names[halfling_female_index] + "\nВаша фамилия: " + halfling_surnames[halfling_surnames_index]
	return halfling_names_var
}

// функция для определения имени вашего персонажа и земель, откуда он родом, если он человек
func HumanNames() string {

	human_name := ""

	human_nationality := []string{"Дамаранец", "Иллюсканец", "Калишит", "Муланец", "Рашемен", "Тетирец", "Турамиец", "Чондатанец", "Шу"}
	human_nationality_index := rand.Intn(len(human_nationality))
	var human_nationality_var string
	human_nationality_var = human_nationality[human_nationality_index]

	if human_nationality_var == "Дамаранец" {
		damar_male_names := []string{"Бор", "Глар", "Григор", "Ивор", "Иган", "Козеф",
			"Миваль", "Орел", "Павел", "Сергор", "Фодель"}
		damar_male_index := rand.Intn(len(damar_male_names))

		damar_female_names := []string{"Алетра", "Зора", "Кара", "Катерина", "Мара",
			"Натали", "Ольма", "Тана"}
		damar_female_index := rand.Intn(len(damar_female_names))

		damar_surname := []string{"Берск", "Доцк", "Куленов", "Марск",
			"Немецк", "Стараг", "Чернин", "Шемов"}
		damar_surname_index := rand.Intn(len(damar_surname))
		var damar_name_var string
		damar_name_var = "Вот имена, подходящие вашему персонажу, если вы Человек " + human_nationality_var + ":\nМужское: " + damar_male_names[damar_male_index] +
			"\nЖенское: " + damar_female_names[damar_female_index] + "\nФамилия: " + damar_surname[damar_surname_index]
		human_name += damar_name_var
	}
	if human_nationality_var == "Иллюсканец" {
		ill_male_names := []string{"Андер", "Блат", "Бран", "Гет",
			"Ландер", "Лут", "Мальсер", "Стор", "Таман", "Фрат", "Урт"}
		ill_male_index := rand.Intn(len(ill_male_names))

		ill_female_names := []string{"Амарфей", "Бета", "Вестра", "Кетра", "Мара",
			"Ольга", "Серфей"}
		ill_female_index := rand.Intn(len(ill_female_names))

		ill_surname := []string{"Брайтвуд", "Виндриввер", "Лекмен",
			"Стормвинд", "Хельдер", "Хорнрейвен"}
		ill_surname_index := rand.Intn(len(ill_surname))

		var ill_name_var string
		ill_name_var = "Вот имена, подходящие вашему персонажу, если вы Человек " + human_nationality_var + ":\nМужское: " + ill_male_names[ill_male_index] +
			"\nЖенское: " + ill_female_names[ill_female_index] + "\nФамилия: " + ill_surname[ill_surname_index]
		human_name += ill_name_var
	}
	if human_nationality_var == "Калишит" {
		kal_male_names := []string{"Азеир", "Бардеид", "Зашеир", "Мехмен",
			"Судейман", "Хасеид", "Хемен"}
		kal_male_index := rand.Intn(len(kal_male_names))

		kal_female_names := []string{"Атала", "Жасмаль", "Зашеида", "Мейлиль", "Сеидиль",
			"Сейпора", "Хама", "Яшеира"}
		kal_female_index := rand.Intn(len(kal_female_names))

		kal_surname := []string{"Баша", "Думеин", "Жассан",
			"Мостана", "Пашар", "Реин", "Халид"}
		kal_surname_index := rand.Intn(len(kal_surname))

		var kal_name_var string
		kal_name_var = "Вот имена, подходящие вашему персонажу, если вы Человек " + human_nationality_var + ":\nМужское: " + kal_male_names[kal_male_index] +
			"\nЖенское: " + kal_female_names[kal_female_index] + "\nФамилия: " + kal_surname[kal_surname_index]
		human_name += kal_name_var
	}
	if human_nationality_var == "Муланец" {
		mul_male_names := []string{"Аот", "Барерис", "Кетот", "Мумед",
			"Рамас", "Со-Кехур", "Тазар-Де", "Урхур", "Эпут-Ки"}
		mul_male_index := rand.Intn(len(mul_male_names))

		mul_female_names := []string{"Аризима", "Золида", "Мурити", "Нефида", "Нулара",
			"Сефрида", "Тола", "Умара", "Чати"}
		mul_female_index := rand.Intn(len(mul_female_names))

		mul_surname := []string{"Анскульд", "Анхалаб", "Натандем",
			"Сепрет", "Уутракт", "Фезим", "Хапет"}
		mul_surname_index := rand.Intn(len(mul_surname))

		var mul_name_var string
		mul_name_var = "Вот имена, подходящие вашему персонажу, если вы Человек " + human_nationality_var + ":\nМужское: " + mul_male_names[mul_male_index] +
			"\nЖенское: " + mul_female_names[mul_female_index] + "\nФамилия: " + mul_surname[mul_surname_index]
		human_name += mul_name_var
	}

	if human_nationality_var == "Рашемен" {
		rash_male_names := []string{"Боривик", "Владислак", "Джандар", "Канитар",
			"Мадислак", "Ральмевик", "Фаугар", "Шаумар"}
		rash_male_index := rand.Intn(len(rash_male_names))

		rash_female_names := []string{"Имзель", "Иммит", "Наварра", "Таммит", "Фьеварра",
			"Хульмарра", "Шеварра", "Юльдара"}
		rash_female_index := rand.Intn(len(rash_female_names))

		rash_surname := []string{"Дьернина", "Ильтазяра", "Мурнетара",
			"Стаянога", "Ульмокина", "Чергоба"}
		rash_surname_index := rand.Intn(len(rash_surname))

		var rash_name_var string
		rash_name_var = "Вот имена, подходящие вашему персонажу, если вы Человек " + human_nationality_var + ":\nМужское: " + rash_male_names[rash_male_index] +
			"\nЖенское: " + rash_female_names[rash_female_index] + "\nФамилия: " + rash_surname[rash_surname_index]
		human_name += rash_name_var
	}

	if human_nationality_var == "Тетирец" {
		tet_male_names := []string{"Антон", "Диеро", "Маркон", "Пьерон",
			"Римардо", "Ромеро", "Салазар", "Умберо"}
		tet_male_index := rand.Intn(len(tet_male_names))

		tet_female_names := []string{"Балама", "Вонда", "Дона", "Жалана", "Квара",
			"Марта", "Селиз", "Фаила"}
		tet_female_index := rand.Intn(len(tet_female_names))

		tet_surname := []string{"Агусто", "Асторио", "Домине",
			"Калабра", "Маривальди", "Пизакар", "Рамондо", "Фалоне"}
		tet_surname_index := rand.Intn(len(tet_surname))

		var tet_name_var string
		tet_name_var = "Вот имена, подходящие вашему персонажу, если вы Человек " + human_nationality_var + ":\nМужское: " + tet_male_names[tet_male_index] +
			"\nЖенское: " + tet_female_names[tet_female_index] + "\nФамилия: " + tet_surname[tet_surname_index]
		human_name += tet_name_var
	}

	if human_nationality_var == "Чондатанец" {
		chon_male_names := []string{"Дарвин", "Дорн", "Горстаг", "Грим",
			"Маларк", "Морн", "Рендал", "Стедд", "Хельм", "Эвендур"}
		chon_male_index := rand.Intn(len(chon_male_names))

		chon_female_names := []string{"Арвин", "Джеззейль", "Кэрри", "Люрин", "Мири",
			"Ровен", "Тессель", "Шандри", "Эсвель"}
		chon_female_index := rand.Intn(len(chon_female_names))

		chon_surname := []string{"Бакмен", "Грейскал", "Дандрагон",
			"Толлстаг", "Эвенвуд", "Эмблкроун"}
		chon_surname_index := rand.Intn(len(chon_surname))

		var chon_name_var string
		chon_name_var = "Вот имена, подходящие вашему персонажу, если вы Человек " + human_nationality_var + ":\nМужское: " + chon_male_names[chon_male_index] +
			"\nЖенское: " + chon_female_names[chon_female_index] + "\nФамилия: " + chon_surname[chon_surname_index]
		human_name += chon_name_var
	}

	if human_nationality_var == "Шу" {
		shu_male_names := []string{"Ань", "Вэнь", "Лон", "Лянь",
			"Мен", "Онь", "Фай", "Цзян", "Чень", "Чжунь", "Чи", "Шань", "Шуй"}
		shu_male_index := rand.Intn(len(shu_male_names))

		shu_female_names := []string{"Бай", "Лэй", "Мэй", "Тай", "Чао",
			"Цзя", "Цяо", "Шуй"}
		shu_female_index := rand.Intn(len(shu_female_names))

		shu_surname := []string{"Вань", "Као", "Кун",
			"Лин", "Ляо", "Мэй", "Пинь", "Сум", "Тань", "Чен", "Хуан", "Шинь"}
		shu_surname_index := rand.Intn(len(shu_surname))

		var shu_name_var string
		shu_name_var = "Вот имена, подходящие вашему персонажу, если вы Человек " + human_nationality_var + ":\nМужское: " + shu_male_names[shu_male_index] +
			"\nЖенское: " + shu_female_names[shu_female_index] + "\nФамилия: " + shu_surname[shu_surname_index]
		human_name += shu_name_var
	}

	return human_name
}

// функция для определения имени вашего персонажа если он Эльф
func ElveNames() string {
	elv_male_names := []string{"Адран", "Арамиль", "Араннис", "Ауст", "Аэлар", "Бейро", // слайс с мужскими именами
		"Берриан", "Варис", "Галиндан", "Ивеллиос", "Иммераль", "Каррик", "Кварион", "Лайсиан", "Миндартис",
		"Паэлиас", "Перен", "Риардон", "Ролен", "Совелисс", "Тамиор", "Таривол", "Терен", "Хадарай",
		"Хейян", "Химо", "Энималис", "Эрдан", "Эреван"}
	elv_male_index := rand.Intn(len(elv_male_names)) // рандомайзер мужских имён

	elv_female_names := []string{"Адри", "Альтея", "Анастрианна", "Андрасте", "Антинуя", // слайс с женскими именами
		"Бетринна", "Бирель", "Вадания", "Валанте", "Друзилия", "Желаннет", "Занафия", "Йеления",
		"Квеленна", "Каиллатэ", "Кейлет", "Кейлинн", "Лешанна", "Лия", "Миали", "Мэриэль",
		"Найвара", "Сариэль", "Силакви", "Тэйрастра", "Тия", "Фелосиаль", "Шава", "Шанэйра",
		"Энна"}
	elv_female_index := rand.Intn(len(elv_female_names)) //рандомайзер женских имён

	elv_surname := []string{"Амакиир (Дагоценный Цветок)", "Амастасия (Звёздный Цветок)", "галанодэль (Лунный Шопот)",
		"Зилосциент (Золотой Лепесток)", "Ильфелькиир (Драгоценный Блеск)", "Лиадон (Серебряный Лист)", // слайс с фамилиями
		"Мелиамнэ (Дубовый Корень)", "Найло (Ночной Ветерок)", "Сианнодэль (Лучнный Ручей)", "Холимион (Алмазная Роса)"}
	elv_surname_index := rand.Intn(len(elv_surname)) //рандомайзер фамилий

	var elv_names_var string //вывод результатов рандомайзера
	elv_names_var = "Вот имена, подходящие вашему персонажу, если вы Высший Эльф, Лесной Эльф или Дроу: \nMужское: " + elv_male_names[elv_male_index] +
		"\nЖенское: " + elv_female_names[elv_female_index] + "\nВаша фамилия: " + elv_surname[elv_surname_index]
	return elv_names_var
}

// функция для определения имени вашего персонажа если он Гном
func GnomeNames() string {
	gnome_male_names := []string{"Альвин", "Альстон", "Бёрджелл", "Боддинок", "Брокк", "Варрин", // слайс с мужскими именами
		"Вренн", "Гербо", "Гимбл", "Глим", "Джебеддо", "Димбл", "Зук", "Келлен", "Намфудль",
		"Оррин", "Рундар", "Сибо", "Синдри", "Фонкин", "Фраг", "Эльдон", "Эрки"}
	gnome_male_index := rand.Intn(len(gnome_male_names)) // рандомайзер мужских имён

	gnome_female_names := []string{"Бимпноттин", "Брина", "Вейвокет", "Донелла", "Дувамиль", // слайс с женскими именами
		"Занна", "Карамип", "Карлин", "Лилли", "Лорилла", "Лунпоттин", "Марднап", "Никс",
		"Нисса", "Ода", "Орла", "Ройвин", "Тана", "Шамилла", "Элла", "Элливик",
		"Эллидджобелл"}
	gnome_female_index := rand.Intn(len(gnome_female_names)) //рандомайзер женских имён

	gnome_tribe := []string{"Берен", "Гаррик", "Дергель", "Мёрнинг", "Некль", "Нингель", "Рольнор", "Тимберс", "Турен", "Фолькор",
		"Щеппен"}
	gnome_tribe_index := rand.Intn(len(gnome_tribe)) //рандомайзер кланов

	gnome_nicknames := []string{"Барсук", "Блестяшка", "Двухзамок", "Ку", "Ним", "Одна-туфля", "Оспина", "Пепельник", "Пивохлёб", "Плащик",
		"Тестохват", "Утковал", "Фниппер"}
	gnome_nicknames_index := rand.Intn(len(gnome_nicknames)) //рандомайзер прозвище

	var gnome_names_var string //вывод результатов рандомайзера
	gnome_names_var = "Вот имена, подходящие вашему персонажу, если вы Лсеной или Скальный Гном: \nMужское: " + gnome_male_names[gnome_male_index] +
		"\nЖенское: " + gnome_female_names[gnome_female_index] + "\nВаш клан: " + gnome_tribe[gnome_tribe_index] +
		"\nВаше прозвище: " + gnome_nicknames[gnome_nicknames_index]
	return gnome_names_var
}

// функция для определения имени вашего персонажа если он Драконорождённый
func DroganbornNames() string {
	dragonborn_male_names := []string{"Аржхан", "Балазар", "Бхараш", "Геш", "Донаар", "Крив", // слайс с мужскими именами
		"Медраж", "Мехен", "Надарр", "Панжед", "Патрин", "Рогар", "Тархун", "Торинн", "Хескан",
		"Шамаш", "Шединн"}
	dragonborn_male_index := rand.Intn(len(dragonborn_male_names)) // рандомайзер мужских имён

	dragonborn_female_names := []string{"Акра", "Бири", "Даар", "Жери", "Кава", // слайс с женскими именами
		"Коринн", "Мишанн", "Нала", "Перра", "Райянн", "Сора", "Сурина", "Тава",
		"Уаджит", "Фарида", "Хавилар", "Харанн"}
	dragonborn_female_index := rand.Intn(len(dragonborn_female_names)) //рандомайзер женских имён

	dragonborn_tribe := []string{"Версисатургиеш", "Даардендриан", "Делмирев", "Драчедандион", "Кепешкмолик", "Керрилон", "Кимбатуул", "Клестинсиаллор",
		"Линксакасендалор", "Мястан", "Неммонис", "Нориксиус", "Офиншталажир", "Прексижандилин", "Турнурот", "Фенкенкабрадон", "Шестенделиат", "Яржерит"}
	dragonborn_tribe_index := rand.Intn(len(dragonborn_tribe)) //рандомайзер кланов

	var dragonborn_names_var string //вывод результатов рандомайзера
	dragonborn_names_var = "Вот имена, подходящие вашему персонажу, если вы Драконорождённый: \nMужское: " + dragonborn_male_names[dragonborn_male_index] +
		"\nЖенское: " + dragonborn_female_names[dragonborn_female_index] + "\nВаш клан: " + dragonborn_tribe[dragonborn_tribe_index]
	return dragonborn_names_var
}

// функция для определения имени вашего персонажа если он Полуорк
func HalforcNames() string {
	halfork_male_names := []string{"Гел", "Денч", "Имш", "Кет", "Краск",
		"Муррен", "Ронт", "Токк", "Фенг", "Хенк", "Холг", "Шамп"} // слайс с мужскими именами
	halfork_male_index := rand.Intn(len(halfork_male_names)) // рандомайзер мужских имён

	halfork_female_names := []string{"Багги", "Вола", "Волен", "Евельда", "Кансиф",
		"Мев", "Нига", "Овак", "Оунка", "Сута", "Шаута", "Эмен", "Энгонг"}
	halfork_female_index := rand.Intn(len(halfork_female_names)) //рандомайзер женских имён

	var halfork_names_var string //вывод результатов рандомайзера
	halfork_names_var = "Вот имена, подходящие вашему персонажу, если вы Полуорк: \nMужское: " + halfork_male_names[halfork_male_index] +
		"\nЖенское: " + halfork_female_names[halfork_female_index]
	return halfork_names_var
}

// функция для определения имени вашего персонажа, если он Тифлинг
func TieflingNames() string {
	tiefling_male_names := []string{"Акменос", "Амнон", "Баракас", "Дамакос", "Йадос", "Кайрон", "Люцис", "Мелех", "Мордай", "Мортос", "Пелайос", "Скамос", "Терай", "Экемон"} // слайс с мужскими именами
	tiefling_male_index := rand.Intn(len(tiefling_male_names))                                                                                                                 // рандомайзер мужских имён

	tiefling_female_names := []string{"Акта", "Анакис", "Брисеис", "Дамая", "Каллиста", "Криелла", "Лерисса", "Макария", "Немея", "Орианна", "Риета", "Фелая", "Эа"}
	tiefling_female_index := rand.Intn(len(tiefling_female_names)) //рандомайзер женских имён

	tiefling_idea_name := []string{
		"Безрассудство", "Вера", "Идеал", "Искусство", "Музыка", "Мука", "Надежда", "Напев", "Нигде", "Открытость", "Отчаяние",
		"Падаль", "Поиск", "Почтение", "Поэзия", "Превосходство", "Скорбь", "Слава", "Случайность", "Страх", "Усталость"}
	tiefling_idea_index := rand.Intn(len(tiefling_idea_name))

	var tiefling_names_var string //вывод результатов рандомайзера
	tiefling_names_var = "Вот имена, подходящие вашему персонажу, если вы Тифлинг: \nMужское инфернальное имя: " + tiefling_male_names[tiefling_male_index] +
		"\nЖенское инфернальное имя: " + tiefling_female_names[tiefling_female_index] + "\nИдейное имя: " + tiefling_idea_name[tiefling_idea_index]
	return tiefling_names_var
}
