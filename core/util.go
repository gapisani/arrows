package core

// Internal thing, point
type point struct {
    x, y uint
}

type Direction uint
const (
    NORTH Direction = iota
    SOUTH
    EAST
    WEST
)

type LDirection uint
const (
    RIGHT LDirection = iota
    LEFT
    BACK
)

// (1,1) -> center
// (0,0) (1,0) (2,0)
// (0,1) (1,1) (2,1)
// (0,2) (1,2) (2,2)

// Converts direction to point(or vector)
func dir2point(dir Direction, p point) point {
    switch dir {
    case NORTH:
        return point{p.x, p.y-1}
    case WEST:
        return point{p.x-1, p.y}
    case SOUTH:
        return point{p.x, p.y+1}
    case EAST:
        return point{p.x+1, p.y}
    }
    return point{}
}

// Rotates direction
func rotateDir(dir Direction, ldir LDirection) Direction {
    switch ldir {
    case RIGHT:
        dir = (dir + 1) % 4
    case LEFT:
        dir = (dir + 3) % 4
    case BACK:
        dir = (dir + 2) % 4
    }
    return dir
}
