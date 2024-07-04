package models

import (
	"fmt"
	"strconv"

	"github.com/araxiaonline/endgame-item-generator/utils"
	"github.com/thoas/go-funk"
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

var SpellEffectStatMap = map[int][]int{
	10: {41},
	17: {38, 39},
	22: {14},
	23: {15},
	58: {38, 39},
	67: {1},
}

// result of a stat conversion from spell to raw stats on item
type SpellStat struct {
	StatType  int
	StatValue int
	Budget    int
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

	return spell, nil
}

func (s Spell) effectNeedsScaled() bool {
	if s.Effect1 == 0 {
		return false
	}

	for _, effect := range SpellEffects {
		if s.Effect1 == effect || s.Effect2 == effect || s.Effect3 == effect {
			return true
		}
	}
	return false
}

func (s Spell) auraEffectNeedsScaled() bool {
	if s.EffectAura1 == 0 {
		return false
	}

	for _, effect := range SpellAuraEffects {
		if s.EffectAura1 == effect || s.EffectAura2 == effect || s.EffectAura3 == effect {
			return true
		}
	}
	return false
}

func (s Spell) HasAuraEffect() bool {
	return s.EffectAura1 != 0 || s.EffectAura2 != 0 || s.EffectAura3 != 0
}

// If a spell effect can be converted over to a primary stat addition
func EffectCanBeConv(effect int) bool {
	statMods := [...]int{17, 10, 22, 23, 58}
	return funk.Contains(statMods, effect)
}

func AuraEffectCanBeConv(effect int) bool {

	statMods := [...]int{13, 22, 34, 85, 99, 107, 115, 124, 135, 189}
	return funk.Contains(statMods, effect)
}

func (s Spell) ConvertToStats() ([]SpellStat, error) {
	stats := []SpellStat{}

	if s.Effect1 == 0 && s.EffectAura1 == 0 {
		fmt.Print("Spell does not have an effect1 or autaEffect1")
		return stats, nil
	}

	if EffectCanBeConv(s.Effect1) {
		stats = append(stats, SpellStat{
			StatType:  SpellEffectStatMap[s.Effect1][0],
			StatValue: s.EffectBasePoints1,
			Budget:    s.EffectBonusMultiplier1,
		})
	}

	if s.Effect2 != 0 && EffectCanBeConv(s.Effect2) {
		stats = append(stats, SpellStat{
			StatType:  SpellEffectStatMap[s.Effect2][0],
			StatValue: s.EffectBasePoints2,
			Budget:    s.EffectBonusMultiplier2,
		})
	}

	if s.Effect3 != 0 && EffectCanBeConv(s.Effect3) {
		stats = append(stats, SpellStat{
			StatType:  SpellEffectStatMap[s.Effect3][0],
			StatValue: s.EffectBasePoints3,
			Budget:    s.EffectBonusMultiplier3,
		})
	}

	return stats, nil
}
