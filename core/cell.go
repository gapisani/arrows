package core

const (
    _LGRID_SIZE = 3
)
type _lgrid [_LGRID_SIZE][_LGRID_SIZE](*Cell)

type Cell interface {
    forcedUpdate() bool
    Update(grid _lgrid) []point
    Check() bool
    Power()
    Dir() Direction
    SetDir(Direction)
}

/* Empty cell */
type None struct {}

// In fact, it should never be updated
func (None) Update(_lgrid) []point { return []point{} }

// Doesn't forces updates on other cells
func (None) forcedUpdate() bool {
    return false
}

func (None) Dir() Direction { return NORTH }
func (None) SetDir(dir Direction) { }

func (None) Power() {}
func (None) Check() bool { return false }
// ------------

/* Wire */
type Wire struct {
    dir Direction
    lit bool
}

func (a Wire) Dir() Direction { return a.dir }
func (a *Wire) SetDir(dir Direction) { a.dir = dir }

func (w Wire) Check() bool { return w.lit }

func (w *Wire) Power() {
    w.lit = true
}

// Doesn't forces updates on other cells
func (Wire) forcedUpdate() bool {
    return false
}

// Pass signal to a cell that it faced with
func (w *Wire) Update(grid _lgrid) []point {
    if(!w.lit) {return []point{}}
    p := dir2point(w.dir, point{1,1})
    cell := grid[p.x][p.y]
    (*cell).Power()
    w.lit = false
    return []point{p}
}
// ------------

/* Forward-Left Wire */
type FrwdLeft struct {
    lit bool
    dir Direction
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
func (fd *FrwdLeft) Update(grid _lgrid) []point {
    if(!fd.lit) { return []point{} }
    p1 := dir2point(rotateDir(fd.dir, LEFT), point{1,1})
    cell := grid[p1.x][p1.y]
    (*cell).Power()
    p2 := dir2point(fd.dir, point{1,1})
    cell = grid[p2.x][p2.y]
    (*cell).Power()
    fd.lit = false
    return []point{p1, p2}
}
// ------------

/* Forward-Right Wire */
type FrwdRight struct {
    lit bool
    dir Direction
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
func (fr *FrwdRight) Update(grid _lgrid) []point {
    if(!fr.lit) { return []point{} }
    p1 := dir2point(fr.dir, point{1,1})
    cell := grid[p1.x][p1.y]
    (*cell).Power()
    p2 := dir2point(rotateDir(fr.dir, RIGHT), point{1,1})
    cell = grid[p2.x][p2.y]
    (*cell).Power()
    fr.lit = false
    return []point{p1, p2}
}
// ------------

/* Cross Wire */
type Cross struct {
    lit bool
    dir Direction
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
func (c *Cross) Update(grid _lgrid) []point {
    if(!c.lit) { return []point{} }
    p1 := dir2point(rotateDir(c.dir, LEFT), point{1,1})
    cell := grid[p1.x][p1.y]
    (*cell).Power()
    p2 := dir2point(c.dir, point{1,1})
    cell = grid[p2.x][p2.y]
    (*cell).Power()
    p3 := dir2point(rotateDir(c.dir, RIGHT), point{1,1})
    cell = grid[p3.x][p3.y]
    (*cell).Power()
    c.lit = false
    return []point{p1, p2, p3}
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

// Pass signal to a cell on the top left relatively of the that it faced
//[O][.][.] X - arrow; I - input; O - output
//[.][X][.]
//[.][I][I]
//P.S. IDK where input is
func (a *Angled) Update(grid _lgrid) []point {
    if(!a.lit) { return []point{} }
    p := dir2point(rotateDir(a.dir, LEFT), dir2point(a.dir, point{1,1}))
    cell := grid[p.x][p.y]
    (*cell).Power()
    a.lit = false
    return []point{p}
}
// ------------

/* Source */
type Source struct {
    dir Direction
}

func (a Source) Dir() Direction { return a.dir }

func (a *Source) SetDir(dir Direction) { a.dir = dir }

func (Source) Check() bool { return true }

func (Source) Power() {}

// Updates other cells
func (Source) forcedUpdate() bool {
    return true
}

// Powers the next cell
func (s *Source) Update(grid _lgrid) []point {
    p := dir2point(s.dir, point{1, 1})
    (*grid[p.x][p.y]).Power()
    return []point{p}
}
// ------------

/* Memory Cell */
type MemCell struct {
    dir Direction
    state bool    // State -> On/Off
}

func (a MemCell) Dir() Direction { return a.dir }

func (a *MemCell) SetDir(dir Direction) { a.dir = dir }

func (mc MemCell) Check() bool { return mc.state }

// Depends, it could work as source
func (mc MemCell) forcedUpdate() bool {
    // return mc.state
    return true // Not sure
}

func (mc *MemCell) Power() {
    mc.state = !mc.state
}

// Works as source that can be turned off or on
func (mc *MemCell) Update(grid _lgrid) []point {
    if(!mc.state) {return []point{}}
    p := dir2point(mc.dir, point{1, 1})
    cell := grid[p.x][p.y]
    (*cell).Power()
    return []point{p}
}
// ------------

/* Flash */
type Flash struct {
    used bool
    dir Direction
}

// Same as memcell
func (f Flash) forcedUpdate() bool { return !f.used}

func (Flash) Power() {}

func (a Flash) Dir() Direction { return a.dir }

func (a *Flash) SetDir(dir Direction) { a.dir = dir }

func (f Flash) Check() bool { return !f.used }

func (flash *Flash) Update(grid _lgrid) []point {
    if(flash.used) {
        return []point{}
    } else { flash.used = true }
    p := dir2point(flash.dir, point{1,1})
    cell := grid[p.x][p.y]
    (*cell).Power()
    return []point{p}
}
// ------------

/* Not */
type Not struct {
    dir Direction
}

func (a Not) Dir() Direction { return a.dir }

func (a *Not) SetDir(dir Direction) { a.dir = dir }

// Kinda same as source
func (Not) forcedUpdate() bool { return true }

// FIXME: ^ it's always updated, so there is some logic issue may-be?

// It should work ONLY when it's not updated, so if it's not updated it probably doesn't have signal
func (Not) Check() bool { return true }

func (Not) Power() {}

func (not *Not) Update(grid _lgrid) []point {
    // p1 -[NOT]-> p2
    p1 := dir2point(rotateDir(not.dir, BACK), point{1, 1})
    p2 := dir2point(not.dir, point{1, 1})

    b1 := grid[p1.x][p1.y]
    b2 := grid[p2.x][p2.y]
    if(!(*b1).Check()) {
        (*b2).Power()
    }
    return []point{p2}
}
// ------------

/* Xor */
type Xor struct {
    dir Direction
}

func (a Xor) Dir() Direction { return a.dir }

func (a *Xor) SetDir(dir Direction) { a.dir = dir }

// When it's not updated probably it doesn't have signal
func (Xor) Check() bool { return false }

func (Xor) Power() {}

func (Xor) forcedUpdate() bool { return false }
func (xor *Xor) Update(grid _lgrid) []point {
    p1 := dir2point(rotateDir(xor.dir, LEFT), point{1, 1})
    p2 := dir2point(rotateDir(xor.dir, RIGHT), point{1, 1})
    rp := dir2point(xor.dir, point{1, 1})
    b1 := grid[p1.x][p1.y]
    b2 := grid[p2.x][p2.y]
    rb := grid[rp.x][rp.y]
    if((*b1).Check() != (*b2).Check()) {
        (*rb).Power()
    }
    return []point{rp}
}
// ------------

/* And */
type And struct {
    dir Direction
}

// Same as xor
func (And) Check() bool { return false }

func (a And) Dir() Direction { return a.dir }

func (a *And) SetDir(dir Direction) { a.dir = dir }

func (And) Power() {}
func (And) forcedUpdate() bool { return false }
func (and *And) Update(grid _lgrid) []point {
    p1 := dir2point(rotateDir(and.dir, LEFT), point{1, 1})
    p2 := dir2point(rotateDir(and.dir, RIGHT), point{1, 1})
    rp := dir2point(and.dir, point{1, 1})
    b1 := grid[p1.x][p1.y]
    b2 := grid[p2.x][p2.y]
    rb := grid[rp.x][rp.y]
    if((*b1).Check() && (*b2).Check()) {
        (*rb).Power()
    }
    return []point{rp}
}
// ------------

/* Block */
type Block struct {
    dir Direction
}

func (a Block) Dir() Direction { return a.dir }

func (a *Block) SetDir(dir Direction) { a.dir = dir }

func (Block) Check() bool { return false }
func (Block) Power() {}
func (Block) forcedUpdate() bool { return false }
func (block *Block) Update(grid _lgrid) []point {
    // FIXME: This version of block won't handle Cross, Angled, etc
    //  p1 -[block]?-> p2
    p1 := dir2point(rotateDir(block.dir, BACK), point{1,1})
    p2 := dir2point(block.dir, point{1,1})
    b1 := grid[p1.x][p1.y]
    b2 := grid[p2.x][p2.y]
    if((*b1).Check()) {
        switch t := (*b2).(type) {
        case *Wire:
            t.lit = false
        case *MemCell:
            t.state = false
        }
    }
    return []point{p2}
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

// It won't be loaded directly so I'll update it forced
func (Get) forcedUpdate() bool { return true }

func (get *Get) Update(grid _lgrid) []point {
    p1 := dir2point(rotateDir(get.dir, BACK), point{1, 1})
    p2 := dir2point(get.dir, point{1, 1})
    b1 := grid[p1.x][p1.y]
    b2 := grid[p2.x][p2.y]
    if((*b1).Check()) {
        (*b2).Power()
    }
    return []point{p2}
}
// ------------

// TODO More types (i guess)
