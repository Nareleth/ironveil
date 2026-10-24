package main

// Templates for creating streamlined Entities


// Player
func AddEntityPlayer(name string, sprite rune, x,y int) Entity {
    e := AddEntity()
    e.SetName(name)
    e.SetSprite(sprite)
    e.SetPosition(x, y)
    e.SetVelocity(0, 0)

    return e
}
