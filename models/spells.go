package models

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/araxiaonline/endgame-item-generator/utils"
	"github.com/thoas/go-funk"
)

// list of spell effects that require scaling
var SpellEffects = [...]int{
	2,  // School Damage
	6,  // AppyAura
	9,  // HealthLeech
	10, // Heal
	30, // Restores Mana
	35, // Apply Area Aura
}

// list of spell aura effects that require scaling
var SpellAuraEffects = [...]int{
	3,   // DOT
	8,   // HOT
	13,  // Modifies Spell Damage Done
	15,  // Modifies Damage Shield
	22,  // Modifies Resistance
	34,  // Modifies HEalth
	85,  // Modifies Mana Regen
	99,  // Modifies Attack Power
	124, // Modifies Range Attack Power
	135, // Modifies Healing Done
	189, // Modifies Critical Strike
}

var AuraEffectsStatMap = map[int]int{
	8:   46,
	13:  45,
	85:  43,
	99:  38,
	124: 38,
	135: 45,
}

// Usually in EffectMiscValueA to describe what the Aura modifies
var SpellModifiers = [...]int{
	0,  // damage
	1,  // duration
	2,  // threat
	3,  // effect1
	4,  // charges
	5,  // range
	6,  // radius
	7,  // crit chance
	8,  // all effects
	9,  // No pushback
	10, // Cast Time
	11, // CD
	12, // effect2
	13, // ignore armor
	14, // cost
	15, // crit damage bonus
	16, // resist miss chance
	17, // jump targets
	18, // Chance of success
	19, // Amplitude
	20, // Dmg multiplier
	21, // GCD
	22, // DoT
	23, // effect3
	24, // bonus multiplier
	26, // PPM
	27, // value multiplier
	28, // resist dispel chance
	29, // crit damage bonus 2
}

// result of a stat conversion from spell to raw stats on item
type ConvItemStat struct {
	StatType  int
	StatValue int
	Budget    int
}

// Spell Effect with max value for effect storage
type SpellEffect struct {
	Effect        int
	BasePoints    int
	DieSides      int
	CalculatedMax int
}

// DB Mapping from spell_dbc
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

func calcMaxValue(base int, sides int) int {
	if base < 0 {
		return base - sides
	}

	return base + sides
}

// get a List of the spell effects (not auras) that need to be scaled
func (s Spell) GetSpellEffects() []SpellEffect {
	effects := make([]SpellEffect, 0)

	effects = append(effects, SpellEffect{
		Effect:        s.Effect1,
		BasePoints:    s.EffectBasePoints1,
		DieSides:      s.EffectDieSides1,
		CalculatedMax: calcMaxValue(s.EffectBasePoints1, s.EffectDieSides1),
	})

	effects = append(effects, SpellEffect{
		Effect:        s.Effect2,
		BasePoints:    s.EffectBasePoints2,
		DieSides:      s.EffectDieSides2,
		CalculatedMax: calcMaxValue(s.EffectBasePoints2, s.EffectDieSides2),
	})

	effects = append(effects, SpellEffect{
		Effect:        s.Effect3,
		BasePoints:    s.EffectBasePoints3,
		DieSides:      s.EffectDieSides3,
		CalculatedMax: calcMaxValue(s.EffectBasePoints3, s.EffectDieSides3),
	})

	return effects
}

// Get he effects and calculate the max value for the a
func (s Spell) GetAuraEffects() []SpellEffect {
	effects := make([]SpellEffect, 0)

	effects = append(effects, SpellEffect{
		Effect:        s.EffectAura1,
		BasePoints:    s.EffectBasePoints1,
		DieSides:      s.EffectDieSides1,
		CalculatedMax: calcMaxValue(s.EffectBasePoints1, s.EffectDieSides1),
	})

	effects = append(effects, SpellEffect{
		Effect:        s.EffectAura2,
		BasePoints:    s.EffectBasePoints2,
		DieSides:      s.EffectDieSides2,
		CalculatedMax: calcMaxValue(s.EffectBasePoints2, s.EffectDieSides2),
	})

	effects = append(effects, SpellEffect{
		Effect:        s.EffectAura3,
		BasePoints:    s.EffectBasePoints3,
		DieSides:      s.EffectDieSides3,
		CalculatedMax: calcMaxValue(s.EffectBasePoints3, s.EffectDieSides3),
	})

	return effects
}

// this spell effect has stats or effects that need to be scaled
func (s Spell) SpellEffectsNeedsScaled() bool {
	if s.Effect1 == 0 {
		return false
	}

	needsScaled := false
	effects := s.GetSpellEffects()
	for _, e := range effects {

		if !funk.Contains(SpellEffects, e.Effect) || e.Effect == 6 {
			continue
		}
		needsScaled = true
	}

	return needsScaled
}

// this aura effect has stats or effects that need to be scaled
func (s Spell) AuraEffectNeedsScaled() bool {
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

func AuraEffectCanBeConv(effect int) bool {
	statMods := [...]int{8, 13, 22, 34, 85, 99, 124, 135, 189}
	return funk.Contains(statMods, effect)
}

// Lookup details about the effect and return the stat type -1 indicates not found
func convertAuraEffect(effect int) int {
	if !funk.Contains(AuraEffectsStatMap, effect) {
		log.Printf("effect %v not found in SpellEffectStatMap skipping", effect)
		return -1
	}

	return AuraEffectsStatMap[effect]
}

// Converts spell buffs to item stats making it easier to convert and normalize
func (s Spell) ConvertToStats() ([]ConvItemStat, error) {
	stats := []ConvItemStat{}

	if s.Effect1 == 0 && s.EffectAura1 == 0 {
		return stats, fmt.Errorf("spell does not have an effect1 or auraEffect1")
	}

	effects := s.GetAuraEffects()
	var seen []int
	for _, e := range effects {
		if !AuraEffectCanBeConv(e.Effect) {
			continue
		}

		statId := convertAuraEffect(e.Effect)
		if statId == -1 {
			continue
		}

		if funk.Contains(seen, statId) {
			continue
		}

		// keep track if we have already seen this stat so we do not duplicate
		// Wotlk changed everything to spell power so might as well do the same in
		// scaling process.
		seen = append(seen, statId)
		statMod := float64(StatModifiers[statId])
		stats = append(stats, ConvItemStat{
			StatType:  statId,
			StatValue: e.CalculatedMax,
			Budget:    int(math.Ceil(float64(e.CalculatedMax) * statMod)),
		})
	}

	return stats, nil
}
