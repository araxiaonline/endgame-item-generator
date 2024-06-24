package models

import (
	"errors"

	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Boss struct {
	entry        int
	name         string
	scriptName   string
	xpMultiplier int
}

func (db Database) GetBosses(mapId int) ([]Boss, error) {

	if mapId == 0 {
		return nil, errors.New("mapId cannot be 0")
	}

	var bosses []Boss
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

	rows, err := db.client.Query(sql, mapId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var boss Boss
		err := rows.Scan(&boss.entry, &boss.name, &boss.scriptName, &boss.xpMultiplier)
		if err != nil {
			return nil, err
		}
		bosses = append(bosses, boss)
	}

	print(len(bosses))

	return bosses, nil
}
