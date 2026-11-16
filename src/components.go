package main

// Components are the logical data used by entities for systems.
// Components are not unique and could be universally used.
// Many components are built around a component type variable, and functions to get and set the value of the component.
// Components follow a strict naming convention:
// c_id = the component type
// m_id = the hashmap for entity compontent relationship lookup


// Component - Camera
var m_Camera = make(map[Entity]*Camera)
func (e Entity) SetCamera(camera *Camera) { m_Camera[e] = camera }
/*
func (e Entity) MoveCamera(dx, dy int) { 
    m_Camera[e].ViewportX = m_Camera[e].ViewportX + dx
    m_Camera[e].ViewportY = m_Camera[e].ViewportY + dy
}
*/


// Component - Choppable
type c_Choppable bool
var m_Choppable = make(map[Entity]c_Choppable)
func (e Entity) GetChoppable() bool { return bool(m_Choppable[e]) }
func (e Entity) SetChoppable(choppable bool) { m_Choppable[e] = c_Choppable(choppable) }


// Component - Collidable
type c_Collidable bool
var m_Collidable = make(map[Entity]c_Collidable)
func (e Entity) GetCollidable() bool { return bool(m_Collidable[e]) }
func (e Entity) SetCollidable(collidable bool) { m_Collidable[e] = c_Collidable(collidable) }


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


// Component - State
// The current action or mode of an entity
type c_State string
const (
    state_Idle c_State = "Idle"
    state_Chop c_State = "Chopping"
)
var m_State = make(map[Entity]c_State)
func (e Entity) GetState() c_State { return c_State(m_State[e]) }
func (e Entity) SetState(state c_State) { m_State[e] = c_State(state) }


// Component - Velocity
type c_Velocity struct {
    dx, dy int
}
var m_Velocity = make(map[Entity]c_Velocity)
func (e Entity) GetVelocity() (int, int) { return m_Velocity[e].dx, m_Velocity[e].dy}
func (e Entity) SetVelocity(x,y int) {
    m_Velocity[e] = c_Velocity{ dx: x, dy: y }
}

