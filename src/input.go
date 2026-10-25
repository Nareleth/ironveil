package main

// Handles the input actions.

import (
    "os"
    "github.com/Nareleth/kyanbasu"
)


// Read Key Input
// TEMP!!!!!!!!!!!! remove the direct player passing here. should pass world
func ReadInput(t *kyanbasu.Term, Player Entity) {
    for {
        // Read key input
        key := t.ReadKey()


        // Parse key input
        switch(key) {
            // Close Game
            case kyanbasu.Key_q, kyanbasu.Key_Q:
                t.CloseTerm()
                os.Exit(0)

            // Movement
            case kyanbasu.Key_w:
                Player.SetVelocity(0, -1)
            case kyanbasu.Key_a:
                Player.SetVelocity(-1, 0)
            case kyanbasu.Key_s:
                Player.SetVelocity(0, 1)
            case kyanbasu.Key_d:
                Player.SetVelocity(1, 0)
        }
    }
}
