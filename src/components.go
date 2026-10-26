package main

// Components are the logical data used by entities for systems.
// Components are not unique and could be universally used.
// Components follow a strict naming convention:
// c_id = the component type
// m_id = the hashmap for entity compontent relationship lookup

// Component - Glyph
var m_Glyph = make(map[Entity]Glyph)
func (e Entity) GetGlyph() Glyph { return m_Glyph[e] }
func (e Entity) SetGlyph(glyph Glyph) { m_Glyph[e] = glyph }


// Component - Name
type c_Name string
var m_Names = make(map[Entity]c_Name)
func (e Entity) GetName() string { return string(m_Names[e]) }
func (e Entity) SetName(name string) { m_Names[e] = c_Name(name) }


// Component - Position
type c_Position struct {
    X, Y int
}
var m_Position = make(map[Entity]c_Position)
func (e Entity) GetPosition() (int, int) { return m_Position[e].X, m_Position[e].Y}

func (e Entity) SetPosition(x,y int) { 
    m_Position[e] = c_Position{ X: x, Y: y }
}


// Component - Sprite
// Will remove but keeping for now. We render glyphs not sprites
/*
type c_Sprite rune
var m_Sprites = make(map[Entity]c_Sprite)
func (e Entity) GetSprite() rune { return rune(m_Sprites[e]) }
func (e Entity) SetSprite(sprite rune) { m_Sprites[e] = c_Sprite(sprite) }
*/

// Component - Velocity
type c_Velocity struct {
    dx, dy int
}
var m_Velocity = make(map[Entity]c_Velocity)

func (e Entity) SetVelocity(x,y int) {
    m_Velocity[e] = c_Velocity{ dx: x, dy: y }
}

