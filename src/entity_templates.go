package main

// Templates for creating streamlined Entities


// Player
func AddEntityPlayer(w *World, name string, sprite rune, x,y int, camera *Camera) Entity {
    e := AddEntity(w)
    e.SetName(name)
    e.SetGlyph(glyph_e_player)
    e.SetPosition(x, y)
    e.SetVelocity(0, 0)
    e.SetCamera(camera)

    return e
}


// Tree
func AddEntityTree(w *World, x,y int) Entity {
    e := AddEntity(w)
    e.SetGlyph(glyph_e_tree)
    e.SetPosition(x, y)

    return e
}
