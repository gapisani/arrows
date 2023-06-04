package core

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

type Cell interface {
    forcedUpdate() bool
    Update(grid [3][3](*Cell)) []point
    Check() bool
    Power()
}

/* Empty cell */
type None struct {}

// In fact, it should never be updated
func (None) Update([3][3](*Cell)) []point { return []point{} }

// Doesn't forces updates on other cells
func (None) forcedUpdate() bool {
    return false
}

func (None) Power() {}
func (None) Check() bool { return false }
// ------------


/* Wire */
type Wire struct {
    dir Direction
    lit bool
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
func (w *Wire) Update(grid [3][3](*Cell)) []point {
    if(!w.lit) {return []point{}}
    p := dir2point(w.dir, point{1,1})
    cell := grid[p.x][p.y]
    (*cell).Power()
    return []point{p}
}
// ------------

//Cross Wire    ==================== @dikiy_opezdal's exam code ^_^ ====================
type Cross struct {
    lit bool
    dir Direction
}

func (Cross) forcedUpdate() bool { return false }
func (c *Cross) Power() { c.lit = true }
func (c Cross) Check() bool { return c.lit }
func (c *Cross) Update(grid [3][3](*Cell)) []point {
    if(c.lit) {
        p := dir2point(rotateDir(c.dir, LEFT), point{1,1})
        cell := grid[p.x][p.y]
        (*cell).Power()

        p = dir2point(c.dir, point{1,1})
        cell = grid[p.x][p.y]
        (*cell).Power()

        p = dir2point(rotateDir(c.dir, RIGHT), point{1,1})
        cell = grid[p.x][p.y]
        (*cell).Power()

        return []point{p}
    }
    return []point{}
}
// ------------      ==================== @dikiy_opezdal's exam code ^_^ ====================

/* Source */
type Source struct {
    dir Direction
}

func (Source) Check() bool { return true }

func (Source) Power() {}

// Updates other cells
func (Source) forcedUpdate() bool {
    return true
}

// Powers the next cell
func (s *Source) Update(grid [3][3](*Cell)) []point {
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

func (mc MemCell) Check() bool { return mc.state }

// Depends, it could work as source
func (mc MemCell) forcedUpdate() bool {
    return mc.state
}

func (mc *MemCell) Power() {
    mc.state = !mc.state
}

// Works as source that can be turned off or on
func (mc *MemCell) Update(grid [3][3](*Cell)) []point {
    if(!mc.state) {return []point{}}
    p := dir2point(mc.dir, point{1, 1})
    cell := grid[p.x][p.y]
    switch t := (*cell).(type) {
    case *Wire:
        t.lit = true
    case *MemCell:
        t.state = !t.state
    }
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

func (f Flash) Check() bool { return !f.used }

func (flash *Flash) Update(grid [3][3](*Cell)) []point {
    if(flash.used) {
        return []point{}
    } else { flash.used = false }
    p := dir2point(flash.dir, point{1,1})
    cell := grid[p.x][p.y]
    switch t := (*cell).(type) {
    case *Wire:
        t.lit = true
    case *MemCell:
        t.state = !t.state
    }
    return []point{p}
}
// ------------

/* Not */
type Not struct {
    dir Direction
}

// Kinda same as source
func (Not) forcedUpdate() bool { return true }

// It should work ONLY when it's not updated, so if it's not updated it probably doesn't have signal
func (Not) Check() bool { return true }

func (Not) Power() {}

func (not *Not) Update(grid [3][3](*Cell)) []point {
    // p1 -[NOT]-> p2
    p1 := dir2point(rotateDir(not.dir, BACK), point{1, 1})
    p2 := dir2point(not.dir, point{1, 1})

    b1 := grid[p1.x][p1.y]
    b2 := grid[p2.x][p2.y]
    enabled := true
    switch t := (*b1).(type) {
    case *Wire:
        enabled = !t.lit
    case *Source:
        enabled = false
    case *MemCell:
        enabled = !t.state
    }
    if(enabled) {
        (*b2).Power()
    }
    return []point{p2}
}
// ------------

/* Xor */
type Xor struct {
    dir Direction
}

// When it's not updated probably it doesn't have signal
func (Xor) Check() bool { return false }

func (Xor) Power() {}

func (Xor) forcedUpdate() bool { return false }
func (xor *Xor) Update(grid [3][3](*Cell)) []point {
    p1 := dir2point(rotateDir(xor.dir, LEFT), point{1, 1})
    p2 := dir2point(rotateDir(xor.dir, LEFT), point{1, 1})
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

func (And) Power() {}
func (And) forcedUpdate() bool { return false }
func (and *And) Update(grid [3][3](*Cell)) []point {
    p1 := dir2point(rotateDir(and.dir, LEFT), point{1, 1})
    p2 := dir2point(rotateDir(and.dir, LEFT), point{1, 1})
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

func (Block) Check() bool { return false }
func (Block) Power() {}
func (Block) forcedUpdate() bool { return false }
func (block *Block) Update(grid [3][3](*Cell)) []point {
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
func (g Get) Check() bool { return g.state }
func (Get) Power() {}

// It won't be loaded directly so I'll update it forced
func (Get) forcedUpdate() bool { return true }

func (get *Get) Update(grid [3][3](*Cell)) []point {
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
