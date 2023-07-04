package core

/* Not */
type Not struct {
    Direction Direction
    Lit bool
}

func (a Not) Dir() Direction { return a.Direction }

func (a Not) UpdateQueue() []Point {
    return []Point{
        dir2point(a.Direction, Point{1, 1}),
    }
}
func (a *Not) SetDir(dir Direction) { a.Direction = dir }

func (a Not) forcedUpdate() bool { return true }

func (a Not) Check() bool { return !a.Lit }

func (a *Not) Power() {
    a.Lit = true
}

func (not *Not) Update(grid _lgrid) {
    // p1 -[NOT]-> p2
    p := not.UpdateQueue()[0]

    b := grid[p.X][p.Y]
    if(!not.Lit) {
        (*b).Power()
    }
    not.Lit = false
}
// ------------

/* Xor */
type Xor struct {
    Direction Direction
    Lit bool
    lcount uint
}

func (a Xor) Dir() Direction { return a.Direction }

func (a Xor) UpdateQueue() []Point {
    return []Point{
        dir2point(a.Direction, Point{1, 1}),
    }
}
func (a *Xor) SetDir(dir Direction) { a.Direction = dir }

func (a Xor) Check() bool { return a.Lit }

func (a *Xor) Power() {
    a.lcount++
    if(a.lcount == 1) {
        a.Lit = true
    } else {
        a.Lit = false
    }
}

func (Xor) forcedUpdate() bool { return false }

func (xor *Xor) Update(grid _lgrid) {
    rp := xor.UpdateQueue()[0]
    rb := grid[rp.X][rp.Y]
    if(xor.Lit) {
        (*rb).Power()
        xor.Lit = false
    }
    xor.lcount = 0
}
// ------------

/* And */
type And struct {
    Direction Direction
    Lit bool

    // lcount isn't global because it's temporary state and should be reseted every tick any way
    lcount uint
}

func (a And) UpdateQueue() []Point {
    return []Point{
        dir2point(a.Direction, Point{1, 1}),
    }
}
func (a And) Check() bool { return a.Lit }

func (a And) Dir() Direction { return a.Direction }

func (a *And) SetDir(dir Direction) { a.Direction = dir }

func (a *And) Power() {
    a.lcount++
    if(a.lcount >= 2) {
        a.Lit = true
    }
}

func (And) forcedUpdate() bool { return false }

func (and *And) Update(grid _lgrid) {
    rp := and.UpdateQueue()[0]
    rb := grid[rp.X][rp.Y]
    if(and.Lit) {
        (*rb).Power()
        and.Lit = false
    }
    and.lcount = 0
}
// ------------

/* Block */
type Block struct {
    Direction Direction
    Lit bool
}
// TODO: redo logic

func (a Block) UpdateQueue() []Point {
    return []Point{
        dir2point(a.Direction, Point{1,1}),
    }
}

func (a Block) Dir() Direction { return a.Direction }

func (a *Block) SetDir(dir Direction) { a.Direction = dir }

func (a Block) Check() bool { return a.Lit}
func (Block) Power() {}
func (Block) forcedUpdate() bool { return false }
func (block *Block) Update(grid _lgrid) {
    // FIXME: This version of block won't handle Cross, Angled, etc
    //  p1 -[block]?-> p2
    p1 := dir2point(rotateDir(block.Direction, BACK), Point{1,1})
    p2 := block.UpdateQueue()[0]
    b1 := grid[p1.X][p1.Y]
    b2 := grid[p2.X][p2.Y]
    if(block.Lit) {
        switch t := (*b2).(type) {
        case *Wire:
            t.Lit = false
        case *MemCell:
            t.State = false
        }
        block.Lit = false
    }
    if((*b1).Check()) {
        block.Lit = true
    }
}
// ------------


