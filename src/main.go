package main

import (
    "time"
    "github.com/Nareleth/kyanbasu"
)


// Config
const (
    FPS         = 30
    world_size  = 200
    seed uint64 = 12345
    resolutionX = 100
    resolutionY = 40
)

var (
    Debug       = true
)


func main() {
    // Initialize term controls
    t := kyanbasu.NewTerm()
    defer t.CloseTerm()

    // Initialize screen flushing
    screen := kyanbasu.NewScreen(t)

    // Initialize the main game canvas window
    win_canvas := kyanbasu.NewWindow(screen, 0, 1, resolutionX, resolutionY)

    // Initialize the alert prompt window
    win_alert := kyanbasu.NewWindow(screen, 0, 0, resolutionX, 3)

    // Initialize the main game UI window
    win_ui := kyanbasu.NewWindow(screen, resolutionX, 1, resolutionX, resolutionY-7)

    // Initialize the game logs window
    win_logs := kyanbasu.NewWindow(screen, resolutionX, resolutionY-6, resolutionX, 7)

    // Initialize debug window
    win_debug := kyanbasu.NewWindow(screen, resolutionX + resolutionX, 1, 80, resolutionY)
    if Debug == false {
        win_debug.Hide()
    }


    // Initialize the world
    world := NewWorld(world_size)
    world.GenerateMap()

    // Initialize the viewport
    camera := NewCamera(t)

    // Initialize Player
    Player := AddEntityPlayer(world, "Jeff", '@', 1, 1, camera)



    // Create a channel to hold 16 commands before flushing
    command_queue := make(chan Command, 16)

    // Go routine read key input
    go ReadInput(t, command_queue)

    // Handle input
    go CommandQueue(t, command_queue, Player)

    // Game Loop
    for {
        // Temp just for UI
        /*
        PlayerPos   := m_Position[Player]
        PlayerState := m_State[Player]
        */
        //test temporary !!!
        win_canvas.WriteString(0, 0, "Game")
        win_ui.WriteString(0, 0, "UI")
        win_logs.WriteString(0, 0, "Logs")
        win_alert.WriteString(1, -1, "Hey, Listen!")
        win_debug.WriteString(0, 0, "Debug: %t", Debug)

        // Systems
        s_Movement(world, Player)

        // Render Layers
        win_canvas.DrawBorder()
        win_ui.DrawBorder()
        win_logs.DrawBorder()
        win_debug.DrawBorder()
        //RenderTiles(screen, world, camera)
        //RenderEntities(screen, world, camera)

        // Print HUD
        /*
        t.Move(0,0)
        t.WriteString("Pos: (%d,%d)", PlayerPos.X, PlayerPos.Y)
        t.WriteString("State: %s", PlayerState)
        */

        // Flush render instructions
        screen.Flush()

        // Frame limit
        time.Sleep(time.Second / time.Duration(FPS))
    }
}
