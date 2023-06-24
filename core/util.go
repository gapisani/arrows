package core

type Point struct {
    X, Y uint
}

type Direction uint8
const (
    NORTH Direction = iota
    EAST
    SOUTH
    WEST
)

type LDirection uint8
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
func dir2point(dir Direction, p Point) Point {
    switch dir {
    case NORTH:
        return Point{p.X, p.Y-1}
    case WEST:
        return Point{p.X-1, p.Y}
    case SOUTH:
        return Point{p.X, p.Y+1}
    case EAST:
        return Point{p.X+1, p.Y}
    }
    return Point{}
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
