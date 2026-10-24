package main

// Handles the input actions.

import (
    "os"
    "github.com/Nareleth/kyanbasu"
)


// Read Key Input
func ReadInput(t *kyanbasu.Term) {
    for {
        // Read key input
        key := t.ReadKey()


        // Parse key input
        switch(key) {
            // Close Game
            case kyanbasu.Key_q, kyanbasu.Key_Q:
                t.CloseTerm()
                os.Exit(0)
        }
    }
}
