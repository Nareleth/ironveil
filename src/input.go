package main

// Handles the input actions.

import (
    "os"
    "github.com/Nareleth/kyanbasu"
)



// Initialize an input command type for handling key commands
type Command int

// Initialize commands
const (
    cmd_none Command = iota
    cmd_state_quit
    cmd_move_up
    cmd_move_down
    cmd_move_left
    cmd_move_right
)


// Bindable hotkeys
// Later convert to reading from conf file
var keybinds = map[Command]kyanbasu.Key {
    cmd_state_quit: kyanbasu.Key_q,
    cmd_move_up:    kyanbasu.Key_w,
    cmd_move_down:  kyanbasu.Key_s,
    cmd_move_left:  kyanbasu.Key_a,
    cmd_move_right: kyanbasu.Key_d,
}


// Read Key Input and pass it to the command queue
func ReadInput(t *kyanbasu.Term, queue chan Command) {
    for {
        // Read key input
        key := t.ReadKey()

        // Parse key input
        switch(key) {
            //State
            // Close Game
            case keybinds[cmd_state_quit]:
                queue <- cmd_state_quit

            // Movement
            // Up
            case keybinds[cmd_move_up]:
                queue <- cmd_move_up

            // Down
            case keybinds[cmd_move_down]:
                queue <- cmd_move_down

            // Left
            case keybinds[cmd_move_left]:
                queue <- cmd_move_left

            // Right
            case keybinds[cmd_move_right]:
                queue <- cmd_move_right
            
        }
    }
}



// Queues received commands and handles them
func CommandQueue(t *kyanbasu.Term, queue chan Command, Player Entity) {
    for cmd := range queue {
        switch cmd {
            // Game state
            case cmd_state_quit:
                t.CloseTerm()
                os.Exit(0)

            // Movement
            case cmd_move_up:
                Player.SetVelocity(0, -1)

            case cmd_move_down:
                Player.SetVelocity(0, 1)

            case cmd_move_left:
                Player.SetVelocity(-1, 0)

            case cmd_move_right:
                Player.SetVelocity(1, 0)
        }
    }
}
