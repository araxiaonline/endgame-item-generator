package main

import (
	"fmt"

	"github.com/araxiaonline/endgame-item-generator/models"
)

// convert an item to a create sql statement
func ItemToSql(item models.Item, reqLevel int, difficulty int) string {

	entryBump := 20000000
	spellBump := 30000000
	if difficulty == 4 {
		entryBump = 21000000
	}
	if difficulty == 5 {
		entryBump = 22000000
	}

	if *item.Quality == 4 {
		spellBump = 31000000
	}
	if *item.Quality == 5 {
		spellBump = 32000000
	}

	spells := ""
	if len(item.Spells) > 0 {
		for i, spell := range item.Spells {
			spells += SpellToSql(spell, *item.Quality)
			item.UpdateField(fmt.Sprintf("SpellId%v", i), spellBump+spell.ID)
		}
	}

	delete := fmt.Sprintf("DELETE FROM acore_world.item_template WHERE entry = %v;", entryBump+item.Entry)

	clone := fmt.Sprintf(`
	INSERT INTO acore_world.item_template  (
		entry, class, subclass, SoundOverrideSubclass, name, displayid, Quality, Flags, FlagsExtra, BuyCount, 
		BuyPrice, SellPrice, InventoryType, AllowableClass, AllowableRace, ItemLevel, RequiredLevel, 
		RequiredSkill, RequiredSkillRank, requiredspell, requiredhonorrank, RequiredCityRank, 
		RequiredReputationFaction, RequiredReputationRank, maxcount, stackable, ContainerSlots, StatsCount, 
		stat_type1, stat_value1, stat_type2, stat_value2, stat_type3, stat_value3, stat_type4, stat_value4, 
		stat_type5, stat_value5, stat_type6, stat_value6, stat_type7, stat_value7, stat_type8, stat_value8, 
		stat_type9, stat_value9, stat_type10, stat_value10, ScalingStatDistribution, ScalingStatValue, 
		dmg_min1, dmg_max1, dmg_type1, dmg_min2, dmg_max2, dmg_type2, armor, holy_res, fire_res, nature_res, 
		frost_res, shadow_res, arcane_res, delay, ammo_type, RangedModRange, spellid_1, spelltrigger_1, 
		spellcharges_1, spellppmRate_1, spellcooldown_1, spellcategory_1, spellcategorycooldown_1, spellid_2, 
		spelltrigger_2, spellcharges_2, spellppmRate_2, spellcooldown_2, spellcategory_2, spellcategorycooldown_2, 
		spellid_3, spelltrigger_3, spellcharges_3, spellppmRate_3, spellcooldown_3, spellcategory_3, 
		spellcategorycooldown_3, spellid_4, spelltrigger_4, spellcharges_4, spellppmRate_4, spellcooldown_4, 
		spellcategory_4, spellcategorycooldown_4, spellid_5, spelltrigger_5, spellcharges_5, spellppmRate_5, 
		spellcooldown_5, spellcategory_5, spellcategorycooldown_5, bonding, description, PageText, LanguageID, 
		PageMaterial, startquest, lockid, Material, sheath, RandomProperty, RandomSuffix, block, itemset, 
		MaxDurability, area, Map, BagFamily, TotemCategory, socketColor_1, socketContent_1, socketColor_2, 
		socketContent_2, socketColor_3, socketContent_3, socketBonus, GemProperties, RequiredDisenchantSkill, 
		ArmorDamageModifier, duration, ItemLimitCategory, HolidayId, ScriptName, DisenchantID, FoodType, 
		minMoneyLoot, maxMoneyLoot, flagsCustom, VerifiedBuild
	  )
	  SELECT 
		entry + %v, class, subclass, SoundOverrideSubclass, name, displayid, Quality, Flags, FlagsExtra, BuyCount, 
		BuyPrice, SellPrice, InventoryType, AllowableClass, AllowableRace, ItemLevel, RequiredLevel, 
		RequiredSkill, RequiredSkillRank, requiredspell, requiredhonorrank, RequiredCityRank, 
		RequiredReputationFaction, RequiredReputationRank, maxcount, stackable, ContainerSlots, StatsCount, 
		stat_type1, stat_value1, stat_type2, stat_value2, stat_type3, stat_value3, stat_type4, stat_value4, 
		stat_type5, stat_value5, stat_type6, stat_value6, stat_type7, stat_value7, stat_type8, stat_value8, 
		stat_type9, stat_value9, stat_type10, stat_value10, ScalingStatDistribution, ScalingStatValue, 
		dmg_min1, dmg_max1, dmg_type1, dmg_min2, dmg_max2, dmg_type2, armor, holy_res, fire_res, nature_res, 
		frost_res, shadow_res, arcane_res, delay, ammo_type, RangedModRange, spellid_1, spelltrigger_1, 
		spellcharges_1, spellppmRate_1, spellcooldown_1, spellcategory_1, spellcategorycooldown_1, spellid_2, 
		spelltrigger_2, spellcharges_2, spellppmRate_2, spellcooldown_2, spellcategory_2, spellcategorycooldown_2, 
		spellid_3, spelltrigger_3, spellcharges_3, spellppmRate_3, spellcooldown_3, spellcategory_3, 
		spellcategorycooldown_3, spellid_4, spelltrigger_4, spellcharges_4, spellppmRate_4, spellcooldown_4, 
		spellcategory_4, spellcategorycooldown_4, spellid_5, spelltrigger_5, spellcharges_5, spellppmRate_5, 
		spellcooldown_5, spellcategory_5, spellcategorycooldown_5, bonding, description, PageText, LanguageID, 
		PageMaterial, startquest, lockid, Material, sheath, RandomProperty, RandomSuffix, block, itemset, 
		MaxDurability, area, Map, BagFamily, TotemCategory, socketColor_1, socketContent_1, socketColor_2, 
		socketContent_2, socketColor_3, socketContent_3, socketBonus, GemProperties, RequiredDisenchantSkill, 
		ArmorDamageModifier, duration, ItemLimitCategory, HolidayId, ScriptName, DisenchantID, FoodType, 
		minMoneyLoot, maxMoneyLoot, flagsCustom, VerifiedBuild
	  FROM acore_world.item_template as src
	  WHERE src.entry = %v ON DUPLICATE KEY UPDATE entry = src.entry + %v;	  
	`, entryBump, item.Entry, entryBump)

	update := fmt.Sprintf(`
	UPDATE acore_world.item_template
	SET 
	  Quality = %v,
	  ItemLevel = %v,
	  RequiredLevel = %v,
	  dmg_min1 = %v,
	  dmg_max1 = %v,
	  dmg_min2 = %v,
	  dmg_max2 = %v,
	  StatsCount = %v,
	  stat_type1 = %v,
	  stat_value1 = %v,
	  stat_type2 = %v,
	  stat_value2 = %v,
	  stat_type3 = %v,
	  stat_value3 = %v,
	  stat_type4 = %v,
	  stat_value4 = %v,
	  stat_type5 = %v,
	  stat_value5 = %v,
	  stat_type6 = %v,
	  stat_value6 = %v,
	  stat_type7 = %v,
	  stat_value7 = %v,
	  stat_type8 = %v,
	  stat_value8 = %v,
	  stat_type9 = %v,
	  stat_value9 = %v,
	  stat_type10 = %v,
	  stat_value10 = %v,
	  spellid_1 = %v,
	  spellid_2 = %v,
	  spellid_3 = %v,
	  RequiredDisenchantSkill = %v,
	  DisenchantID = %v,
	  SellPrice = FLOOR(100000 + (RAND() * 400001)),
	  Armor = %v
	WHERE entry = %v;
	`, *item.Quality, *item.ItemLevel, reqLevel, *item.MinDmg1, *item.MaxDmg1, *item.MinDmg2, *item.MaxDmg2, *item.StatsCount,
		*item.StatType1, *item.StatValue1, *item.StatType2, *item.StatValue2, *item.StatType3, *item.StatValue3, *item.StatType4, *item.StatValue4,
		*item.StatType5, *item.StatValue5, *item.StatType6, *item.StatValue6, *item.StatType7, *item.StatValue7, *item.StatType8, *item.StatValue8,
		*item.StatType9, *item.StatValue9, *item.StatType10, *item.StatValue10, *item.SpellId1, *item.SpellId2, *item.SpellId3, 375,
		68, *item.Armor, entryBump+item.Entry)

	return fmt.Sprintf("%s %s \n %s \n %s", spells, delete, clone, update)
}

func SpellToSql(spell models.Spell, quality int) string {

	entryBump := 30000000
	if quality == 4 {
		entryBump = 31000000
	}
	if quality == 5 {
		entryBump = 32000000
	}

	insert := fmt.Sprintf(`
	INSERT IGNORE INTO acore_world.spell_dbc (
		ID, Category, DispelType, Mechanic, Attributes, AttributesEx, AttributesEx2, AttributesEx3, AttributesEx4,
		AttributesEx5, AttributesEx6, AttributesEx7, ShapeshiftMask, unk_320_2, ShapeshiftExclude, unk_320_3, Targets,
		TargetCreatureType, RequiresSpellFocus, FacingCasterFlags, CasterAuraState, TargetAuraState, ExcludeCasterAuraState,
		ExcludeTargetAuraState, CasterAuraSpell, TargetAuraSpell, ExcludeCasterAuraSpell, ExcludeTargetAuraSpell, CastingTimeIndex,
		RecoveryTime, CategoryRecoveryTime, InterruptFlags, AuraInterruptFlags, ChannelInterruptFlags, ProcTypeMask, ProcChance,
		ProcCharges, MaxLevel, BaseLevel, SpellLevel, DurationIndex, PowerType, ManaCost, ManaCostPerLevel, ManaPerSecond,
		ManaPerSecondPerLevel, RangeIndex, Speed, ModalNextSpell, CumulativeAura, Totem_1, Totem_2, Reagent_1, Reagent_2, Reagent_3,
		Reagent_4, Reagent_5, Reagent_6, Reagent_7, Reagent_8, ReagentCount_1, ReagentCount_2, ReagentCount_3, ReagentCount_4,
		ReagentCount_5, ReagentCount_6, ReagentCount_7, ReagentCount_8, EquippedItemClass, EquippedItemSubclass, EquippedItemInvTypes,
		Effect_1, Effect_2, Effect_3, EffectDieSides_1, EffectDieSides_2, EffectDieSides_3, EffectRealPointsPerLevel_1,
		EffectRealPointsPerLevel_2, EffectRealPointsPerLevel_3, EffectBasePoints_1, EffectBasePoints_2, EffectBasePoints_3,
		EffectMechanic_1, EffectMechanic_2, EffectMechanic_3, ImplicitTargetA_1, ImplicitTargetA_2, ImplicitTargetA_3, ImplicitTargetB_1,
		ImplicitTargetB_2, ImplicitTargetB_3, EffectRadiusIndex_1, EffectRadiusIndex_2, EffectRadiusIndex_3, EffectAura_1,
		EffectAura_2, EffectAura_3, EffectAuraPeriod_1, EffectAuraPeriod_2, EffectAuraPeriod_3, EffectMultipleValue_1, EffectMultipleValue_2,
		EffectMultipleValue_3, EffectChainTargets_1, EffectChainTargets_2, EffectChainTargets_3, EffectItemType_1, EffectItemType_2,
		EffectItemType_3, EffectMiscValue_1, EffectMiscValue_2, EffectMiscValue_3, EffectMiscValueB_1, EffectMiscValueB_2, EffectMiscValueB_3,
		EffectTriggerSpell_1, EffectTriggerSpell_2, EffectTriggerSpell_3, EffectPointsPerCombo_1, EffectPointsPerCombo_2, EffectPointsPerCombo_3,
		EffectSpellClassMaskA_1, EffectSpellClassMaskA_2, EffectSpellClassMaskA_3, EffectSpellClassMaskB_1, EffectSpellClassMaskB_2,
		EffectSpellClassMaskB_3, EffectSpellClassMaskC_1, EffectSpellClassMaskC_2, EffectSpellClassMaskC_3, SpellVisualID_1, SpellVisualID_2,
		SpellIconID, ActiveIconID, SpellPriority, Name_Lang_enUS, Name_Lang_enGB, Name_Lang_koKR, Name_Lang_frFR, Name_Lang_deDE,
		Name_Lang_enCN, Name_Lang_zhCN, Name_Lang_enTW, Name_Lang_zhTW, Name_Lang_esES, Name_Lang_esMX, Name_Lang_ruRU, Name_Lang_ptPT,
		Name_Lang_ptBR, Name_Lang_itIT, Name_Lang_Unk, Name_Lang_Mask, NameSubtext_Lang_enUS, NameSubtext_Lang_enGB, NameSubtext_Lang_koKR,
		NameSubtext_Lang_frFR, NameSubtext_Lang_deDE, NameSubtext_Lang_enCN, NameSubtext_Lang_zhCN, NameSubtext_Lang_enTW, NameSubtext_Lang_zhTW,
		NameSubtext_Lang_esES, NameSubtext_Lang_esMX, NameSubtext_Lang_ruRU, NameSubtext_Lang_ptPT, NameSubtext_Lang_ptBR, NameSubtext_Lang_itIT,
		NameSubtext_Lang_Unk, NameSubtext_Lang_Mask, Description_Lang_enUS, Description_Lang_enGB, Description_Lang_koKR, Description_Lang_frFR,
		Description_Lang_deDE, Description_Lang_enCN, Description_Lang_zhCN, Description_Lang_enTW, Description_Lang_zhTW, Description_Lang_esES,
		Description_Lang_esMX, Description_Lang_ruRU, Description_Lang_ptPT, Description_Lang_ptBR, Description_Lang_itIT, Description_Lang_Unk,
		Description_Lang_Mask, AuraDescription_Lang_enUS, AuraDescription_Lang_enGB, AuraDescription_Lang_koKR, AuraDescription_Lang_frFR,
		AuraDescription_Lang_deDE, AuraDescription_Lang_enCN, AuraDescription_Lang_zhCN, AuraDescription_Lang_enTW, AuraDescription_Lang_zhTW,
		AuraDescription_Lang_esES, AuraDescription_Lang_esMX, AuraDescription_Lang_ruRU, AuraDescription_Lang_ptPT, AuraDescription_Lang_ptBR,
		AuraDescription_Lang_itIT, AuraDescription_Lang_Unk, AuraDescription_Lang_Mask, ManaCostPct, StartRecoveryCategory, StartRecoveryTime,
		MaxTargetLevel, SpellClassSet, SpellClassMask_1, SpellClassMask_2, SpellClassMask_3, MaxTargets, DefenseType, PreventionType, StanceBarOrder,
		EffectChainAmplitude_1, EffectChainAmplitude_2, EffectChainAmplitude_3, MinFactionID, MinReputation, RequiredAuraVision, RequiredTotemCategoryID_1,
		RequiredTotemCategoryID_2, RequiredAreasID, SchoolMask, RuneCostID, SpellMissileID, PowerDisplayID, EffectBonusMultiplier_1, EffectBonusMultiplier_2,
		EffectBonusMultiplier_3, SpellDescriptionVariableID, SpellDifficultyID
	) SELECT 
	ID + %v, Category, DispelType, Mechanic, Attributes, AttributesEx, AttributesEx2, AttributesEx3, AttributesEx4,
	AttributesEx5, AttributesEx6, AttributesEx7, ShapeshiftMask, unk_320_2, ShapeshiftExclude, unk_320_3, Targets,
	TargetCreatureType, RequiresSpellFocus, FacingCasterFlags, CasterAuraState, TargetAuraState, ExcludeCasterAuraState,
	ExcludeTargetAuraState, CasterAuraSpell, TargetAuraSpell, ExcludeCasterAuraSpell, ExcludeTargetAuraSpell, CastingTimeIndex,
	RecoveryTime, CategoryRecoveryTime, InterruptFlags, AuraInterruptFlags, ChannelInterruptFlags, ProcTypeMask, ProcChance,
	ProcCharges, MaxLevel, BaseLevel, SpellLevel, DurationIndex, PowerType, ManaCost, ManaCostPerLevel, ManaPerSecond,
	ManaPerSecondPerLevel, RangeIndex, Speed, ModalNextSpell, CumulativeAura, Totem_1, Totem_2, Reagent_1, Reagent_2, Reagent_3,
	Reagent_4, Reagent_5, Reagent_6, Reagent_7, Reagent_8, ReagentCount_1, ReagentCount_2, ReagentCount_3, ReagentCount_4,
	ReagentCount_5, ReagentCount_6, ReagentCount_7, ReagentCount_8, EquippedItemClass, EquippedItemSubclass, EquippedItemInvTypes,
	Effect_1, Effect_2, Effect_3, EffectDieSides_1, EffectDieSides_2, EffectDieSides_3, EffectRealPointsPerLevel_1,
	EffectRealPointsPerLevel_2, EffectRealPointsPerLevel_3, EffectBasePoints_1, EffectBasePoints_2, EffectBasePoints_3,
	EffectMechanic_1, EffectMechanic_2, EffectMechanic_3, ImplicitTargetA_1, ImplicitTargetA_2, ImplicitTargetA_3, ImplicitTargetB_1,
	ImplicitTargetB_2, ImplicitTargetB_3, EffectRadiusIndex_1, EffectRadiusIndex_2, EffectRadiusIndex_3, EffectAura_1,
	EffectAura_2, EffectAura_3, EffectAuraPeriod_1, EffectAuraPeriod_2, EffectAuraPeriod_3, EffectMultipleValue_1, EffectMultipleValue_2,
	EffectMultipleValue_3, EffectChainTargets_1, EffectChainTargets_2, EffectChainTargets_3, EffectItemType_1, EffectItemType_2,
	EffectItemType_3, EffectMiscValue_1, EffectMiscValue_2, EffectMiscValue_3, EffectMiscValueB_1, EffectMiscValueB_2, EffectMiscValueB_3,
	EffectTriggerSpell_1, EffectTriggerSpell_2, EffectTriggerSpell_3, EffectPointsPerCombo_1, EffectPointsPerCombo_2, EffectPointsPerCombo_3,
	EffectSpellClassMaskA_1, EffectSpellClassMaskA_2, EffectSpellClassMaskA_3, EffectSpellClassMaskB_1, EffectSpellClassMaskB_2,
	EffectSpellClassMaskB_3, EffectSpellClassMaskC_1, EffectSpellClassMaskC_2, EffectSpellClassMaskC_3, SpellVisualID_1, SpellVisualID_2,
	SpellIconID, ActiveIconID, SpellPriority, Name_Lang_enUS, Name_Lang_enGB, Name_Lang_koKR, Name_Lang_frFR, Name_Lang_deDE,
	Name_Lang_enCN, Name_Lang_zhCN, Name_Lang_enTW, Name_Lang_zhTW, Name_Lang_esES, Name_Lang_esMX, Name_Lang_ruRU, Name_Lang_ptPT,
	Name_Lang_ptBR, Name_Lang_itIT, Name_Lang_Unk, Name_Lang_Mask, NameSubtext_Lang_enUS, NameSubtext_Lang_enGB, NameSubtext_Lang_koKR,
	NameSubtext_Lang_frFR, NameSubtext_Lang_deDE, NameSubtext_Lang_enCN, NameSubtext_Lang_zhCN, NameSubtext_Lang_enTW, NameSubtext_Lang_zhTW,
	NameSubtext_Lang_esES, NameSubtext_Lang_esMX, NameSubtext_Lang_ruRU, NameSubtext_Lang_ptPT, NameSubtext_Lang_ptBR, NameSubtext_Lang_itIT,
	NameSubtext_Lang_Unk, NameSubtext_Lang_Mask, Description_Lang_enUS, Description_Lang_enGB, Description_Lang_koKR, Description_Lang_frFR,
	Description_Lang_deDE, Description_Lang_enCN, Description_Lang_zhCN, Description_Lang_enTW, Description_Lang_zhTW, Description_Lang_esES,
	Description_Lang_esMX, Description_Lang_ruRU, Description_Lang_ptPT, Description_Lang_ptBR, Description_Lang_itIT, Description_Lang_Unk,
	Description_Lang_Mask, AuraDescription_Lang_enUS, AuraDescription_Lang_enGB, AuraDescription_Lang_koKR, AuraDescription_Lang_frFR,
	AuraDescription_Lang_deDE, AuraDescription_Lang_enCN, AuraDescription_Lang_zhCN, AuraDescription_Lang_enTW, AuraDescription_Lang_zhTW,
	AuraDescription_Lang_esES, AuraDescription_Lang_esMX, AuraDescription_Lang_ruRU, AuraDescription_Lang_ptPT, AuraDescription_Lang_ptBR,
	AuraDescription_Lang_itIT, AuraDescription_Lang_Unk, AuraDescription_Lang_Mask, ManaCostPct, StartRecoveryCategory, StartRecoveryTime,
	MaxTargetLevel, SpellClassSet, SpellClassMask_1, SpellClassMask_2, SpellClassMask_3, MaxTargets, DefenseType, PreventionType, StanceBarOrder,
	EffectChainAmplitude_1, EffectChainAmplitude_2, EffectChainAmplitude_3, MinFactionID, MinReputation, RequiredAuraVision, RequiredTotemCategoryID_1,
	RequiredTotemCategoryID_2, RequiredAreasID, SchoolMask, RuneCostID, SpellMissileID, PowerDisplayID, EffectBonusMultiplier_1, EffectBonusMultiplier_2,
	EffectBonusMultiplier_3, SpellDescriptionVariableID, SpellDifficultyID from acore_world.spell_dbc as src
	WHERE src.ID = %v ON DUPLICATE KEY UPDATE ID = src.ID + %v;`, entryBump, spell.ID, entryBump)

	update := fmt.Sprintf(`
	UPDATE acore_world.spell_dbc
	SET EffectBasePoints_1 = %v, EffectBasePoints_2 = %v
	WHERE ID = %v;`, spell.EffectBasePoints1, spell.EffectBasePoints2, entryBump+spell.ID)

	return fmt.Sprintf("\n %s \n %s \n", insert, update)
}
