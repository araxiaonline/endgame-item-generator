package models

import (
	"errors"

	"github.com/araxiaonline/endgame-item-generator/utils"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
)

type Boss struct {
	Entry              int
	Name               string
	ScriptName         string `db:"ScriptName"`
	ExperienceModifier int    `db:"ExperienceModifier"`
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

	// This will first find items that are not in the reference boss loot table
	items := []Item{}
	sql := `
	SELECT ` + utils.GetItemFields() + `
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
