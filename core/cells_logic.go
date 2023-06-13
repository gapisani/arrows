package core

/* Not */
type Not struct {
    dir Direction
    lit bool
}

func (a Not) Dir() Direction { return a.dir }

func (a Not) updateQueue() []Point {
    return []Point{
        dir2point(a.dir, Point{1, 1}),
    }
}
func (a *Not) SetDir(dir Direction) { a.dir = dir }

func (a Not) forcedUpdate() bool { return true }

func (a Not) Check() bool { return a.lit }

func (Not) Power() {}

func (not *Not) Update(grid _lgrid) {
    // p1 -[NOT]-> p2
    p1 := dir2point(rotateDir(not.dir, BACK), Point{1, 1})
    p2 := not.updateQueue()[0]

    b1 := grid[p1.X][p1.Y]
    b2 := grid[p2.X][p2.Y]
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

func (a Xor) updateQueue() []Point {
    return []Point{
        dir2point(a.dir, Point{1, 1}),
    }
}
func (a *Xor) SetDir(dir Direction) { a.dir = dir }

func (a Xor) Check() bool { return a.lit }

func (a *Xor) Power() {
    a.lcount++
    if(a.lcount == 1) {
        a.lit = true
    } else {
        a.lit = false
    }
}

func (Xor) forcedUpdate() bool { return false }

func (xor *Xor) Update(grid _lgrid) {
    rp := xor.updateQueue()[0]
    rb := grid[rp.X][rp.Y]
    if(xor.lit) {
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

func (a And) updateQueue() []Point {
    return []Point{
        dir2point(a.dir, Point{1, 1}),
    }
}
func (a And) Check() bool { return a.lit }

func (a And) Dir() Direction { return a.dir }

func (a *And) SetDir(dir Direction) { a.dir = dir }

func (a *And) Power() {
    a.lcount++
    if(a.lcount >= 2) {
        a.lit = true
    }
}

func (And) forcedUpdate() bool { return false }

func (and *And) Update(grid _lgrid) {
    rp := and.updateQueue()[0]
    rb := grid[rp.X][rp.Y]
    if(and.lit) {
        (*rb).Power()
        and.lit = false
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

func (a Block) updateQueue() []Point {
    return []Point{
        dir2point(a.dir, Point{1,1}),
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
    p1 := dir2point(rotateDir(block.dir, BACK), Point{1,1})
    p2 := block.updateQueue()[0]
    b1 := grid[p1.X][p1.Y]
    b2 := grid[p2.X][p2.Y]
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


