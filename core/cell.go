package core

import "math/rand"

const (
    _LGRID_SIZE = 3
)
type _lgrid [_LGRID_SIZE][_LGRID_SIZE](*Cell)

type Cell interface {
    forcedUpdate() bool
    Update(grid _lgrid)
    Check() bool
    Power()
    Dir() Direction
    SetDir(Direction)
    updateQueue() []Point
}

/* Empty cell */
type None struct {}

// In fact, it should never be updated
func (None) Update(_lgrid) {}

// Doesn't forces updates on other cells
func (None) forcedUpdate() bool {
    return false
}

func (None) Dir() Direction { return NORTH }
func (None) SetDir(dir Direction) { }

func (None) Power() {}
func (None) Check() bool { return false }
func (None) updateQueue() []Point { return []Point{} }
// ------------

/* Source */
type Source struct {}

func (a Source) updateQueue() []Point {
    return []Point{
        dir2point(NORTH, Point{1, 1}),
        dir2point(EAST,  Point{1, 1}),
        dir2point(SOUTH, Point{1, 1}),
        dir2point(WEST,  Point{1, 1}),
    }
}
func (a Source) Dir() Direction { return NORTH }

func (Source) SetDir(dir Direction) { }

func (Source) Check() bool { return true }

func (Source) Power() {}

// Updates other cells
func (Source) forcedUpdate() bool {
    return true
}

// Powers the next cell
func (a *Source) Update(grid _lgrid) {
    for _, p := range(a.updateQueue()) {
        (*grid[p.X][p.Y]).Power()
    }
}
// ------------

/* Memory Cell */
type MemCell struct {
    dir Direction
    state bool    // State -> On/Off
}
func (a MemCell) updateQueue() []Point {
    return []Point{
        dir2point(a.dir, Point{1, 1}),
    }
}

func (a MemCell) Dir() Direction { return a.dir }

func (a *MemCell) SetDir(dir Direction) { a.dir = dir }

func (mc MemCell) Check() bool { return mc.state }

// Depends, it could work as source
func (mc MemCell) forcedUpdate() bool {
    return mc.state
}

func (mc *MemCell) Power() {
    mc.state = !mc.state
}

// Works as source that can be turned off or on
func (a *MemCell) Update(grid _lgrid) {
    if(!a.state) {return}
    p := a.updateQueue()[0]
    cell := grid[p.X][p.Y]
    (*cell).Power()
}
// ------------

/* Flash */
type Flash struct {
    used bool
}

// Same as memcell
func (f Flash) forcedUpdate() bool { return !f.used}

func (a Flash) updateQueue() []Point {
    return []Point{
        dir2point(NORTH, Point{1, 1}),
        dir2point(EAST,  Point{1, 1}),
        dir2point(SOUTH, Point{1, 1}),
        dir2point(WEST,  Point{1, 1}),
    }
}

func (Flash) Power() {}

func (Flash) Dir() Direction { return NORTH }

func (Flash) SetDir(Direction) {}

func (f Flash) Check() bool { return !f.used }

func (flash *Flash) Update(grid _lgrid) {
    if(flash.used) {
        return
    } else { flash.used = true }
    for _, p := range(flash.updateQueue()) {
        (*grid[p.X][p.Y]).Power()
    }
}
// ------------

/* Get */
type Get struct {
    dir Direction
    state bool
}

func (a Get) Dir() Direction { return a.dir }
func (a *Get) SetDir(dir Direction) { a.dir = dir }
func (g Get) Check() bool { return g.state }
func (Get) Power() {}

func (a Get) updateQueue() []Point {
    return []Point{
        dir2point(a.dir, Point{1, 1}),
    }
}

// It won't be loaded directly so I'll update it forced
func (Get) forcedUpdate() bool { return true }

func (get *Get) Update(grid _lgrid) {
    p1 := dir2point(rotateDir(get.dir, BACK), Point{1, 1})
    b1 := grid[p1.X][p1.Y]
    if(get.state) {
        p2 := get.updateQueue()[0]
        b2 := grid[p2.X][p2.Y]
        (*b2).Power()
        get.state = false
    }
    if((*b1).Check()) {
        get.state = true
    }
}
// ------------


// Random
type Random struct {
    dir Direction
    lit uint8
}

func (a *Random) Update(g _lgrid) {
    switch a.lit {
    case 0:
        return
    case 1:
        a.lit = 0
    case 2:
        p := a.updateQueue()[0]
        (*g[p.X][p.Y]).Power()
        a.lit = 0
    }
}

func (Random) forcedUpdate() bool {
    return false
}

func (a Random) Dir() Direction { return a.dir }
func (a *Random) SetDir(dir Direction) { a.dir = dir }

func (a *Random) Power() {
    a.lit = uint8(rand.Int()) % 2
}
func (a Random) Check() bool {
    return a.lit == 2
}
func (a Random) updateQueue() []Point {
    return []Point{
        dir2point(a.dir, Point{1, 1}),
    }
}
//-----------

/* Double Memory Cell */
type DoubleMemCell struct {
    dir Direction
    lcount uint
    state bool    // State -> On/Off
}
func (a DoubleMemCell) updateQueue() []Point {
    return []Point{
        dir2point(a.dir, Point{1, 1}),
    }
}

func (a DoubleMemCell) Dir() Direction { return a.dir }

func (a *DoubleMemCell) SetDir(dir Direction) { a.dir = dir }

func (a DoubleMemCell) Check() bool { return a.state }

func (a DoubleMemCell) forcedUpdate() bool { return a.state }

func (a *DoubleMemCell) Power() {
    a.lcount++
}

// Works as source that can be turned off or on
func (a *DoubleMemCell) Update(grid _lgrid) {
    if(a.state) {
        p := a.updateQueue()[0]
        cell := grid[p.X][p.Y]
        (*cell).Power()
    }
    if(a.lcount >= 2) {
        a.state = true
    } else if(a.state && a.lcount == 1) {
        a.state = false
    }
    a.lcount = 0
}
// ------------

// TODO More types (i guess)
