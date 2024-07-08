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

	bosses, err := models.DB.GetBosses(229)
	if err != nil {
		log.Fatal("failed to get bosses")
	}

	for _, boss := range bosses {

		items, err := models.DB.GetBossLoot(boss.Entry)
		log.Printf("Boss: %s Entry: %v has %v items\n", boss.Name, boss.Entry, len(items))
		if err != nil {
			log.Fatal(err)
		}

		for _, item := range items {

			log.Printf("\nItem %v Entry: %v ItemLevel %v \n", item.Name, item.Entry, *item.ItemLevel)

			_, error := item.ScaleItem(320, 3)
			fmt.Print(ItemToSql(item, 80, 3))
			if error != nil {
				log.Printf("Failed to scale item: %v", error)
			}

			// stat, value, err := item.GetPrimaryStat()
			if err != nil {
				log.Fatal(err)
			}

		}
	}

	defer models.DB.Close()
}
