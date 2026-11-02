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


// Clamping is used to constraint an object within the viewport of the camera
// It takes input of a value and compares against the camera range
// If the value is not within the range, a false is returned
func (c *Camera) Clamp(value int) bool {
    if value < c.ViewportX || value >= c.ViewportW || value < c.ViewportY || value >= c.ViewportH {
        return false
    }

    return true
}
