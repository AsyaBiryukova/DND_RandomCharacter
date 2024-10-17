package charactercard

import (
	r "github.com/AsyaBiryukova/DND_RandomCharacter/Randoms"
)

func ReadyHeroInfo() string {
	hero_race := r.RandomRace()
	hero_class := r.RandomClass()
	hero_backgroung := r.RandomBackground()
	hero_al := r.RandomAlignment()

	total_hero := hero_race + hero_class + hero_backgroung + hero_al
	return total_hero
}
