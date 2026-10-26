package main

// The world holds all the entities and maps that make up... the world.
// Thus, all funcs that change the world should provide the world as an argument.

type World struct {
    Size        int
    Map         [][]Tile
    Entities    []Entity
}


// Tile
// A tile represents a space in the world
type Tile struct {
    Glyph   Glyph
    X, Y    int
}


var (
    tile_grass = Tile{ Glyph: glyph_e_grass }
)

// Initialize a new world
func NewWorld(size int) *World {
    return &World{
        Size:   size,
        Map:    make([][]Tile, size),
    }
}


// Methods
// Generate the world map
func (w *World) GenerateMap() {
    // Iterate the world dimensions
    for y := 0; y < w.Size; y++ {
        // Create a column of tiles
        w.Map[y] = make([]Tile, w.Size)

        for x := 0; x < w.Size; x++ {
            // Set this tile
            w.Map[y][x] = tile_grass
            w.Map[y][x].X = x
            w.Map[y][x].Y = y
        }
    }
}
