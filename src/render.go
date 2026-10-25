package main

import (
    "github.com/Nareleth/kyanbasu"
)


// Layer 2 - Render the Entities
func s_RenderEntities(t *kyanbasu.Term, w *World) {
    for i := 0; i < len(w.Entities); i++ {
        e := Entity(i)
        Pos := m_Position[e]
        t.Move(Pos.X, Pos.Y)
        t.WriteRune(e.GetSprite())
    }
}
