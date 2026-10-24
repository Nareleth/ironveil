package main


// An array to store all entities.
var Entities []Entity


// An entity is an integer ID with a set of components.
// Each entity has a number of components and systems.
type Entity uint64


// Functions

// Create a new Entity
func AddEntity() Entity{
    newID := len(Entities)
    e := Entity(newID)
    Entities = append(Entities, e)
    return e
}

// Methods

// Get Entity ID
func (e Entity) GetID() Entity {
    return e
}

