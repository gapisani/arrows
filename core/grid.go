package core

type Grid struct {
    // Contains all cells with arrows
    cells []Cell

    width, height uint

    // Used for smart block loading, could be changed in future
    updatePoints []point
}

func (grid *Grid) Init(w, h uint) {
    grid.width  = w
    grid.height = h
    for x := uint(0); x < w; x++ {
        for y := uint(0); y < h; y++ {
            grid.cells = append(grid.cells, Cell(None{}))
        }
    }
}

func (grid *Grid) GetCell(x, y uint) *Cell {
    index := y*grid.width + x
    return &grid.cells[index]
}

func (grid *Grid) Update() {
    for x := uint(0); x < grid.width; x++ {
        for y := uint(0); y < grid.height; y++ {
            if (*grid.GetCell(x, y)).forcedUpdate() {
                grid.updatePoints = append(grid.updatePoints, point{x, y})
            }
        }
    }
    newUpdate := []point{}
    for _, p := range(grid.updatePoints) {
        cell := grid.GetCell(p.x, p.y)

        // Passing grid 3x3 around cell to Update method
        var g [3][3](*Cell)
        var i, j uint = 0, 0
        for x := p.x-1; x <= p.x+1; x++ {
            for y := p.y-1; y <= p.y+1; y++ {
                if x > grid.width || y > grid.height { continue }
                g[i][j] = grid.GetCell(x, y)
                i++
            }
            i = 0
            j++
        }
        for _, rp := range((*cell).Update(g)) {
            newUpdate = append(newUpdate, point{rp.x+p.x, rp.y+p.y})
        }
    }
    grid.updatePoints = newUpdate
}
