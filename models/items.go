package models

import (
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"reflect"

	"github.com/araxiaonline/endgame-item-generator/utils"
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
	SpellId1       *int `db:"spellid_1"`
	SpellId2       *int `db:"spellid_2"`
	SpellId3       *int `db:"spellid_3"`
	SpellTrigger1  *int `db:"spelltrigger_1"`
	SpellTrigger2  *int `db:"spelltrigger_2"`
	SpellTrigger3  *int `db:"spelltrigger_3"`
}

var armorModifiers = map[int]float64{
	1:  0.813, // Head
	2:  1.0,   // Neck
	3:  0.75,  // Shoulder
	4:  1.0,   // Shirt (using the same as Chest for simplicity)
	5:  1.0,   // Chest
	6:  0.562, // Waist
	7:  0.875, // Legs
	8:  0.688, // Feet
	9:  0.437, // Wrists
	10: 0.625, // Hands
	11: 1.0,   // Finger
	20: 1.0,   // Robe (using the same as Chest for simplicity)
}

var weaponModifiers = map[int]float64{
	13: 0.42, // One-Hand (not to confuse with Off-Hand = 22)
	14: 0.56, // Shield (class = armor, not weapon even if in weapon slot)
	15: 0.32, // Ranged (Bows) (see also Ranged right = 26)
	16: 0.56, // Back
	17: 1.0,  // Two-Hand
	18: 1.0,  // Bag (assuming same as Chest for simplicity)
	19: 1.0,  // Tabard (assuming same as Chest for simplicity)
	20: 1.0,  // Robe (see also Chest = 5)
	21: 1.0,  // Main hand
	22: 0.42, // Off Hand weapons (see also One-Hand = 13)
	23: 0.56, // Held in Off-Hand (class = armor, not weapon even if in weapon slot)
	24: 1.0,  // Ammo (assuming same as Chest for simplicity)
	25: 0.32, // Thrown
	26: 0.32, // Ranged right (Wands, Guns) (see also Ranged = 15)
	27: 1.0,  // Quiver (assuming same as Chest for simplicity)
}

var qualityModifiers = map[int]float64{
	0: 1.0, // Common
	1: 1.1, // Uncommon
	2: 1.2, // Rare
	3: 1.3, // Epic
	4: 1.4, // Legendary
	5: 1.5, // Artifact
}

var statModifiers = map[int]float64{
	0:  1.0,  // ITEM_MOD_MANA
	1:  1.0,  // ITEM_MOD_HEALTH
	3:  1.0,  // ITEM_MOD_AGILITY
	4:  1.0,  // ITEM_MOD_STRENGTH
	5:  1.0,  // ITEM_MOD_INTELLECT
	6:  1.0,  // ITEM_MOD_SPIRIT
	7:  1.0,  // ITEM_MOD_STAMINA
	12: 1.0,  // ITEM_MOD_DEFENSE_SKILL_RATING
	13: 1.0,  // ITEM_MOD_DODGE_RATING
	14: 1.0,  // ITEM_MOD_PARRY_RATING
	15: 1.0,  // ITEM_MOD_BLOCK_RATING
	16: 1.0,  // ITEM_MOD_HIT_MELEE_RATING
	17: 1.0,  // ITEM_MOD_HIT_RANGED_RATING
	18: 1.0,  // ITEM_MOD_HIT_SPELL_RATING
	19: 1.0,  // ITEM_MOD_CRIT_MELEE_RATING
	20: 1.0,  // ITEM_MOD_CRIT_RANGED_RATING
	21: 1.0,  // ITEM_MOD_CRIT_SPELL_RATING
	22: 1.0,  // ITEM_MOD_HIT_TAKEN_MELEE_RATING
	23: 1.0,  // ITEM_MOD_HIT_TAKEN_RANGED_RATING
	24: 1.0,  // ITEM_MOD_HIT_TAKEN_SPELL_RATING
	25: 1.0,  // ITEM_MOD_CRIT_TAKEN_MELEE_RATING
	26: 1.0,  // ITEM_MOD_CRIT_TAKEN_RANGED_RATING
	27: 1.0,  // ITEM_MOD_CRIT_TAKEN_SPELL_RATING
	28: 1.0,  // ITEM_MOD_HASTE_MELEE_RATING
	29: 1.0,  // ITEM_MOD_HASTE_RANGED_RATING
	30: 1.0,  // ITEM_MOD_HASTE_SPELL_RATING
	31: 1.0,  // ITEM_MOD_HIT_RATING
	32: 1.0,  // ITEM_MOD_CRIT_RATING
	33: 1.0,  // ITEM_MOD_HIT_TAKEN_RATING
	34: 1.0,  // ITEM_MOD_CRIT_TAKEN_RATING
	35: 1.0,  // ITEM_MOD_RESILIENCE_RATING
	36: 1.0,  // ITEM_MOD_HASTE_RATING
	37: 1.0,  // ITEM_MOD_EXPERTISE_RATING
	38: 0.5,  // ITEM_MOD_ATTACK_POWER
	39: 1.0,  // ITEM_MOD_RANGED_ATTACK_POWER
	40: 1.0,  // ITEM_MOD_FERAL_ATTACK_POWER (not used as of 3.3)
	41: 1.0,  // ITEM_MOD_SPELL_HEALING_DONE
	42: 1.0,  // ITEM_MOD_SPELL_DAMAGE_DONE
	43: 2.5,  // ITEM_MOD_MANA_REGENERATION
	44: 1.0,  // ITEM_MOD_ARMOR_PENETRATION_RATING
	45: 0.5,  // ITEM_MOD_SPELL_POWER
	46: 1.0,  // ITEM_MOD_HEALTH_REGEN
	47: 2.0,  // ITEM_MOD_SPELL_PENETRATION
	48: 0.65, // ITEM_MOD_BLOCK_VALUE
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

// Scales and items dps damage numbers based on a desired item level.
func (item *Item) ScaleDPS(level int) (float64, error) {
	if item.ItemLevel == nil {
		return 0, fmt.Errorf("ItemLevel is not set")
	}

	if item.Delay == nil {
		return 0, fmt.Errorf("delay is not set")
	}

	modifier, err := item.GetDpsModifier()
	if err != nil {
		log.Fatalf("Error getting DPS modifier: %v", err)
		return 0.0, err
	}

	dps := modifier * float64(level)
	adjDps := (dps * (*item.Delay / 1000) / 100)

	//(((Y8*Y4)/100))*((100 - Y5)) Forumula from Weapon Item Genertor
	minimum := adjDps * float64(100-(rand.IntN(15)+25))
	maximum := adjDps * float64(100+(rand.IntN(15)+25))

	// If the weapon has secondary damage, scale that as well based on the ratio of the primary damage
	if item.MinDmg2 != nil && item.MaxDmg2 != nil {
		ratioMin := float64(*item.MinDmg2) / float64(*item.MinDmg1)
		ratioMax := float64(*item.MaxDmg2) / float64(*item.MaxDmg1)
		minimum2 := int(ratioMin * float64(minimum))
		maximum2 := int(ratioMax * float64(maximum))

		item.MinDmg2 = &minimum2
		item.MaxDmg2 = &maximum2

		// In order to balance the original scale up remove have of the secondary damage from primary
		minimum = minimum - float64((minimum2 / 2))
		maximum = maximum - float64((maximum2 / 2))
	}

	// item.MinDmg1 = &minimum
	var min int = int(minimum)
	var max int = int(maximum)
	item.MinDmg1 = &min
	item.MaxDmg1 = &max

	return dps, nil
}

func (d Database) GetItem(entry int) (Item, error) {
	if entry == 0 {
		return Item{}, fmt.Errorf("entry cannot be 0")
	}

	item := Item{}
	sql := "SELECT " + utils.GetItemFields() + " FROM item_template WHERE entry = ?"
	err := d.client.Get(&item, sql, entry)
	if err != nil {
		return Item{}, err
	}

	return item, nil
}

// Stat Formula scaler
// Ceiling of ((ItemLevel * QualityModifier * ItemTypeModifier)^1.7095 * %ofStats) ^ (1/1.7095)) / StatModifier
// i.e)   Green Strength Helmet  (((100 * 1.1 * 1.0)^1.705) * 1)^(1/1.7095) / 1.0 = 110 Strength on item

// Create a Map of stat percentages based on the current stat and how budgets are caluated
func (item Item) GetStatPercents() map[int]int64 {

	statMap := make(map[int]int64)
	statBudget := 0.0

	values := reflect.ValueOf(item)
	for i := 1; i < 11; i++ {
		var statValue = values.FieldByName(fmt.Sprintf("StatValue%v", i)).Elem().Int()
		var statType = values.FieldByName(fmt.Sprintf("StatType%v", i)).Elem().Int()
		if statValue == 0 {
			continue
		}

		statBudget += math.Round(float64(statValue) / statModifiers[int(statType)])
		statMap[int(statType)] = statValue
	}

	fmt.Printf("Stat Budget: %v\n", statBudget)
	fmt.Printf("Stat Map: %v\n", statMap)

	return map[int]int64{}
}
