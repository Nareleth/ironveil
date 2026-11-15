package main

// The world holds all the entities and maps that make up... the world.
// Thus, all funcs that change the world should provide the world as an argument.

import (
    "math/rand/v2"
)

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


// Create static tiles mapped to glyphs
var (
    tile_grass  = Tile{ Glyph: glyph_e_grass }
)

// Initialize a new world
func NewWorld(size int) *World {
    return &World{
        Size:   size,
        Map:    make([][]Tile, size),
    }
}


// Methods

// Get Entity from a coordinate
func (w *World) GetEntityAt(x, y int) (Entity, bool) {
    // Loop through the position hashmap
    for e, pos := range m_Position {
        // Check if given coordinates match an entity
        if x == pos.X && y == pos.Y {
            // Return entity and true
            return e, true
        }
    }

    // Return else
    return 0, false
}


// Generate the world map
func (w *World) GenerateMap() {
    // Initialize seed rng
    src := rand.NewPCG(seed, 0)
    rng := rand.New(src)


    // Iterate the world dimensions
    for y := 0; y < w.Size; y++ {
        // Create a column of tiles
        w.Map[y] = make([]Tile, w.Size)

        for x := 0; x < w.Size; x++ {
            // Set this tile

            // RNG seeded tile
            randTile := rng.IntN(100)

            switch {
            case randTile < 5:
                this_tree := AddEntityTree(w, x, y)
                w.Map[y][x] = Tile{ Glyph: this_tree.GetGlyph() }
            default:
                w.Map[y][x] = tile_grass
            }
            w.Map[y][x].X = x
            w.Map[y][x].Y = y
        }
    }
}
