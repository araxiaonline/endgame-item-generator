package utils

func GetItemFields(prefix string) string {
	pre := ""
	if prefix != "" {
		pre = prefix + "."
	}

	return `	
	` + pre + `entry, ` + pre + `name, displayid,
	quality, ItemLevel, class, subclass, inventoryType,
	allowableClass, allowableRace,
	armor,material,
	requiredSkill, requiredLevel,
	dmg_min1, dmg_max1,
	dmg_min2,dmg_max2,
	dmg_type1, dmg_type2,
	delay, sheath, MaxDurability,
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
	stat_type10, stat_value10,
	spellid_1, spellid_2, spellid_3, 
	spelltrigger_1, spelltrigger_2, spelltrigger_3`
}

func GetSpellFields() string {
	return `
	ID,
    Name_Lang_enUS,
    Description_Lang_enUS,
    AuraDescription_Lang_enUS,
    ProcChance,
    SpellLevel,
    Effect_1,
    Effect_2,
    Effect_3,
    EffectDieSides_1,
    EffectDieSides_2,
    EffectDieSides_3,
    EffectRealPointsPerLevel_1,
    EffectRealPointsPerLevel_2,
    EffectRealPointsPerLevel_3,
    EffectBasePoints_1,
    EffectBasePoints_2,
    EffectBasePoints_3,
    EffectAura_1,
    EffectAura_2,
    EffectAura_3,
    EffectBonusMultiplier_1,
    EffectBonusMultiplier_2,
    EffectBonusMultiplier_3	
	`
}
