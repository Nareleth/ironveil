package main

import (
    "github.com/Nareleth/kyanbasu"
)

// A camera is used to represent the viewport and rendering area
// The camera is centered on and moves with the player
type Camera struct {
    ViewportX, ViewportY int
    ViewportW, ViewportH int
}


// Initialize a new camera using the kyanbasu terminal size
func NewCamera(term *kyanbasu.Term) *Camera {
    w, h := term.GetSize()

    return &Camera {
        ViewportX: 0,
        ViewportY: 0,
        ViewportW: w, 
        ViewportH: h,
    }
}
