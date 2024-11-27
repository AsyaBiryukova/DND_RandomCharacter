package personality

import (
	"math/rand"

	r "github.com/AsyaBiryukova/DND_RandomCharacter/MainHeroInfo"
)

func HeroPersonality() string {

	h_back := r.RandomBackground()
	var tpers string

	if h_back == "Артист" {
		pers := map[int]string{
			1: "Почти для любого случая у меня есть подходящая байка.",
			2: "Попав в новое место, я собираю местные слухи и сам сплетничаю.",
			3: "Я безнадёжный романтик и всегда ищу 'ту самую'.",
			4: "На меня невозможно долго злиться, я умею разряжать обстановку.",
			5: "Обожаю мелкие подколки, даже если они в мой адрес.",
			6: "Ненавижу, когда в центре внимания кто-то другой, а не я.",
			7: "Я во всём стремлюсь к совершентсву.",
			8: "Моё настроение и мнение меняются также легко, как мелодии, которые я играю.",
		}
		rpers := rand.Intn(8) + 1
		tpers = pers[rpers]

	} else if h_back == "Благородный" {
		pers := map[int]string{
			1: "Моя лесть заставляет собеседников чувствовать себя важными и особенными.",
			2: "Простой народ любит меня за доброту и честность.",
			3: "По моему цартсвенному виду издалека заметно, что я отличаюсь от немытой толпы.",
			4: "Я изо всех сил стараюсь выглядеть идеально и следовать новейшей моде.",
			5: "Я ненавижу грязь и скорее умру, чем буду спать в некомфортных условиях.",
			6: "Несмторя на происхождение, я не гляжу на простой народ свысока. Мы все одной крови.",
			7: "Если кто-то потерял моё расположение - это навсегда.",
			8: "Того, кто мне навредит, я сокрушу, смешаю с грязью, а поля его посыплю солью.",
		}
		rpers := rand.Intn(8) + 1
		tpers = pers[rpers]

	} else if h_back == "Бродяга" {
		pers := map[int]string{
			1: "Я прячу по карманам кусочки еды и безделушки.",
			2: "Я всегда задаю кучу вопросов.",
			3: "Мне нравится забираться в такие места, где меня никто не найдёт.",
			4: "Я всегда сплю спиной к сетене или к дереву, прижав свои пожитки к себе.",
			5: "Я ем как свинья, да и манеры у меня дрянные.",
			6: "Я подозреваю, что те, кто добр ко мне, на деле замышляют зло.",
			7: "Ненавижу мыться!",
			8: "О чём другой промолчал бы или намекнул бы, я скажу прямо в лоб.",
		}
		rpers := rand.Intn(8) + 1
		tpers = pers[rpers]

	} else if h_back == "Моряк" {
		pers := map[int]string{
			1: "Друзья могут на меня положиться при любых обстоятельствах.",
			2: "Я впахиваю по полной, чтобы после раюоты оторваться по полной.",
			3: "Обожаю заходить в новые порты и заводить новых друзей за кружкой эля.",
			4: "Лучший способ узнать новый город - это подраться в трактире.",
			5: "Я всегда готов биться об заклад и делать ставки.",
			6: "Я могу слегка приврать, чтобы рассказ получился интереснее.",
			7: "Выражения у меня грязнее, чем логово отьюга.",
			8: "Люблю, когда работа сделана хорошо, - особенно если кто-то делает её за меня.",
		}
		rpers := rand.Intn(8) + 1
		tpers = pers[rpers]

	} else if h_back == "Мудрец" {
		pers := map[int]string{
			1: "Я использую многосложные слова, создающие впечатление образованности.",
			2: "Я прочитал все книги в величайших библиотеках мира — или, по крайней мере, так я люблю хвастать.",
			3: "Я привык помогать тем, кто не так умён как я, и терпеливо объясняю им всё.",
			4: "Нет ничего лучше хорошей тайны.",
			5: "Я стараюсь выслушать все стороны в споре, прежде чем приму решение.",
			6: "Я ... говорю ... ме-е-едленно ... когда общаюсь ... с идио-о-отами, ... то есть ... почти со все-е-еми, кто меня окружает. ",
			7: "Я ужасно, ужасно неловко себя чувствую, когда общаюсь с другими.",
			8: "Я уверен, что другие постоянно пытаются украсть мои секреты.",
		}
		rpers := rand.Intn(8) + 1
		tpers = pers[rpers]

	} else if h_back == "Народный герой" {
		pers := map[int]string{
			1: "Я сужу о людях по их делам, а не по словам.",
			2: "Если кто-то в беде, я всегда рад помочь.",
			3: "Если уж я что-то решил, сделаю обязательно, чтобы ни стояло у меня на пути.",
			4: "У меня обострённое чувство справедливости, я всегда ищу самое честное решение проблемы.",
			5: "Я верю в свои силы и стараюсь внушить такую же веру другим.",
			6: "Пусть другие думают, я буду действовать.",
			7: "Я использую длинные слова, смысла которых не знаю, чтобы казаться умнее.",
			8: "Я быстро начинаю скучать. Каогда я уже смогу воплотить своё предназначение?",
		}
		rpers := rand.Intn(8) + 1
		tpers = pers[rpers]

	} else if h_back == "Отшельник" {
		pers := map[int]string{
			1: "Я так долго жил один, что почти не разговариваю. Чаще я общаюсь жестами и изредка ворчу.",
			2: "Я совершенно спокоен даже перед лицом опасности.",
			3: "У моего духовного учителя были мудрые изречения на любую тему, и я стараюсь поделиться его мудростью.",
			4: "Я очень сочуствую всем, кто страдает.",
			5: "Я не имею представление об этикете и социальных нормах.",
			6: "Всё, что происходит со мной, я считаю частью великого вселенского плана.",
			7: "Я часто погружаюсь в размышления и не замечаю, что происходит вокруг.",
			8: "Я работаю над грандиозной философской теорией и люблю делиться своими идеями.",
		}
		rpers := rand.Intn(8) + 1
		tpers = pers[rpers]

	} else if h_back == "Послушник" {
		pers := map[int]string{
			1: "Я восхощаюсь героем моей веры и постоянно ставлю его подвиги себе в пример.",
			2: "Я могу найти общий язык даже заклятым врагам, понять их и побудить заключить мир.",
			3: "Я вижу знамения в каждом событии и действии. Боги пытаются говорить с нами, надо только слушать их.",
			4: "Ничто не может нарушить мой оптимизм.",
			5: "Я часто (или неточно) цитирую священные тксты и притчи по любому поводу.",
			6: "Я терпимо (или нетерпимо) отношусь к чужой вере и уважаю (или осуждаю) поклонение другим богам.",
			7: "Я был частью элиты своего храма и привык к хорошей еде, выпивке и высшему обществу. Лишения в дороге меня выматывают.",
			8: "Я так долго жил в храме, что у меня нет опыта в повседневном общении с людьми.",
		}
		rpers := rand.Intn(8) + 1
		tpers = pers[rpers]

	} else if h_back == "Преступник" {
		pers := map[int]string{
			1: "У меня есть план на любой случай неприятности.",
			2: "Я спокоен в любой ситуации. Я никогда не повышаю голос и не поддаюсь эмоциям.",
			3: "На новом месте я первым делом подмечаю, что тут ценного - или где оно могло бы быть.",
			4: "Я лучше заведу нового друга, чем нового врага.",
			5: "Моё доверие трудно завоевать. Те, кто кажутся честными, скрывают больше всех.",
			6: "Я плевать хотел на риски и опасности. Не говори мне об опасностях.",
			7: "Лучший способ заставить меня что-то сделать - взять на слабо.",
			8: "При малйшем оскорблении я могу взорваться.",
		}
		rpers := rand.Intn(8) + 1
		tpers = pers[rpers]

	} else if h_back == "Ремесленник из гильдии" {
		pers := map[int]string{
			1: "Я так считаю: берёшься за дело - делай его с толком. Да, я перфекционист.",
			2: "Я сноб и презираю тех, кто не ценит моё мастерство.",
			3: "Я всегда стараюсь знать, как делаются делаются дела и что нужно людям.",
			4: "У меня полно остроумных фразочек и пословиц на все случаи жизни.",
			5: "Я груб с людьми, которые не разделяют мою преверженность к труду и честности.",
			6: "Я люблю подолгу рассказывать о своей профессии.",
			7: "Я нелегко расстаюсь с деньгами и готов торговаться до упора, чтобы добиться быгодной сделки.",
			8: "Моя работа меня прославила, и я стараюсь, чтобы все её ценили. Люди, которые обо мне не слышали, меня поражают.",
		}
		rpers := rand.Intn(8) + 1
		tpers = pers[rpers]

	} else if h_back == "Солдат" {
		pers := map[int]string{
			1: "Я всегда вежлив и уважителен.",
			2: "Меня мучают воспоминания о войне. Я не могу забыть ужасы, которые повидал.",
			3: "Я потерял многих друзей, а новых не спешу заводить.",
			4: "У меня полно вдохновляющих и поучительных историй из моего опыта, связанных с битвами.",
			5: "Я могу смотреть в глаза адской гончей и не дрогуть.",
			6: "Я силач и люблю ломать вещи, чтобы это показать.",
			7: "Шутки у меня порой грубоватые.",
			8: "Проблемы я решаю просто и напрямую - это лучший путь к успеху.",
		}
		rpers := rand.Intn(8) + 1
		tpers = pers[rpers]

	} else if h_back == "Чужеземец" {
		pers := map[int]string{
			1: "Мною движет тяга к странствиям, из-за неё я ушёл из дома.",
			2: "Я присматриваю за своими друзьями, как за неопытными щенками.",
			3: "Однажды я пробежал 25 миль без передышки, чтобы предупредить свой клан о наступлении орочьей орды. Если надо, я снова так поступлю.",
			4: "Для каждой ситуации у меня есть урок, который я почерпнул из наблюдений за природой.",
			5: "Я не поставлю ни гроша на богатых и воспитанных. Деньги и хорошие манеры не спасут ват от голодного медвесыча.",
			6: "Я постоянно подбираю вещи, верчу их в руках и иногда по неостарожности ломаю.",
			7: "Мне больше по нраву общество зверей, чем людей.",
			8: "По правде говоря, меня вырастили волки.",
		}
		rpers := rand.Intn(8) + 1
		tpers = pers[rpers]

	} else if h_back == "Шарлатан" {
		pers := map[int]string{
			1: "Я легко влюбляюсь и расстаюсь, и я всегда 'в активном поиске'.",
			2: "Я шучу по любому поводу, даже когда юмор неуместен.",
			3: "Я помощью лести я добиваюсь чего угодно.",
			4: "Я азартный игрок и не могу удержаться от риска, если награда велика.",
			5: "Я вру как дышу, даже когда в этом нет необходимости.",
			6: "Моё любимое оружие - сарказм и подколки.",
			7: "Я таскаю с собой символы разных богов и обращаюсь к тому из них, кто подходит в данных момент.",
			8: "Я тащу в карманы всё ценное, что плохо лежит.",
		}
		rpers := rand.Intn(8) + 1
		tpers = pers[rpers]
	}
	return "Ваша черта характера" + tpers
}
