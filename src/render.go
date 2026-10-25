package main

import (
    "github.com/Nareleth/kyanbasu"
)

// Layer 2 - Render the Entities
// TEMP!!!!!!! ONLY RENDERING PLAYER ATM. NEED TO RENDER WORLD ENTITIES
func s_RenderEntities(t *kyanbasu.Term, e Entity) {
    // loop through entities goes here
    Pos := m_Position[e]
    t.Move(Pos.X, Pos.Y)
    t.WriteRune(e.GetSprite())
}
