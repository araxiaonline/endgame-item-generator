package models

import (
	"errors"
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
	Armor          *int `db:"armor"`
	Material       *int `db:"material"`
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
	StatsMap       map[int]*ItemStat
	ConvStatCount  int
	Spells         []Spell
}

var InvTypeModifiers = map[int]float64{
	1:  0.813, // Head
	2:  1.0,   // Neck
	3:  0.75,  // Shoulder
	5:  1.0,   // Chest
	6:  0.562, // Waist
	7:  0.875, // Legs
	8:  0.688, // Feet
	9:  0.437, // Wrists
	10: 0.625, // Hands
	11: 1.0,   // Finger
	13: 0.42,  // One-Hand (not to confuse with Off-Hand = 22)
	14: 0.56,  // Shield (class = armor, not weapon even if in weapon slot)
	15: 0.32,  // Ranged (Bows) (see also Ranged right = 26)
	16: 0.56,  // Back
	17: 1.0,   // Two-Hand
	18: 1.0,   // Bag (assuming same as Chest for simplicity)
	19: 1.0,   // Tabard (assuming same as Chest for simplicity)
	20: 1.0,   // Robe (see also Chest = 5)
	21: 1.0,   // Main hand
	22: 0.42,  // Off Hand weapons (see also One-Hand = 13)
	23: 0.56,  // Held in Off-Hand (class = armor, not weapon even if in weapon slot)
	24: 1.0,   // Ammo (assuming same as Chest for simplicity)
	25: 0.32,  // Thrown
	26: 0.32,  // Ranged right (Wands, Guns) (see also Ranged = 15)
	27: 1.0,   // Quiver (assuming same as Chest for simplicity)
}

var QualityModifiers = map[int]float64{
	0: 1.0, // Common
	1: 1.1, // Uncommon
	2: 1.2, // Rare
	3: 1.3, // Epic
	4: 1.5, // Legendary
	5: 1.7, // Artifact
}

var MaterialModifiers = map[int]float64{
	5: 5,   // Mail
	6: 9.0, // Plate
	7: 1.2, // Cloth
	8: 2.2, // Leather
	1: 20,  // Shield
}

var StatModifiers = map[int]float64{
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
	39: 0.5,  // ITEM_MOD_RANGED_ATTACK_POWER
	40: 0.5,  // ITEM_MOD_FERAL_ATTACK_POWER (not used as of 3.3)
	41: 0.5,  // ITEM_MOD_SPELL_HEALING_DONE
	42: 0.5,  // ITEM_MOD_SPELL_DAMAGE_DONE
	43: 2.5,  // ITEM_MOD_MANA_REGENERATION
	44: 1.0,  // ITEM_MOD_ARMOR_PENETRATION_RATING
	45: 0.5,  // ITEM_MOD_SPELL_POWER
	46: 1.0,  // ITEM_MOD_HEALTH_REGEN
	47: 2.0,  // ITEM_MOD_SPELL_PENETRATION
	48: 0.65, // ITEM_MOD_BLOCK_VALUE
}

type ItemStat struct {
	Value    int
	Percent  float64
	Type     string
	AdjValue float64
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
		typeModifier = 0.68
	}

	// Is a Two-Handed Weapon
	if *i.Subclass == 1 || *i.Subclass == 5 || *i.Subclass == 6 || *i.Subclass == 8 || *i.Subclass == 10 || *i.Subclass == 17 {
		typeModifier = 0.95
	}

	// Ranged Weapons
	if *i.Subclass == 2 || *i.Subclass == 3 || *i.Subclass == 16 || *i.Subclass == 18 {
		typeModifier = 0.75
	}

	// Wands
	if *i.Subclass == 19 {
		typeModifier = 0.70
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
		return 0, fmt.Errorf("Item subclass is not a weapon %v", *i.Subclass)
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

	dps := math.Round((float64(*item.MinDmg1+*item.MaxDmg1)/2.0)/float64(*item.Delay/1000)*100) / 100
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
	sql := "SELECT " + utils.GetItemFields("") + " FROM item_template WHERE entry = ?"
	err := d.client.Get(&item, sql, entry)
	if err != nil {
		return Item{}, err
	}

	return item, nil
}

// Create a Map of stat percentages based on the current stat and how budgets are caluated
func (item Item) GetStatPercents(spellStats []ConvItemStat) map[int]*ItemStat {

	statMap := make(map[int]*ItemStat)
	statBudget := 0.0

	values := reflect.ValueOf(item)
	for i := 1; i < 11; i++ {
		var statValue = values.FieldByName(fmt.Sprintf("StatValue%v", i)).Elem().Int()
		var statType = values.FieldByName(fmt.Sprintf("StatType%v", i)).Elem().Int()
		if statValue == 0 {
			continue
		}

		adjValue := float64(statValue) / StatModifiers[int(statType)]
		statBudget += adjValue
		statMap[int(statType)] = &ItemStat{
			Value:    int(statValue),
			Percent:  0.0,
			Type:     "Item",
			AdjValue: adjValue,
		}
	}

	// Calculate the total budget for the spell stats if we have some
	for _, spellStat := range spellStats {
		statBudget += float64(spellStat.Budget)
		statMap[spellStat.StatType] = &ItemStat{
			Value:    spellStat.StatValue,
			Percent:  0.0,
			Type:     "Spell",
			AdjValue: float64(spellStat.Budget),
		}
	}

	// Combine all stats and calculate percentages for each stat
	for statId, stat := range statMap {
		statMap[statId].Percent = math.Round(float64(stat.AdjValue)/statBudget*100) / 100
	}

	return statMap
}

// get an array of all the spells set on the item
func (item *Item) GetSpells() ([]Spell, error) {
	// dont reload for the same item .
	if len(item.Spells) > 0 {
		return item.Spells, nil
	}

	spells := []Spell{}
	values := reflect.ValueOf(item)
	for i := 1; i < 4; i++ {
		spellId := values.Elem().FieldByName(fmt.Sprintf("SpellId%v", i)).Elem().Int()
		if spellId == 0 {
			continue
		}

		spell, err := DB.GetSpell(int(spellId))
		if err != nil {
			log.Printf("failed to get the spell: %v error: %v", spellId, err)
			continue
		}

		spells = append(spells, spell)
	}
	item.Spells = spells
	return spells, nil
}

// Stat Formula scaler
// Ceiling of ((ItemLevel * QualityModifier * ItemTypeModifier)^1.7095 * %ofStats) ^ (1/1.7095)) / StatModifier
// i.e)   Green Strength Helmet  (((100 * 1.1 * 1.0)^1.705) * 1)^(1/1.7095) / 1.0 = 110 Strength on item
func (item *Item) ScaleItem(itemLevel int, itemQuality int) (bool, error) {
	var allSpellStats []ConvItemStat
	if item.ItemLevel == nil {
		return false, errors.New("field itemLevel is not set")
	}

	if item.Quality == nil {
		return false, errors.New("field quality is not set")
	}

	// Get all the spell Stats on the item we can convert
	spells, err := item.GetSpells()
	if err != nil {
		log.Printf("Failed to get spells for item: %v", err)
		return false, err
	}

	for i := 0; i < len(spells); i++ {
		convStats, err := spells[i].ConvertToStats()
		if err != nil {
			log.Printf("Failed to convert spell to stats: %v for spell %v", err, spells[i].Name)
			continue
		}

		allSpellStats = append(allSpellStats, convStats...)
	}

	allStats := item.GetStatPercents(allSpellStats)
	for statId, stat := range allStats {
		log.Printf(">>>>>> StatId: %v Type: %s Value: %v Percent: %v", statId, stat.Type, stat.Value, stat.Percent)
	}

	// Scale Armor Stats
	if *item.Class == 4 && *item.Armor > 0 {
		preArmor := *item.Armor
		// Handle shield subclass = 6
		if *item.Subclass == 6 {
			*item.Armor = int(math.Ceil(float64(itemLevel) * QualityModifiers[itemQuality] * MaterialModifiers[1]))
		} else {
			*item.Armor = int(math.Ceil(float64(itemLevel) * QualityModifiers[itemQuality] * MaterialModifiers[*item.Material]))
		}

		log.Printf("New Armor: %v scaled up from previous armor %v", *item.Armor, preArmor)
	}

	// If the item is a weapon scale the DPS
	if *item.Class == 2 && *item.MinDmg1 > 0 {
		predps, err := item.GetDPS()
		if err != nil {
			log.Printf("Failed to get DPS: %v", err)
		}

		dps, err := item.ScaleDPS(itemLevel)
		if err != nil {
			log.Printf("Failed to scale DPS: %v", err)
			return false, err
		}

		log.Printf("DPS: %.1f scaled up from previous dps %v", dps, predps)
	}

	return true, nil

	// // Calculate the Armor or Weapon modifier
	// var modifier float64
	// if *item.Class == 4 {
	// 	modifier = ArmorModifiers[*item.InventoryType]
	// } else {
	// 	modifier = WeaponModifiers[*item.InventoryType]
	// }

	// // Calculate the Quality Modifier
	// qualityModifier := QualityModifiers[*item.Quality]

	// // Calculate the Stat Modifier
	// statModifier := StatModifiers[primaryStat]

	// // Calculate the budget for the item
	// budget := math.Pow(float64(*item.ItemLevel)*qualityModifier*modifier, 1.7095)

	// // Calculate the percentage of stats
	// percent := math.Pow(budget, 1/1.7095) / statModifier

	// // Calculate the budget for the primary stat
	// primaryBudget := percent * float64(primaryVal)

	// // Calculate the budget for the other stats
	// otherBudget := percent - primaryBudget

	// // Calculate the budget for each stat
	// statBudget := otherBudget / float64(*item.StatsCount)

	// // Calculate the adjusted value for the primary stat
	// adjPrimaryVal := primaryBudget * statModifier

	// // Calculate the adjusted value for the other stats
	// adjStatVal := statBudget * statModifier

	// // Update the primary stat value
	// primaryVal = int(adjPrimaryVal)

	// // Update the other stats
	// values := reflect.ValueOf(item)
	// for i := 1; i < 11; i++ {
	// 	statType := values.FieldByName(fmt.Sprintf("StatType%v", i)).Elem().Int()
	// 	if statType == primaryStat {
	// 		continue
	// 	}

	// 	statValue := values.FieldByName(fmt.Sprintf("StatValue

}

// Scale formula ((ItemLevel * QualityModifier * ItemTypeModifier)^1.7095 * %ofStats) ^ (1/1.7095)) / StatModifier
func scaleStat(itemLevel int, itemType int, itemQuality int, percOfStat float64, statModifier float64) int {
	scaledUp := math.Pow((float64(itemLevel)*QualityModifiers[itemQuality]*InvTypeModifiers[itemType]), 1.7095) * percOfStat
	return int(math.Ceil(math.Pow(scaledUp, 1/1.7095) / statModifier)) // normalized
}
