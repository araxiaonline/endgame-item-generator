package models

import "fmt"

type Dungeon struct {
	Id          int    `db:"Id"`
	Name        string `db:"Name"`
	Level       int
	ExpansionId int `db:"ExpansionId"`
}

// dungeon instance id : avg level
var dungeons = map[int]int{

	// Classic WoW dungeons
	389: 18, // Ragefire Chasm
	43:  25, // Wailing Caverns
	36:  23, // The Deadmines
	33:  30, // Shadowfang Keep
	34:  30, // The Stockade
	48:  32, // Blackfathom Deeps
	90:  38, // Gnomeregan
	47:  40, // Razorfen Kraul
	189: 45, // Scarlet Monastery (Graveyard)
	289: 60, // Scholomance
	329: 60, // Stratholme
	229: 60, // Blackrock Spire (Lower)
	230: 60, // Blackrock Spire (Upper)
	429: 60, // Dire Maul
	209: 54, // Zul'Farrak
	349: 55, // Maraudon
	269: 57, // Temple of Atal'Hakkar

	// The Burning Crusade dungeons
	540: 62, // Hellfire Citadel: Hellfire Ramparts
	542: 63, // Hellfire Citadel: The Blood Furnace
	547: 64, // Coilfang Reservoir: The Slave Pens
	546: 65, // Coilfang Reservoir: The Underbog
	557: 66, // Auchindoun: Mana-Tombs
	558: 67, // Auchindoun: Auchenai Crypts
	560: 68, // Caverns of Time: Old Hillsbrad Foothills
	556: 69, // Auchindoun: Sethekk Halls
	545: 70, // Coilfang Reservoir: The Steamvault
	555: 70, // Auchindoun: Shadow Labyrinth
	543: 70, // Hellfire Citadel: The Shattered Halls
	553: 70, // Tempest Keep: The Botanica
	554: 70, // Tempest Keep: The Mechanar
	552: 70, // Tempest Keep: The Arcatraz
	585: 70, // Magisters' Terrace

	// Wrath of the Lich King dungeons
	70:  72, // Utgarde Keep
	129: 74, // Azjol-Nerub
	619: 75, // Ahn'kahet: The Old Kingdom
	576: 73, // The Nexus
	600: 76, // Drak'Tharon Keep
	608: 77, // The Violet Hold
	604: 78, // Gundrak
	599: 79, // Halls of Stone
	602: 80, // Halls of Lightning
	578: 79, // The Oculus
	575: 80, // Utgarde Pinnacle
	595: 80, // The Culling of Stratholme
	650: 80, // Trial of the Champion
	632: 80, // The Forge of Souls
	658: 80, // Pit of Saron
	668: 80, // Halls of Reflection
}

func (db Database) GetDungeons(expansionId int) ([]Dungeon, error) {
	dungeons := []Dungeon{}

	sql := `
		SELECT ID as Id, MapName_Lang_enUS as Name, ExpansionID as ExpansionId 
		FROM map_dbc 
		WHERE InstanceType = 1 AND Name NOT LIKE '%unused%';
	`
	var err error
	if expansionId != -1 {
		sql = sql + "AND ExpansionID = ?"
		err = db.client.Select(&dungeons, sql, expansionId)
	} else {
		err = db.client.Select(&dungeons, sql)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get dungeons %v", err)
	}

	return dungeons, nil
}
