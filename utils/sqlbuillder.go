package utils

func GetItemFields() string {
	return `	
	entry, name, displayid,
	quality, ItemLevel, class, subclass, inventoryType,
	allowableClass, allowableRace,
	requiredSkill, requiredLevel,
	dmg_min1, dmg_max1,
	dmg_min2,dmg_max2,
	dmg_type1, dmg_type2,
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
	stat_type10, stat_value10`
}
