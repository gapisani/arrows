package core_test

import (
	"fmt"
	"strings"
	"testing"
    "runtime"

	"github.com/gapisani/arrows/core"
)

func rotArr(arr byte, dir core.Direction) byte {
    switch dir {
    case core.WEST:
        return arr
    case core.NORTH:
        return arr+1
    case core.EAST:
        return arr+2
    case core.SOUTH:
        return arr+3
    }
    return arr
}

func render(g core.Grid, points []core.Point) {
    w, h := g.Dimensions()
    screen := []byte(strings.Repeat(strings.Repeat(".", int(w)) + "\n", int(h)))
    for _, p := range(points) {
        cell := g.GetCell(p.X, p.Y)
        if(cell == nil) { continue }
        lit := (*cell).Check()
        chr := byte('x')
        switch (*cell).(type) {
        case *core.Wire:
            if(lit) {
                chr = '^'
            } else {
                chr = '>'
            }
        case *core.Source:
            chr = '@'
        case *core.None:
            chr = '.'
        case *core.MemCell:
            if(lit) {
                chr = '#'
            } else {
                chr = 'O'
            }
        }
        screen[p.Y*(w+1)+p.X] = chr
    }
    fmt.Println(string(screen), "---------")
}

func TestSerpinski(t *testing.T) {
    g := core.Grid{}
    g.Init(50, 50)
    w, h := g.Dimensions()
    for i := uint(0); i < h; i++ {
        *g.GetCell(0, i) = core.Cell(&core.MemCell{})
        *g.GetCell(1, i) = core.Cell(&core.Get{})
        (*g.GetCell(1, i)).SetDir(core.EAST)
        for j := uint(2); j < w; j++ {
            *g.GetCell(j, i) = core.Cell(&core.Wire{})
            (*g.GetCell(j, i)).SetDir(core.EAST)
        }
    }
    *g.GetCell(0, h-1) = core.Cell(&core.Source{})
    g.RecountUpdate()
    for t := 0; t <= 30; t++ {
        p := g.Update()
        render(g, p)
    }
}

func TestWire(t *testing.T) {
    g := core.Grid{}
    g.Init(10, 10)
    _, h := g.Dimensions()
    for i := uint(1); i < h; i++ {
        (*g.GetCell(5, i)) = &core.Wire{}
        (*g.GetCell(5, i)).SetDir(core.NORTH)
        g.RecountUpdate()
    }
    *g.GetCell(5, h-2) = core.Cell(&core.Source{})
    g.RecountUpdate()
    for t := 0; t <= 5; t++ {
        g.Update()
        // render(g)
        // time.Sleep(time.Millisecond * 300)
    }
    *g.GetCell(1, h-1) = core.Cell(&core.None{})
    for t := 0; t <= 10; t++ {
        g.Update()
        // render(g)
        // time.Sleep(time.Millisecond * 300)
    }
}

func _TestEdges(t *testing.T) {
    g := core.Grid{}
    g.Init(3, 3)
    *g.GetCell(0, 0) = &core.Get{}
    *g.GetCell(2, 2) = &core.Xor{}
    p := g.Update()
    render(g, p)
}

func _TestUpdate(t *testing.T) {
    g := core.Grid{}
    g.Init(5, 5)
    *g.GetCell(3, 4) = &core.Source{}
    g.RecountUpdate()
    *g.GetCell(3, 3) = &core.Wire{}
    g.RecountUpdate()
    *g.GetCell(3, 2) = &core.Wire{}
    g.RecountUpdate()
    for t:=0; t<3; t++ {
        g.Update()
    }
}

func _Bench_markSerpinski(b *testing.B) {
    g := core.Grid{}
    g.Init(5000, 5000)
    w, h := g.Dimensions()
    for i := uint(h)-1; i > 1; i-- {
        *g.GetCell(0, i) = core.Cell(&core.MemCell{})
        *g.GetCell(1, i) = core.Cell(&core.Get{})
        (*g.GetCell(1, i)).SetDir(core.EAST)
        for j := uint(2); j < w-1; j++ {
            *g.GetCell(j, i) = core.Cell(&core.Wire{})
            (*g.GetCell(j, i)).SetDir(core.EAST)
        }
    }
    *g.GetCell(1, h-1) = core.Cell(&core.Source{})
    g.RecountUpdate()
    for t := 0; t <= 100; t++ {
        g.Update()
    }
}

func TestRam(t *testing.T) {
    var m1, m2 runtime.MemStats
    g := core.Grid{}
    runtime.GC()
    runtime.ReadMemStats(&m1)
    g.Init(5000, 5000);
    runtime.ReadMemStats(&m2)
    fmt.Println("total:", m2.TotalAlloc - m1.TotalAlloc)
    fmt.Println("mallocs:", m2.Mallocs - m1.Mallocs)
}

// func BenchmarkNormal(b *testing.B) {
//     g := core.Grid{}
//     g.Init(50, 50)
//     w, h := g.Dimensions()
//     for i := uint(h)-1; i > 0; i-- {
//         *g.GetCell(1, i) = core.Cell(&core.MemCell{})
//         *g.GetCell(2, i) = core.Cell(&core.Get{})
//         (*g.GetCell(2, i)).SetDir(core.EAST)
//         for j := uint(3); j < w; j++ {
//             *g.GetCell(j, i) = core.Cell(&core.Wire{})
//             (*g.GetCell(j, i)).SetDir(core.EAST)
//         }
//     }
//     *g.GetCell(1, h-1) = core.Cell(&core.Source{})
//     g.FAST = false
//     for t := 0; t <= 100; t++ {
//         g.Update()
//     }
// }
