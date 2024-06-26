package models

import (
	"fmt"
	"log"
	"math/rand/v2"
	"reflect"
)

/**
 * For details about values of item int values use link below
 * @link https://www.azerothcore.org/wiki/item_template
 */
type Item struct {
	Entry          int
	Name           string
	DisplayId      int `db:"displayid"`
	Quality        *int
	ItemLevel      *int `db:"ItemLevel"`
	Class          *int
	Subclass       *int
	InventoryType  *int `db:"inventoryType"`
	AllowableClass *int `db:"allowableClass"`
	AllowableRace  *int `db:"allowableRace"`
	RequiredSkill  *int `db:"requiredSkill"`
	RequiredLevel  *int `db:"requiredLevel"`
	Durability     *int `db:"MaxDurability"`
	MinDmg1        *int `db:"dmg_min1"`
	MaxDmg1        *int `db:"dmg_max1"`
	MinDmg2        *int `db:"dmg_min2"`
	MaxDmg2        *int `db:"dmg_max2"`
	DmgType1       *int `db:"dmg_type1"`
	DmgType2       *int `db:"dmg_type2"`
	Delay          *float64
	Material       *int
	Sheath         *int
	StatsCount     *int `db:"statsCount"`
	StatType1      *int `db:"stat_type1"`
	StatValue1     *int `db:"stat_value1"`
	StatType2      *int `db:"stat_type2"`
	StatValue2     *int `db:"stat_value2"`
	StatType3      *int `db:"stat_type3"`
	StatValue3     *int `db:"stat_value3"`
	StatType4      *int `db:"stat_type4"`
	StatValue4     *int `db:"stat_value4"`
	StatType5      *int `db:"stat_type5"`
	StatValue5     *int `db:"stat_value5"`
	StatType6      *int `db:"stat_type6"`
	StatValue6     *int `db:"stat_value6"`
	StatType7      *int `db:"stat_type7"`
	StatValue7     *int `db:"stat_value7"`
	StatType8      *int `db:"stat_type8"`
	StatValue8     *int `db:"stat_value8"`
	StatType9      *int `db:"stat_type9"`
	StatValue9     *int `db:"stat_value9"`
	StatType10     *int `db:"stat_type10"`
	StatValue10    *int `db:"stat_value10"`
}

func (item Item) GetPrimaryStat() (int, int, error) {
	var primaryStat int64
	var primaryVal int64

	values := reflect.ValueOf(item)
	for i := 1; i < 11; i++ {
		statType := values.FieldByName(fmt.Sprintf("StatType%v", i)).Elem().Int()
		// first check if the stat type is not in the primary stats str, agi, intellect, spirit, stamina
		if statType < 3 || statType > 7 {
			continue
		}

		statValue := values.FieldByName(fmt.Sprintf("StatValue%v", i)).Elem().Int()
		if statValue > primaryVal {
			primaryVal = statValue
			primaryStat = statType
		}
	}

	return int(primaryStat), int(primaryVal), nil
}

func (i Item) GetDpsModifier() (float64, error) {
	if i.Subclass == nil {
		return 0, fmt.Errorf("subclass on the item is not set")
	}

	if i.Quality == nil {
		return 0, fmt.Errorf("quality is not set")
	}

	typeModifier := 0.0
	// Is a One-Handed Weapon
	if *i.Subclass == 0 || *i.Subclass == 4 || *i.Subclass == 13 || *i.Subclass == 15 {
		typeModifier = 0.55
	}

	// Is a Two-Handed Weapon
	if *i.Subclass == 1 || *i.Subclass == 5 || *i.Subclass == 6 || *i.Subclass == 8 || *i.Subclass == 10 {
		typeModifier = 0.65
	}

	// Ranged Weapons
	if *i.Subclass == 2 || *i.Subclass == 3 || *i.Subclass == 16 || *i.Subclass == 18 {
		typeModifier = 0.60
	}

	// Wands
	if *i.Subclass == 17 {
		typeModifier = 0.65
	}

	qualityModifier := 1.0

	// Add the quality modifier for the DPS calculation
	if *i.Quality == 2 {
		qualityModifier = 1.25
	}
	if *i.Quality == 3 {
		qualityModifier = 1.38
	}
	if *i.Quality == 4 {
		qualityModifier = 1.5
	}

	if typeModifier == 0 {
		return 0, fmt.Errorf("Item subclass is not a weapon")
	}

	return (qualityModifier * typeModifier), nil
}

// Get the current expected DPS of the item bsed on the min and max damage and delay
func (item Item) GetDPS() (float64, error) {

	if item.MinDmg1 == nil || item.MaxDmg1 == nil {
		return 0, fmt.Errorf("MinDmg1 or MaxDmg1 is not set")
	}

	if item.Delay == nil {
		return 0, fmt.Errorf("delay is not set")
	}

	dps := (float64(*item.MinDmg1+*item.MaxDmg1) / 2.0) / float64(*item.Delay)
	return dps, nil
}

func (item *Item) ScaleDPS() (float64, error) {
	if item.ItemLevel == nil {
		return 0, fmt.Errorf("ItemLevel is not set")
	}

	if item.Delay == nil {
		return 0, fmt.Errorf("delay is not set")
	}

	modifier, err := item.GetDpsModifier()
	if err != nil {
		log.Fatalf("Error getting DPS modifier: %v", err)
	}

	dps := float64(*item.ItemLevel) * modifier
	minimum := int(dps * float64(*item.Delay) / 100 * float64(100-(rand.IntN(15)+15)))
	maximum := int(dps * float64(*item.Delay) / 100 * float64(100+(rand.IntN(15)+15)))

	item.MinDmg1 = &minimum
	item.MaxDmg1 = &maximum

	log.Printf("Item %s MinDPS: %v", item.Name, minimum)
	log.Printf("Item %s MaxDPS: %v", item.Name, maximum)

	return dps, nil
}
