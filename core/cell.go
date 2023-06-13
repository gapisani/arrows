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
    updateQueue() []point
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
func (None) updateQueue() []point { return []point{} }
// ------------

/* Wire */
type Wire struct {
    dir Direction
    lit bool
}

func (a Wire) Dir() Direction { return a.dir }
func (a *Wire) SetDir(dir Direction) { a.dir = dir }
func (a Wire) updateQueue() []point {
    return []point{
        dir2point(a.dir, point{1,1}),
    }
}

func (w Wire) Check() bool { return w.lit }

func (w *Wire) Power() {
    w.lit = true
}

// Doesn't forces updates on other cells
func (Wire) forcedUpdate() bool {
    return false
}

// Pass signal to a cell that it faced with
func (a *Wire) Update(grid _lgrid) {
    if(!a.lit) {return}
    p := a.updateQueue()[0]
    cell := grid[p.x][p.y]
    (*cell).Power()
    a.lit = false
}
// ------------

/* Forward-Left Wire */
type FrwdLeft struct {
    lit bool
    dir Direction
}

func (a FrwdLeft) updateQueue() []point {
    if(!a.lit) { return []point{} }
    return []point{
        dir2point(rotateDir(a.dir, LEFT), point{1,1}),
        dir2point(a.dir, point{1,1}),
    }
}
func (a FrwdLeft) Dir() Direction { return a.dir }

func (a *FrwdLeft) SetDir(dir Direction) { a.dir = dir }

func (fd FrwdLeft) Check() bool { return fd.lit }

func (fd *FrwdLeft) Power() { fd.lit = true }

// Doesn't forces updates on other cells
func (FrwdLeft) forcedUpdate() bool { return false }

// Pass signal to a cell that it faced with, as well as on the left side
//[.][O][.] X - arrow; I - input; O - output
//[O][X][.]
//[.][I][.]
func (a *FrwdLeft) Update(grid _lgrid) {
    if(!a.lit) { return }
    p1 := a.updateQueue()[0]
    p2 := a.updateQueue()[1]
    cell := grid[p1.x][p1.y]
    (*cell).Power()
    cell = grid[p2.x][p2.y]
    (*cell).Power()
    a.lit = false
}
// ------------

/* Forward-Right Wire */
type FrwdRight struct {
    lit bool
    dir Direction
}

func (a FrwdRight) updateQueue() []point {
    return []point{
        dir2point(a.dir, point{1,1}),
        dir2point(rotateDir(a.dir, RIGHT), point{1,1}),
    }
}

func (a FrwdRight) Dir() Direction { return a.dir }

func (a *FrwdRight) SetDir(dir Direction) { a.dir = dir }

func (fr FrwdRight) Check() bool {
    return fr.lit
}

func (fr *FrwdRight) Power() {
    fr.lit = true
}

// Doesn't forces updates on other cells
func (FrwdRight) forcedUpdate() bool {
    return false
}

// Pass signal to a cell that it faced with, as well as on the right side
//[.][O][.] X - arrow; I - input; O - output
//[.][X][O]
//[.][I][.]
func (a *FrwdRight) Update(grid _lgrid) {
    if(!a.lit) { return }
    p1 := a.updateQueue()[0]
    p2 := a.updateQueue()[1]
    cell := grid[p1.x][p1.y]
    (*cell).Power()
    cell = grid[p2.x][p2.y]
    (*cell).Power()
    a.lit = false
}
// ------------

/* Cross Wire */
type Cross struct {
    lit bool
    dir Direction
}

func (a Cross) updateQueue() []point {
    return []point{
        dir2point(rotateDir(a.dir, LEFT), point{1,1}),
        dir2point(a.dir, point{1,1}),
        dir2point(rotateDir(a.dir, RIGHT), point{1,1}),
    }
}

func (c Cross) Check() bool {
    return c.lit
}

func (c *Cross) Power() {
    c.lit = true
}

func (a Cross) Dir() Direction { return a.dir }

func (a *Cross) SetDir(dir Direction) { a.dir = dir }

// Doesn't forces updates on other cells
func (Cross) forcedUpdate() bool {
    return false
}

// Pass signal to a cell that it faced with, as well as on the left and right side
//[.][O][.] X - arrow; I - input; O - output
//[O][X][O]
//[.][I][.]
func (a *Cross) Update(grid _lgrid) {
    if(!a.lit) { return }
    for _, p := range(a.updateQueue()) {
        (*grid[p.x][p.y]).Power()
    }
    a.lit = false
}
// ------------

/* Angled Wire */
type Angled struct {
    lit bool
    dir Direction
}

func (a Angled) Check() bool { return a.lit }

func (a *Angled) Power() { a.lit = true }

func (a Angled) Dir() Direction { return a.dir }

func (a *Angled) SetDir(dir Direction) { a.dir = dir }

// Doesn't forces updates on other cells
func (Angled) forcedUpdate() bool { return false }
func (a Angled) updateQueue() []point {
    return []point{
        dir2point(rotateDir(a.dir, LEFT), dir2point(a.dir, point{1,1})),
    }
}

// Pass signal to a cell on the top left relatively of the that it faced
//[O][.][.] X - arrow; I - input; O - output
//[.][X][.]
//[.][I][I]
//P.S. IDK where input is
func (a *Angled) Update(grid _lgrid) {
    if(!a.lit) { return }
    p := a.updateQueue()[0]
    cell := grid[p.x][p.y]
    (*cell).Power()
    a.lit = false
}
// ------------

/* Source */
type Source struct {}

func (a Source) updateQueue() []point {
    return []point{
        dir2point(NORTH, point{1, 1}),
        dir2point(EAST,  point{1, 1}),
        dir2point(SOUTH, point{1, 1}),
        dir2point(WEST,  point{1, 1}),
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
        (*grid[p.x][p.y]).Power()
    }
}
// ------------

/* Memory Cell */
type MemCell struct {
    dir Direction
    state bool    // State -> On/Off
}
func (a MemCell) updateQueue() []point {
    return []point{
        dir2point(a.dir, point{1, 1}),
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
    cell := grid[p.x][p.y]
    (*cell).Power()
}
// ------------

/* Flash */
type Flash struct {
    used bool
    dir Direction
}

// Same as memcell
func (f Flash) forcedUpdate() bool { return !f.used}

func (a Flash) updateQueue() []point {
    return []point{
        dir2point(a.dir, point{1,1}),
    }
}

func (Flash) Power() {}

func (a Flash) Dir() Direction { return a.dir }

func (a *Flash) SetDir(dir Direction) { a.dir = dir }

func (f Flash) Check() bool { return !f.used }

func (flash *Flash) Update(grid _lgrid) {
    if(flash.used) {
        return
    } else { flash.used = true }
    p := flash.updateQueue()[0]
    (*grid[p.x][p.y]).Power()
}
// ------------

/* Not */
type Not struct {
    dir Direction
    lit bool
}

func (a Not) Dir() Direction { return a.dir }

func (a Not) updateQueue() []point {
    return []point{
        dir2point(a.dir, point{1, 1}),
    }
}
func (a *Not) SetDir(dir Direction) { a.dir = dir }

func (a Not) forcedUpdate() bool { return a.lit }

func (a Not) Check() bool { return a.lit }

func (Not) Power() {}

func (not *Not) Update(grid _lgrid) {
    // p1 -[NOT]-> p2
    p1 := dir2point(rotateDir(not.dir, BACK), point{1, 1})
    p2 := not.updateQueue()[0]

    b1 := grid[p1.x][p1.y]
    b2 := grid[p2.x][p2.y]
    if(not.lit) {
        (*b2).Power()
        not.lit = true
    }
    if(!(*b1).Check()) {
        not.lit = true
    }
}
// ------------

/* Xor */
type Xor struct {
    dir Direction
    lit bool
    lcount uint
}

func (a Xor) Dir() Direction { return a.dir }

func (a Xor) updateQueue() []point {
    return []point{
        dir2point(a.dir, point{1, 1}),
    }
}
func (a *Xor) SetDir(dir Direction) { a.dir = dir }

// When it's not updated probably it doesn't have signal
func (Xor) Check() bool { return false }

func (a *Xor) Power() {
    a.lcount++
}

func (Xor) forcedUpdate() bool { return false }

func (xor *Xor) Update(grid _lgrid) {
    rp := xor.updateQueue()[0]
    rb := grid[rp.x][rp.y]
    if(xor.lcount == 1) {
        (*rb).Power()
        xor.lit = false
    }
    xor.lcount = 0
}
// ------------

/* And */
type And struct {
    dir Direction
    lit bool
    lcount uint
}

func (a And) updateQueue() []point {
    return []point{
        dir2point(a.dir, point{1, 1}),
    }
}
func (a And) Check() bool { return a.lit }

func (a And) Dir() Direction { return a.dir }

func (a *And) SetDir(dir Direction) { a.dir = dir }

func (a *And) Power() {
    a.lcount++
}

func (And) forcedUpdate() bool { return false }

func (and *And) Update(grid _lgrid) {
    rp := and.updateQueue()[0]
    rb := grid[rp.x][rp.y]
    if(and.lit) {
        (*rb).Power()
        and.lit = false
    }
    if(and.lcount >= 2) {
        and.lit = true
    }
    and.lcount = 0
}
// ------------

/* Block */
type Block struct {
    dir Direction
    lit bool
}
// TODO: redo logic

func (a Block) updateQueue() []point {
    return []point{
        dir2point(a.dir, point{1,1}),
    }
}

func (a Block) Dir() Direction { return a.dir }

func (a *Block) SetDir(dir Direction) { a.dir = dir }

func (a Block) Check() bool { return a.lit}
func (Block) Power() {}
func (Block) forcedUpdate() bool { return false }
func (block *Block) Update(grid _lgrid) {
    // FIXME: This version of block won't handle Cross, Angled, etc
    //  p1 -[block]?-> p2
    p1 := dir2point(rotateDir(block.dir, BACK), point{1,1})
    p2 := block.updateQueue()[0]
    b1 := grid[p1.x][p1.y]
    b2 := grid[p2.x][p2.y]
    if(block.lit) {
        switch t := (*b2).(type) {
        case *Wire:
            t.lit = false
        case *MemCell:
            t.state = false
        }
        block.lit = false
    }
    if((*b1).Check()) {
        block.lit = true
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

func (a Get) updateQueue() []point {
    return []point{
        dir2point(a.dir, point{1, 1}),
    }
}
// It won't be loaded directly so I'll update it forced
func (Get) forcedUpdate() bool { return true }

func (get *Get) Update(grid _lgrid) {
    p1 := dir2point(rotateDir(get.dir, BACK), point{1, 1})
    b1 := grid[p1.x][p1.y]
    if(get.state) {
        p2 := get.updateQueue()[0]
        b2 := grid[p2.x][p2.y]
        (*b2).Power()
        get.state = false
    }
    if((*b1).Check()) {
        get.state = true
    }
}
// ------------


// Double
type Double struct {
    dir Direction
    lit uint8
}

func (a *Double) Update(g _lgrid) {
    switch a.lit {
    case 0:
        return
    case 1:
        a.lit++
    case 2:
        p := a.updateQueue()[0]
        (*g[p.x][p.y]).Power()
        a.lit = 0
    }
}

func (Double) forcedUpdate() bool {
    return false
}

func (a Double) Dir() Direction { return a.dir }
func (a *Double) SetDir(dir Direction) { a.dir = dir }

func (a *Double) Power() {
    a.lit = 1
}
func (a Double) Check() bool {
    return a.lit == 2
}
func (a Double) updateQueue() []point {
    return []point{
        dir2point(a.dir, point{1, 1}),
    }
}
//-----------

// Double
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
        (*g[p.x][p.y]).Power()
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
func (a Random) updateQueue() []point {
    return []point{
        dir2point(a.dir, point{1, 1}),
    }
}
//-----------

/* Double Memory Cell */
type DoubleMemCell struct {
    dir Direction
    lcount uint
    state bool    // State -> On/Off
}
func (a DoubleMemCell) updateQueue() []point {
    return []point{
        dir2point(a.dir, point{1, 1}),
    }
}

func (a DoubleMemCell) Dir() Direction { return a.dir }

func (a *DoubleMemCell) SetDir(dir Direction) { a.dir = dir }

func (a DoubleMemCell) Check() bool { return a.state }

func (a DoubleMemCell) forcedUpdate() bool {
    return a.state
}

func (a *DoubleMemCell) Power() {
    a.lcount++
}

// Works as source that can be turned off or on
func (a *DoubleMemCell) Update(grid _lgrid) {
    if(a.state) {
        p := a.updateQueue()[0]
        cell := grid[p.x][p.y]
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
