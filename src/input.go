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
var keybinds = map[kyanbasu.Key]Command {
    kyanbasu.Key_c: cmd_move_down,
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
            case kyanbasu.Key_q, kyanbasu.Key_Q:
                queue <- cmd_state_quit

            // Movement
            case kyanbasu.Key_w:
                queue <- cmd_move_up

            case kyanbasu.Key_a:
                queue <- cmd_move_left

            case kyanbasu.Key_s:
                queue <- cmd_move_down

            case kyanbasu.Key_d:
                queue <- cmd_move_right
            

            // Commands
            case kyanbasu.Key_c:
                queue <- cmd_move_down
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
