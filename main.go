package main

import (
	"fmt"
	"log"

	"github.com/araxiaonline/endgame-item-generator/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	godotenv.Load()
	models.Connect()

	weapon, err := models.DB.GetItem(13982)
	weapon.ScaleDPS(350)
	// log.Printf("Weapon: %v-%v", *weapon.MinDmg1, *weapon.MaxDmg1)
	// log.Printf("Weapon: %v-%v", *weapon.MinDmg2, *weapon.MaxDmg2)
	if err != nil {
		log.Fatal(err)
	}

	bosses, err := models.DB.GetBosses(540)
	if err != nil {
		log.Fatal(err)
	}

	for _, boss := range bosses {

		items, err := models.DB.GetBossLoot(boss.Entry)
		log.Printf("Boss: %s Entry: %v has %v items\n", boss.Name, boss.Entry, len(items))
		if err != nil {
			log.Fatal(err)
		}

		for _, item := range items {

			fmt.Printf("\nItem %v Entry: %v ItemLevel %v \n", item.Name, item.Entry, *item.ItemLevel)
			// item.GetStatPercents()

			// if *item.SpellId1 != 0 {
			// 	spell, err := models.DB.GetSpell(*item.SpellId1)
			// 	if err != nil {
			// 		log.Printf("failed to get the spell: %v error: %v", *item.SpellId1, err)
			// 	}

			// log.Printf("Spell %v Spell Effects 1: %v 2: %v, 3: %v \n", spell.Name, spell.Effect1, spell.Effect2, spell.Effect3)
			// log.Printf("Spell Aura 1: %v 2: %v, 3: %v \n", spell.EffectAura1, spell.EffectAura2, spell.EffectAura3)

			// convStats, err := spell.ConvertToStats()
			// if err != nil {
			// 	log.Printf("Failed to convert spell to stats: %v", err)
			// }

			// scaleItemStats := item.GetStatPercents(convStats)
			// for statId, stat := range scaleItemStats {
			// 	log.Printf("StatId: %v Type: %s Value: %v Percent: %v", statId, stat.Type, stat.Value, stat.Percent)
			// }
			//				log.Printf("Scaled Spell Stats: %v\n", convStats)

			// }

			_, error := item.ScaleItem()
			if error != nil {
				log.Printf("Failed to scale item: %v", error)
			}

			// stat, value, err := item.GetPrimaryStat()
			if err != nil {
				log.Fatal(err)
			}
			// fmt.Println(stat, value)

		}
	}

	// iLevel := 219
	// qual := 3
	// delay := 2.60
	// sub := 0
	// myItem := models.Item{Name: "Hypnotic Blade", ItemLevel: &iLevel, Quality: &qual, Delay: &delay, Subclass: &sub}
	// dps, err := myItem.ScaleDPS()

	// log.Printf("Item %s DPS: %.1f", myItem.Name, dps)
	defer models.DB.Close()
}
