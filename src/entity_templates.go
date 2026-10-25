package main

// Templates for creating streamlined Entities


// Player
func AddEntityPlayer(w *World, name string, sprite rune, x,y int) Entity {
    e := AddEntity(w)
    e.SetName(name)
    e.SetSprite(sprite)
    e.SetGlyph(glyph_e_player)
    e.SetPosition(x, y)
    e.SetVelocity(0, 0)

    return e
}
