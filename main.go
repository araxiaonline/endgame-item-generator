package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/araxiaonline/endgame-item-generator/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	godotenv.Load()
	models.Connect()

	debug := flag.Bool("debug", false, "Enable verbose logging inside generator")
	flag.Parse()

	if *debug {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(io.Discard)
	}

	// main loop
	dungeons, err := models.DB.GetDungeons(-1)
	if err != nil {
		log.Panicf("failed to get dungeons for expansion %v error: %v", 0, err)
	}

	for _, dungeon := range dungeons {
		log.Printf("+++++Dungeon:  %s ID: %v level %v \n", dungeon.Name, dungeon.Id, dungeon.Level)

		bosses, err := models.DB.GetBosses(dungeon.Id)
		if err != nil {
			log.Fatal("failed to get bosses")
		}

		scaleValue := 305

		if dungeon.Level == 60 {
			scaleValue = 315
		}

		if dungeon.ExpansionId == 1 && dungeon.Level <= 70 {
			scaleValue = 308
		}

		if dungeon.ExpansionId == 1 && dungeon.Level == 70 {
			scaleValue = 315

		}

		if dungeon.ExpansionId == 2 && dungeon.Level <= 80 {
			scaleValue = 309
		}

		if dungeon.ExpansionId == 2 && dungeon.Level == 80 {
			scaleValue = 320
		}

		for _, boss := range bosses {

			items, err := models.DB.GetBossLoot(boss.Entry)
			log.Printf("++++++++++ Boss: %s Entry: %v has %v items\n", boss.Name, boss.Entry, len(items))
			if err != nil {
				log.Fatal(err)
				continue
			}

			for _, item := range items {

				_, error := item.ScaleItem(scaleValue, 4)
				if error != nil {
					log.Printf("Failed to scale item: %v", error)
					continue
				}

				fmt.Printf("\n-- Item %v Entry: %v ItemLevel %v \n", item.Name, item.Entry, *item.ItemLevel)
				if scaleValue >= 315 {
					fmt.Print(ItemToSql(item, 81, 3))
				} else {
					fmt.Print(ItemToSql(item, 80, 3))
				}

			}

		}

		items2, err := models.DB.GetAddlDungeonDrops(dungeon.Id)
		if err != nil {
			log.Printf("failed to get additional loot for dungeon %v - err: %v", dungeon.Id, err)
		}

		for _, item := range items2 {

			// reduce item level of dungeon random drops since they are not boss fights
			adjScaleValue := scaleValue - 4

			_, error := item.ScaleItem(adjScaleValue, 3)
			if error != nil {
				log.Printf("Failed to scale item: %v", error)
				continue
			}

			fmt.Printf("\n-- Item %v Entry: %v ItemLevel %v \n", item.Name, item.Entry, *item.ItemLevel)

			if scaleValue >= 315 {
				fmt.Print(ItemToSql(item, 81, 3))
			} else {
				fmt.Print(ItemToSql(item, 80, 3))
			}

		}
		log.Printf("++++++++++ AdditionalLoot Count: %v\n", len(items2))
	}

	defer models.DB.Close()
}
