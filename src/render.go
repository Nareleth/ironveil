package main

import (
    "github.com/Nareleth/kyanbasu"
)

// Glyph type
// A glyph represents the rune printed to screen. It contains a character, FG, and BG.
// Since we are using the kyanbasu framework, Glyphs follow the Cell type:
// Char rune, FG Color, BG Color

type Glyph struct {
    Char    rune
    FG      kyanbasu.Color
    BG      kyanbasu.Color
}

// Layer 2 - Render the Entities
func s_RenderEntities(t *kyanbasu.Term, w *World) {
    // Loop through all world entities
    for i := 0; i < len(w.Entities); i++ {
        // Extract entity information
        e := Entity(i)
        Pos := m_Position[e]
        glyph := e.GetGlyph()

        // Use terminal methods to render
        t.Move(Pos.X, Pos.Y)
        t.SetColor(glyph.FG)
        t.WriteRune(glyph.Char)
        t.SetColor(kyanbasu.Color_Reset)
    }
}
