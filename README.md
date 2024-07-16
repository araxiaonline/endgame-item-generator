## WoW EndGame Item Generator

This is a project I started to learn golang.  The goal is to enable scaling of base items in World Warcraft, based on a few parameters. 
The idea is combining  **mod-autobalance** + some ets-modules I have built to extend endgame content, specifically 5-man dungeon instances. 

This project creates scaled items based on math from a google sheets item generator, but also with some personal preferences. 

### Prebuilt Items 
The easiest path is to just download the prebuilt items I have generated.  The intention was for these items to beyond current achievable items 280+ so they are high stats.  Items with few stats ...get kind of insane when scaled. 

`mythic-scaled.sql` : ILevels 301 - 320  epic and rares item_template ids are original entry + 2000000

`legend-scaled.sql` : ILevels 316 - 340  epics only item_template ids are original entry + 2100000

`ascendant-scaled.sql` : ILevels 336 - 360  legendaries and epics item_template ids are original entry + 2200000

Additionally, this script adds spells to the spell_dbc table starting at 3000000, these spells are server side with effects that can be scaled done so and added to the original weapon. 

For instance `Cobalt Hammer` level 29 rare Chance on Hit Cold damage scales from 110 > 2800 as a level 80 Epic.   

### Cli Usage
IF you want to use the script to create your own I have added a few options to enable you to do so.  I am not big into write lots of documentation so if you have questions drop them in discord. 

If you have golang installed you can simply clone the repo and run the script
```
go run . --help
```

Otherwise you can download a binary from the releases page

Generate new items with defaults
```
./item-gen -ilvl 300 -difficulty 3 > myitems.sql
```

Generate items that require level up on boss drops in end game dungeons (strath, brd, HoL, shatterd halls..etc)
```
./item-gen -ilvl 320 -difficulty 4 -baselevel 83 > legendary.sql
```

Generate Items for a crazy ass PvP slaughter fest
```
./item-gen -ilevel 400 -baselevel 1 > overpowered.sql
```

The sql does not do anything without the additional autobalance mod that enables them to drop, unless you add a way to get them yourself in the game. 
