package main


// An entity is an integer ID with a set of components.
// Each entity has a number of components and systems.
type Entity uint64


// Functions

// Create a new Entity
func AddEntity(w *World) Entity{
    newID := len(w.Entities)
    e := Entity(newID)
    w.Entities = append(w.Entities, e)
    return e
}

// Methods

// Get Entity ID
func (e Entity) GetID() Entity {
    return e
}

