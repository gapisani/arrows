package core

import (
    "fmt"
)

type Direction uint
const (
    NORD Direction = iota
    SOUTH
    EAST
    WEST
)

type LDirection uint
const (
    RIGHT LDirection = iota
    LEFT
)

type Cell interface {
    forcedUpdate() bool
    Update(grid [3][3](*Cell)) []point
}

/* Empty cell */
type None struct {}

// No direction but technically always faced nord
func (None) Dir() Direction {
    return NORD
}

// In fact, it should never be updated
func (None) Update([3][3](*Cell)) []point { return []point{} }

// Doesn't forces updates on other cells
func (None) forcedUpdate() bool {
    return false
}
// ------------


/* Wire */
type Wire struct {
    dir Direction
    lit bool
}

// Has a direction
func (w Wire) Dir() Direction {
    return w.dir
}

// Doesn't forces updates on other cells
func (Wire) forcedUpdate() bool {
    return false
}

// Pass signal to a cell that it faced with
func (w Wire) Update(grid [3][3](*Cell)) []point {
    if(!w.lit) {return []point{}}
    // TODO
    return []point{}
}
// ------------

/* Source */
type Source struct {
    dir Direction
}

// Has a direction
func (s Source) Direction() Direction {
    return s.dir
}

// Updates other cells
func (Source) forcedUpdate() bool {
    return true
}

// Powers the next cell
func (s Source) Update(grid [3][3](*Cell)) []point {
    p := dir2point(s.dir, point{1, 1})
    switch t := (*grid[p.x][p.y]).(type) {
    default:
        fmt.Printf("%T\n", t)
    }
    return []point{}
    // TODO: This is just testing code
}
// ------------

/* Memory Cell */
type MemCell struct {
    dir Direction
    state bool    // State -> On/Off
}

func (mc MemCell) Dir() Direction {
    return mc.dir
}

func (MemCell) forcedUpdate() bool {
    return false
}

// Works as source that can be turned off or on
func (mc MemCell) Update(grid [3][3](*Cell)) []point {
    if(!mc.state) {return []point{}}
    p := dir2point(mc.dir, point{1, 1})
    // TODO
    _ = p
    return []point{}
}
// ------------
