package core

type Grid struct {
    // Contains all cells with arrows
    cells []Cell

    width, height uint

    // Used for smart cells loading, could be changed in future
    updateQueue []Point
}

// Get size of the grid
func (grid Grid) Dimensions() (uint, uint) {
    return grid.width, grid.height
}

// Initialization of the grid
func (grid *Grid) Init(w, h uint) {
    grid.width  = w
    grid.height = h
    grid.cells = make([]Cell, w*h)
}

// Returns pointer to cell at x, y
func (grid *Grid) GetCell(x, y uint) *Cell {
    index := y*grid.width + x
    if(index >= uint(len(grid.cells))) {
        return nil
    }
    if(grid.cells[index] == nil) {
        grid.cells[index] = &None{}
    }
    return &grid.cells[index]
}

func (grid *Grid) AddUpdate(x, y uint) {
    grid.updateQueue = append(grid.updateQueue, Point{x, y})
    grid.updateQueueClean()
}

func (grid *Grid) RecountUpdate() {
    for x := uint(0); x < grid.width; x++ {
        for y := uint(0); y < grid.height; y++ {
            if (*grid.GetCell(x, y)).forcedUpdate() {
                grid.updateQueue = append(grid.updateQueue, Point{x, y})
            }
        }
    }
    grid.updateQueueClean()
}

func (grid *Grid) updateQueueClean() {
    encountered := map[Point]bool{}
    cells := []Point{}
    for _, v := range grid.updateQueue {
        if !encountered[v] {
            encountered[v] = true
            cells = append(cells, v)
        }
    }
    grid.updateQueue = cells
}

// Updates the grid
func (grid *Grid) Update() []Point {
    if(len(grid.updateQueue) == 0) { return []Point{} }
    // Gets list of points with forced update
    grid.updateQueueClean()

    // New list of update points for that points that are not forced
    newUpdate := []Point{}
    for _, p := range(grid.updateQueue) {
        cell := grid.GetCell(p.X, p.Y)
        if(cell == nil) { continue }

        // Passing grid around cell to Update method
        var g _lgrid

        // Local x and y in grid 3x3
        var i, j uint = 0, 0

        // TODO: make it flexible
        // Loop over area 3x3 around cell
        for x := int(p.X-1); x <= int(p.X+1); x++ {
            for y := int(p.Y-1); y <= int(p.Y+1); y++ {
                if(x < 0 || x >= int(grid.width) || y < 0 || y >= int(grid.height)) {
                    c := Cell(None{})
                    g[j][i] = &c
                } else {
                    g[j][i] = grid.GetCell(uint(x), uint(y))
                }
                i++
            }
            i = 0
            j++
        }

        (*cell).Update(g)
        // list of points that needs to be loaded in next tick
        for _, rp := range((*cell).updateQueue()) {
            newUpdate = append(newUpdate, Point{rp.X+p.X-1, rp.Y+p.Y-1})
        }
        if((*cell).forcedUpdate()) {
            newUpdate = append(newUpdate, p)
        }
    }
    // Swapping buffers instead of copying
    grid.updateQueue, newUpdate = newUpdate, grid.updateQueue
    // Now it returns cells that was updated at this tick
    return newUpdate
}
