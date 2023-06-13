package core

/* Wire */
type Wire struct {
    dir Direction
    lit bool
}

func (a Wire) Dir() Direction { return a.dir }
func (a *Wire) SetDir(dir Direction) { a.dir = dir }
func (a Wire) updateQueue() []Point {
    return []Point{
        dir2point(a.dir, Point{1,1}),
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
    cell := grid[p.X][p.Y]
    (*cell).Power()
    a.lit = false
}
// ------------

/* Forward-Side Wire */
type FrwdSide struct {
    lit bool
    dir Direction
}

func (a FrwdSide) updateQueue() []Point {
    return []Point{
        dir2point(a.dir, Point{1,1}),
        dir2point(rotateDir(a.dir, RIGHT), Point{1,1}),
    }
}

func (a FrwdSide) Dir() Direction { return a.dir }

func (a *FrwdSide) SetDir(dir Direction) { a.dir = dir }

func (fr FrwdSide) Check() bool {
    return fr.lit
}

func (fr *FrwdSide) Power() {
    fr.lit = true
}

// Doesn't forces updates on other cells
func (FrwdSide) forcedUpdate() bool {
    return false
}

// Pass signal to a cell that it faced with, as well as on the right side
//[.][O][.] X - arrow; I - input; O - output
//[.][X][O]
//[.][I][.]
func (a *FrwdSide) Update(grid _lgrid) {
    if(!a.lit) { return }
    p1 := a.updateQueue()[0]
    p2 := a.updateQueue()[1]
    cell := grid[p1.X][p1.Y]
    (*cell).Power()
    cell = grid[p2.X][p2.Y]
    (*cell).Power()
    a.lit = false
}
// ------------

/* Cross Wire */
type Cross struct {
    lit bool
    dir Direction
}

func (a Cross) updateQueue() []Point {
    return []Point{
        dir2point(rotateDir(a.dir, LEFT), Point{1,1}),
        dir2point(a.dir, Point{1,1}),
        dir2point(rotateDir(a.dir, RIGHT), Point{1,1}),
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
        (*grid[p.X][p.Y]).Power()
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
func (a Angled) updateQueue() []Point {
    return []Point{
        dir2point(rotateDir(a.dir, LEFT), dir2point(a.dir, Point{1,1})),
    }
}

// Pass signal to a cell on the top left relatively of the that it faced
//[O][.][.] X - arrow; I - input; O - output
//[.][X][.]
//[.][I][I]
//P.S. IDK where input is
//P.S.S anywhere
func (a *Angled) Update(grid _lgrid) {
    if(!a.lit) { return }
    p := a.updateQueue()[0]
    cell := grid[p.X][p.Y]
    (*cell).Power()
    a.lit = false
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
        (*g[p.X][p.Y]).Power()
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
func (a Double) updateQueue() []Point {
    return []Point{
        dir2point(a.dir, Point{1, 1}),
    }
}
//-----------

