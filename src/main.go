package main

import (
    "time"
    "github.com/Nareleth/kyanbasu"
)


// Config
const (
    FPS = 30
    world_size = 20
)


func main() {
    // Initialize term controls
    t := kyanbasu.NewTerm()
    defer t.CloseTerm()

    // Initialize screen rendering
    screen := kyanbasu.NewScreen(t)


    // Initialize the world
    world := NewWorld(world_size)
    world.GenerateMap()

    // Initialize Player
    Player := AddEntityPlayer(world, "Jeff", '@', 1, 1)


    // Go routine read key input
    go ReadInput(t, Player)


    // Game Loop
    for {
        PlayerPos := m_Position[Player]
        
        // Systems
        s_Movement(Player)

        // Render Layers
        RenderTiles(screen, world)
        RenderEntities(screen, world)

        // Print HUD
        t.Move(0,0)
        t.WriteString("Pos: (%d,%d)", PlayerPos.X, PlayerPos.Y)

        // Flush render instructions
        screen.Flush()

        // Frame limit
        time.Sleep(time.Second / time.Duration(FPS))
    }
}
