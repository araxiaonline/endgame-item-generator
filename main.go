package main

import (
	"log"

	"github.com/araxiaonline/endgame-item-generator/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	models.Connect()

	bosses, err := models.DB.GetBosses(189)
	if err != nil {
		log.Fatal(err)
	}

	for _, boss := range bosses {

		log.Printf("Getting loot for Boss: %s\n", boss.Name)

		items, err := models.DB.GetBossLoot(boss.Entry)
		if err != nil {
			log.Fatal(err)
		}

		for _, item := range items {
			log.Println(item)
		}

	}

	defer models.DB.Close()

	// theItem := models.Item{}
	// s := reflect.ValueOf(&theItem).Elem()
	// numCols := s.NumField()
	// log.Println(numCols)
	// columns := make([]interface{}, numCols)
	// fmt.Print(columns)
	// for i := 0; i < numCols; i++ {
	// 	field := s.Field(i)
	// 	columns[i] = field.Addr().Interface()
	// }
	// fmt.Println(columns...)

	// var rows *sql.Rows
	// rows, err = DB.Query("SELECT name, entry FROM item_template where name like 'Hypnotic B%';")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for rows.Next() {
	// 	var name string
	// 	var entry int

	// 	err := rows.Scan(&name, &entry)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(name, entry)
	// }

	// rows.Close()

	// defer db.Close()
}
