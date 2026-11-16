# qudlike runescape sandbox grind game thing idk 
Built using ECS

## TO DO:
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
### Input:
- chop adjacent tree

### Entity:
- squirrel
    - AddEntitySquirrel
    - Add States: idle

### Components:
- choppable
- hp

### Systems:
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
-  new gamestate
- character control state (for input)

### Movement:
- fix the movement system
- add gamestate control to the input

