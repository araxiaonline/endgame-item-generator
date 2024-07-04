package models

import (
	"fmt"
	"strconv"

	"github.com/araxiaonline/endgame-item-generator/utils"
)

var SpellEffects = [...]int{
	2,   // School Damage
	6,   // AppyAura
	9,   // HealthLeech
	10,  // Heal
	17,  // WeaponDamage
	22,  // Parry
	23,  // Block
	31,  // WeaponDmg % Increase
	35,  // Apply Area Aura
	58,  // Weapon Damange Increase
	67,  // Max Health
	121, // Normalized Weapon Damage
}

var SpellAuraEffects = [...]int{
	3,   // DOT
	8,   // HOT
	13,  // Modifies Damage Done
	14,  // Modifies Damange Taken
	15,  // Modifies Damage Shield
	22,  // Modifies Resistance
	29,  // Modifies a Stage
	34,  // Modifies HEalth
	85,  // Modifies Power Regen
	99,  // Modifies Attack Power
	107, // Modifier Flat
	115, // Modifies Healing
	124, // Modifies Range Attack Power
	127, // Melee Attack Bonus Dmg
	135, // Modifies Healing Done
	189, // Modifies Armor Rating
}

type Spell struct {
	ID                        int    `db:"ID"`
	Name                      string `db:"Name_Lang_enUS"`
	Description               string `db:"Description_Lang_enUS"`
	AuraDescription           string `db:"AuraDescription_Lang_enUS"`
	ProcChance                int    `db:"ProcChance"`
	SpellLevel                int    `db:"SpellLevel"`
	Effect1                   int    `db:"Effect_1"`
	Effect2                   int    `db:"Effect_2"`
	Effect3                   int    `db:"Effect_3"`
	EffectDieSides1           int    `db:"EffectDieSides_1"`
	EffectDieSides2           int    `db:"EffectDieSides_2"`
	EffectDieSides3           int    `db:"EffectDieSides_3"`
	EffectRealPointsPerLevel1 int    `db:"EffectRealPointsPerLevel_1"`
	EffectRealPointsPerLevel2 int    `db:"EffectRealPointsPerLevel_2"`
	EffectRealPointsPerLevel3 int    `db:"EffectRealPointsPerLevel_3"`
	EffectBasePoints1         int    `db:"EffectBasePoints_1"`
	EffectBasePoints2         int    `db:"EffectBasePoints_2"`
	EffectBasePoints3         int    `db:"EffectBasePoints_3"`
	EffectAura1               int    `db:"EffectAura_1"`
	EffectAura2               int    `db:"EffectAura_2"`
	EffectAura3               int    `db:"EffectAura_3"`
	EffectBonusMultiplier1    int    `db:"EffectBonusMultiplier_1"`
	EffectBonusMultiplier2    int    `db:"EffectBonusMultiplier_2"`
	EffectBonusMultiplier3    int    `db:"EffectBonusMultiplier_3"`
}

func (db Database) GetSpell(id int) (Spell, error) {

	if id == 0 {
		return Spell{}, fmt.Errorf("id cannot be 0")
	}

	spell := Spell{}
	sql := "SELECT " + utils.GetSpellFields() + " FROM `spell_dbc` WHERE ID = ? -- " + strconv.Itoa(id)

	err := db.client.Get(&spell, sql, id)
	if err != nil {
		return Spell{}, fmt.Errorf("failed to get spell: %v", err)
	}

	// log.Printf("%s was found for id %d", spell.Name, spell.ID)

	return spell, nil
}
