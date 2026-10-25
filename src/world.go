package main

// The world holds all the entities and maps that make up... the world.
// Thus, all funcs that change the world should provide the world as an argument.

type World struct {
    Entities    []Entity
}


// Initialize a new world
func NewWorld() *World {
    return &World{}
}
