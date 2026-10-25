package main

import (
    "time"
    "github.com/Nareleth/kyanbasu"
)

// Conf
const (
    FPS = 30
)

func main() {
    // Initialize term controls
    t := kyanbasu.NewTerm()
    defer t.CloseTerm()

    // Initialize screen rendering
    //screen := kyanbasu.NewScreen(t)


    // Initialize the world
    world := NewWorld()

    // Initialize Player
    Player := AddEntityPlayer(world, "Jeff", '@', 1, 1)


    // Go routine read key input
    go ReadInput(t, Player)


    // Game Loop
    for {
        t.Clear()

        PlayerPos := m_Position[Player]
        
        // Systems
        s_Movement(Player)
        s_RenderEntities(t, Player)

        // Print HUD
        t.Move(0,0)
        t.WriteString("Pos: (%d,%d)", PlayerPos.X, PlayerPos.Y)

        // Flush render instructions
        t.Flush()

        // Frame limit
        time.Sleep(time.Second / time.Duration(FPS))
    }
}
