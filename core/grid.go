package core

type Grid struct {
    // Contains all cells with arrows
    cells []Cell

    width, height uint

    // Used for smart cells loading, could be changed in future
    updatePoints []point

    // enables fast mode, increases FPS but could be unstable
    FAST bool
}

// Get size of the grid
func (grid Grid) Dimensions() (uint, uint) {
    return grid.width, grid.height
}

// Initialization of the grid
func (grid *Grid) Init(w, h uint) {
    grid.width  = w
    grid.height = h
    for y := uint(0); y < h; y++ {
        for x := uint(0); x < w; x++ {
            grid.cells = append(grid.cells, Cell(None{}))
        }
    }
}

// Returns pointer to cell at x, y
func (grid *Grid) GetCell(x, y uint) *Cell {
    index := y*grid.width + x
    return &grid.cells[index]
}

func (grid *Grid) RecountUpdate() {
    for x := uint(0); x < grid.width; x++ {
        for y := uint(0); y < grid.height; y++ {
            if (*grid.GetCell(x, y)).forcedUpdate() {
                grid.updatePoints = append(grid.updatePoints, point{x, y})
            }
        }
    }
}

// Updates the grid
func (grid *Grid) Update() {
    // Gets list of points with forced update
    // XXX: Could be moved to SetCell and Init methods for a better perfomance?
    // XXX: I'd like to make FAST mode stable
    if(!grid.FAST) {
        grid.RecountUpdate()
    }

    // New list of update points for that points that are not forced
    newUpdate := []point{}
    for _, p := range(grid.updatePoints) {
        cell := grid.GetCell(p.x, p.y)

        // Passing grid 3x3 around cell to Update method
        // TODO: 5x5?
        var g _lgrid

        // Local x and y in grid 3x3
        var i, j uint = 0, 0

        // TODO: make it flexible
        // Loop over area 3x3 around cell
        for x := p.x-1; x <= p.x+1; x++ {
            for y := p.y-1; y <= p.y+1; y++ {
                // Because x & y are uint, they can't be below 0, so just check
                // is it greater than width || height
                if x > grid.width || y > grid.height { continue }
                g[j][i] = grid.GetCell(x, y)
                i++
            }
            i = 0
            j++
        }

        // Cell returns a list of points that needs to be loaded in next tick
        points := (*cell).Update(g)
        for _, rp := range(points) {
            newUpdate = append(newUpdate, point{rp.x+p.x-1, rp.y+p.y-1})
        }
        if(!grid.FAST) {
            if((*cell).forcedUpdate()) {
                newUpdate = append(newUpdate, p)
            }
        }
    }
    grid.updatePoints = newUpdate
}
