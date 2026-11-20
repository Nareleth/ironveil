# Open world survival adventure rpg
Built using ECS
Themes like stalker and qud
Life sim

## TO DO:
Phase Z2:
### Debug:
- also add run flag param to run game in debug mode
- bugfix why even when disabled the window still appears

### logs:
- create the slice for holding logs
- create func to generate logs
- read logs on screen

### hud:
- show player name
- some hotkey help
- expand a window to show all hotkeys and help

### canvas:
- redraw the game tiles + entities inside it

------
Phase Z3:
### Cursor mode:
- hotkey to detach the camera and follow a cursor
- cursor can hover and order commands
- can issue a move to order
- show interactions on highlighted entity

### Interactions
- when pressing key with nothing nearby nothing happens
- when pressing interaction key with only one thing nearby do the thing
- when pressing interaction with multiple things nearby, pause and confirm focus target
- highlight focus target
- tab switch between focus targets
------

Phase 4:
- add input for chopping adjacent tree
- Create chop tree system
- remove entity method
------
Phase 3: visualization
### Camera:
- it follows the player centered

### Bugfix:
- Clamp printing within boundaries

### Tilemap/Sprites:
- make a JSON file
- read and parse json file

------
Phase 4: Labor and Hunting
### Chop:
- Create chop command

### Attack/Hunt:
- directional attack command
- cursor interaction attack order





### Input:
- chop adjacent tree

### Entity:
- sticks
    - ground pickup
- rocks
    - ground pickup
- squirrel
    - AddEntitySquirrel
    - Add States: idle

### Components:
- choppable
- hp

### Systems:
- crafting
- Chop Tree
- idle wandering
- melee

### World:
- Remove entity

------
Phase 5: HUD
### kyanbasu:
- window
- draw border

### Stats/Components:
- HP
- Woodcutting
- Melee Combat

### HUD:
- Player name
- Stats
- Actions
- Event log

### Event log:
- make it

### Systems:
- Experience Gain

------
Phase 6: Building and Crafting
### Components:
- Inventory

### Items:
- logs
- meat
- squirrel fur

### Entity:
- Crafting Bench

### Components: 
- Interactable

### Menu:
- Inventory
- Crafting

### Crafting:
- Wooden Wall

### Entity:
- Wooden Wall

------
Phase Z: Polish
### Components:
- Playable?

### GameState:
- character control state (for input)

### Config:
- sprite + tilemap file
- setting conf file
- keybinds file
- default debug mode is hidden 


## Other things
- can find comraderr with over npc adventurers (like stalker)
- world feel is very weird and alien
- some towns have really weird quirks (like trading in teeth, some speak in tongues but friendly)
- world is somewhat eldritch and decaying
- maybe some major storms like the stalker emissions and psy storms + acid rain?
- seasons change come and go
- grass and trees can regrow
- will have dungeons to get loot and things to further advance your character or home
