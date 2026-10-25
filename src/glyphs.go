package main

// To eventually be replaced with a json file.
// Glyphs are the runes for tiles, entities, and effects.
// Glyphs follow a naming format: glyph_type_id
// Types: 
//    t: tile
//    e: entity

import (
    "github.com/Nareleth/kyanbasu"
)

var (
    glyph_e_player = Glyph{ Char: '@', FG: kyanbasu.Color_Blue, BG: kyanbasu.Color_Reset, }
    glyph_e_grass = Glyph{ Char: '~', FG: kyanbasu.Color_Green, BG: kyanbasu.Color_Reset, }
)
