# qudlike runescape sandbox grind game thing idk 
Built using ECS

## TO DO:
------
Phase 3: visualization
### World:
- world size variable
- generate a tilemap(just grass tiles for now)
- render tilemap layer

### Camera:
- map grid separate from rendering grid
- create it
- it follows the player centered

### Tilemap/Sprites:
- make a JSON file
- read and parse json file

------
Phase 4: Labor and Hunting
### Entity:
- trees
- squirrel

### Components:
- choppable
- collision
- hp

### Systems:
- collision
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
-  new gamestate
- character control state (for input)

### Movement:
- fix the movement system
- add gamestate control to the input

