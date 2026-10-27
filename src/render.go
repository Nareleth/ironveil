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
func RenderEntities(s *kyanbasu.Screen, w *World) {
    // Loop through all world entities
    for i := 0; i < len(w.Entities); i++ {
        // Extract entity information
        e := Entity(i)
        Pos := m_Position[e]
        glyph := e.GetGlyph()

        // Convert this entity glyph into a Cell for kyanbasu rendering
        this_cell := kyanbasu.Cell{
            Char:   glyph.Char,
            FG:     glyph.FG,
            BG:     glyph.BG,
        }

        // Use terminal methods to render
        s.SetCell(Pos.X, Pos.Y, this_cell)
    }
}

