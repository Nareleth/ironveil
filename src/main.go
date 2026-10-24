package main

import (
    "fmt"
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

    // Go routine read key input
    go ReadInput(t)

    Player := AddEntityPlayer("Jeff", '@', 1, 1)
    PlayerPos := m_Position[Player]

    fmt.Printf("Player: %s\nSprite: %c\nPos: (%d, %d)\n", Player.GetName(), Player.GetSprite(), PlayerPos.X, PlayerPos.Y)
    fmt.Println("Number of Entities: ", len(Entities))
    

    // Game Loop
    for {

        // Frame limit
        time.Sleep(time.Second / time.Duration(FPS))
    }
}
