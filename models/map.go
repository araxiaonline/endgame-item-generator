package models

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
)

type Boss struct {
	Entry              int
	Name               string
	ScriptName         string `db:"ScriptName"`
	ExperienceModifier int    `db:"ExperienceModifier"`
}

/**
 * For details about values of item int values use link below
 * @link https://www.azerothcore.org/wiki/item_template
 */
type Item struct {
	Entry          int
	Name           string
	DisplayId      int `db:"displayid"`
	Quality        int
	ItemLevel      int `db:"ItemLevel"`
	Class          int
	Subclass       int
	InventoryType  int `db:"inventoryType"`
	AllowableClass int `db:"allowableClass"`
	AllowableRace  int `db:"allowableRace"`
	RequiredSkill  int `db:"requiredSkill"`
	RequiredLevel  int `db:"requiredLevel"`
	Durability     int `db:"MaxDurability"`
	MinDmg1        int `db:"dmg_min1"`
	MaxDmg1        int `db:"dmg_max1"`
	MinDmg2        int `db:"dmg_min2"`
	MaxDmg2        int `db:"dmg_max2"`
	Delay          int
	Material       int
	Sheath         int
	StatsCount     int `db:"statsCount"`
	StatType1      int `db:"stat_type1"`
	StatValue1     int `db:"stat_value1"`
	StatType2      int `db:"stat_type2"`
	StatValue2     int `db:"stat_value2"`
	StatType3      int `db:"stat_type3"`
	StatValue3     int `db:"stat_value3"`
	StatType4      int `db:"stat_type4"`
	StatValue4     int `db:"stat_value4"`
	StatType5      int `db:"stat_type5"`
	StatValue5     int `db:"stat_value5"`
	StatType6      int `db:"stat_type6"`
	StatValue6     int `db:"stat_value6"`
	StatType7      int `db:"stat_type7"`
	StatValue7     int `db:"stat_value7"`
	StatType8      int `db:"stat_type8"`
	StatValue8     int `db:"stat_value8"`
	StatType9      int `db:"stat_type9"`
	StatValue9     int `db:"stat_value9"`
	StatType10     int `db:"stat_type10"`
	StatValue10    int `db:"stat_value10"`
}

func (db Database) GetBosses(mapId int) ([]Boss, error) {

	if mapId == 0 {
		return nil, errors.New("mapId cannot be 0")
	}

	bosses := []Boss{}
	var sql string

	// 540 is pre-classic dungeons so XP Multiplier is best way to determine bosses / rare mobs
	if mapId < 540 {
		sql = `
			SELECT ct.entry, ct.name, ct.ScriptName, ct.ExperienceModifier from acore_world.creature c
			join acore_world.creature_template ct ON(c.id1 = ct.entry)  where map = ? and ExperienceModifier >= 2;
		`
	} else {
		sql = `
			SELECT ct.entry, ct.name, ct.ScriptName, ct.ExperienceModifier from acore_world.creature c
    		join acore_world.creature_template ct ON(c.id1 = ct.entry)  where map = ? and ct.ScriptName Like 'boss_%'
		`
	}

	err := db.client.Select(&bosses, sql, mapId)
	if err != nil {
		return nil, err
	}

	return bosses, nil
}

func (db Database) GetBossLoot(bossId int) ([]Item, error) {
	if bossId == 0 {
		return nil, errors.New("bossId cannot be 0")
	}

	items := []Item{}
	sql := `
	SELECT
		entry, name, displayid,
		quality, ItemLevel, class, subclass, inventoryType,
		allowableClass, allowableRace,
		requiredSkill, requiredLevel,
		dmg_min1, dmg_max1,
		dmg_min2,dmg_max2,
		delay, material, sheath, MaxDurability,
		statsCount,
		stat_type1, stat_value1,
		stat_type2, stat_value2,
		stat_type3, stat_value3,
		stat_type4, stat_value4,
		stat_type5, stat_value5,
		stat_type6, stat_value6,
		stat_type7, stat_value7,
		stat_type8, stat_value8,
		stat_type9, stat_value9,
		stat_type10, stat_value10

	from acore_world.item_template 
	where 
	entry in
		(SELECT item from acore_world.creature_loot_template where entry = ? and GroupId != 0 and Reference = 0)
	and Quality > 2
	and StatsCount > 0
	`

	udb := db.client.Unsafe()
	err := udb.Select(&items, sql, bossId)
	if err != nil {
		return nil, err
	}

	return items, nil
}
