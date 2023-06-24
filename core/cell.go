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
func (None) SetDir(Direction) { }

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

func (Source) SetDir(Direction) { }

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
    Direction Direction
    State bool    // State -> On/Off
}
func (a MemCell) updateQueue() []Point {
    return []Point{
        dir2point(a.Direction, Point{1, 1}),
    }
}

func (a MemCell) Dir() Direction { return a.Direction }

func (a *MemCell) SetDir(dir Direction) { a.Direction = dir }

func (mc MemCell) Check() bool { return mc.State }

// Depends, it could work as source
func (mc MemCell) forcedUpdate() bool {
    return mc.State
}

func (mc *MemCell) Power() {
    mc.State = !mc.State
}

// Works as source that can be turned off or on
func (a *MemCell) Update(grid _lgrid) {
    if(!a.State) {return}
    p := a.updateQueue()[0]
    cell := grid[p.X][p.Y]
    (*cell).Power()
}
// ------------

/* Flash */
type Flash struct {
    Used bool
}

// Same as memcell
func (f Flash) forcedUpdate() bool { return !f.Used}

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

func (f Flash) Check() bool { return !f.Used }

func (flash *Flash) Update(grid _lgrid) {
    if(flash.Used) {
        return
    } else { flash.Used = true }
    for _, p := range(flash.updateQueue()) {
        (*grid[p.X][p.Y]).Power()
    }
}
// ------------

/* Get */
type Get struct {
    Direction Direction
    State bool
}

func (a Get) Dir() Direction { return a.Direction }
func (a *Get) SetDir(dir Direction) { a.Direction = dir }
func (g Get) Check() bool { return g.State }
func (Get) Power() {}

func (a Get) updateQueue() []Point {
    return []Point{
        dir2point(a.Direction, Point{1, 1}),
    }
}

// It won't be loaded directly so I'll update it forced
func (Get) forcedUpdate() bool { return true }

func (get *Get) Update(grid _lgrid) {
    p1 := dir2point(rotateDir(get.Direction, BACK), Point{1, 1})
    b1 := grid[p1.X][p1.Y]
    if(get.State) {
        p2 := get.updateQueue()[0]
        b2 := grid[p2.X][p2.Y]
        (*b2).Power()
        get.State = false
    }
    if((*b1).Check()) {
        get.State = true
    }
}
// ------------


// Random
type Random struct {
    Direction Direction
    Lit uint8
}

func (a *Random) Update(g _lgrid) {
    switch a.Lit {
    case 0:
        return
    case 2:
        p := a.updateQueue()[0]
        (*g[p.X][p.Y]).Power()
    }
    a.Lit = 0
}

func (Random) forcedUpdate() bool {
    return false
}

func (a Random) Dir() Direction { return a.Direction }
func (a *Random) SetDir(dir Direction) { a.Direction = dir }

func (a *Random) Power() {
    a.Lit = uint8(rand.Int()) % 2
}
func (a Random) Check() bool {
    return a.Lit == 2
}
func (a Random) updateQueue() []Point {
    return []Point{
        dir2point(a.Direction, Point{1, 1}),
    }
}
//-----------

/* Double Memory Cell */
type DoubleMemCell struct {
    Direction Direction
    lcount uint
    State bool    // State -> On/Off
}
func (a DoubleMemCell) updateQueue() []Point {
    return []Point{
        dir2point(a.Direction, Point{1, 1}),
    }
}

func (a DoubleMemCell) Dir() Direction { return a.Direction }

func (a *DoubleMemCell) SetDir(dir Direction) { a.Direction = dir }

func (a DoubleMemCell) Check() bool { return a.State }

func (a DoubleMemCell) forcedUpdate() bool { return a.State }

func (a *DoubleMemCell) Power() {
    a.lcount++
}

// Works as source that can be turned off or on
func (a *DoubleMemCell) Update(grid _lgrid) {
    if(a.State) {
        p := a.updateQueue()[0]
        cell := grid[p.X][p.Y]
        (*cell).Power()
    }
    if(a.lcount >= 2) {
        a.State = true
    } else if(a.State && a.lcount == 1) {
        a.State = false
    }
    a.lcount = 0
}
// ------------

// TODO More types (i guess)
