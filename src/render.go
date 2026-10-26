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


// Layer 1 - Render the tiles
// More specifically this converts our map into a kyanbasu Cell to be handled natively with those rendering functions.
func RenderTiles(s *kyanbasu.Screen, w *World) {
    // Loop through the world tiles
    for y := 0; y < w.Size; y++ {
        for x := 0; x < w.Size; x++ {
            // Initialize this specific tile
            this_tile := w.Map[y][x]

            // Convert this tile into a Cell for kyanbasu rendering
            this_cell := kyanbasu.Cell{
                Char:   this_tile.Glyph.Char,
                FG:     this_tile.Glyph.FG,
                BG:     this_tile.Glyph.BG,
            }

            s.SetCell(x, y, this_cell)
        }
    }
}


// Layer 2 - Render the Entities
func RenderEntities(t *kyanbasu.Term, w *World) {
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
