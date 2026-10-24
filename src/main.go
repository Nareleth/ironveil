package main

import (
    "fmt"
)


func main() {
    Player := AddEntityPlayer("Jeff", '@', 1, 1)
    PlayerPos := m_Position[Player]

    fmt.Printf("Player: %s\nSprite: %c\nPos: (%d, %d)\n", Player.GetName(), Player.GetSprite(), PlayerPos.X, PlayerPos.Y)
    fmt.Println("Number of Entities: ", len(Entities))
}
