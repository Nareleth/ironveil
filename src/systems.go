package main

// The global systems for entities to do things.


// Collision
func Collided(w *World, e Entity, ex, ey, vx, vy int) bool {
    // Detect if a colliding entity exists in target tile
    target, collided := w.GetEntityAt(vx, vy)            

    if collided {
        if target.GetCollidable() == true {
            return true
        }
    }
    return false
}


// Movement
// TEMP!!!!! setting to just player - need to set to all world entities
func s_Movement(w *World, e Entity) {
    // loop through world entities with a velocity goes here

    // Current Position
    ex := m_Position[e].X
    ey := m_Position[e].Y

    // Velocity
    dx := m_Velocity[e].dx
    dy := m_Velocity[e].dy

    // Target position
    vx := ex + dx
    vy := ey + dy

    // Detect collision
    if Collided(w, e, ex, ey, vx, vy) {
        // Stop movement
        e.SetPosition(ex, ey)
        e.SetVelocity(0, 0)
    } else {
        // Move entity
        e.SetPosition(vx, vy)
        e.SetVelocity(0, 0)
    }

}
