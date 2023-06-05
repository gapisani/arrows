package core_test

import (
	"github.com/gapisani/arrows/core"
	"testing"
    "fmt"
)

func rotArr(arr rune, dir core.Direction) string {
    switch dir {
    case core.WEST:
        return string(arr)
    case core.NORTH:
        return string(arr+1)
    case core.EAST:
        return string(arr+2)
    case core.SOUTH:
        return string(arr+3)
    }
    return string(arr)
}

func render(g core.Grid) {
    w, h := g.Dimensions()
    for x := uint(0); x < w; x++ {
        for y := uint(0); y < h; y++ {
            cell := g.GetCell(x, y)
            dir := (*cell).Dir()
            lit := (*cell).Check()
            switch (*cell).(type) {
            case *core.Wire:
                if(lit) {
                    fmt.Print(rotArr('⬅', dir))
                } else {
                    fmt.Print(rotArr('←', dir))
                }
            case *core.Source:
                fmt.Print("@")
            case *core.None:
                fmt.Print(" ")
            default:
                fmt.Print("x")
            }
        }
        fmt.Println()
    }
}

func TestMain(t *testing.T) {
    g := core.Grid{}
    g.Init(10, 10)
    w, h := g.Dimensions()
    for y := uint(0); y < h; y++ {
        for x := uint(0); x < w; x++ {
            *g.GetCell(x, y) = core.Cell(&core.None{})
        }
    }
    *g.GetCell(5, 5) = core.Cell(&core.Source{})
    *g.GetCell(4, 5) = core.Cell(&core.Wire{})
    *g.GetCell(3, 5) = core.Cell(&core.Wire{})
    render(g)
    g.Update()
    render(g)
    g.Update()
    render(g)
    g.Update()
    render(g)
    g.Update()
    render(g)
}
