package main

// The global systems for entities to do things.


// Movement
// TEMP!!!!! setting to just player - need to set to all world entities
func s_Movement(e Entity) {
    // loop through world entities with a velocity goes here
    vx := m_Position[e].X + m_Velocity[e].dx
    vy := m_Position[e].Y + m_Velocity[e].dy
    e.SetPosition(vx, vy)
    e.SetVelocity(0, 0)
}
