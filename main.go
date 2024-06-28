package main

import (
	"fmt"
	"log"

	"github.com/araxiaonline/endgame-item-generator/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	models.Connect()

	weapon, err := models.DB.GetItem(13982)
	weapon.ScaleDPS(350)
	log.Printf("Weapon: %v-%v", *weapon.MinDmg1, *weapon.MaxDmg1)
	log.Printf("Weapon: %v-%v", *weapon.MinDmg2, *weapon.MaxDmg2)
	if err != nil {
		log.Fatal(err)
	}

	bosses, err := models.DB.GetBosses(540)
	if err != nil {
		log.Fatal(err)
	}

	for _, boss := range bosses {

		// log.Printf("Getting loot for Boss: %s\n", boss.Name)

		items, err := models.DB.GetBossLoot(boss.Entry)
		if err != nil {
			log.Fatal(err)
		}

		for _, item := range items {

			fmt.Printf("Item %v ItemLevel %v \n", item.Name, item.ItemLevel)

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
